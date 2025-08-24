package chart

import (
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

type lineSeries struct {
	ser *series.LineSeries
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
func (ls NumericalLineSeries) DeleteDataInRange(min float64, max float64) (c int, err error) {
	if ls.ser == nil {
		return
	}
	c, err = ls.ser.DeleteNumericalDataInRange(min, max)
	return
}

// AddData adds data points to the series.
// data does not need to be sorted. It will be sorted by X by the method.
// The method does not check for duplicates (i.e. data points with same X)
// The range of X and Val is not restricted
func (ls NumericalLineSeries) AddData(input []data.NumericalDataPoint) (err error) {
	if ls.ser == nil {
		return
	}
	err = ls.ser.AddNumericalData(input)
	return
}

// TemporalLineSeries represents a line series over a temporal t-axis
type TemporalLineSeries struct {
	lineSeries
}

// DeleteDataInRange deletes all data points with a t-coordinate after min and before max.
// The return value gives the number of data points that have been removed
// An error is returned if min after max
func (ls TemporalLineSeries) DeleteDataInRange(min time.Time, max time.Time) (c int, err error) {
	if ls.ser == nil {
		return
	}
	c, err = ls.ser.DeleteTemporalDataInRange(min, max)
	return
}

// AddData adds data points to the series.
// data does not need to be sorted. It will be sorted by T by the method.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T is not restricted. The range of Val is not restricted in a cartesian chart, but Val>=0 in a polar chart
func (ls TemporalLineSeries) AddData(input []data.TemporalDataPoint) (err error) {
	if ls.ser == nil {
		return
	}
	err = ls.ser.AddTemporalData(input)
	return
}

// AngularLineSeries represents a line series over an angular phi-axis
type AngularLineSeries struct {
	lineSeries
}

// DeleteDataInRange deletes all data points with an a-coordinate greater than min and smaller than max
// The return value gives the number of data points that have been removed
// An error is returned if min>max
func (als AngularLineSeries) DeleteDataInRange(min float64, max float64) (c int, err error) {
	if als.ser == nil {
		return
	}
	c, err = als.ser.DeleteAngularDataInRange(min, max)
	return
}

// AddData adds data points to the series.
// data does not need to be sorted. It will be sorted by A by the method.
// The method does not check for duplicates (i.e. data points with same A)
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
func (als AngularLineSeries) AddData(input []data.AngularDataPoint) (err error) {
	if als.ser == nil {
		return
	}
	err = als.ser.AddAngularData(input)
	return
}
