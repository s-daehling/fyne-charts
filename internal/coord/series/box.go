package series

import (
	"errors"
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/renderer"
	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type boxPoint struct {
	c            string
	t            time.Time
	n            float64
	max          float64
	thirdQuart   float64
	median       float64
	firstQuart   float64
	min          float64
	outlier      []float64
	maxLine      *canvas.Line
	upperWhisker *canvas.Line
	medianLine   *canvas.Line
	lowerWhisker *canvas.Line
	minLine      *canvas.Line
	outlierDots  []*canvas.Circle
	box          *canvas.Rectangle
	width        float64
}

func emptyBoxPoint(nOutliers int, col color.Color) (point *boxPoint) {
	point = &boxPoint{
		maxLine:      canvas.NewLine(col),
		upperWhisker: canvas.NewLine(col),
		medianLine:   canvas.NewLine(col),
		lowerWhisker: canvas.NewLine(col),
		minLine:      canvas.NewLine(col),
	}
	point.box = canvas.NewRectangle(color.RGBA{0x00, 0x00, 0x00, 0x00})
	point.box.StrokeColor = col
	point.box.StrokeWidth = 1
	for range nOutliers {
		p := canvas.NewCircle(col)
		p.Resize(fyne.NewSize(5, 5))
		point.outlierDots = append(point.outlierDots, p)
	}
	return
}

func (point *boxPoint) hide() {
	point.maxLine.Hide()
	point.upperWhisker.Hide()
	point.medianLine.Hide()
	point.lowerWhisker.Hide()
	point.minLine.Hide()
	point.box.Hide()
	for i := range point.outlierDots {
		point.outlierDots[i].Hide()
	}
}

func (point *boxPoint) show() {
	point.maxLine.Show()
	point.upperWhisker.Show()
	point.medianLine.Show()
	point.lowerWhisker.Show()
	point.minLine.Show()
	point.box.Show()
	for i := range point.outlierDots {
		point.outlierDots[i].Show()
	}
}

func (point *boxPoint) setColor(col color.Color) {
	point.maxLine.StrokeColor = col
	point.upperWhisker.StrokeColor = col
	point.medianLine.StrokeColor = col
	point.lowerWhisker.StrokeColor = col
	point.minLine.StrokeColor = col
	for i := range point.outlierDots {
		point.outlierDots[i].FillColor = col
	}
	point.box.FillColor = col
}

func (point *boxPoint) setLineWidth(lw float32) {
	point.maxLine.StrokeWidth = lw
	point.upperWhisker.StrokeWidth = lw
	point.medianLine.StrokeWidth = lw
	point.lowerWhisker.StrokeWidth = lw
	point.minLine.StrokeWidth = lw
	point.box.StrokeWidth = lw
}

func (point *boxPoint) setOutlierSize(os float32) {
	for i := range point.outlierDots {
		point.outlierDots[i].Resize(fyne.NewSize(os, os))
	}
}

func (point *boxPoint) setWidth(width float64) {
	point.width = width
}

func (point *boxPoint) cartesianNodes(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []renderer.CartesianNode) {
	if point.n < xMin || point.n > xMax || point.min < yMin || point.max > yMax {
		return
	}
	for i := range point.outlier {
		if point.outlier[i] < yMin || point.outlier[i] > yMax {
			continue
		}
		n := renderer.CartesianNode{
			X:   point.n,
			Y:   point.outlier[i],
			Dot: point.outlierDots[i],
		}
		ns = append(ns, n)
	}
	return
}

func (point *boxPoint) cartesianEdges(xMin float64, xMax float64, yMin float64,
	yMax float64) (es []renderer.CartesianEdge) {
	if point.n < xMin || point.n > xMax || point.min < yMin || point.max > yMax {
		return
	}
	e1 := renderer.CartesianEdge{
		X1:   point.n - (point.width / 2),
		Y1:   point.max,
		X2:   point.n + (point.width / 2),
		Y2:   point.max,
		Line: point.maxLine,
	}
	es = append(es, e1)
	e2 := renderer.CartesianEdge{
		X1:   point.n,
		Y1:   point.thirdQuart,
		X2:   point.n,
		Y2:   point.max,
		Line: point.upperWhisker,
	}
	es = append(es, e2)
	e3 := renderer.CartesianEdge{
		X1:   point.n - (point.width / 2),
		Y1:   point.median,
		X2:   point.n + (point.width / 2),
		Y2:   point.median,
		Line: point.medianLine,
	}
	es = append(es, e3)
	e4 := renderer.CartesianEdge{
		X1:   point.n,
		Y1:   point.min,
		X2:   point.n,
		Y2:   point.firstQuart,
		Line: point.lowerWhisker,
	}
	es = append(es, e4)
	e5 := renderer.CartesianEdge{
		X1:   point.n - (point.width / 2),
		Y1:   point.min,
		X2:   point.n + (point.width / 2),
		Y2:   point.min,
		Line: point.minLine,
	}
	es = append(es, e5)
	return
}

func (point *boxPoint) cartesianRects(xMin float64, xMax float64,
	yMin float64, yMax float64) (as []renderer.CartesianRect) {
	if point.n < xMin || point.n > xMax || point.min < yMin || point.max > yMax {
		return
	}
	a := renderer.CartesianRect{
		X1:   point.n - (point.width / 2),
		Y1:   point.firstQuart,
		X2:   point.n + (point.width / 2),
		Y2:   point.thirdQuart,
		Rect: point.box,
	}
	as = append(as, a)
	return
}

type BoxSeries struct {
	baseSeries
	data []*boxPoint
}

func EmptyBoxSeries(name string, color color.Color) (ser *BoxSeries) {
	ser = &BoxSeries{}
	ser.baseSeries = emptyBaseSeries(name, color, ser.toggleView)
	return
}

func (ser *BoxSeries) CRange() (cs []string) {
	for i := range ser.data {
		cs = append(cs, ser.data[i].c)
	}
	return
}

func (ser *BoxSeries) TRange() (isEmpty bool, min time.Time, max time.Time) {
	isEmpty = false
	if len(ser.data) == 0 {
		isEmpty = true
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
	return
}

func (ser *BoxSeries) NRange() (isEmpty bool, min float64, max float64) {
	min = 0
	max = 0
	isEmpty = false
	if len(ser.data) == 0 {
		isEmpty = true
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
	return
}
func (ser *BoxSeries) ValRange() (isEmpty bool, min float64, max float64) {
	min = 0
	max = 0
	isEmpty = false
	if len(ser.data) == 0 {
		isEmpty = true
		return
	}
	min = ser.data[0].min
	max = ser.data[0].max
	for i := range ser.data {
		if ser.data[i].min < min {
			min = ser.data[i].min
		}
		if ser.data[i].max > max {
			max = ser.data[i].max
		}
		for j := range ser.data[i].outlier {
			if ser.data[i].outlier[j] < min {
				min = ser.data[i].outlier[j]
			}
			if ser.data[i].outlier[j] > max {
				max = ser.data[i].outlier[j]
			}
		}
	}
	return
}

func (ser *BoxSeries) ConvertCtoN(cToN func(c string) (n float64)) {
	for i := range ser.data {
		ser.data[i].n = cToN(ser.data[i].c)
	}
}

func (ser *BoxSeries) ConvertTtoN(tToN func(t time.Time) (n float64)) {
	for i := range ser.data {
		ser.data[i].n = tToN(ser.data[i].t)
	}
}

func (ser *BoxSeries) CartesianNodes(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []renderer.CartesianNode) {
	for i := range ser.data {
		ns = append(ns, ser.data[i].cartesianNodes(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *BoxSeries) CartesianEdges(xMin float64, xMax float64, yMin float64,
	yMax float64) (es []renderer.CartesianEdge) {
	for i := range ser.data {
		es = append(es, ser.data[i].cartesianEdges(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *BoxSeries) CartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (as []renderer.CartesianRect) {
	for i := range ser.data {
		as = append(as, ser.data[i].cartesianRects(xMin, xMax, yMin, yMax)...)
	}
	return
}

// setWidth sets width of boxes for this series
func (ser *BoxSeries) SetWidth(width float64) {
	for i := range ser.data {
		ser.data[i].setWidth(width)
	}
}

func (ser *BoxSeries) NumberOfPoints() (n int) {
	n = len(ser.data)
	return
}

// Show makes all elements of the bar series visible
func (ser *BoxSeries) Show() {
	ser.visible = true
	for i := range ser.data {
		ser.data[i].show()
	}
}

// Hide hides all elements of the bar series
func (ser *BoxSeries) Hide() {
	ser.visible = false
	for i := range ser.data {
		ser.data[i].hide()
	}
}

func (ser *BoxSeries) toggleView() {
	if ser.visible {
		ser.Hide()
	} else {
		ser.Show()
	}
}

// SetColor changes the color of the bar series
func (ser *BoxSeries) SetColor(col color.Color) {
	ser.color = col
	ser.legendButton.SetRectColor(col)
	for i := range ser.data {
		ser.data[i].setColor(col)
	}
}

// SetLineWidth changes the width of the Line
// Standard value is 1
// The provided width must be greater than zero for this method to take effect
func (ser *BoxSeries) SetLineWidth(lw float32) {
	if lw < 0 {
		return
	}
	for i := range ser.data {
		ser.data[i].setLineWidth(lw)
	}
}

// SetOutlierSize changes the size of the outlier dots
// Standard value is 5
// The provided size must be greater than zero for this method to take effect
func (ser *BoxSeries) SetOutlierSize(os float32) {
	if os < 0 {
		return
	}
	for i := range ser.data {
		ser.data[i].setOutlierSize(os)
	}
}

func (ser *BoxSeries) Clear() {
	ser.data = []*boxPoint{}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
}

// DeleteDataInRange deletes all boxes with a x-coordinate greater than min and smaller than max
// The return value gives the number of boxes that have been removed
func (ser *BoxSeries) DeleteNumericalDataInRange(min float64, max float64) (c int) {
	c = 0
	if min > max {
		return
	}
	finalData := []*boxPoint{}
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

// AddData adds boxes to the series.
// The method does not check for duplicates (i.e. boxes with same X)
func (ser *BoxSeries) AddNumericalData(input []data.NumericalBox) (err error) {
	if len(input) == 0 {
		return
	}
	for i := range input {
		if input[i].Minimum > input[i].FirstQuartile || input[i].FirstQuartile > input[i].Median ||
			input[i].Median > input[i].ThirdQuartile || input[i].ThirdQuartile > input[i].Maximum {
			err = errors.New("invalid data")
			return
		}
	}
	for i := range input {
		bPoint := emptyBoxPoint(len(input[i].Outlier), ser.color)
		bPoint.n = input[i].N
		bPoint.max = input[i].Maximum
		bPoint.thirdQuart = input[i].ThirdQuartile
		bPoint.median = input[i].Median
		bPoint.firstQuart = input[i].FirstQuartile
		bPoint.min = input[i].Minimum
		bPoint.outlier = append(bPoint.outlier, input[i].Outlier...)
		ser.data = append(ser.data, bPoint)
	}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}

// DeleteDataInRange deletes all boxes with a t-coordinate after min and before max.
// The return value gives the number of boxes that have been removed
func (ser *BoxSeries) DeleteTemporalDataInRange(min time.Time, max time.Time) (c int) {
	c = 0
	if min.After(max) {
		return
	}
	finalData := []*boxPoint{}
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

// AddData adds boxes to the series.
// The method does not check for duplicates (i.e. boxes with same T)
func (ser *BoxSeries) AddTemporalData(input []data.TemporalBox) (err error) {
	if len(input) == 0 {
		return
	}
	for i := range input {
		if input[i].Minimum > input[i].FirstQuartile || input[i].FirstQuartile > input[i].Median ||
			input[i].Median > input[i].ThirdQuartile || input[i].ThirdQuartile > input[i].Maximum {
			err = errors.New("invalid data")
			return
		}
	}
	for i := range input {
		bPoint := emptyBoxPoint(len(input[i].Outlier), ser.color)
		bPoint.t = input[i].T
		bPoint.max = input[i].Maximum
		bPoint.thirdQuart = input[i].ThirdQuartile
		bPoint.median = input[i].Median
		bPoint.firstQuart = input[i].FirstQuartile
		bPoint.min = input[i].Minimum
		bPoint.outlier = append(bPoint.outlier, input[i].Outlier...)
		ser.data = append(ser.data, bPoint)
	}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}

// DeleteDataInRange deletes all boxes with one of the given category
// The return value gives the number of boxes that have been removed
func (ser *BoxSeries) DeleteCategoricalDataInRange(cat []string) (c int) {
	c = 0
	if len(cat) == 0 {
		return
	}
	finalData := []*boxPoint{}
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

// AddData adds boxes to the series.
// The method checks for duplicates (i.e. boxes with same C).
// Boxes with a C that already exists, will be ignored.
func (ser *BoxSeries) AddCategoricalData(input []data.CategoricalBox) (err error) {
	if len(input) == 0 {
		return
	}
	for i := range input {
		if input[i].Minimum > input[i].FirstQuartile || input[i].FirstQuartile > input[i].Median ||
			input[i].Median > input[i].ThirdQuartile || input[i].ThirdQuartile > input[i].Maximum {
			err = errors.New("invalid data")
			return
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
		bPoint := emptyBoxPoint(len(input[i].Outlier), ser.color)
		bPoint.c = input[i].C
		bPoint.max = input[i].Maximum
		bPoint.thirdQuart = input[i].ThirdQuartile
		bPoint.median = input[i].Median
		bPoint.firstQuart = input[i].FirstQuartile
		bPoint.min = input[i].Minimum
		bPoint.outlier = append(bPoint.outlier, input[i].Outlier...)
		ser.data = append(ser.data, bPoint)
	}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}
