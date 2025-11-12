package coord

import (
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/coord/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

type lollipopSeries struct {
	ser *series.DataPointSeries
}

// Name returns the name of the series
func (ls lollipopSeries) Name() (n string) {
	if ls.ser == nil {
		return
	}
	n = ls.ser.Name()
	return
}

// Show makes the elements of the series visible
func (ls lollipopSeries) Show() {
	if ls.ser == nil {
		return
	}
	ls.ser.Show()
}

// Hide makes the elements of the series invisible
func (ls lollipopSeries) Hide() {
	if ls.ser == nil {
		return
	}
	ls.ser.Hide()
}

// SetColor changes the color of series elements
func (ls lollipopSeries) SetColor(col color.Color) {
	if ls.ser == nil {
		return
	}
	ls.ser.SetColor(col)
}

// SetLineWidth sets the width of the line
func (ls lollipopSeries) SetLineWidth(lw float32) {
	if ls.ser == nil {
		return
	}
	ls.ser.SetLineWidth(lw)
}

// SetDotSize sets the size of the dots at series data points
func (ls lollipopSeries) SetDotSize(ds float32) {
	if ls.ser == nil {
		return
	}
	ls.ser.SetDotSize(ds)
}

// Clear deletes all data
func (ls lollipopSeries) Clear() (err error) {
	if ls.ser == nil {
		return
	}
	err = ls.ser.Clear()
	return
}

// NumericalLollipopSeries represents a lollipop series over a numerical x-axis
type NumericalLollipopSeries struct {
	lollipopSeries
}

// DeleteDataInRange deletes all data points with a x-coordinate greater than min and smaller than max
// The return value gives the number of data points that have been removed
// An error is returned if min>max
func (nls NumericalLollipopSeries) DeleteDataInRange(min float64, max float64) (c int, err error) {
	if nls.ser == nil {
		return
	}
	c, err = nls.ser.DeleteNumericalDataInRange(min, max)
	if err != nil {
		return
	}
	return
}

// AddData adds data points to the series.
// The method does not check for duplicates (i.e. data points with same X)
// The range of X and Val is not restricted
func (nls NumericalLollipopSeries) AddData(input []data.NumericalDataPoint) (err error) {
	if nls.ser == nil {
		return
	}
	err = nls.ser.AddNumericalData(input)
	if err != nil {
		return
	}
	return
}

// TemporalLollipopSeries represents a lollipop series over a temporal t-axis
type TemporalLollipopSeries struct {
	lollipopSeries
}

// DeleteDataInRange deletes all data points with a t-coordinate after min and before max.
// The return value gives the number of data points that have been removed
// An error is returned if min after max
func (tls TemporalLollipopSeries) DeleteDataInRange(min time.Time, max time.Time) (c int, err error) {
	if tls.ser == nil {
		return
	}
	c, err = tls.ser.DeleteTemporalDataInRange(min, max)
	if err != nil {
		return
	}
	return
}

// AddData adds data points to the series.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T is not restricted. The range of Val is not restricted in a cartesian chart, but Val>=0 in a polar chart
func (tls TemporalLollipopSeries) AddData(input []data.TemporalDataPoint) (err error) {
	if tls.ser == nil {
		return
	}
	err = tls.ser.AddTemporalData(input)
	if err != nil {
		return
	}
	return
}

// CategoricalLollipopSeries represents a lollipop series over a categorical c-axis
type CategoricalLollipopSeries struct {
	lollipopSeries
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
// An error is returned if cat is empty
func (cls CategoricalLollipopSeries) DeleteDataInRange(cat []string) (c int, err error) {
	if cls.ser == nil {
		return
	}
	c, err = cls.ser.DeleteCategoricalDataInRange(cat)
	if err != nil {
		return
	}
	return
}

// AddData adds data points to the series.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is not restricted in a cartesian chart, but Val>=0 in a polar chart
func (cls CategoricalLollipopSeries) AddData(input []data.CategoricalDataPoint) (err error) {
	if cls.ser == nil {
		return
	}
	err = cls.ser.AddCategoricalData(input)
	if err != nil {
		return
	}
	return
}
