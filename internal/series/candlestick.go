package series

import (
	"errors"
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type candleStickPoint struct {
	tStart    time.Time
	tEnd      time.Time
	nStart    float64
	nEnd      float64
	open      float64
	close     float64
	high      float64
	low       float64
	upperLine *canvas.Line
	lowerLine *canvas.Line
	candle    *canvas.Rectangle
}

func emptyCandleStickPoint() (point *candleStickPoint) {
	point = &candleStickPoint{
		upperLine: canvas.NewLine(theme.Color(theme.ColorNameForeground)),
		lowerLine: canvas.NewLine(theme.Color(theme.ColorNameForeground)),
		candle:    canvas.NewRectangle(color.Black),
	}
	point.candle.CornerRadius = 2
	return
}

func (point *candleStickPoint) hide() {
	point.upperLine.Hide()
	point.lowerLine.Hide()
	point.candle.Hide()
}

func (point *candleStickPoint) show() {
	point.upperLine.Show()
	point.lowerLine.Show()
	point.candle.Show()
}

func (point *candleStickPoint) setLineWidth(lw float32) {
	point.upperLine.StrokeWidth = lw
	point.lowerLine.StrokeWidth = lw
}

func (point *candleStickPoint) cartesianEdges(xMin float64, xMax float64, yMin float64,
	yMax float64) (es []CartesianEdge) {
	if point.nEnd > xMax || point.nStart < xMin || point.high > yMax || point.low < yMin {
		// point out of range
		return
	}
	cMax := point.open
	cMin := point.close
	if point.open < point.close {
		cMax = point.close
		cMin = point.open
	}
	e1 := CartesianEdge{
		X1:   (point.nEnd + point.nStart) / 2,
		Y1:   cMax,
		X2:   (point.nEnd + point.nStart) / 2,
		Y2:   point.high,
		Line: point.upperLine,
	}
	es = append(es, e1)
	e2 := CartesianEdge{
		X1:   (point.nEnd + point.nStart) / 2,
		Y1:   point.low,
		X2:   (point.nEnd + point.nStart) / 2,
		Y2:   cMin,
		Line: point.lowerLine,
	}
	es = append(es, e2)
	return
}

func (point *candleStickPoint) cartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (as []CartesianRect) {
	if point.nEnd > xMax || point.nStart < xMin || point.high > yMax || point.low < yMin {
		// point out of range
		return
	}
	cMax := point.open
	cMin := point.close
	point.candle.FillColor = color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}
	if point.open < point.close {
		cMax = point.close
		cMin = point.open
		point.candle.FillColor = color.RGBA{R: 0x00, G: 0x88, B: 0x00, A: 0xff}
	}
	a := CartesianRect{
		X1:   point.nStart,
		Y1:   cMin,
		X2:   point.nEnd,
		Y2:   cMax,
		Rect: point.candle,
	}
	as = append(as, a)
	return
}

type CandleStickSeries struct {
	baseSeries
	data []*candleStickPoint
}

func EmptyCandleStickSeries(chart chart, name string, polar bool) (ser *CandleStickSeries) {
	ser = &CandleStickSeries{}
	ser.baseSeries = emptyBaseSeries(chart, name, theme.Color(theme.ColorNameForeground), polar, ser.toggleView)
	return
}

func (ser *CandleStickSeries) TRange() (isEmpty bool, min time.Time, max time.Time) {
	isEmpty = false
	if len(ser.data) == 0 {
		isEmpty = true
		return
	}
	min = ser.data[0].tStart
	max = ser.data[0].tEnd
	for i := range ser.data {
		if ser.data[i].tStart.Before(min) {
			min = ser.data[i].tStart
		}
		if ser.data[i].tEnd.After(max) {
			max = ser.data[i].tEnd
		}
	}
	return
}

func (ser *CandleStickSeries) NRange() (isEmpty bool, min float64, max float64) {
	min = 0
	max = 0
	isEmpty = false
	if len(ser.data) == 0 {
		isEmpty = true
		return
	}
	min = ser.data[0].nStart
	max = ser.data[0].nEnd
	for i := range ser.data {
		if ser.data[i].nStart < min {
			min = ser.data[i].nStart
		}
		if ser.data[i].nEnd > max {
			max = ser.data[i].nEnd
		}
	}
	return
}
func (ser *CandleStickSeries) ValRange() (isEmpty bool, min float64, max float64) {
	min = 0
	max = 0
	isEmpty = false
	if len(ser.data) == 0 {
		isEmpty = true
		return
	}
	min = ser.data[0].low
	max = ser.data[0].high
	for i := range ser.data {
		if ser.data[i].low < min {
			min = ser.data[i].low
		}
		if ser.data[i].high > max {
			max = ser.data[i].high
		}
	}
	return
}

func (ser *CandleStickSeries) ConvertTtoN(tToN func(t time.Time) (n float64)) {
	for i := range ser.data {
		ser.data[i].nStart = tToN(ser.data[i].tStart)
		ser.data[i].nEnd = tToN(ser.data[i].tEnd)
	}
}

func (ser *CandleStickSeries) CartesianEdges(xMin float64, xMax float64, yMin float64,
	yMax float64) (es []CartesianEdge) {
	for i := range ser.data {
		es = append(es, ser.data[i].cartesianEdges(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *CandleStickSeries) CartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []CartesianRect) {
	for i := range ser.data {
		fs = append(fs, ser.data[i].cartesianRects(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *CandleStickSeries) RefreshThemeColor() {
	ser.legendLabel.Color = theme.Color(theme.ColorNameForeground)
	ser.color = theme.Color(theme.ColorNameForeground)
	ser.legendButton.SetColor(theme.Color(theme.ColorNameForeground))
}

// Show makes all elements of the series visible
func (ser *CandleStickSeries) Show() {
	ser.visible = true
	for i := range ser.data {
		ser.data[i].show()
	}
}

// Hide hides all elements of the series
func (ser *CandleStickSeries) Hide() {
	ser.visible = false
	for i := range ser.data {
		ser.data[i].hide()
	}
}

func (ser *CandleStickSeries) toggleView() {
	if ser.visible {
		ser.Hide()
	} else {
		ser.Show()
	}
}

// SetLineWidth changes the width of the upper and lower line
// Standard value is 1
// The provided width must be greater than zero for this method to take effect
func (ser *CandleStickSeries) SetLineWidth(lw float32) {
	if lw < 0 {
		return
	}
	for i := range ser.data {
		ser.data[i].setLineWidth(lw)
	}
}

func (ser *CandleStickSeries) Clear() (err error) {
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
	chart := ser.chart
	ser.data = []*candleStickPoint{}
	chart.DataChange()
	return
}

// DeleteDataInRange deletes all candles with a nEnd greater than min and a nStart smaller than max
// The return value gives the number of candles that have been removed
func (ser *CandleStickSeries) DeleteNumericalDataInRange(min float64, max float64) (c int, err error) {
	c = 0
	if min > max {
		err = errors.New("invalid range")
		return
	}
	finalData := []*candleStickPoint{}
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
	for i := range ser.data {
		if ser.data[i].nStart > min && ser.data[i].nEnd < max {
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

// AddData adds candles to the series.
// The method does not check for duplicates (i.e. candles with same XStart or XEnd)
func (ser *CandleStickSeries) AddNumericalData(input []data.NumericalCandleStick) (err error) {
	if len(input) == 0 {
		err = errors.New("no input data")
		return
	}
	for i := range input {
		if input[i].NEnd < input[i].NStart || input[i].Low > input[i].High {
			err = errors.New("invalid data")
			return
		}
	}
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
	for i := range input {
		csPoint := emptyCandleStickPoint()
		csPoint.nStart = input[i].NStart
		csPoint.nEnd = input[i].NEnd
		csPoint.open = input[i].Open
		csPoint.close = input[i].Close
		csPoint.high = input[i].High
		csPoint.low = input[i].Low
		ser.data = append(ser.data, csPoint)
	}
	ser.chart.DataChange()
	return
}

// DeleteDataInRange deletes all candles with a tEnd after min and a tStart before max.
// The return value gives the number of candles that have been removed
func (ser *CandleStickSeries) DeleteTemporalDataInRange(min time.Time, max time.Time) (c int, err error) {
	c = 0
	if min.After(max) {
		err = errors.New("invalid range")
		return
	}
	finalData := []*candleStickPoint{}
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
	for i := range ser.data {
		if ser.data[i].tStart.After(min) && ser.data[i].tEnd.Before(max) {
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

// AddData adds candles to the series.
// The method does not check for duplicates (i.e. candles with same TStart or TEnd)
func (ser *CandleStickSeries) AddTemporalData(input []data.TemporalCandleStick) (err error) {
	if len(input) == 0 {
		err = errors.New("no input data")
		return
	}
	for i := range input {
		if input[i].TEnd.Before(input[i].TStart) || input[i].Low > input[i].High {
			err = errors.New("invalid data")
			return
		}
	}
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
	for i := range input {
		csPoint := emptyCandleStickPoint()
		csPoint.tStart = input[i].TStart
		csPoint.tEnd = input[i].TEnd
		csPoint.open = input[i].Open
		csPoint.close = input[i].Close
		csPoint.high = input[i].High
		csPoint.low = input[i].Low
		ser.data = append(ser.data, csPoint)
	}
	ser.chart.DataChange()
	return
}
