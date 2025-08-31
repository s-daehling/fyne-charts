package chart

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/s-daehling/fyne-charts/internal/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

type lineSeries struct {
	ser *series.LineSeries
	wid *widget.BaseWidget
}

// Name returns the name of the series
func (ls lineSeries) Name() (n string) {
	if ls.ser == nil {
		return
	}
	n = ls.ser.Name()
	return
}

// Show makes the elements of the series visible
func (ls lineSeries) Show() {
	if ls.ser == nil {
		return
	}
	ls.ser.Show()
}

// Hide makes the elements of the series invisible
func (ls lineSeries) Hide() {
	if ls.ser == nil {
		return
	}
	ls.ser.Hide()
}

// SetColor changes the color of series elements
func (ls lineSeries) SetColor(col color.Color) {
	if ls.ser == nil {
		return
	}
	ls.ser.SetColor(col)
}

// SetLineWidth sets the width of the line
func (ls lineSeries) SetLineWidth(lw float32) {
	if ls.ser == nil {
		return
	}
	ls.ser.SetLineWidth(lw)
}

// SetDotSize sets the size of the dots at series data points
func (ls lineSeries) SetDotSize(ds float32) {
	if ls.ser == nil {
		return
	}
	ls.ser.SetDotSize(ds)
}

// NumericalLineSeries represents a line series over a numerical x-axis
type NumericalLineSeries struct {
	lineSeries
}

// DeleteDataInRange deletes all data points with a x-coordinate greater than min and smaller than max
// The return value gives the number of data points that have been removed
// An error is returned if min>max
func (nls NumericalLineSeries) DeleteDataInRange(min float64, max float64) (c int, err error) {
	if nls.ser == nil || nls.wid == nil {
		return
	}
	c, err = nls.ser.DeleteNumericalDataInRange(min, max)
	if err != nil {
		return
	}

	nls.wid.Refresh()
	return
}

// AddData adds data points to the series.
// data does not need to be sorted. It will be sorted by X by the method.
// The method does not check for duplicates (i.e. data points with same X)
// The range of X and Val is not restricted
func (nls NumericalLineSeries) AddData(input []data.NumericalDataPoint) (err error) {
	if nls.ser == nil || nls.wid == nil {
		return
	}
	err = nls.ser.AddNumericalData(input)
	if err != nil {
		return
	}
	nls.wid.Refresh()
	return
}

// TemporalLineSeries represents a line series over a temporal t-axis
type TemporalLineSeries struct {
	lineSeries
}

// DeleteDataInRange deletes all data points with a t-coordinate after min and before max.
// The return value gives the number of data points that have been removed
// An error is returned if min after max
func (tls TemporalLineSeries) DeleteDataInRange(min time.Time, max time.Time) (c int, err error) {
	if tls.ser == nil || tls.wid == nil {
		return
	}
	c, err = tls.ser.DeleteTemporalDataInRange(min, max)
	if err != nil {
		return
	}
	tls.wid.Refresh()
	return
}

// AddData adds data points to the series.
// data does not need to be sorted. It will be sorted by T by the method.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T is not restricted. The range of Val is not restricted in a cartesian chart, but Val>=0 in a polar chart
func (tls TemporalLineSeries) AddData(input []data.TemporalDataPoint) (err error) {
	if tls.ser == nil || tls.wid == nil {
		return
	}
	err = tls.ser.AddTemporalData(input)
	if err != nil {
		return
	}
	tls.wid.Refresh()
	return
}
