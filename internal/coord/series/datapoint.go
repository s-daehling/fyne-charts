package series

import (
	"errors"
	"image/color"
	"math"
	"sort"
	"time"

	"github.com/s-daehling/fyne-charts/internal/renderer"
	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type catOffset struct {
	c         string
	valOffset float64
}

type dataPoint struct {
	c         string
	t         time.Time
	n         float64
	val       float64
	valBase   float64
	nBarWidth float64
	tBarWidth time.Duration
	nBarShift float64
	tBarShift time.Duration
	// nBarStart           float64
	// nBarEnd             float64
	// valBarStart         float64
	// valBarEnd           float64
	dot                 *canvas.Circle
	fromValBase         *canvas.Line
	fromPrev            *canvas.Line
	bar                 *canvas.Rectangle
	showDot             bool
	showFromValBaseLine bool
	showFromPrevLine    bool
	showBar             bool
}

func emptyDataPoint(color color.Color, showDot bool, showFromBase bool, showFromPrev bool,
	showBar bool) (point *dataPoint) {
	point = &dataPoint{
		dot:                 canvas.NewCircle(color),
		fromValBase:         canvas.NewLine(color),
		fromPrev:            canvas.NewLine(color),
		bar:                 canvas.NewRectangle(color),
		showDot:             showDot,
		showFromValBaseLine: showFromBase,
		showFromPrevLine:    showFromPrev,
		showBar:             showBar,
		valBase:             0,
		nBarWidth:           0,
		tBarWidth:           0,
		nBarShift:           0,
		tBarShift:           0,
	}
	point.dot.Resize(fyne.NewSize(5, 5))
	return
}

func (point *dataPoint) refresh() {
	point.dot.Refresh()
	point.fromValBase.Refresh()
	point.fromPrev.Refresh()
	point.bar.Refresh()
}

func (point *dataPoint) hide() {
	point.dot.Hide()
	point.fromValBase.Hide()
	point.fromPrev.Hide()
	point.bar.Hide()
}

func (point *dataPoint) show() {
	point.dot.Show()
	point.fromValBase.Show()
	point.fromPrev.Show()
	point.bar.Show()
}

func (point *dataPoint) setColor(col color.Color) {
	point.dot.FillColor = col
	point.fromValBase.StrokeColor = col
	point.fromPrev.StrokeColor = col
	point.bar.FillColor = col
	point.refresh()
}

func (point *dataPoint) setLineWidth(lw float32) {
	point.fromValBase.StrokeWidth = lw
	point.fromPrev.StrokeWidth = lw
	point.refresh()
}

func (point *dataPoint) setDotSize(ds float32) {
	point.dot.Resize(fyne.NewSize(ds, ds))
	point.refresh()
}

func (point *dataPoint) setValBase(vb float64) {
	point.valBase = vb
}

func (point *dataPoint) setNBarWidthAndShift(bw float64, bs float64) {
	if bw < 0 {
		bw = 0
	}
	point.nBarWidth = bw
	point.nBarShift = bs
}

func (point *dataPoint) setTBarWidthAndShift(bw time.Duration, bs time.Duration) {
	point.tBarWidth = bw
	point.tBarShift = bs
}

// func (point *dataPoint) calculateBarRange() {
// 	point.nBarStart = point.n + point.nBarShift - (point.nBarWidth / 2)
// 	point.valBarStart = math.Min(point.valBase, point.valBase+point.val)
// 	point.nBarEnd = point.n + point.nBarShift + (point.nBarWidth / 2)
// 	point.valBarEnd = math.Max(point.valBase, point.valBase+point.val)
// }

func (point *dataPoint) cartesianNodes(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []renderer.CartesianNode) {
	if !point.showDot || point.n < xMin || point.n > xMax || point.val < yMin || point.val > yMax {
		return
	}
	ns = append(ns, renderer.CartesianNode{
		X:   point.n,
		Y:   point.val,
		Dot: point.dot,
	})
	return
}

func (point *dataPoint) polarNodes(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ns []renderer.PolarNode) {
	if !point.showDot || point.val > rMax || point.val < rMin || point.n < phiMin ||
		point.n > phiMax {
		return
	}
	ns = append(ns, renderer.PolarNode{
		Phi: point.n,
		R:   point.val,
		Dot: point.dot,
	})
	return
}

func (point *dataPoint) cartesianEdges(firstPoint bool, prevX float64, prevY float64, xMin float64,
	xMax float64, yMin float64, yMax float64) (es []renderer.CartesianEdge) {
	if point.showFromValBaseLine && !(point.n > xMax || point.n < xMin) {
		es = append(es, renderer.CartesianEdge{
			X1:   point.n,
			Y1:   math.Max(yMin, point.valBase),
			X2:   point.n,
			Y2:   math.Min(yMax, point.val),
			Line: point.fromValBase,
		})
	}
	if !point.showFromPrevLine || firstPoint || prevX > xMax || point.n < xMin {
		return
	}
	x1 := prevX
	y1 := prevY
	x2 := point.n
	y2 := point.val
	if x1 < xMin {
		// X1 outside range -> X1 to xmin
		y1 = y1 + ((xMin - x1) * ((y2 - y1) / (x2 - x1)))
		x1 = xMin
	}
	if x2 > xMax {
		// X2 outside range -> X2 toxmax
		y2 = y1 + ((xMax - x1) * ((y2 - y1) / (x2 - x1)))
		x2 = xMax
	}
	if (y1 < yMin && y2 < yMin) ||
		(y1 > yMax && y2 > yMax) {
		// both points out of range
		return
	}
	if y1 < yMin {
		// Y1 to ymin
		x1 = x1 + ((yMin - y1) * ((x2 - x1) / (y2 - y1)))
		y1 = yMin
	}
	if y1 > yMax {
		// Y1 to ymax
		x1 = x1 + ((yMax - y1) * ((x2 - x1) / (y2 - y1)))
		y1 = yMax
	}
	if y2 < yMin {
		// Y2 to ymin
		x2 = x1 + ((yMin - y1) * ((x2 - x1) / (y2 - y1)))
		y2 = yMin
	}
	if y2 > yMax {
		// Y2 to ymax
		x2 = x1 + ((yMax - y1) * ((x2 - x1) / (y2 - y1)))
		y2 = yMax
	}
	es = append(es, renderer.CartesianEdge{
		X1:   x1,
		Y1:   y1,
		X2:   x2,
		Y2:   y2,
		Line: point.fromPrev,
	})
	return
}

func (point *dataPoint) polarEdges(firstPoint bool, prevPhi float64, prevR float64, phiMin float64,
	phiMax float64, rMin float64, rMax float64) (es []renderer.PolarEdge) {
	if point.showFromValBaseLine && !(point.n < phiMin || point.n > phiMax || point.val < rMin) {
		es = append(es, renderer.PolarEdge{
			Phi1: point.n,
			R1:   math.Max(rMin, point.valBase),
			Phi2: point.n,
			R2:   math.Min(rMax, point.val),
			Line: point.fromValBase,
		})
	}
	if !point.showFromPrevLine || firstPoint || prevR > rMax || point.val > rMax || prevR < rMin ||
		point.val < rMin || prevPhi < phiMin || point.n < phiMin || prevPhi > phiMax ||
		point.n > phiMax {
		return
	}
	es = append(es, renderer.PolarEdge{
		Phi1: prevPhi,
		R1:   prevR,
		Phi2: point.n,
		R2:   point.val,
		Line: point.fromPrev,
	})
	return
}

func (point *dataPoint) cartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64, stacked bool) (rs []renderer.CartesianRect) {
	if !point.showBar || point.n < xMin || point.n > xMax {
		return
	}
	// if stacked {
	rs = append(rs, renderer.CartesianRect{
		X1:   point.n + point.nBarShift - (point.nBarWidth / 2),
		Y1:   math.Max(math.Min(point.valBase, point.valBase+point.val), yMin),
		X2:   point.n + point.nBarShift + (point.nBarWidth / 2),
		Y2:   math.Min(math.Max(point.valBase, point.valBase+point.val), yMax),
		Rect: point.bar,
	})
	// } else {
	// 	rs = append(rs, renderer.CartesianRect{
	// 		X1:   point.n + point.nBarShift - (point.nBarWidth / 2),
	// 		Y1:   math.Max(yMin, math.Min(point.valBase, point.val)),
	// 		X2:   point.n + point.nBarShift + (point.nBarWidth / 2),
	// 		Y2:   math.Min(yMax, math.Max(point.valBase, point.val)),
	// 		Rect: point.bar,
	// 	})
	// }
	return
}

func (point *dataPoint) RasterColorPolar(phi float64, r float64) (col color.Color, useColor bool) {
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	useColor = false
	if !point.showBar || phi < point.n+point.nBarShift-(point.nBarWidth/2) ||
		phi > point.n+point.nBarShift+(point.nBarWidth/2) ||
		r < point.valBase || r > point.val+point.valBase {
		return
	}
	col = point.bar.FillColor
	useColor = true
	return
}

type PointSeries struct {
	baseSeries
	data                []*dataPoint
	valBase             float64
	nBarWidth           float64
	tBarWidth           time.Duration
	nBarShift           float64
	tBarShift           time.Duration
	showDot             bool
	showFromValBaseLine bool
	showFromPrevLine    bool
	showBar             bool
	showArea            bool
	sortPoints          bool
	isStacked           bool
	valMin              float64
	valMax              float64
}

func EmptyPointSeries(name string, color color.Color) (ser *PointSeries) {
	ser = &PointSeries{
		valBase:             0,
		nBarWidth:           0,
		tBarWidth:           0,
		nBarShift:           0,
		tBarShift:           0,
		showDot:             false,
		showFromValBaseLine: false,
		showFromPrevLine:    false,
		showBar:             false,
		showArea:            false,
		isStacked:           false,
		sortPoints:          true,
	}
	ser.baseSeries = emptyBaseSeries(name, color, ser.toggleView)
	return
}

func (ser *PointSeries) MakeBar() {
	ser.showBar = true
	for i := range ser.data {
		ser.data[i].showBar = true
	}
}

func (ser *PointSeries) MakeArea(showDot bool) {
	ser.showDot = showDot
	ser.showFromPrevLine = true
	ser.showArea = true
	for i := range ser.data {
		ser.data[i].showDot = showDot
		ser.data[i].showFromPrevLine = true
	}
}

func (ser *PointSeries) MakeLine(showDot bool) {
	ser.showDot = showDot
	ser.showFromPrevLine = true
	for i := range ser.data {
		ser.data[i].showDot = showDot
		ser.data[i].showFromPrevLine = true
	}
}

func (ser *PointSeries) MakeLollipop() {
	ser.showDot = true
	ser.showFromValBaseLine = true
	for i := range ser.data {
		ser.data[i].showDot = true
		ser.data[i].showFromValBaseLine = true
	}
}

func (ser *PointSeries) MakeScatter() {
	ser.showDot = true
	for i := range ser.data {
		ser.data[i].showDot = true
	}
}

func (ser *PointSeries) CRange() (cs []string) {
	for i := range ser.data {
		cs = append(cs, ser.data[i].c)
	}
	return
}

func (ser *PointSeries) TRange() (isEmpty bool, min time.Time, max time.Time) {
	isEmpty = false
	if len(ser.data) == 0 {
		isEmpty = true
		return
	}
	min = ser.data[0].t
	max = ser.data[0].t
	if ser.showBar {
		min = min.Add(ser.data[0].tBarShift - (ser.data[0].tBarWidth / 2))
		max = max.Add(ser.data[0].tBarShift + (ser.data[0].tBarWidth / 2))
	}
	for i := range ser.data {
		pMin := ser.data[i].t
		pMax := ser.data[i].t
		if ser.showBar {
			pMin = pMin.Add(ser.data[i].tBarShift - (ser.data[i].tBarWidth / 2))
			pMax = pMax.Add(ser.data[i].tBarShift + (ser.data[i].tBarWidth / 2))
		}
		if pMin.Before(min) {
			min = pMin
		}
		if pMax.After(max) {
			max = pMax
		}
	}
	return
}

func (ser *PointSeries) NRange() (isEmpty bool, min float64, max float64) {
	min = 0
	max = 0
	isEmpty = false
	if len(ser.data) == 0 {
		isEmpty = true
		return
	}
	min = ser.data[0].n
	max = ser.data[0].n
	if ser.showBar {
		min += ser.data[0].nBarShift - (ser.data[0].nBarWidth / 2)
		max += ser.data[0].nBarShift + (ser.data[0].nBarWidth / 2)
	}
	for i := range ser.data {
		pMin := ser.data[i].n
		pMax := ser.data[i].n
		if ser.showBar {
			pMin += ser.data[i].nBarShift - (ser.data[i].nBarWidth / 2)
			pMax += ser.data[i].nBarShift + (ser.data[i].nBarWidth / 2)
		}
		if pMin < min {
			min = pMin
		}
		if pMax > max {
			max = pMax
		}
	}
	return
}

func (ser *PointSeries) ValRange() (isEmpty bool, min float64, max float64) {
	min = 0
	max = 0
	isEmpty = false
	if len(ser.data) == 0 {
		isEmpty = true
		return
	}
	min = ser.data[0].val
	max = ser.data[0].val
	if ser.showBar {
		min = math.Min(ser.data[0].valBase, ser.data[0].val+ser.data[0].valBase)
		max = math.Max(ser.data[0].valBase, ser.data[0].val+ser.data[0].valBase)
	}
	if ser.showFromValBaseLine {
		min = math.Min(ser.data[0].valBase, ser.data[0].val)
		max = math.Max(ser.data[0].valBase, ser.data[0].val)
	}
	for i := range ser.data {
		pMin := ser.data[i].val
		pMax := ser.data[i].val
		if ser.showBar {
			pMin = math.Min(ser.data[i].valBase, ser.data[i].val+ser.data[i].valBase)
			pMax = math.Max(ser.data[i].valBase, ser.data[i].val+ser.data[i].valBase)
		}
		if ser.showFromValBaseLine {
			pMin = math.Min(ser.data[i].valBase, ser.data[i].val)
			pMax = math.Max(ser.data[i].valBase, ser.data[i].val)
		}
		if pMin < min {
			min = pMin
		}
		if pMax > max {
			max = pMax
		}
	}
	ser.valMin = min
	ser.valMax = max
	return
}

func (ser *PointSeries) ConvertCtoN(cToN func(c string) (n float64)) {
	for i := range ser.data {
		ser.data[i].n = cToN(ser.data[i].c)
	}
}

func (ser *PointSeries) ConvertTtoN(tToN func(t time.Time) (n float64)) {
	for i := range ser.data {
		ser.data[i].n = tToN(ser.data[i].t)
		if ser.showBar {
			ser.data[i].nBarWidth = tToN(ser.data[i].t.Add(ser.data[i].tBarWidth)) - ser.data[i].n
			ser.data[i].nBarShift = tToN(ser.data[i].t.Add(ser.data[i].tBarShift)) - ser.data[i].n
		}
	}
}

func (ser *PointSeries) CartesianNodes(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []renderer.CartesianNode) {
	for i := range ser.data {
		ns = append(ns, ser.data[i].cartesianNodes(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *PointSeries) CartesianEdges(xMin float64, xMax float64, yMin float64,
	yMax float64) (es []renderer.CartesianEdge) {
	for i := range ser.data {
		if i == 0 {
			es = append(es, ser.data[i].cartesianEdges(true, 0, 0, xMin, xMax, yMin, yMax)...)
		} else {
			es = append(es, ser.data[i].cartesianEdges(false, ser.data[i-1].n, ser.data[i-1].val,
				xMin, xMax, yMin, yMax)...)
		}
	}
	return
}

func (ser *PointSeries) CartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []renderer.CartesianRect) {
	for i := range ser.data {
		fs = append(fs, ser.data[i].cartesianRects(xMin, xMax, yMin, yMax, ser.isStacked)...)
	}
	return
}

func (ser *PointSeries) RasterColorCartesian(x float64, y float64) (col color.Color) {
	col = ser.baseSeries.RasterColorCartesian(x, y)
	if !ser.visible || !ser.showArea || y < ser.valMin || y > ser.valMax {
		return
	}
	// find first data point with x higher
	for i := range ser.data {
		if ser.data[i].n > x {
			if i == 0 {
				break
			}
			x1 := ser.data[i-1].n
			x2 := ser.data[i].n
			y1 := ser.data[i-1].val
			y2 := ser.data[i].val
			// interpolate
			yS := y1 + (((x - x1) / (x2 - x1)) * (y2 - y1))
			if yS > ser.valBase && y > ser.valBase && y < yS {
				r, g, b, _ := ser.color.RGBA()
				col = color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: 0x8888}

			} else if yS < ser.valBase && y < ser.valBase && y > yS {
				r, g, b, _ := ser.color.RGBA()
				col = color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: 0x8888}
			}
			break
		}
	}
	return
}

func (ser *PointSeries) PolarNodes(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ns []renderer.PolarNode) {
	for i := range ser.data {
		ns = append(ns, ser.data[i].polarNodes(phiMin, phiMax, rMin, rMax)...)
	}
	return
}

func (ser *PointSeries) PolarEdges(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (es []renderer.PolarEdge) {
	for i := range ser.data {
		if i == 0 {
			es = append(es, ser.data[i].polarEdges(true, 0, 0, phiMin, phiMax, rMin, rMax)...)
		} else {
			es = append(es, ser.data[i].polarEdges(false, ser.data[i-1].n, ser.data[i-1].val,
				phiMin, phiMax, rMin, rMax)...)
		}
	}
	return
}

func (ser *PointSeries) RasterColorPolar(phi float64, r float64, x float64,
	y float64) (col color.Color) {
	col = ser.baseSeries.RasterColorPolar(phi, r, x, y)
	if !ser.visible || (!ser.showBar && !ser.showArea) || r > ser.valMax {
		return
	}
	if ser.showBar {
		for i := range ser.data {
			pCol, useColor := ser.data[i].RasterColorPolar(phi, r)
			if useColor {
				col = pCol
				break
			}
		}
	} else if ser.showArea {
		red, green, blue, _ := ser.color.RGBA()
		colArea := color.RGBA64{R: uint16(red), G: uint16(green), B: uint16(blue), A: 0x8888}
		// find first data point with x higher
		for i := range ser.data {
			if ser.data[i].n > phi {
				if i == 0 {
					break
				}
				phi1 := ser.data[i-1].n
				phi2 := ser.data[i].n
				r1 := ser.data[i-1].val
				r2 := ser.data[i].val
				R := r1 + (((phi - phi1) / (phi2 - phi1)) * (r2 - r1))
				if r < R {
					col = colArea
				}
				break
			}
		}
	}
	return
}

func (ser *PointSeries) IsPartOfChartRaster() (b bool) {
	b = false
	if ser.cont == nil || !ser.visible {
		return
	}
	if (ser.cont.IsPolar() && !ser.showBar && !ser.showArea) ||
		(!ser.cont.IsPolar() && !ser.showArea) {
		return
	}
	b = true
	return
}

func (ser *PointSeries) SetAndUpdateValBaseCategorical(in []catOffset) (out []catOffset) {
	for i := range ser.data {
		catInOffsetList := false
		for j := range in {
			if in[j].c == ser.data[i].c {
				ser.data[i].setValBase(in[j].valOffset)
				catInOffsetList = true
				break
			}
		}
		if !catInOffsetList {
			ser.data[i].setValBase(0)
		}
	}
	copy(out, in)
	if !ser.visible {
		return
	}
	for i := range ser.data {
		catExist := false
		for j := range out {
			if ser.data[i].c == out[j].c {
				catExist = true
				out[j].valOffset += ser.data[i].val
				break
			}
		}
		if !catExist {
			out = append(out, catOffset{c: ser.data[i].c, valOffset: ser.data[i].val})
		}
	}
	return
}

func (ser *PointSeries) SetValBaseNumerical(vb float64) {
	ser.valBase = vb
	for i := range ser.data {
		ser.data[i].setValBase(vb)
	}
}

// Show makes all elements of the series visible
func (ser *PointSeries) Show() {
	ser.visible = true
	for i := range ser.data {
		ser.data[i].show()
	}
	ser.legendEntry.Show()
	if ser.showBar && ser.cont != nil {
		ser.cont.DataChange()
	}
}

// Hide hides all elements of the series
func (ser *PointSeries) Hide() {
	ser.visible = false
	for i := range ser.data {
		ser.data[i].hide()
	}
	ser.legendEntry.Hide()
	if ser.showBar && ser.cont != nil {
		ser.cont.DataChange()
	}
}

func (ser *PointSeries) toggleView() {
	if ser.visible {
		ser.Hide()
	} else {
		ser.Show()
	}
	if ser.showArea && ser.cont != nil {
		ser.cont.RasterRefresh()
	}
}

func (ser *PointSeries) SetColor(col color.Color) {
	ser.color = col
	ser.legendEntry.SetColor(col)
	for i := range ser.data {
		ser.data[i].setColor(col)
	}
}

func (ser *PointSeries) SetLineWidth(lw float32) {
	if lw < 0 {
		return
	}
	for i := range ser.data {
		ser.data[i].setLineWidth(lw)
	}
}

func (ser *PointSeries) SetDotSize(ds float32) {
	if ds < 0 {
		return
	}
	for i := range ser.data {
		ser.data[i].setDotSize(ds)
	}
}

func (ser *PointSeries) SetNumericalBarWidthAndShift(width float64, shift float64) (err error) {
	if width < 0 {
		err = errors.New("invalid width")
		return
	}
	ser.nBarWidth = width
	ser.nBarShift = shift
	for i := range ser.data {
		ser.data[i].setNBarWidthAndShift(width, shift)
	}
	return
}

func (ser *PointSeries) SetNumericalBarWidth(width float64) (err error) {
	if width < 0 {
		err = errors.New("invalid width")
		return
	}
	ser.nBarWidth = width
	for i := range ser.data {
		ser.data[i].setNBarWidthAndShift(width, ser.nBarShift)
	}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}

func (ser *PointSeries) SetTemporalBarWidth(width time.Duration) (err error) {
	if width < 0 {
		err = errors.New("invalid width")
		return
	}
	ser.tBarWidth = width
	for i := range ser.data {
		ser.data[i].setTBarWidthAndShift(width, ser.tBarShift)
	}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}

func (ser *PointSeries) IsAreaSeries() (b bool) {
	b = ser.showArea
	return
}

func (ser *PointSeries) IsBarSeries() (b bool) {
	b = ser.showBar
	return
}

func (ser *PointSeries) IsLollipopSeries() (b bool) {
	b = ser.showFromValBaseLine && ser.showDot
	return
}

func (ser *PointSeries) BindToChart(ch container) (err error) {
	if ch.IsPolar() {
		for i := range ser.data {
			if ser.data[i].val < 0 {
				err = errors.New("invalid data, negative val not allowed for polar charts")
				return
			}
		}
	}
	err = ser.baseSeries.BindToChart(ch)
	return
}

func (ser *PointSeries) BindToStack(stack *StackedSeries) (err error) {
	for i := range ser.data {
		if ser.data[i].val < 0 {
			err = errors.New("invalid data, negative val not allowed for stacked series")
			return
		}
	}
	ser.legendEntry.SetSuper(stack.name)
	ser.super = stack.name
	err = ser.baseSeries.BindToChart(stack)
	if err != nil {
		return
	}
	ser.isStacked = true
	return
}

func (ser *PointSeries) Release() {
	ser.baseSeries.Release()
	ser.showDot = false
	ser.showFromValBaseLine = false
	ser.showFromPrevLine = false
	ser.showBar = false
	ser.showArea = false
	ser.isStacked = false
	ser.legendEntry.SetSuper("")
	ser.super = ""
	for i := range ser.data {
		ser.data[i].showDot = false
		ser.data[i].showFromPrevLine = false
		ser.data[i].showFromValBaseLine = false
		ser.data[i].showBar = false
	}
}

func (ser *PointSeries) Clear() {
	ser.data = []*dataPoint{}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
}

func (ser *PointSeries) DeleteNumericalDataInRange(min float64, max float64) (c int) {
	c = 0
	if min > max {
		return
	}
	finalData := []*dataPoint{}
	for i := range ser.data {
		if ser.data[i].n > min && ser.data[i].n < max {
			c++
		} else {
			finalData = append(finalData, ser.data[i])
		}
	}
	if c == 0 {
		return
	}
	ser.data = nil
	ser.data = finalData
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}

func (ser *PointSeries) AddNumericalData(input []data.NumericalPoint) (err error) {
	if len(input) == 0 {
		return
	}
	if ser.cont != nil {
		if ser.cont.IsPolar() || ser.isStacked {
			for i := range input {
				if input[i].Val < 0 {
					err = errors.New("negative val not allowed")
					return
				}
			}
		}
	}
	var newData []data.NumericalPoint
	if ser.sortPoints {
		newData = append(make([]data.NumericalPoint, 0, len(input)), input...)
		for i := range ser.data {
			newData = append(newData, data.NumericalPoint{N: ser.data[i].n, Val: ser.data[i].val})
		}
		sort.Sort(data.DpByNValue(newData))
		ser.data = []*dataPoint{}
	} else {
		newData = input
	}
	for i := range newData {
		dPoint := emptyDataPoint(ser.color, ser.showDot, ser.showFromValBaseLine,
			ser.showFromPrevLine, ser.showBar)
		dPoint.n = newData[i].N
		dPoint.val = newData[i].Val
		if ser.showBar {
			dPoint.setNBarWidthAndShift(ser.nBarWidth, ser.nBarWidth)
		}
		if ser.showFromValBaseLine {
			dPoint.setValBase(ser.valBase)
		}
		ser.data = append(ser.data, dPoint)
	}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}

func (ser *PointSeries) DeleteTemporalDataInRange(min time.Time, max time.Time) (c int) {
	c = 0
	if min.After(max) {
		return
	}
	finalData := []*dataPoint{}
	for i := range ser.data {
		if ser.data[i].t.After(min) && ser.data[i].t.Before(max) {
			c++
		} else {
			finalData = append(finalData, ser.data[i])
		}
	}
	if c == 0 {
		return
	}
	ser.data = nil
	ser.data = finalData
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}

func (ser *PointSeries) AddTemporalData(input []data.TemporalPoint) (err error) {
	if len(input) == 0 {
		return
	}
	if ser.cont != nil {
		if ser.cont.IsPolar() || ser.isStacked {
			for i := range input {
				if input[i].Val < 0 {
					err = errors.New("negative val not allowed")
					return
				}
			}
		}
	}
	var newData []data.TemporalPoint
	if ser.sortPoints {
		newData = append(make([]data.TemporalPoint, 0, len(input)), input...)
		for i := range ser.data {
			newData = append(newData, data.TemporalPoint{T: ser.data[i].t, Val: ser.data[i].val})
		}
		sort.Sort(data.DpByTValue(newData))
		ser.data = []*dataPoint{}
	} else {
		newData = input
	}
	for i := range newData {
		dPoint := emptyDataPoint(ser.color, ser.showDot, ser.showFromValBaseLine,
			ser.showFromPrevLine, ser.showBar)
		dPoint.t = newData[i].T
		dPoint.val = newData[i].Val
		if ser.showBar {
			dPoint.setTBarWidthAndShift(ser.tBarWidth, ser.tBarWidth)
		}
		if ser.showFromValBaseLine {
			dPoint.setValBase(ser.valBase)
		}
		ser.data = append(ser.data, dPoint)
	}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}

func (ser *PointSeries) DeleteCategoricalDataInRange(cat []string) (c int) {
	c = 0
	if len(cat) == 0 {
		return
	}
	finalData := []*dataPoint{}
	for i := range ser.data {
		del := false
		for j := range cat {
			if ser.data[i].c == cat[j] {
				del = true
				break
			}
		}
		if del {
			c++
		} else {
			finalData = append(finalData, ser.data[i])
		}
	}
	if c == 0 {
		return
	}
	ser.data = nil
	ser.data = finalData
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}

func (ser *PointSeries) AddCategoricalData(input []data.CategoricalPoint) (err error) {
	if len(input) == 0 {
		return
	}
	if ser.cont != nil {
		if ser.cont.IsPolar() || ser.isStacked {
			for i := range input {
				if input[i].Val < 0 {
					err = errors.New("negative val not allowed")
					return
				}
			}
		}
	}
	for i := range input {
		catExist := false
		for j := range ser.data {
			if input[i].C == ser.data[j].c {
				catExist = true
				break
			}
		}
		if catExist {
			continue
		}
		dPoint := emptyDataPoint(ser.color, ser.showDot, ser.showFromValBaseLine,
			ser.showFromPrevLine, ser.showBar)
		dPoint.c = input[i].C
		dPoint.val = input[i].Val
		if ser.showFromValBaseLine {
			dPoint.setValBase(ser.valBase)
		}
		ser.data = append(ser.data, dPoint)
	}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}
