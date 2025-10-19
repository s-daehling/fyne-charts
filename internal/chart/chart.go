package chart

import (
	"image/color"
	"math"

	"github.com/s-daehling/fyne-charts/internal/axis"
	"github.com/s-daehling/fyne-charts/internal/renderer"
	"github.com/s-daehling/fyne-charts/internal/series"

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
		base.render = renderer.EmptyCartesianRenderer(base)
	} else {
		base.fromAx = axis.EmptyAxis("", axis.PolarPhiAxis)
		base.toAx = axis.EmptyAxis("", axis.PolarRAxis)
		base.rast = canvas.NewRasterWithPixels(base.PixelGenPolar)
		base.render = renderer.EmptyPolarRenderer(base)
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

func (base *BaseChart) Title() (ct *canvas.Text) {
	ct = base.label
	return
}

func (base *BaseChart) CartesianObjects() (canObj []fyne.CanvasObject) {
	// objects will be drawn in the same order as added here

	// first get all objects from the series
	if base.legendVisible {
		lEntries := base.LegendEntries()
		for i := range lEntries {
			canObj = append(canObj, lEntries[i].Button, lEntries[i].Label)
		}
	}
	canObj = append(canObj, base.rast)
	rects := base.CartesianRects()
	for i := range rects {
		canObj = append(canObj, rects[i].Rect)
	}
	edges := base.CartesianEdges()
	for i := range edges {
		canObj = append(canObj, edges[i].Line)
	}
	nodes := base.CartesianNodes()
	for i := range nodes {
		canObj = append(canObj, nodes[i].Dot)
	}
	texts := base.CartesianTexts()
	for i := range texts {
		canObj = append(canObj, texts[i].Text)
	}

	// add axis elements
	canObj = append(canObj, base.fromAx.Objects()...)
	canObj = append(canObj, base.toAx.Objects()...)

	// add chart title and axis titles
	if base.label.Text != "" {
		canObj = append(canObj, base.label)
	}
	return
}

func (base *BaseChart) CartesianNodes() (ns []renderer.CartesianNode) {
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	for i := range base.series {
		ns = append(ns, base.series[i].CartesianNodes(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (base *BaseChart) CartesianEdges() (es []renderer.CartesianEdge) {
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	for i := range base.series {
		es = append(es, base.series[i].CartesianEdges(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (base *BaseChart) CartesianRects() (as []renderer.CartesianRect) {
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	for i := range base.series {
		as = append(as, base.series[i].CartesianRects(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (base *BaseChart) CartesianTexts() (ts []renderer.CartesianText) {
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	for i := range base.series {
		ts = append(ts, base.series[i].CartesianTexts(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (base *BaseChart) PolarObjects() (canObj []fyne.CanvasObject) {
	// objects will be drawn in the same order as added here

	// first get all objects from the series
	if base.legendVisible {
		lEntries := base.LegendEntries()
		for i := range lEntries {
			canObj = append(canObj, lEntries[i].Button, lEntries[i].Label)
		}
	}
	canObj = append(canObj, base.rast)
	edges := base.PolarEdges()
	for i := range edges {
		canObj = append(canObj, edges[i].Line)
	}
	nodes := base.PolarNodes()
	for i := range nodes {
		canObj = append(canObj, nodes[i].Dot)
	}
	texts := base.PolarTexts()
	for i := range texts {
		canObj = append(canObj, texts[i].Text)
	}

	// add axis elements
	canObj = append(canObj, base.fromAx.Objects()...)
	canObj = append(canObj, base.toAx.Objects()...)

	// add chart title and axis titles
	if base.label.Text != "" {
		canObj = append(canObj, base.label)
	}
	return
}

func (base *BaseChart) PolarNodes() (ns []renderer.PolarNode) {
	phiMin, phiMax := base.fromAx.NRange()
	rMin, rMax := base.toAx.NRange()
	for i := range base.series {
		ns = append(ns, base.series[i].PolarNodes(phiMin, phiMax, rMin, rMax)...)
	}
	return
}

func (base *BaseChart) PolarEdges() (es []renderer.PolarEdge) {
	phiMin, phiMax := base.fromAx.NRange()
	rMin, rMax := base.toAx.NRange()
	for i := range base.series {
		es = append(es, base.series[i].PolarEdges(phiMin, phiMax, rMin, rMax)...)
	}
	return
}

func (base *BaseChart) PolarTexts() (ts []renderer.PolarText) {
	phiMin, phiMax := base.fromAx.NRange()
	rMin, rMax := base.toAx.NRange()
	for i := range base.series {
		ts = append(ts, base.series[i].PolarTexts(phiMin, phiMax, rMin, rMax)...)
	}
	return
}

func (base *BaseChart) Raster() (rs *canvas.Raster) {
	rs = base.rast
	return
}

func (base *BaseChart) LegendEntries() (les []renderer.LegendEntry) {
	for i := range base.series {
		les = append(les, base.series[i].LegendEntries()...)
	}
	return
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

func (base *BaseChart) SetTitle(l string) {
	base.label.Text = l
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
