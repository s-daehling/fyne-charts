package series

import (
	"errors"
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/elements"
	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type boxPoint struct {
	c           string
	t           time.Time
	n           float64
	max         float64
	thirdQuart  float64
	median      float64
	firstQuart  float64
	min         float64
	outlier     []float64
	outlierDots []*elements.Dot
	box         *elements.Box
	width       float64
}

func emptyBoxPoint(nOutliers int, col color.Color) (point *boxPoint) {
	point = &boxPoint{
		box: elements.NewBox(col),
	}
	for range nOutliers {
		p := elements.NewDot(col, 5)
		p.Resize(fyne.NewSize(5, 5))
		point.outlierDots = append(point.outlierDots, p)
	}
	return
}

func (point *boxPoint) refresh() {
	point.box.Refresh()
	for i := range point.outlierDots {
		point.outlierDots[i].Refresh()
	}
}

func (point *boxPoint) hide() {
	point.box.Hide()
	for i := range point.outlierDots {
		point.outlierDots[i].Hide()
	}
}

func (point *boxPoint) show() {
	point.box.Show()
	for i := range point.outlierDots {
		point.outlierDots[i].Show()
	}
}

func (point *boxPoint) setColor(col color.Color) {
	for i := range point.outlierDots {
		point.outlierDots[i].SetColor(col)
	}
	point.box.SetColor(col)
}

func (point *boxPoint) setLineWidth(lw float32) {
	point.box.SetLineWidth(lw)
}

func (point *boxPoint) setOutlierSize(os float32) {
	for i := range point.outlierDots {
		point.outlierDots[i].SetMinSize(os)
		point.outlierDots[i].Resize(fyne.NewSize(os, os))
	}
}

func (point *boxPoint) setWidth(width float64) {
	point.width = width
}

func (point *boxPoint) cartesianDots(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []*elements.Dot) {
	if point.n < xMin || point.n > xMax || point.min < yMin || point.max > yMax {
		return
	}
	for i := range point.outlier {
		if point.outlier[i] < yMin || point.outlier[i] > yMax {
			continue
		}
		point.outlierDots[i].N = point.n
		point.outlierDots[i].Val = point.outlier[i]
		ns = append(ns, point.outlierDots[i])
	}
	return
}

func (point *boxPoint) cartesianBoxes(xMin float64, xMax float64,
	yMin float64, yMax float64) (bs []*elements.Box) {
	if point.n < xMin || point.n > xMax || point.min < yMin || point.max > yMax {
		return
	}
	point.box.N1 = point.n - (point.width / 2)
	point.box.N2 = point.n + (point.width / 2)
	point.box.Max = point.max
	point.box.ThirdQuart = point.thirdQuart
	point.box.Median = point.median
	point.box.FirstQuart = point.firstQuart
	point.box.Min = point.min
	bs = append(bs, point.box)
	return
}

type BoxSeries struct {
	baseSeries
	data []*boxPoint
}

func EmptyBoxSeries(name string, colName fyne.ThemeColorName) (ser *BoxSeries) {
	ser = &BoxSeries{}
	ser.baseSeries = emptyBaseSeries(name, colName, ser.toggleView)
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

func (ser *BoxSeries) CartesianDots(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []*elements.Dot) {
	for i := range ser.data {
		ns = append(ns, ser.data[i].cartesianDots(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *BoxSeries) CartesianBoxes(xMin float64, xMax float64, yMin float64,
	yMax float64) (bs []*elements.Box) {
	for i := range ser.data {
		bs = append(bs, ser.data[i].cartesianBoxes(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *BoxSeries) RefreshTheme() {
	ser.col = theme.Color(ser.colName)
	for i := range ser.data {
		ser.data[i].setColor(ser.col)
	}
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
	ser.legendEntry.Show()
}

// Hide hides all elements of the bar series
func (ser *BoxSeries) Hide() {
	ser.visible = false
	for i := range ser.data {
		ser.data[i].hide()
	}
	ser.legendEntry.Hide()
}

func (ser *BoxSeries) toggleView() {
	if ser.visible {
		ser.Hide()
	} else {
		ser.Show()
	}
}

// SetColor changes the color of the bar series
func (ser *BoxSeries) SetColor(colName fyne.ThemeColorName) {
	ser.colName = colName
	ser.col = theme.Color(ser.colName)
	ser.legendEntry.SetColor(colName)
	for i := range ser.data {
		ser.data[i].setColor(ser.col)
		ser.data[i].refresh()
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
		ser.data[i].refresh()
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
		bPoint := emptyBoxPoint(len(input[i].Outlier), ser.col)
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
		bPoint := emptyBoxPoint(len(input[i].Outlier), ser.col)
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
		bPoint := emptyBoxPoint(len(input[i].Outlier), ser.col)
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
