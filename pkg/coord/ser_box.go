package coord

import (
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/coord/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

type boxSeries struct {
	ser *series.BoxSeries
}

// Name returns the name of the series
func (bs boxSeries) Name() (n string) {
	if bs.ser == nil {
		return
	}
	n = bs.ser.Name()
	return
}

// Show makes the elements of the series visible
func (bs boxSeries) Show() {
	if bs.ser == nil {
		return
	}
	bs.ser.Show()
}

// Hide makes the elements of the series invisible
func (bs boxSeries) Hide() {
	if bs.ser == nil {
		return
	}
	bs.ser.Hide()
}

// SetColor changes the color of series elements
func (bs boxSeries) SetColor(col color.Color) {
	if bs.ser == nil {
		return
	}
	bs.ser.SetColor(col)
}

// SetLineWidth sets the width of the line
func (bs boxSeries) SetLineWidth(lw float32) {
	if bs.ser == nil {
		return
	}
	bs.ser.SetLineWidth(lw)
}

// SetOutlierSize sets the size of the dots at outlier points
func (bs boxSeries) SetOutlierSize(os float32) {
	if bs.ser == nil {
		return
	}
	bs.ser.SetOutlierSize(os)
}

// Clear deletes all data
func (bs boxSeries) Clear() (err error) {
	if bs.ser == nil {
		return
	}
	err = bs.ser.Clear()
	return
}

// NumericalBoxSeries represents a box series over a numerical x-axis
type NumericalBoxSeries struct {
	boxSeries
}

// DeleteDataInRange deletes all data points with a x-coordinate greater than min and smaller than max
// The return value gives the number of data points that have been removed
// An error is returned if min>max
func (nbs NumericalBoxSeries) DeleteDataInRange(min float64, max float64) (c int, err error) {
	if nbs.ser == nil {
		return
	}
	c, err = nbs.ser.DeleteNumericalDataInRange(min, max)
	if err != nil {
		return
	}
	return
}

// AddData adds data points to the series.
// The method does not check for duplicates (i.e. data points with same X)
// The range of X and Val is not restricted
func (nbs NumericalBoxSeries) AddData(input []data.NumericalBox) (err error) {
	if nbs.ser == nil {
		return
	}
	err = nbs.ser.AddNumericalData(input)
	if err != nil {
		return
	}
	return
}

// TemporalBoxSeries represents a box series over a temporal t-axis
type TemporalBoxSeries struct {
	boxSeries
}

// DeleteDataInRange deletes all data points with a t-coordinate after min and before max.
// The return value gives the number of data points that have been removed
// An error is returned if min after max
func (tbs TemporalBoxSeries) DeleteDataInRange(min time.Time, max time.Time) (c int, err error) {
	if tbs.ser == nil {
		return
	}
	c, err = tbs.ser.DeleteTemporalDataInRange(min, max)
	if err != nil {
		return
	}
	return
}

// AddData adds data points to the series.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and values is not restricted
func (tbs TemporalBoxSeries) AddData(input []data.TemporalBox) (err error) {
	if tbs.ser == nil {
		return
	}
	err = tbs.ser.AddTemporalData(input)
	if err != nil {
		return
	}
	return
}

// CategoricalBoxSeries represents a box series over a categorical c-axis
type CategoricalBoxSeries struct {
	boxSeries
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
// An error is returned if cat is empty
func (cbs CategoricalBoxSeries) DeleteDataInRange(cat []string) (c int, err error) {
	if cbs.ser == nil {
		return
	}
	c, err = cbs.ser.DeleteCategoricalDataInRange(cat)
	if err != nil {
		return
	}
	return
}

// AddData adds data points to the series.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C and values is not restricted.
func (cbs CategoricalBoxSeries) AddData(input []data.CategoricalBox) (err error) {
	if cbs.ser == nil {
		return
	}
	err = cbs.ser.AddCategoricalData(input)
	if err != nil {
		return
	}
	return
}
