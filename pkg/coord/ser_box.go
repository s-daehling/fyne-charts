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
func (bs *boxSeries) Name() (n string) {
	if bs.ser == nil {
		return
	}
	n = bs.ser.Name()
	return
}

// Show makes the elements of the series visible
func (bs *boxSeries) Show() {
	if bs.ser == nil {
		return
	}
	bs.ser.Show()
}

// Hide makes the elements of the series invisible
func (bs *boxSeries) Hide() {
	if bs.ser == nil {
		return
	}
	bs.ser.Hide()
}

// SetColor changes the color of series elements
func (bs *boxSeries) SetColor(col color.Color) {
	if bs.ser == nil {
		return
	}
	bs.ser.SetColor(col)
}

// SetLineWidth sets the width of the line
func (bs *boxSeries) SetLineWidth(lw float32) {
	if bs.ser == nil {
		return
	}
	bs.ser.SetLineWidth(lw)
}

// SetOutlierSize sets the size of the dots at outlier points
func (bs *boxSeries) SetOutlierSize(os float32) {
	if bs.ser == nil {
		return
	}
	bs.ser.SetOutlierSize(os)
}

// Clear deletes all data
func (bs *boxSeries) Clear() {
	if bs.ser == nil {
		return
	}
	bs.ser.Clear()
}

// NumericalBoxSeries represents a box series over a numerical x-axis
type NumericalBoxSeries struct {
	boxSeries
}

// NewNumericalBoxSeries creates a new NumericalBoxSeries and populates it with input data
// An error is returned if the input data is invalid
func NewNumericalBoxSeries(name string, col color.Color, input []data.NumericalBox) (nbs *NumericalBoxSeries, err error) {
	nbs = &NumericalBoxSeries{
		boxSeries: boxSeries{
			ser: series.EmptyBoxSeries(name, col),
		},
	}
	err = nbs.AddData(input)
	if err != nil {
		nbs = nil
	}
	return
}

// DeleteDataInRange deletes all data points with a x-coordinate greater than min and smaller than max
// The return value gives the number of data points that have been removed
func (nbs *NumericalBoxSeries) DeleteDataInRange(min float64, max float64) (c int) {
	if nbs.ser == nil {
		return
	}
	c = nbs.ser.DeleteNumericalDataInRange(min, max)
	return
}

// AddData adds data points to the series.
// An error is returned if the input data is invalid
func (nbs *NumericalBoxSeries) AddData(input []data.NumericalBox) (err error) {
	if nbs.ser == nil {
		return
	}
	err = nbs.ser.AddNumericalData(input)
	return
}

// TemporalBoxSeries represents a box series over a temporal t-axis
type TemporalBoxSeries struct {
	boxSeries
}

// NewTemporalBoxSeries creates a new TemporalBoxSeries and populates it with input data
// An error is returned if the input data is invalid
func NewTemporalBoxSeries(name string, col color.Color, input []data.TemporalBox) (tbs *TemporalBoxSeries, err error) {
	tbs = &TemporalBoxSeries{
		boxSeries: boxSeries{
			ser: series.EmptyBoxSeries(name, col),
		},
	}
	err = tbs.AddData(input)
	if err != nil {
		tbs = nil
	}
	return
}

// DeleteDataInRange deletes all data points with a t-coordinate after min and before max.
// The return value gives the number of data points that have been removed
func (tbs *TemporalBoxSeries) DeleteDataInRange(min time.Time, max time.Time) (c int) {
	if tbs.ser == nil {
		return
	}
	c = tbs.ser.DeleteTemporalDataInRange(min, max)
	return
}

// AddData adds data points to the series.
// An error is returned if the input data is invalid
func (tbs *TemporalBoxSeries) AddData(input []data.TemporalBox) (err error) {
	if tbs.ser == nil {
		return
	}
	err = tbs.ser.AddTemporalData(input)
	return
}

// CategoricalBoxSeries represents a box series over a categorical c-axis
type CategoricalBoxSeries struct {
	boxSeries
}

// NewCategoricalBoxSeries creates a new CategoricalBoxSeries and populates it with input data
// The method checks for duplicates (i.e. entries with same C).
// If multiple entries with the same C exist only the first is added to the series
// An error is returned if the input data is invalid
func NewCategoricalBoxSeries(name string, col color.Color, input []data.CategoricalBox) (cbs *CategoricalBoxSeries, err error) {
	cbs = &CategoricalBoxSeries{
		boxSeries: boxSeries{
			ser: series.EmptyBoxSeries(name, col),
		},
	}
	err = cbs.AddData(input)
	if err != nil {
		cbs = nil
	}
	return
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
func (cbs *CategoricalBoxSeries) DeleteDataInRange(cat []string) (c int) {
	if cbs.ser == nil {
		return
	}
	c = cbs.ser.DeleteCategoricalDataInRange(cat)
	return
}

// AddData adds data points to the series.
// The method checks for duplicates (i.e. data points with same C).
// If multiple entries with the same C exist only the first is added to the series
// An error is returned if the input data is invalid
func (cbs *CategoricalBoxSeries) AddData(input []data.CategoricalBox) (err error) {
	if cbs.ser == nil {
		return
	}
	err = cbs.ser.AddCategoricalData(input)
	return
}
