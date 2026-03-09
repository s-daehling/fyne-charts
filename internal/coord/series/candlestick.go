package series

import (
	"errors"
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/elements"
	"github.com/s-daehling/fyne-charts/internal/style"
	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type candleStickPoint struct {
	tStart      time.Time
	tEnd        time.Time
	nStart      float64
	nEnd        float64
	open        float64
	close       float64
	high        float64
	low         float64
	candle      *elements.Candle
	colDown     color.Color
	colUp       color.Color
	highlighted bool
	ser         *CandleStickSeries
}

func emptyCandleStickPoint(lineCol color.Color, colDown color.Color, colUp color.Color, ser *CandleStickSeries) (point *candleStickPoint) {
	point = &candleStickPoint{
		colDown:     colDown,
		colUp:       colUp,
		highlighted: false,
		ser:         ser,
	}
	point.candle = elements.NewCandle(lineCol, point.highlight, point.unhighlight)
	return
}

func (point *candleStickPoint) refresh() {
	point.candle.Refresh()
}

func (point *candleStickPoint) hide() {
	point.candle.Hide()
}

func (point *candleStickPoint) show() {
	point.candle.Show()
}

func (point *candleStickPoint) setLineWidth(lw float32) {
	point.candle.SetLineWidth(lw)
}

func (point *candleStickPoint) setColor(line, down, up color.Color) {
	point.colDown = down
	point.colUp = up
	point.candle.SetLineColor(line)
}

func (point *candleStickPoint) highlight() {
	point.highlighted = true
	point.ser.highlight()
}

func (point *candleStickPoint) unhighlight() {
	point.highlighted = false
	point.ser.unhighlight()
}

func (point *candleStickPoint) cartesianCandles(xMin float64, xMax float64, yMin float64,
	yMax float64) (cs []*elements.Candle) {
	if point.nEnd > xMax || point.nStart < xMin || point.high > yMax || point.low < yMin {
		// point out of range
		return
	}
	point.candle.SetCandleColor(point.colDown)
	if point.open < point.close {
		point.candle.SetCandleColor(point.colUp)
	}
	point.candle.N1 = point.nStart
	point.candle.N2 = point.nEnd
	point.candle.High = point.high
	point.candle.Low = point.low
	point.candle.Open = point.open
	point.candle.Close = point.close
	cs = append(cs, point.candle)
	return
}

type CandleStickSeries struct {
	baseSeries
	data         []*candleStickPoint
	colNameDown  fyne.ThemeColorName
	colDown      color.Color
	colDownFaded color.Color
	colNameUp    fyne.ThemeColorName
	colUp        color.Color
	colUpFaded   color.Color
}

func EmptyCandleStickSeries(name string) (ser *CandleStickSeries) {
	ser = &CandleStickSeries{
		colNameDown: theme.ColorNameError,
		colNameUp:   theme.ColorNameSuccess,
		colDown:     theme.Color(theme.ColorNameError),
		colUp:       theme.Color(theme.ColorNameSuccess),
	}
	ser.baseSeries = emptyBaseSeries(name, theme.ColorNameForeground, ser.toggleView, ser.highlight, ser.unhighlight)
	ser.colDownFaded = style.MakeFaded(ser.colDown)
	ser.colUpFaded = style.MakeFaded(ser.colUp)
	// ser.legendButton.UseGradient(color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, color.RGBA{R: 0x00, G: 0x88, B: 0x00, A: 0xff})
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

func (ser *CandleStickSeries) CartesianCandles(xMin float64, xMax float64, yMin float64,
	yMax float64) (cs []*elements.Candle) {
	for i := range ser.data {
		cs = append(cs, ser.data[i].cartesianCandles(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *CandleStickSeries) RefreshTheme() {
	ser.col = theme.Color(ser.colName)
	ser.colFaded = style.MakeFaded(ser.col)
	ser.colDown = theme.Color(ser.colNameDown)
	ser.colDownFaded = style.MakeFaded(ser.colDown)
	ser.colUp = theme.Color(ser.colNameUp)
	ser.colUpFaded = style.MakeFaded(ser.colUp)
	col := ser.col
	colDown := ser.colDown
	colUp := ser.colUp
	if ser.isFaded {
		col = ser.colFaded
		colDown = ser.colDownFaded
		colUp = ser.colUpFaded
	}
	for i := range ser.data {
		ser.data[i].setColor(col, colDown, colUp)
	}
	ser.legendEntry.SetColor(col)
}

func (ser *CandleStickSeries) FadeUnhighlighted() {
	if ser.highlighted {
		return
	}
	ser.isFaded = true
	for i := range ser.data {
		ser.data[i].setColor(ser.colFaded, ser.colDownFaded, ser.colUpFaded)
	}
}

func (ser *CandleStickSeries) UnFade() {
	ser.isFaded = false
	for i := range ser.data {
		ser.data[i].setColor(ser.col, ser.colDown, ser.colUp)
	}
}

func (ser *CandleStickSeries) highlight() {
	ser.highlighted = true
	if ser.cont != nil {
		ser.cont.Highlight()
	}
}

func (ser *CandleStickSeries) unhighlight() {
	ser.highlighted = false
	if ser.cont != nil {
		ser.cont.Unhighlight()
	}
}

// Show makes all elements of the series visible
func (ser *CandleStickSeries) Show() {
	ser.visible = true
	for i := range ser.data {
		ser.data[i].show()
	}
	ser.legendEntry.Show()
}

// Hide hides all elements of the series
func (ser *CandleStickSeries) Hide() {
	ser.visible = false
	for i := range ser.data {
		ser.data[i].hide()
	}
	ser.legendEntry.Hide()
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
		ser.data[i].refresh()
	}
}

func (ser *CandleStickSeries) Clear() {
	ser.data = []*candleStickPoint{}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
}

// DeleteDataInRange deletes all candles with a nEnd greater than min and a nStart smaller than max
// The return value gives the number of candles that have been removed
func (ser *CandleStickSeries) DeleteNumericalDataInRange(min float64, max float64) (c int) {
	c = 0
	if min > max {
		return
	}
	finalData := []*candleStickPoint{}
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
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}

// AddData adds candles to the series.
// The method does not check for duplicates (i.e. candles with same XStart or XEnd)
func (ser *CandleStickSeries) AddNumericalData(input []data.NumericalCandleStick) (err error) {
	if len(input) == 0 {
		return
	}
	for i := range input {
		if input[i].NEnd < input[i].NStart || input[i].Low > input[i].High {
			err = errors.New("invalid data")
			return
		}
	}
	for i := range input {
		csPoint := emptyCandleStickPoint(ser.col, ser.colDown, ser.colUp, ser)
		csPoint.nStart = input[i].NStart
		csPoint.nEnd = input[i].NEnd
		csPoint.open = input[i].Open
		csPoint.close = input[i].Close
		csPoint.high = input[i].High
		csPoint.low = input[i].Low
		ser.data = append(ser.data, csPoint)
	}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}

// DeleteDataInRange deletes all candles with a tEnd after min and a tStart before max.
// The return value gives the number of candles that have been removed
func (ser *CandleStickSeries) DeleteTemporalDataInRange(min time.Time, max time.Time) (c int) {
	c = 0
	if min.After(max) {
		return
	}
	finalData := []*candleStickPoint{}
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
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}

// AddData adds candles to the series.
// The method does not check for duplicates (i.e. candles with same TStart or TEnd)
func (ser *CandleStickSeries) AddTemporalData(input []data.TemporalCandleStick) (err error) {
	if len(input) == 0 {
		return
	}
	for i := range input {
		if input[i].TEnd.Before(input[i].TStart) || input[i].Low > input[i].High {
			err = errors.New("invalid data")
			return
		}
	}
	for i := range input {
		csPoint := emptyCandleStickPoint(ser.col, ser.colDown, ser.colUp, ser)
		csPoint.tStart = input[i].TStart
		csPoint.tEnd = input[i].TEnd
		csPoint.open = input[i].Open
		csPoint.close = input[i].Close
		csPoint.high = input[i].High
		csPoint.low = input[i].Low
		ser.data = append(ser.data, csPoint)
	}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}
