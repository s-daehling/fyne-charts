package coord

import (
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/coord/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

type scatterSeries struct {
	ser *series.ScatterSeries
}

// Name returns the name of the series
func (ss scatterSeries) Name() (n string) {
	if ss.ser == nil {
		return
	}
	n = ss.ser.Name()
	return
}

// Show makes the elements of the series visible
func (ss scatterSeries) Show() {
	if ss.ser == nil {
		return
	}
	ss.ser.Show()
}

// Hide makes the elements of the series invisible
func (ss scatterSeries) Hide() {
	if ss.ser == nil {
		return
	}
	ss.ser.Hide()
}

// SetColor changes the color of series elements
func (ss scatterSeries) SetColor(col color.Color) {
	if ss.ser == nil {
		return
	}
	ss.ser.SetColor(col)
}

// SetDotSize sets the size of the dots at series data points
func (ss scatterSeries) SetDotSize(ds float32) {
	if ss.ser == nil {
		return
	}
	ss.ser.SetDotSize(ds)
}

// Clear deletes all data
func (ss scatterSeries) Clear() (err error) {
	if ss.ser == nil {
		return
	}
	err = ss.ser.Clear()
	return
}

// NumericalScatterSeries represents a scatter series over a numerical x-axis
type NumericalScatterSeries struct {
	scatterSeries
}

// DeleteDataInRange deletes all data points with a x-coordinate greater than min and smaller than max
// The return value gives the number of data points that have been removed
// An error is returned if min>max
func (nss NumericalScatterSeries) DeleteDataInRange(min float64, max float64) (c int, err error) {
	if nss.ser == nil {
		return
	}
	c, err = nss.ser.DeleteNumericalDataInRange(min, max)
	if err != nil {
		return
	}
	return
}

// AddData adds data points to the series.
// The method does not check for duplicates (i.e. data points with same X)
// The range of X and Val is not restricted
func (nss NumericalScatterSeries) AddData(input []data.NumericalDataPoint) (err error) {
	if nss.ser == nil {
		return
	}
	err = nss.ser.AddNumericalData(input)
	if err != nil {
		return
	}
	return
}

// TemporalScatterSeries represents a scatter series over a temporal t-axis
type TemporalScatterSeries struct {
	scatterSeries
}

// DeleteDataInRange deletes all data points with a t-coordinate after min and before max.
// The return value gives the number of data points that have been removed
// An error is returned if min after max
func (tss TemporalScatterSeries) DeleteDataInRange(min time.Time, max time.Time) (c int, err error) {
	if tss.ser == nil {
		return
	}
	c, err = tss.ser.DeleteTemporalDataInRange(min, max)
	if err != nil {
		return
	}
	return
}

// AddData adds data points to the series.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T is not restricted. The range of Val is not restricted in a cartesian chart, but Val>=0 in a polar chart
func (tss TemporalScatterSeries) AddData(input []data.TemporalDataPoint) (err error) {
	if tss.ser == nil {
		return
	}
	err = tss.ser.AddTemporalData(input)
	if err != nil {
		return
	}
	return
}

// CategoricalScatterSeries represents a scatter series over a categorical c-axis
type CategoricalScatterSeries struct {
	scatterSeries
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
// An error is returned if cat is empty
func (css CategoricalScatterSeries) DeleteDataInRange(cat []string) (c int, err error) {
	if css.ser == nil {
		return
	}
	c, err = css.ser.DeleteCategoricalDataInRange(cat)
	if err != nil {
		return
	}
	return
}

// AddData adds data points to the series.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is not restricted in a cartesian chart, but Val>=0 in a polar chart
func (css CategoricalScatterSeries) AddData(input []data.CategoricalDataPoint) (err error) {
	if css.ser == nil {
		return
	}
	err = css.ser.AddCategoricalData(input)
	if err != nil {
		return
	}
	return
}
