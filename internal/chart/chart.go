package chart

import (
	"errors"
	"image/color"
	"math"
	"time"

	"github.com/s-daehling/fyne-charts/internal/axis"
	"github.com/s-daehling/fyne-charts/internal/series"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type PlaneType string

const (
	CartesianPlane PlaneType = "Cartesian"
	PolarPlane     PlaneType = "Polar"
)

type FromType string

const (
	Numerical    FromType = "Numerical"
	Temporal     FromType = "Temporal"
	Categorical  FromType = "Categorical"
	Proportional FromType = "Proportional"
)

type BaseChart struct {
	name          string
	label         *canvas.Text
	fromAx        *axis.Axis
	toAx          *axis.Axis
	series        []series.Series
	changed       bool
	autoFromRange bool
	autoToRange   bool
	autoOrigin    bool
	legendVisible bool
	planeType     PlaneType
	fromType      FromType
	rast          *canvas.Raster
	render        fyne.WidgetRenderer
}

func EmptyBaseChart(pType PlaneType, fType FromType) (base *BaseChart) {
	base = &BaseChart{
		label:         canvas.NewText("", theme.Color(theme.ColorNameForeground)),
		changed:       false,
		autoFromRange: true,
		autoToRange:   true,
		autoOrigin:    true,
		legendVisible: true,
		planeType:     pType,
		fromType:      fType,
	}
	if pType == CartesianPlane {
		base.fromAx = axis.EmptyAxis("", axis.CartesianAxis)
		base.toAx = axis.EmptyAxis("", axis.CartesianAxis)
		base.rast = canvas.NewRasterWithPixels(base.PixelGenCartesian)
		base.render = EmptyCartesianRenderer(base)
	} else {
		base.fromAx = axis.EmptyAxis("", axis.PolarPhiAxis)
		base.toAx = axis.EmptyAxis("", axis.PolarRAxis)
		base.rast = canvas.NewRasterWithPixels(base.PixelGenPolar)
		base.render = EmptyPolarRenderer(base)
	}
	if fType == Proportional {
		base.HideFromAxis()
		base.HideToAxis()
	}
	base.updateRangeAndOrigin()
	return
}

func (base *BaseChart) GetRenderer() fyne.WidgetRenderer {
	return base.render
}

func (base *BaseChart) CartesianOrientation() (transposed bool) {
	transposed = false
	return
}

func (base *BaseChart) PolarOrientation() (rot float64, mathPos bool) {
	rot = 0
	mathPos = true
	return
}

func (base *BaseChart) SeriesExist(n string) (exist bool) {
	exist = false
	for i := range base.series {
		if base.series[i].Name() == n {
			exist = true
			break
		}
	}
	return
}

func (base *BaseChart) DeleteSeries(name string) {
	newSeries := make([]series.Series, 0)
	for i := range base.series {
		if base.series[i].Name() != name {
			newSeries = append(newSeries, base.series[i])
		} else {
			base.series[i].Delete()
		}
	}
	base.series = newSeries
	base.DataChange()
}

func (base *BaseChart) fromAxis() (from *axis.Axis) {
	from = base.fromAx
	return
}

func (base *BaseChart) toAxis() (to *axis.Axis) {
	to = base.toAx
	return
}

func (base *BaseChart) title() (ct Title) {
	ct.Name = base.name
	ct.Label = base.label
	return
}

func (base *BaseChart) cartesianObjects() (canObj []fyne.CanvasObject) {
	// objects will be drawn in the same order as added here

	// first get all objects from the series
	if base.legendVisible {
		lEntries := base.legendEntries()
		for i := range lEntries {
			canObj = append(canObj, lEntries[i].Button, lEntries[i].Label)
		}
	}
	canObj = append(canObj, base.rast)
	rects := base.cartesianRects()
	for i := range rects {
		canObj = append(canObj, rects[i].Rect)
	}
	edges := base.cartesianEdges()
	for i := range edges {
		canObj = append(canObj, edges[i].Line)
	}
	nodes := base.cartesianNodes()
	for i := range nodes {
		canObj = append(canObj, nodes[i].Dot)
	}
	texts := base.cartesianTexts()
	for i := range texts {
		canObj = append(canObj, texts[i].Text)
	}

	// add axis elements
	canObj = append(canObj, base.fromAx.Objects()...)
	canObj = append(canObj, base.toAx.Objects()...)

	// add chart title and axis titles
	if base.name != "" {
		canObj = append(canObj, base.label)
	}
	return
}

func (base *BaseChart) cartesianNodes() (ns []series.CartesianNode) {
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	for i := range base.series {
		ns = append(ns, base.series[i].CartesianNodes(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (base *BaseChart) cartesianEdges() (es []series.CartesianEdge) {
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	for i := range base.series {
		es = append(es, base.series[i].CartesianEdges(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (base *BaseChart) cartesianRects() (as []series.CartesianRect) {
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	for i := range base.series {
		as = append(as, base.series[i].CartesianRects(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (base *BaseChart) cartesianTexts() (ts []series.CartesianText) {
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	for i := range base.series {
		ts = append(ts, base.series[i].CartesianTexts(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (base *BaseChart) polarObjects() (canObj []fyne.CanvasObject) {
	// objects will be drawn in the same order as added here

	// first get all objects from the series
	if base.legendVisible {
		lEntries := base.legendEntries()
		for i := range lEntries {
			canObj = append(canObj, lEntries[i].Button, lEntries[i].Label)
		}
	}
	canObj = append(canObj, base.rast)
	edges := base.polarEdges()
	for i := range edges {
		canObj = append(canObj, edges[i].Line)
	}
	nodes := base.polarNodes()
	for i := range nodes {
		canObj = append(canObj, nodes[i].Dot)
	}
	texts := base.polarTexts()
	for i := range texts {
		canObj = append(canObj, texts[i].Text)
	}

	// add axis elements
	canObj = append(canObj, base.fromAx.Objects()...)
	canObj = append(canObj, base.toAx.Objects()...)

	// add chart title and axis titles
	if base.name != "" {
		canObj = append(canObj, base.label)
	}
	return
}

func (base *BaseChart) polarNodes() (ns []series.PolarNode) {
	phiMin, phiMax := base.fromAx.NRange()
	rMin, rMax := base.toAx.NRange()
	for i := range base.series {
		ns = append(ns, base.series[i].PolarNodes(phiMin, phiMax, rMin, rMax)...)
	}
	return
}

func (base *BaseChart) polarEdges() (es []series.PolarEdge) {
	phiMin, phiMax := base.fromAx.NRange()
	rMin, rMax := base.toAx.NRange()
	for i := range base.series {
		es = append(es, base.series[i].PolarEdges(phiMin, phiMax, rMin, rMax)...)
	}
	return
}

func (base *BaseChart) polarTexts() (ts []series.PolarText) {
	phiMin, phiMax := base.fromAx.NRange()
	rMin, rMax := base.toAx.NRange()
	for i := range base.series {
		ts = append(ts, base.series[i].PolarTexts(phiMin, phiMax, rMin, rMax)...)
	}
	return
}

func (base *BaseChart) chartRaster() (rs *canvas.Raster) {
	rs = base.rast
	return
}

func (base *BaseChart) legendEntries() (les []series.LegendEntry) {
	for i := range base.series {
		les = append(les, base.series[i].LegendEntries()...)
	}
	return
}

func (base *BaseChart) refreshThemeColor() {
	base.fromAx.RefreshThemeColor()
	base.toAx.RefreshThemeColor()
	base.label.Color = theme.Color(theme.ColorNameForeground)
	for i := range base.series {
		base.series[i].RefreshThemeColor()
	}
}

// func (base *BaseChart) hasChanged() (c bool) {
// 	return base.changed
// }

// func (base *BaseChart) resetHasChanged() {
// 	base.changed = false
// }

// func (base *BaseChart) widgetSize() (s fyne.Size) {
// 	return base.Size()
// }

func (base *BaseChart) SetTitle(l string) {
	base.name = l
}

func (base *BaseChart) ShowLegend() {
	base.legendVisible = true
	base.DataChange()
}

func (base *BaseChart) HideLegend() {
	base.legendVisible = false
	base.DataChange()
}

func (base *BaseChart) legendVisibility() (v bool) {
	v = base.legendVisible
	return
}

func (base *BaseChart) HideFromAxis() {
	base.fromAx.Hide()
}

func (base *BaseChart) ShowFromAxis() {
	base.fromAx.Show()
}

func (base *BaseChart) SetFromAxisLabel(l string) {
	base.fromAx.SetLabel(l)
}

func (base *BaseChart) HideToAxis() {
	base.toAx.Hide()
}

func (base *BaseChart) ShowToAxis() {
	base.toAx.Show()
}

func (base *BaseChart) SetToAxisLabel(l string) {
	base.toAx.SetLabel(l)
}

func (base *BaseChart) SetFromNRange(min float64, max float64) (err error) {
	if min > max {
		err = errors.New("invalid range")
		return
	}
	if !base.autoOrigin &&
		(base.fromAx.NOrigin() < min || base.fromAx.NOrigin() > max) {
		err = errors.New("previously defined origin not in range")
		return
	}
	base.autoFromRange = false
	base.fromAx.SetNRange(min, max)
	base.DataChange()
	return
}

func (base *BaseChart) SetFromTRange(min time.Time, max time.Time) (err error) {
	if min.After(max) {
		err = errors.New("invalid range")
		return
	}
	if !base.autoOrigin &&
		(base.fromAx.TOrigin().Before(min) || base.fromAx.TOrigin().After(max)) {
		err = errors.New("previously defined origin not in range")
		return
	}
	base.autoFromRange = false
	base.fromAx.SetTRange(min, max)
	base.DataChange()
	return
}

func (base *BaseChart) SetFromCRange(cs []string) (err error) {
	if len(cs) < 1 {
		err = errors.New("invalid range")
		return
	}
	base.autoFromRange = false
	base.fromAx.SetCRange(cs)
	base.DataChange()
	return
}

func (base *BaseChart) SetAutoFromRange() {
	base.autoFromRange = true
	base.DataChange()
}

func (base *BaseChart) calculateAutoFromNRange() {
	var min, max float64
	if !base.autoOrigin {
		// if origin was set by user, init range with x origin
		min = base.fromAx.NOrigin()
		max = min
	}
	init := false
	for i := range base.series {
		isEmpty, sMin, sMax := base.series[i].NRange()
		if isEmpty {
			continue
		}
		if !init {
			if base.autoOrigin {
				// range not inited yet and no user set origin -> init range now
				min = sMin
				max = sMax
			}
			init = true
		}
		if min > sMin {
			min = sMin
		}
		if max < sMax {
			max = sMax
		}
	}

	if !init {
		if base.autoOrigin {
			// range around 0
			min = -1
			max = 1
		} else {
			// range around user specified origin
			min -= 1
			max += 1
		}
	}

	// make sure the min and max are not equal
	absMin := math.Abs(min)
	if math.Abs(max) < absMin {
		absMin = math.Abs(max)
	}
	r := math.Abs(max - min)
	if r*1000 < absMin {
		min = max - 1
		max = max + 1
	}
	base.fromAx.SetNRange(min, max)
}

func (base *BaseChart) calculateAutoFromTRange() {
	var min, max time.Time
	if !base.autoOrigin {
		// if origin was set by user, init range with x origin
		min = base.fromAx.TOrigin()
		max = min
	}
	init := false

	for i := range base.series {
		isEmpty, sMin, sMax := base.series[i].TRange()
		// todo check if min and max are really type time.Time
		if isEmpty {
			continue
		}
		if !init {
			if base.autoOrigin {
				// range not inited yet and no user set origin -> init range now
				min = sMin
				max = sMax
			}
			init = true
		}
		if min.After(sMin) {
			min = sMin
		}
		if max.Before(sMax) {
			max = sMax
		}
	}
	if !init {
		if base.autoOrigin {
			// range around 0
			min = time.Now().Add(-time.Hour)
			max = time.Now().Add(time.Hour)
		} else {
			// range around user specified origin
			min = min.Add(-time.Hour)
			max = min.Add(time.Hour)
		}
	}

	// make sure the min and max are not equal
	if min.Equal(max) {
		min = min.Add(-time.Second)
		max = max.Add(time.Second)
	}

	base.fromAx.SetTRange(min, max)
}

func (base *BaseChart) calculateAutoFromCRange() {
	cs := []string{}
	for i := range base.series {
		scs := base.series[i].CRange()
		for j := range scs {
			exist := false
			for k := range cs {
				if scs[j] == cs[k] {
					exist = true
					break
				}
			}
			if !exist {
				cs = append(cs, scs[j])
			}
		}
	}
	base.fromAx.SetCRange(cs)
}

func (base *BaseChart) SetToRange(min float64, max float64) (err error) {
	if max < 0 {
		err = errors.New("invalid range")
		return
	}
	if !base.autoOrigin &&
		(base.toAx.NOrigin() < min || base.toAx.NOrigin() > max) {
		err = errors.New("previously defined origin not in range")
		return
	}
	base.autoToRange = false
	base.toAx.SetNRange(min, max)
	base.DataChange()
	return
}

func (base *BaseChart) SetAutoToRange() {
	base.autoToRange = true
	base.DataChange()
}

func (base *BaseChart) calculateAutoToRange() {
	var min, max float64
	if !base.autoOrigin {
		min = base.toAx.NOrigin()
		max = min
	}
	init := false
	for i := range base.series {
		isEmpty, sMin, sMax := base.series[i].ValRange()
		if isEmpty {
			continue
		}
		if !init {
			if base.autoOrigin {
				// range not inited yet and no user set origin -> init range now
				min = sMin
				max = sMax
			}
			init = true
		}
		if min > sMin {
			min = sMin
		}
		if max < sMax {
			max = sMax
		}
	}

	if !init {
		if base.autoOrigin {
			// range around 0
			min = -1
			max = 1
		} else {
			// range around user specified origin
			min -= 1
			max += 1
		}
	}

	// make sure the min and max are not equal
	absMin := math.Abs(min)
	if math.Abs(max) < absMin {
		absMin = math.Abs(max)
	}
	r := math.Abs(max - min)
	if r*1000 < absMin {
		min = max - 1
		max = max + 1
	}

	if base.planeType == PolarPlane {
		min = 0.0
	}
	base.toAx.SetNRange(min, max)
}

func (base *BaseChart) SetFromNTicks(ts []data.NumericalTick) {
	if len(ts) < 1 {
		return
	}
	base.fromAx.SetManualTicks()
	min := ts[0].N
	max := ts[0].N
	for i := range ts {
		if ts[i].N < min {
			min = ts[i].N
		}
		if ts[i].N > max {
			max = ts[1].N
		}
	}
	r := max - min
	orderOfMagn := -100
	// find upper limit for orderOfMagn
	for {
		if math.Pow10(orderOfMagn) < r {
			orderOfMagn++
		} else {
			break
		}
	}
	base.fromAx.SetNTicks(ts, orderOfMagn)
	base.DataChange()
}

func (base *BaseChart) SetFromTTicks(ts []data.TemporalTick, format string) {
	base.fromAx.SetTTicks(ts, format)
	base.DataChange()
}

func (base *BaseChart) SetAutoFromTicks(autoSupport bool) {
	base.fromAx.SetAutoTicks(autoSupport)
	base.DataChange()
}

func (base *BaseChart) SetToTicks(ts []data.NumericalTick) {
	if len(ts) < 1 {
		return
	}
	base.toAx.SetManualTicks()
	min := ts[0].N
	max := ts[0].N
	for i := range ts {
		if ts[i].N < min {
			min = ts[i].N
		}
		if ts[i].N > max {
			max = ts[1].N
		}
	}
	r := max - min
	orderOfMagn := -100
	// find upper limit for orderOfMagn
	for {
		if math.Pow10(orderOfMagn) < r {
			orderOfMagn++
		} else {
			break
		}
	}
	base.toAx.SetNTicks(ts, orderOfMagn)
	base.DataChange()
}

func (base *BaseChart) SetAutoToTicks(autoSupport bool) {
	base.toAx.SetAutoTicks(autoSupport)
	base.DataChange()
}

func (base *BaseChart) SetNOrigin(from float64, to float64) (err error) {
	nMinFrom, nMaxFrom := base.fromAx.NRange()
	if !base.autoFromRange && (from > nMaxFrom || from < nMinFrom) {
		err = errors.New("out of user defined range")
		return
	}
	nMinTo, nMaxTo := base.toAx.NRange()
	if !base.autoToRange && (to > nMaxTo || to < nMinTo) {
		err = errors.New("out of user defined range")
		return
	}
	base.autoOrigin = false
	base.toAx.SetNOrigin(to)
	base.fromAx.SetNOrigin(from)
	base.DataChange()
	return
}

func (base *BaseChart) SetTOrigin(from time.Time, to float64) (err error) {
	tMinFrom, tMaxFrom := base.fromAx.TRange()
	if !base.autoFromRange && (from.After(tMaxFrom) || from.Before(tMinFrom)) {
		err = errors.New("t out of user defined range")
		return
	}
	nMinTo, nMaxTo := base.toAx.NRange()
	if !base.autoToRange && (to > nMaxTo || to < nMinTo) {
		err = errors.New("out of user defined range")
		return
	}
	base.autoOrigin = false
	base.toAx.SetNOrigin(to)
	base.fromAx.SetTOrigin(from)
	base.DataChange()
	return
}

func (base *BaseChart) SetAutoOrigin() {
	base.autoOrigin = true
	base.DataChange()
}

func (base *BaseChart) calculateAutoNOrigin() {
	base.fromAx.AutoNOrigin()
	base.toAx.AutoNOrigin()
}

func (base *BaseChart) calculateAutoTOrigin() {
	base.fromAx.AutoTOrigin()
	base.calculateAutoNOrigin()
}

func (base *BaseChart) DataChange() {
	base.updateRangeAndOrigin()
	base.updateAxTicks()
	base.updateSeriesVariables()
	base.render.Refresh()
}

func (base *BaseChart) RasterVisibilityChange() {
	base.rast.Refresh()
}

func (base *BaseChart) resize(fromSpace float32, toSpace float32) {
	base.fromAx.SetSpace(fromSpace)
	base.toAx.SetSpace(toSpace)
	base.updateAxTicks()
}

func (base *BaseChart) updateRangeAndOrigin() {
	switch base.fromType {
	case Numerical:
		if base.autoToRange {
			base.calculateAutoToRange()
		}
		if base.autoFromRange || base.planeType == CartesianPlane {
			base.calculateAutoFromNRange()
		}
		if base.autoOrigin {
			base.calculateAutoNOrigin()
		}
	case Temporal:
		if base.autoToRange {
			base.calculateAutoToRange()
		}
		if base.autoFromRange {
			base.calculateAutoFromTRange()
		}
		if base.autoOrigin {
			base.calculateAutoTOrigin()
		}
	case Categorical:
		if base.autoToRange {
			base.calculateAutoToRange()
		}
		if base.autoFromRange {
			base.calculateAutoFromCRange()
		}
		if base.autoOrigin {
			base.calculateAutoNOrigin()
		}
	case Proportional:
		if base.autoOrigin {
			base.calculateAutoNOrigin()
		}
	}
}

func (base *BaseChart) updateAxTicks() {
	switch base.fromType {
	case Numerical:
		base.fromAx.AutoNTicks()
	case Temporal:
		base.fromAx.AutoTTicks()
		base.fromAx.ConvertTTickstoN()
	case Categorical:
		base.fromAx.AutoCTicks()
		base.fromAx.ConvertCTickstoN()
	case Proportional:
		base.fromAx.AutoNTicks()
	}
	base.toAx.AutoNTicks()
}

func (base *BaseChart) updateSeriesVariables() {
	nBarSeries := 0
	nPropSeries := 0
	maxBoxPoints := 5
	for i := range base.series {
		if _, ok := base.series[i].(*series.BarSeries); ok {
			nBarSeries++
		} else if _, ok := base.series[i].(*series.StackedBarSeries); ok {
			nBarSeries++
		} else if bs, ok := base.series[i].(*series.BoxSeries); ok {
			n := bs.NumberOfPoints()
			if n > maxBoxPoints {
				maxBoxPoints = n
			}
		} else if _, ok := base.series[i].(*series.ProportionalSeries); ok {
			nPropSeries++
		}
	}
	nFromMin, nFromMax := base.fromAx.NRange()
	nToMin, nToMax := base.toAx.NRange()
	catSize := (nFromMax - nFromMin) * 0.9
	numCategories := len(base.fromAxis().CRange())
	if numCategories > 0 {
		catSize = ((nFromMax - nFromMin) / float64(numCategories)) * 0.9
	}
	barWidth := catSize
	if nBarSeries > 0 {
		barWidth = catSize / float64(nBarSeries)
	}
	barOffset := -barWidth * (0.5 * float64(nBarSeries-1))
	propHeight := (nToMax - nToMin) / float64(nPropSeries)
	propOffset := 0.0
	boxWidth := (nFromMax - nFromMin) / float64(maxBoxPoints)
	for i := range base.series {
		if ls, ok := base.series[i].(*series.LollipopSeries); ok {
			if base.planeType == CartesianPlane {
				ls.SetValBaseNumerical(base.toAx.NOrigin())
			}
		} else if bs, ok := base.series[i].(*series.BarSeries); ok {
			if base.fromType == Categorical {
				bs.SetNumericalBarWidthAndShift(barWidth, barOffset)
				barOffset += barWidth
			}
		} else if sbs, ok := base.series[i].(*series.StackedBarSeries); ok {
			if base.fromType == Categorical {
				sbs.SetNumericalBarWidthAndShift(barWidth, barOffset)
				barOffset += barWidth
			}
			sbs.UpdateValOffset()
		} else if bs, ok := base.series[i].(*series.BoxSeries); ok {
			bs.SetWidth(boxWidth)
		} else if as, ok := base.series[i].(*series.AreaSeries); ok {
			as.SetValBaseNumerical(base.toAx.NOrigin())
		} else if ps, ok := base.series[i].(*series.ProportionalSeries); ok {
			ps.SetHeightAndOffset(propHeight*0.9, propOffset)
			propOffset += propHeight
			// ps.UpdateValOffest()
		}
	}

	switch base.fromType {
	case Temporal:
		for i := range base.series {
			base.series[i].ConvertTtoN(base.fromAx.TtoN)
		}
	case Categorical:
		for i := range base.series {
			base.series[i].ConvertCtoN(base.fromAx.CtoN)
		}
	case Proportional:
		for i := range base.series {
			base.series[i].ConvertPtoN(base.fromAx.PtoN)
		}
	}
}

func (base *BaseChart) PixelGenCartesian(pX, pY, w, h int) (col color.Color) {
	x, y := base.PositionToCartesianCoordinates(pX, pY, w, h)
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	for i := range base.series {
		serCol := base.series[i].RasterColorCartesian(x, y)
		r, g, b, _ := serCol.RGBA()
		if r > 0 || g > 0 || b > 0 {
			col = serCol
			break
		}
	}
	return
}

func (base *BaseChart) PixelGenPolar(pX, pY, w, h int) (col color.Color) {
	phi, r, x, y := base.PositionToPolarCoordinates(pX, pY, w, h)
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	_, rMax := base.toAx.NRange()
	if r > rMax {
		return
	}
	for i := range base.series {
		serCol := base.series[i].RasterColorPolar(phi, r, x, y)
		r, g, b, _ := serCol.RGBA()
		if r > 0 || g > 0 || b > 0 {
			col = serCol
			break
		}
	}
	return
}

func (base *BaseChart) PositionToCartesianCoordinates(pX int, pY int, w int, h int) (x float64, y float64) {
	trans := false
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	if trans {
		x = xMin + ((float64(h-pY) / float64(h)) * (xMax - xMin))
		y = yMin + ((float64(pX) / float64(w)) * (yMax - yMin))
	} else {
		x = xMin + ((float64(pX) / float64(w)) * (xMax - xMin))
		y = yMin + ((float64(h-pY) / float64(h)) * (yMax - yMin))
	}
	return
}

func (base *BaseChart) PositionToPolarCoordinates(pX int, pY int, w int, h int) (phi float64,
	r float64, x float64, y float64) {
	_, rMax := base.toAx.NRange()
	rot := 0.0
	mathPos := true
	posToCoord := rMax / (float64(w) / 2.0)
	x = (float64(pX) - (float64(w) / 2.0)) * posToCoord
	y = ((float64(h) / 2.0) - float64(pY)) * posToCoord
	r = math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	phi = math.Acos(x / r)
	if y < 0 {
		phi = -phi + (2 * math.Pi)
	}
	dirCor := 0.0
	if !mathPos {
		dirCor = 1.0
	}
	phi -= (rot + (dirCor * 2 * math.Pi))
	if phi < 0 {
		phi += 2 * math.Pi
	} else if phi > 2*math.Pi {
		phi -= 2 * math.Pi
	}
	return
}
