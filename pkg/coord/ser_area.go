package coord

import (
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/coord/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

type areaSeries struct {
	ser *series.DataPointSeries
}

// Name returns the name of the series
func (as areaSeries) Name() (n string) {
	if as.ser == nil {
		return
	}
	n = as.ser.Name()
	return
}

// Show makes the elements of the series visible
func (as areaSeries) Show() {
	if as.ser == nil {
		return
	}
	as.ser.Show()
}

// Hide makes the elements of the series invisible
func (as areaSeries) Hide() {
	if as.ser == nil {
		return
	}
	as.ser.Hide()
}

// SetColor changes the color of series elements
func (as areaSeries) SetColor(col color.Color) {
	if as.ser == nil {
		return
	}
	as.ser.SetColor(col)
}

// SetLineWidth sets the width of the line
func (as areaSeries) SetLineWidth(lw float32) {
	if as.ser == nil {
		return
	}
	as.ser.SetLineWidth(lw)
}

// SetDotSize sets the size of the dots at series data points
func (as areaSeries) SetDotSize(ds float32) {
	if as.ser == nil {
		return
	}
	as.ser.SetDotSize(ds)
}

// Clear deletes all data
func (as areaSeries) Clear() (err error) {
	if as.ser == nil {
		return
	}
	err = as.ser.Clear()
	return
}

// NumericalAreaSeries represents an area series over a numerical x-axis
type NumericalAreaSeries struct {
	areaSeries
}

// DeleteDataInRange deletes all data points with a x-coordinate greater than min and smaller than max
// The return value gives the number of data points that have been removed
// An error is returned if min>max
func (nas NumericalAreaSeries) DeleteDataInRange(min float64, max float64) (c int, err error) {
	if nas.ser == nil {
		return
	}
	c, err = nas.ser.DeleteNumericalDataInRange(min, max)
	if err != nil {
		return
	}
	return
}

// AddData adds data points to the series.
// data does not need to be sorted. It will be sorted by X by the method.
// The method does not check for duplicates (i.e. data points with same X)
// The range of X and Val is not restricted
func (nas NumericalAreaSeries) AddData(input []data.NumericalDataPoint) (err error) {
	if nas.ser == nil {
		return
	}
	err = nas.ser.AddNumericalData(input)
	if err != nil {
		return
	}
	return
}

// TemporalAreaSeries represents an area series over a temporal t-axis
type TemporalAreaSeries struct {
	areaSeries
}

// DeleteDataInRange deletes all data points with a t-coordinate after min and before max.
// The return value gives the number of data points that have been removed
// An error is returned if min after max
func (tas TemporalAreaSeries) DeleteDataInRange(min time.Time, max time.Time) (c int, err error) {
	if tas.ser == nil {
		return
	}
	c, err = tas.ser.DeleteTemporalDataInRange(min, max)
	if err != nil {
		return
	}
	return
}

// AddData adds data points to the series.
// data does not need to be sorted. It will be sorted by T by the method.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T is not restricted. The range of Val is not restricted in a cartesian chart, but Val>=0 in a polar chart
func (tas TemporalAreaSeries) AddData(input []data.TemporalDataPoint) (err error) {
	if tas.ser == nil {
		return
	}
	err = tas.ser.AddTemporalData(input)
	if err != nil {
		return
	}
	return
}
