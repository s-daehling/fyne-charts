package series

import (
	"errors"
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type scatterPoint struct {
	c         string
	t         time.Time
	n         float64
	val       float64
	dot       *canvas.Circle
	line      *canvas.Line
	valOrigin float64
}

func emptyScatterPoint(color color.Color) (point *scatterPoint) {
	point = &scatterPoint{
		dot:  canvas.NewCircle(color),
		line: canvas.NewLine(color),
	}
	point.dot.Resize(fyne.NewSize(5, 5))
	return
}

func (point *scatterPoint) hide() {
	point.dot.Hide()
	point.line.Hide()
}

func (point *scatterPoint) show() {
	point.dot.Show()
	point.line.Show()
}

func (point *scatterPoint) setColor(col color.Color) {
	point.dot.FillColor = col
	point.line.StrokeColor = col
}

func (point *scatterPoint) setLineWidth(lw float32) {
	point.line.StrokeWidth = lw
}

func (point *scatterPoint) setDotSize(ds float32) {
	point.dot.Resize(fyne.NewSize(ds, ds))
}

func (point *scatterPoint) setValOrigin(vo float64) {
	point.valOrigin = vo
}

func (point *scatterPoint) cartesianNodes(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []CartesianNode) {
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

func (point *scatterPoint) polarNodes(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ns []PolarNode) {
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

func (point *scatterPoint) cartesianEdges(xMin float64, xMax float64, yMin float64,
	yMax float64) (es []CartesianEdge) {
	if point.n > xMax || point.n < xMin {
		// point out of range
		return
	}
	e := CartesianEdge{
		X1:   point.n,
		Y1:   point.valOrigin,
		X2:   point.n,
		Y2:   point.val,
		Line: point.line,
	}
	if point.val > yMax {
		point.val = yMax
	} else if point.val < yMin {
		point.val = yMin
	}
	es = append(es, e)
	return
}

func (point *scatterPoint) polarEdges(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (es []PolarEdge) {
	if point.n < phiMin || point.n > phiMax || point.val < rMin {
		return
	}
	e := PolarEdge{
		Phi1: point.n,
		R1:   0.0,
		Phi2: point.n,
		R2:   point.val,
		Line: point.line,
	}
	if e.R2 > rMax {
		e.R2 = rMax
	}
	es = append(es, e)

	return
}

type ScatterSeries struct {
	baseSeries
	data []*scatterPoint
}

func EmptyScatterSeries(chart chart, name string, color color.Color, polar bool) (ser *ScatterSeries) {
	ser = &ScatterSeries{}
	ser.baseSeries = emptyBaseSeries(chart, name, color, polar, ser.toggleView)
	return
}

func (ser *ScatterSeries) CRange() (cs []string) {
	ser.mutex.Lock()
	for i := range ser.data {
		cs = append(cs, ser.data[i].c)
	}
	ser.mutex.Unlock()
	return
}

func (ser *ScatterSeries) TRange() (isEmpty bool, min time.Time, max time.Time) {
	isEmpty = false
	ser.mutex.Lock()
	if len(ser.data) == 0 {
		isEmpty = true
		ser.mutex.Unlock()
		return
	}
	min = ser.data[0].t
	max = ser.data[0].t
	for i := range ser.data {
		if ser.data[i].t.Before(min) {
			min = ser.data[i].t
		}
		if ser.data[i].t.After(max) {
			max = ser.data[i].t
		}
	}
	ser.mutex.Unlock()
	return
}

func (ser *ScatterSeries) NRange() (isEmpty bool, min float64, max float64) {
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
	max = ser.data[0].n
	for i := range ser.data {
		if ser.data[i].n < min {
			min = ser.data[i].n
		}
		if ser.data[i].n > max {
			max = ser.data[i].n
		}
	}
	ser.mutex.Unlock()
	return
}

func (ser *ScatterSeries) ValRange() (isEmpty bool, min float64, max float64) {
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

func (ser *ScatterSeries) ConvertCtoN(cToN func(c string) (n float64)) {
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].n = cToN(ser.data[i].c)
	}
	ser.mutex.Unlock()
}

func (ser *ScatterSeries) ConvertTtoN(tToN func(t time.Time) (n float64)) {
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].n = tToN(ser.data[i].t)
	}
	ser.mutex.Unlock()
}

func (ser *ScatterSeries) CartesianNodes(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []CartesianNode) {
	ser.mutex.Lock()
	for i := range ser.data {
		ns = append(ns, ser.data[i].cartesianNodes(xMin, xMax, yMin, yMax)...)
	}
	ser.mutex.Unlock()
	return
}

func (ser *ScatterSeries) PolarNodes(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ns []PolarNode) {
	ser.mutex.Lock()
	for i := range ser.data {
		ns = append(ns, ser.data[i].polarNodes(phiMin, phiMax, rMin, rMax)...)
	}
	ser.mutex.Unlock()
	return
}

// Show makes all elements of the series visible
func (ser *ScatterSeries) Show() {
	ser.mutex.Lock()
	ser.visible = true
	for i := range ser.data {
		ser.data[i].show()
	}
	ser.mutex.Unlock()
}

// Hide hides all elements of the series
func (ser *ScatterSeries) Hide() {
	ser.mutex.Lock()
	ser.visible = false
	for i := range ser.data {
		ser.data[i].hide()
	}
	ser.mutex.Unlock()
}

func (ser *ScatterSeries) toggleView() {
	if ser.visible {
		ser.Hide()
	} else {
		ser.Show()
	}
}

// SetColor changes the color of the bar series
func (ser *ScatterSeries) SetColor(col color.Color) {
	ser.mutex.Lock()
	ser.color = col
	ser.legendButton.SetColor(col)
	for i := range ser.data {
		ser.data[i].setColor(col)
	}
	ser.mutex.Unlock()
}

// SetDotSize changes the size of the dots
// Standard value is 5
// The provided size must be greater than zero for this method to take effect
func (ser *ScatterSeries) SetDotSize(ds float32) {
	if ds < 0 {
		return
	}
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].setDotSize(ds)
	}
	ser.mutex.Unlock()
}

func (ser *ScatterSeries) DeleteNumericalDataInRange(min float64, max float64) (c int, err error) {
	c = 0
	if min > max {
		err = errors.New("invalid range")
		return
	}
	finalData := []*scatterPoint{}
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

func (ser *ScatterSeries) AddNumericalData(input []data.NumericalDataPoint) (err error) {
	err = numericalDataPointRangeCheck(input, ser.polar, ser.polar)
	if err != nil {
		return
	}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	for i := range input {
		lPoint := emptyScatterPoint(ser.color)
		lPoint.n = input[i].X
		lPoint.val = input[i].Val
		ser.data = append(ser.data, lPoint)
	}
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *ScatterSeries) DeleteTemporalDataInRange(min time.Time, max time.Time) (c int, err error) {
	c = 0
	if min.After(max) {
		err = errors.New("invalid range")
		return
	}
	finalData := []*scatterPoint{}
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

func (ser *ScatterSeries) AddTemporalData(input []data.TemporalDataPoint) (err error) {
	err = temporalDataPointRangeCheck(input, ser.polar)
	if err != nil {
		return
	}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	for i := range input {
		lPoint := emptyScatterPoint(ser.color)
		lPoint.t = input[i].T
		lPoint.val = input[i].Val
		ser.data = append(ser.data, lPoint)
	}
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *ScatterSeries) DeleteCategoricalDataInRange(cat []string) (c int, err error) {
	c = 0
	if len(cat) == 0 {
		err = errors.New("invalid range")
		return
	}
	finalData := []*scatterPoint{}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
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
		ser.mutex.Unlock()
		return
	}
	ser.data = nil
	ser.data = finalData
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *ScatterSeries) AddCategoricalData(input []data.CategoricalDataPoint) (err error) {
	err = categoricalDataPointRangeCheck(input, ser.polar)
	if err != nil {
		return
	}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
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
		sPoint := emptyScatterPoint(ser.color)
		sPoint.c = input[i].C
		sPoint.val = input[i].Val
		ser.data = append(ser.data, sPoint)
	}
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *ScatterSeries) DeleteAngularDataInRange(min float64, max float64) (c int, err error) {
	c, err = ser.DeleteNumericalDataInRange(min, max)
	return
}

func (ser *ScatterSeries) AddAngularData(input []data.AngularDataPoint) (err error) {
	err = ser.AddNumericalData(angularToNumerical(input))
	return
}
