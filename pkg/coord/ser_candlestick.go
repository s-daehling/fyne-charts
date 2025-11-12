package coord

import (
	"time"

	"github.com/s-daehling/fyne-charts/internal/coord/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

type candleStickSeries struct {
	ser *series.CandleStickSeries
}

// Name returns the name of the series
func (cs candleStickSeries) Name() (n string) {
	if cs.ser == nil {
		return
	}
	n = cs.ser.Name()
	return
}

// Show makes the elements of the series visible
func (cs candleStickSeries) Show() {
	if cs.ser == nil {
		return
	}
	cs.ser.Show()
}

// Hide makes the elements of the series invisible
func (cs candleStickSeries) Hide() {
	if cs.ser == nil {
		return
	}
	cs.ser.Hide()
}

// SetLineWidth sets the width of the high and low line
func (cs candleStickSeries) SetLineWidth(lw float32) {
	if cs.ser == nil {
		return
	}
	cs.ser.SetLineWidth(lw)
}

// Clear deletes all data
func (cs candleStickSeries) Clear() {
	if cs.ser == nil {
		return
	}
	cs.ser.Clear()
}

// NumericalCandleStickSeries represents a candle stick series over a numerical x-axis
type NumericalCandleStickSeries struct {
	candleStickSeries
}

// DeleteDataInRange deletes all candles with a nEnd greater than min and a nStart smaller than max
// The return value gives the number of candles that have been removed
func (ncs NumericalCandleStickSeries) DeleteDataInRange(min float64, max float64) (c int) {
	if ncs.ser == nil {
		return
	}
	c = ncs.ser.DeleteNumericalDataInRange(min, max)
	return
}

// AddData adds candles to the series.
// The method does not check for duplicates (i.e. candles with same XStart or XEnd)
// The range of XStart, XEnd and values is not restricted
func (ncs NumericalCandleStickSeries) AddData(input []data.NumericalCandleStick) (err error) {
	if ncs.ser == nil {
		return
	}
	err = ncs.ser.AddNumericalData(input)
	if err != nil {
		return
	}
	return
}

// TemporalCandleStickSeries represents a candle stick series over a temporal t-axis
type TemporalCandleStickSeries struct {
	candleStickSeries
}

// DeleteDataInRange deletes all candles with a tEnd after min and a tStart before max.
// The return value gives the number of candles that have been removed
func (tcs TemporalCandleStickSeries) DeleteDataInRange(min time.Time, max time.Time) (c int) {
	if tcs.ser == nil {
		return
	}
	c = tcs.ser.DeleteTemporalDataInRange(min, max)
	return
}

// AddData adds candles to the series.
// The method does not check for duplicates (i.e. candles with same TStart or TEnd)
// The range of TStart, TEnd and values is not restricted
func (tcs TemporalCandleStickSeries) AddData(input []data.TemporalCandleStick) (err error) {
	if tcs.ser == nil {
		return
	}
	err = tcs.ser.AddTemporalData(input)
	if err != nil {
		return
	}
	return
}
