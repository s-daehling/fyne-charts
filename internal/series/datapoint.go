package series

import (
	"errors"
	"image/color"
	"math"
	"sort"
	"time"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type catOffset struct {
	c         string
	valOffset float64
}

type dataPoint struct {
	c                   string
	t                   time.Time
	n                   float64
	val                 float64
	valBase             float64
	nBarWidth           float64
	tBarWidth           time.Duration
	nBarShift           float64
	tBarShift           time.Duration
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
}

func (point *dataPoint) setLineWidth(lw float32) {
	point.fromValBase.StrokeWidth = lw
	point.fromPrev.StrokeWidth = lw
}

func (point *dataPoint) setDotSize(ds float32) {
	point.dot.Resize(fyne.NewSize(ds, ds))
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

func (point *dataPoint) cartesianNodes(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []CartesianNode) {
	if !point.showDot || point.n < xMin || point.n > xMax || point.val < yMin || point.val > yMax {
		return
	}
	ns = append(ns, CartesianNode{
		X:   point.n,
		Y:   point.val,
		Dot: point.dot,
	})
	return
}

func (point *dataPoint) polarNodes(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ns []PolarNode) {
	if !point.showDot || point.val > rMax || point.val < rMin || point.n < phiMin ||
		point.n > phiMax {
		return
	}
	ns = append(ns, PolarNode{
		Phi: point.n,
		R:   point.val,
		Dot: point.dot,
	})
	return
}

func (point *dataPoint) cartesianEdges(firstPoint bool, prevX float64, prevY float64, xMin float64,
	xMax float64, yMin float64, yMax float64) (es []CartesianEdge) {
	if point.showFromValBaseLine && !(point.n > xMax || point.n < xMin) {
		es = append(es, CartesianEdge{
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
	es = append(es, CartesianEdge{
		X1:   x1,
		Y1:   y1,
		X2:   x2,
		Y2:   y2,
		Line: point.fromPrev,
	})
	return
}

func (point *dataPoint) polarEdges(firstPoint bool, prevPhi float64, prevR float64, phiMin float64,
	phiMax float64, rMin float64, rMax float64) (es []PolarEdge) {
	if point.showFromValBaseLine && !(point.n < phiMin || point.n > phiMax || point.val < rMin) {
		es = append(es, PolarEdge{
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
	es = append(es, PolarEdge{
		Phi1: prevPhi,
		R1:   prevR,
		Phi2: point.n,
		R2:   point.val,
		Line: point.fromPrev,
	})
	return
}

func (point *dataPoint) cartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (rs []CartesianRect) {
	if !point.showBar || point.n < xMin || point.n > xMax {
		return
	}
	rs = append(rs, CartesianRect{
		X1:   point.n + point.nBarShift - (point.nBarWidth / 2),
		Y1:   math.Max(math.Min(point.valBase, point.valBase+point.val), yMin),
		X2:   point.n + point.nBarShift + (point.nBarWidth / 2),
		Y2:   math.Min(math.Max(point.valBase, point.valBase+point.val), yMax),
		Rect: point.bar,
	})
	return
}

func (point *dataPoint) RasterColorPolar(phi float64, r float64, x float64,
	y float64) (col color.Color) {
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	if !point.showBar || phi < point.n+point.nBarShift-(point.nBarWidth/2) ||
		phi > point.n+point.nBarShift+(point.nBarWidth/2) ||
		r < point.valBase || r > point.val+point.valBase {
		return
	}
	col = point.bar.FillColor
	return
}

type dataPointSeries struct {
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
	sortPoints          bool
}

func EmptyDataPointSeries(chart chart, name string, color color.Color,
	polar bool) (ser dataPointSeries) {
	ser = dataPointSeries{
		valBase:             0,
		nBarWidth:           0,
		tBarWidth:           0,
		nBarShift:           0,
		tBarShift:           0,
		showDot:             false,
		showFromValBaseLine: false,
		showFromPrevLine:    false,
		showBar:             false,
		sortPoints:          false,
	}
	ser.baseSeries = emptyBaseSeries(chart, name, color, polar, ser.toggleView)
	return
}

func (ser *dataPointSeries) CRange() (cs []string) {
	for i := range ser.data {
		cs = append(cs, ser.data[i].c)
	}
	return
}

func (ser *dataPointSeries) TRange() (isEmpty bool, min time.Time, max time.Time) {
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

func (ser *dataPointSeries) NRange() (isEmpty bool, min float64, max float64) {
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

func (ser *dataPointSeries) ValRange() (isEmpty bool, min float64, max float64) {
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
	return
}

func (ser *dataPointSeries) ConvertCtoN(cToN func(c string) (n float64)) {
	for i := range ser.data {
		ser.data[i].n = cToN(ser.data[i].c)
	}
}

func (ser *dataPointSeries) ConvertTtoN(tToN func(t time.Time) (n float64)) {
	for i := range ser.data {
		ser.data[i].n = tToN(ser.data[i].t)
		if ser.showBar {
			ser.data[i].nBarWidth = tToN(ser.data[i].t.Add(ser.data[i].tBarWidth)) - ser.data[i].n
			ser.data[i].nBarShift = tToN(ser.data[i].t.Add(ser.data[i].tBarShift)) - ser.data[i].n
		}
	}
}

func (ser *dataPointSeries) CartesianNodes(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []CartesianNode) {
	for i := range ser.data {
		ns = append(ns, ser.data[i].cartesianNodes(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *dataPointSeries) CartesianEdges(xMin float64, xMax float64, yMin float64,
	yMax float64) (es []CartesianEdge) {
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

func (ser *dataPointSeries) CartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []CartesianRect) {
	for i := range ser.data {
		fs = append(fs, ser.data[i].cartesianRects(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *dataPointSeries) PolarNodes(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ns []PolarNode) {
	for i := range ser.data {
		ns = append(ns, ser.data[i].polarNodes(phiMin, phiMax, rMin, rMax)...)
	}
	return
}

func (ser *dataPointSeries) PolarEdges(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (es []PolarEdge) {
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

func (ser *dataPointSeries) RasterColorPolar(phi float64, r float64, x float64,
	y float64) (col color.Color) {
	col = ser.baseSeries.RasterColorPolar(phi, r, x, y)
	if !ser.visible || !ser.showBar {
		return
	}
	for i := range ser.data {
		pCol := ser.data[i].RasterColorPolar(phi, r, x, y)
		r, g, b, _ := pCol.RGBA()
		if r > 0 || g > 0 || b > 0 {
			col = pCol
			break
		}
	}
	return
}

func (ser *dataPointSeries) SetAndUpdateValBaseCategorical(in []catOffset) (out []catOffset) {
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

func (ser *dataPointSeries) SetValBaseNumerical(vb float64) {
	ser.valBase = vb
	for i := range ser.data {
		ser.data[i].setValBase(vb)
	}
}

// Show makes all elements of the series visible
func (ser *dataPointSeries) Show() {
	ser.visible = true
	for i := range ser.data {
		ser.data[i].show()
	}
	if ser.showBar {
		ser.chart.DataChange()
	}
}

// Hide hides all elements of the series
func (ser *dataPointSeries) Hide() {
	ser.visible = false
	for i := range ser.data {
		ser.data[i].hide()
	}
	if ser.showBar {
		ser.chart.DataChange()
	}
}

func (ser *dataPointSeries) toggleView() {
	if ser.visible {
		ser.Hide()
	} else {
		ser.Show()
	}
}

func (ser *dataPointSeries) SetColor(col color.Color) {
	ser.color = col
	ser.legendButton.SetColor(col)
	for i := range ser.data {
		ser.data[i].setColor(col)
	}
}

func (ser *dataPointSeries) SetLineWidth(lw float32) {
	if lw < 0 {
		return
	}
	for i := range ser.data {
		ser.data[i].setLineWidth(lw)
	}
}

func (ser *dataPointSeries) SetDotSize(ds float32) {
	if ds < 0 {
		return
	}
	for i := range ser.data {
		ser.data[i].setDotSize(ds)
	}
}

func (ser *dataPointSeries) SetNumericalBarWidthAndShift(width float64, shift float64) (err error) {
	if width < 0 || (ser.polar && width > 2*math.Pi) {
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

func (ser *dataPointSeries) SetTemporalBarWidthAndShift(width time.Duration,
	shift time.Duration) (err error) {
	if width < 0 {
		err = errors.New("invalid width")
		return
	}
	ser.tBarWidth = width
	ser.tBarShift = shift
	for i := range ser.data {
		ser.data[i].setTBarWidthAndShift(width, shift)
	}
	return
}

func (ser *dataPointSeries) Clear() (err error) {
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
	chart := ser.chart
	ser.data = []*dataPoint{}
	chart.DataChange()
	return
}

func (ser *dataPointSeries) DeleteNumericalDataInRange(min float64, max float64) (c int, err error) {
	c = 0
	if min > max {
		err = errors.New("invalid range")
		return
	}
	finalData := []*dataPoint{}
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
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
	ser.chart.DataChange()
	return
}

func (ser *dataPointSeries) AddNumericalData(input []data.NumericalDataPoint) (err error) {
	err = numericalDataPointRangeCheck(input, ser.polar, ser.polar)
	if err != nil {
		return
	}
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
	var newData []data.NumericalDataPoint
	if ser.sortPoints {
		newData = append(make([]data.NumericalDataPoint, 0, len(input)), input...)
		for i := range ser.data {
			newData = append(newData, data.NumericalDataPoint{N: ser.data[i].n, Val: ser.data[i].val})
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
	ser.chart.DataChange()
	return
}

func (ser *dataPointSeries) DeleteTemporalDataInRange(min time.Time, max time.Time) (c int, err error) {
	c = 0
	if min.After(max) {
		err = errors.New("invalid range")
		return
	}
	finalData := []*dataPoint{}
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
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
	ser.chart.DataChange()
	return
}

func (ser *dataPointSeries) AddTemporalData(input []data.TemporalDataPoint) (err error) {
	err = temporalDataPointRangeCheck(input, ser.polar)
	if err != nil {
		return
	}
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
	var newData []data.TemporalDataPoint
	if ser.sortPoints {
		newData = append(make([]data.TemporalDataPoint, 0, len(input)), input...)
		for i := range ser.data {
			newData = append(newData, data.TemporalDataPoint{T: ser.data[i].t, Val: ser.data[i].val})
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
	ser.chart.DataChange()
	return
}

func (ser *dataPointSeries) DeleteCategoricalDataInRange(cat []string) (c int, err error) {
	c = 0
	if len(cat) == 0 {
		err = errors.New("invalid range")
		return
	}
	finalData := []*dataPoint{}
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
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
	ser.chart.DataChange()
	return
}

func (ser *dataPointSeries) AddCategoricalData(input []data.CategoricalDataPoint) (err error) {
	err = categoricalDataPointRangeCheck(input, ser.polar)
	if err != nil {
		return
	}
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
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
	ser.chart.DataChange()
	return
}

func numericalDataPointRangeCheck(input []data.NumericalDataPoint, noNegativeVal bool, isPolar bool) (err error) {
	if len(input) == 0 {
		err = errors.New("no input data")
		return
	}
	if isPolar || noNegativeVal {
		for i := range input {
			if (isPolar && (input[i].N < 0 || input[i].N > 2*math.Pi)) ||
				(noNegativeVal && input[i].Val < 0) {
				err = errors.New("invalid data")
				return
			}
		}
	}
	return
}

func temporalDataPointRangeCheck(input []data.TemporalDataPoint, noNegativeVal bool) (err error) {
	if len(input) == 0 {
		err = errors.New("no input data")
		return
	}
	if noNegativeVal {
		for i := range input {
			if input[i].Val < 0 {
				err = errors.New("invalid data")
				return
			}
		}
	}
	return
}

func categoricalDataPointRangeCheck(input []data.CategoricalDataPoint, noNegativeVal bool) (err error) {
	if len(input) == 0 {
		err = errors.New("no input data")
		return
	}
	if noNegativeVal {
		for i := range input {
			if input[i].Val < 0 {
				err = errors.New("invalid data")
				return
			}
		}
	}
	return
}
