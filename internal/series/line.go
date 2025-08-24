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

type linePoint struct {
	t        time.Time
	n        float64
	val      float64
	dot      *canvas.Circle
	fromPrev *canvas.Line
}

func emptyLinePoint(showDots bool, color color.Color) (point *linePoint) {
	point = &linePoint{
		fromPrev: canvas.NewLine(color),
	}
	point.fromPrev.StrokeWidth = 1
	if showDots {
		point.dot = canvas.NewCircle(color)
		point.dot.Resize(fyne.NewSize(5, 5))
	} else {
		point.dot = nil
	}
	return
}

func (point *linePoint) hide() {
	point.dot.Hide()
	point.fromPrev.Hide()
}

func (point *linePoint) show() {
	point.dot.Show()
	point.fromPrev.Show()
}

func (point *linePoint) setColor(col color.Color) {
	if point.dot != nil {
		point.dot.FillColor = col
	}
	point.fromPrev.StrokeColor = col
}

func (point *linePoint) setLineWidth(lw float32) {
	point.fromPrev.StrokeWidth = lw
}

func (point *linePoint) setDotSize(ds float32) {
	if point.dot != nil {
		point.dot.Resize(fyne.NewSize(ds, ds))
	}
}

func (point *linePoint) cartesianNodes(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []CartesianNode) {
	if point.dot == nil {
		return
	}
	if !(point.n < xMin || point.n > xMax || point.val < yMin || point.val > yMax) {
		n := CartesianNode{
			X:   point.n,
			Y:   point.val,
			Dot: point.dot,
		}
		ns = append(ns, n)
	}
	return
}

func (point *linePoint) polarNodes(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ns []PolarNode) {
	if point.dot == nil {
		return
	}
	if !(point.val > rMax || point.val < rMin || point.n < phiMin || point.n > phiMax) {
		n := PolarNode{
			Phi: point.n,
			R:   point.val,
			Dot: point.dot,
		}
		ns = append(ns, n)
	}
	return
}

func (point *linePoint) cartesianEdges(prevX float64, prevY float64, xMin float64,
	xMax float64, yMin float64, yMax float64) (es []CartesianEdge) {

	e := CartesianEdge{
		X1:   prevX,
		Y1:   prevY,
		X2:   point.n,
		Y2:   point.val,
		Line: point.fromPrev,
	}
	if e.X1 > xMax || e.X2 < xMin {
		// both points out of range
		return
	}
	if e.X1 < xMin {
		// X1 outside range -> X1 to xmin
		e.Y1 = e.Y1 + ((xMin - e.X1) * ((e.Y2 - e.Y1) / (e.X2 - e.X1)))
		e.X1 = xMin
	}
	if e.X2 > xMax {
		// X2 outside range -> X2 toxmax
		e.Y2 = e.Y1 + ((xMax - e.X1) * ((e.Y2 - e.Y1) / (e.X2 - e.X1)))
		e.X2 = xMax
	}
	if (e.Y1 < yMin && e.Y2 < yMin) ||
		(e.Y1 > yMax && e.Y2 < yMax) {
		// both points out of range
		return
	}
	if e.Y1 < yMin {
		// Y1 to ymin
		e.X1 = e.X1 + ((yMin - e.Y1) * ((e.X2 - e.X1) / (e.Y2 - e.Y1)))
		e.Y1 = yMin
	}
	if e.Y1 > yMax {
		// Y1 to ymax
		e.X1 = e.X1 + ((yMax - e.Y1) * ((e.X2 - e.X1) / (e.Y2 - e.Y1)))
		e.Y1 = yMax
	}
	if e.Y2 < yMin {
		// Y2 to ymin
		e.X2 = e.X1 + ((yMin - e.Y1) * ((e.X2 - e.X1) / (e.Y2 - e.Y1)))
		e.Y2 = yMin
	}
	if e.Y2 > yMax {
		// Y2 to ymax
		e.X2 = e.X1 + ((yMax - e.Y1) * ((e.X2 - e.X1) / (e.Y2 - e.Y1)))
		e.Y2 = yMax
	}
	es = append(es, e)
	return
}

func (point *linePoint) polarEdges(prevPhi float64, prevR float64, phiMin float64, phiMax float64,
	rMin float64, rMax float64) (es []PolarEdge) {
	e := PolarEdge{
		Phi1: prevPhi,
		R1:   prevR,
		Phi2: point.n,
		R2:   point.val,
		Line: point.fromPrev,
	}
	if e.R1 > rMax || e.R2 > rMax || e.R1 < rMin || e.R2 < rMin || e.Phi1 < phiMin || e.Phi2 < phiMin ||
		e.Phi1 > phiMax || e.Phi2 > phiMax {
		// points out of range
		return
	}

	es = append(es, e)

	return
}

type LineSeries struct {
	baseSeries
	showDots bool
	data     []*linePoint
}

func EmptyLineSeries(chart chart, name string, showDots bool, color color.Color, polar bool) (ser *LineSeries) {
	ser = &LineSeries{
		showDots: showDots,
	}
	ser.baseSeries = emptyBaseSeries(chart, name, color, polar, ser.toggleView)
	return
}

func (ser *LineSeries) TRange() (isEmpty bool, min time.Time, max time.Time) {
	isEmpty = false
	ser.mutex.Lock()
	if len(ser.data) == 0 {
		isEmpty = true
		ser.mutex.Unlock()
		return
	}
	min = ser.data[0].t
	max = ser.data[len(ser.data)-1].t
	ser.mutex.Unlock()
	return
}

func (ser *LineSeries) NRange() (isEmpty bool, min float64, max float64) {
	min = 0
	max = 0
	isEmpty = false
	ser.mutex.Lock()
	if len(ser.data) == 0 {
		isEmpty = true
		ser.mutex.Unlock()
		return
	}
	min = ser.data[0].n
	max = ser.data[len(ser.data)-1].n
	ser.mutex.Unlock()
	return
}

func (ser *LineSeries) ValRange() (isEmpty bool, min float64, max float64) {
	min = 0
	max = 0
	isEmpty = false
	ser.mutex.Lock()
	if len(ser.data) == 0 {
		isEmpty = true
		ser.mutex.Unlock()
		return
	}
	min = ser.data[0].val
	max = ser.data[0].val
	for i := range ser.data {
		if ser.data[i].val < min {
			min = ser.data[i].val
		}
		if ser.data[i].val > max {
			max = ser.data[i].val
		}
	}
	ser.mutex.Unlock()
	return
}

func (ser *LineSeries) ConvertTtoN(tToN func(t time.Time) (n float64)) {
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].n = tToN(ser.data[i].t)
	}
	ser.mutex.Unlock()
}

func (ser *LineSeries) CartesianNodes(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []CartesianNode) {
	ser.mutex.Lock()
	for i := range ser.data {
		ns = append(ns, ser.data[i].cartesianNodes(xMin, xMax, yMin, yMax)...)
	}
	ser.mutex.Unlock()
	return
}

func (ser *LineSeries) CartesianEdges(xMin float64, xMax float64, yMin float64,
	yMax float64) (es []CartesianEdge) {
	ser.mutex.Lock()
	for i := range ser.data {
		if i > 0 {
			es = append(es, ser.data[i].cartesianEdges(ser.data[i-1].n, ser.data[i-1].val,
				xMin, xMax, yMin, yMax)...)
		}
	}
	ser.mutex.Unlock()
	return
}

func (ser *LineSeries) PolarNodes(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ns []PolarNode) {
	ser.mutex.Lock()
	for i := range ser.data {
		ns = append(ns, ser.data[i].polarNodes(phiMin, phiMax, rMin, rMax)...)
	}
	ser.mutex.Unlock()
	return
}

func (ser *LineSeries) PolarEdges(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (es []PolarEdge) {
	ser.mutex.Lock()
	for i := range ser.data {
		if i > 0 {
			es = append(es, ser.data[i].polarEdges(ser.data[i-1].n, ser.data[i-1].val, phiMin, phiMax, rMin, rMax)...)
		}
	}
	ser.mutex.Unlock()
	return
}

// Show makes all elements of the series visible
func (ser *LineSeries) Show() {
	ser.mutex.Lock()
	ser.visible = true
	for i := range ser.data {
		ser.data[i].show()
	}
	ser.mutex.Unlock()
}

// Hide hides all elements of the series
func (ser *LineSeries) Hide() {
	ser.mutex.Lock()
	ser.visible = false
	for i := range ser.data {
		ser.data[i].hide()
	}
	ser.mutex.Unlock()
}

func (ser *LineSeries) toggleView() {
	if ser.visible {
		ser.Hide()
	} else {
		ser.Show()
	}
}

// SetColor changes the color of the bar series
func (ser *LineSeries) SetColor(col color.Color) {
	ser.mutex.Lock()
	ser.color = col
	ser.legendButton.SetColor(col)
	for i := range ser.data {
		ser.data[i].setColor(col)
	}
	ser.mutex.Unlock()
}

// SetLineWidth changes the width of the line
// Standard value is 1
// The provided width must be greater than zero for this method to take effect
func (ser *LineSeries) SetLineWidth(lw float32) {
	if lw < 0 {
		return
	}
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].setLineWidth(lw)
	}
	ser.mutex.Unlock()
}

// SetDotSize changes the size of the dots (if activated)
// Standard value is 5
// The provided size must be greater than zero for this method to take effect
func (ser *LineSeries) SetDotSize(ds float32) {
	if ds < 0 {
		return
	}
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].setDotSize(ds)
	}
	ser.mutex.Unlock()
}

func (ser *LineSeries) DeleteNumericalDataInRange(min float64, max float64) (c int, err error) {
	c = 0
	if min > max {
		err = errors.New("invalid range")
		return
	}
	finalData := []*linePoint{}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	for i := range ser.data {
		if ser.data[i].n > min && ser.data[i].n < max {
			c++
		} else {
			finalData = append(finalData, ser.data[i])
		}
	}
	if c == 0 {
		ser.mutex.Unlock()
		return
	}
	ser.data = nil
	ser.data = finalData
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *LineSeries) AddNumericalData(input []data.NumericalDataPoint) (err error) {
	if len(input) == 0 {
		err = errors.New("no input data")
		return
	}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	newData := append(make([]data.NumericalDataPoint, 0, len(input)), input...)
	for i := range ser.data {
		newData = append(newData, data.NumericalDataPoint{X: ser.data[i].n, Val: ser.data[i].val})
	}
	sort.Sort(data.DpByNValue(newData))
	ser.data = nil

	for i := range newData {
		lPoint := emptyLinePoint(ser.showDots, ser.color)
		lPoint.n = newData[i].X
		lPoint.val = newData[i].Val
		ser.data = append(ser.data, lPoint)
	}
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *LineSeries) DeleteTemporalDataInRange(min time.Time, max time.Time) (c int, err error) {
	c = 0
	if min.After(max) {
		err = errors.New("invalid range")
		return
	}
	finalData := []*linePoint{}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	for i := range ser.data {
		if ser.data[i].t.After(min) && ser.data[i].t.Before(max) {
			c++
		} else {
			finalData = append(finalData, ser.data[i])
		}
	}
	if c == 0 {
		ser.mutex.Unlock()
		return
	}
	ser.data = nil
	ser.data = finalData
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *LineSeries) AddTemporalData(input []data.TemporalDataPoint) (err error) {
	if len(input) == 0 {
		err = errors.New("no input data")
		return
	}
	if ser.polar {
		for i := range input {
			if input[i].Val < 0 {
				err = errors.New("invalid data")
				return
			}
		}
	}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	newData := append(make([]data.TemporalDataPoint, 0, len(input)), input...)
	for i := range ser.data {
		newData = append(newData, data.TemporalDataPoint{T: ser.data[i].t, Val: ser.data[i].val})
	}
	sort.Sort(data.DpByTValue(newData))
	ser.data = nil

	for i := range newData {
		lPoint := emptyLinePoint(ser.showDots, ser.color)
		lPoint.t = newData[i].T
		lPoint.val = newData[i].Val
		ser.data = append(ser.data, lPoint)
	}
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *LineSeries) DeleteAngularDataInRange(min float64, max float64) (c int, err error) {
	c = 0
	if min > max {
		err = errors.New("invalid range")
		return
	}
	finalData := []*linePoint{}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	for i := range ser.data {
		if ser.data[i].n > min && ser.data[i].n < max {
			c++
		} else {
			finalData = append(finalData, ser.data[i])
		}
	}
	if c == 0 {
		ser.mutex.Unlock()
		return
	}
	ser.data = nil
	ser.data = finalData
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *LineSeries) AddAngularData(input []data.AngularDataPoint) (err error) {
	if len(input) == 0 {
		err = errors.New("no input data")
		return
	}

	for i := range input {
		if input[i].Val < 0 || input[i].A < 0 || input[i].A > 2*math.Pi {
			err = errors.New("invalid data")
			return
		}
	}

	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	newData := append(make([]data.AngularDataPoint, 0, len(input)), input...)
	for i := range ser.data {
		newData = append(newData, data.AngularDataPoint{A: ser.data[i].n, Val: ser.data[i].val})
	}
	sort.Sort(data.DpByAValue(newData))
	ser.data = nil

	for i := range newData {
		lPoint := emptyLinePoint(ser.showDots, ser.color)
		lPoint.n = newData[i].A
		lPoint.val = newData[i].Val
		ser.data = append(ser.data, lPoint)
	}
	ser.mutex.Unlock()
	chart.DataChange()
	return
}
