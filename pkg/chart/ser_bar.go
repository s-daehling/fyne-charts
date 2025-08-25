package chart

import (
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

type barSeries struct {
	ser *series.BarSeries
}

// Name returns the name of the series
func (bs barSeries) Name() (n string) {
	if bs.ser == nil {
		return
	}
	n = bs.ser.Name()
	return
}

// Show makes the elements of the series visible
func (bs barSeries) Show() {
	if bs.ser == nil {
		return
	}
	bs.ser.Show()
}

// Hide makes the elements of the series invisible
func (bs barSeries) Hide() {
	if bs.ser == nil {
		return
	}
	bs.ser.Hide()
}

// SetColor changes the color of series elements
func (bs barSeries) SetColor(col color.Color) {
	if bs.ser == nil {
		return
	}
	bs.ser.SetColor(col)
}

// NumericalBarSeries represents a scatter series over a numerical x-axis
type NumericalBarSeries struct {
	barSeries
}

// DeleteDataInRange deletes all data points with a x-coordinate greater than min and smaller than max
// The return value gives the number of data points that have been removed
// An error is returned if min>max
func (nbs NumericalBarSeries) DeleteDataInRange(min float64, max float64) (c int, err error) {
	if nbs.ser == nil {
		return
	}
	c, err = nbs.ser.DeleteNumericalDataInRange(min, max)
	return
}

// AddData adds data points to the series.
// The method does not check for duplicates (i.e. data points with same X)
// The range of X and Val is not restricted
func (nbs NumericalBarSeries) AddData(input []data.NumericalDataPoint) (err error) {
	if nbs.ser == nil {
		return
	}
	err = nbs.ser.AddNumericalData(input)
	return
}

// SetBarWidth sets the width of the bars. The bars are centered around their X value of the data points
// An error is returned in w < 0
func (nbs NumericalBarSeries) SetBarWidth(w float64) (err error) {
	if nbs.ser == nil {
		return
	}
	nbs.ser.SetNumericalWidthAndOffset(w, 0)
	return
}

// TemporalBarSeries represents a scatter series over a temporal t-axis
type TemporalBarSeries struct {
	barSeries
}

// DeleteDataInRange deletes all data points with a t-coordinate after min and before max.
// The return value gives the number of data points that have been removed
// An error is returned if min after max
func (tbs TemporalBarSeries) DeleteDataInRange(min time.Time, max time.Time) (c int, err error) {
	if tbs.ser == nil {
		return
	}
	c, err = tbs.ser.DeleteTemporalDataInRange(min, max)
	return
}

// AddData adds data points to the series.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T is not restricted. The range of Val is not restricted in a cartesian chart, but Val>=0 in a polar chart
func (tbs TemporalBarSeries) AddData(input []data.TemporalDataPoint) (err error) {
	if tbs.ser == nil {
		return
	}
	err = tbs.ser.AddTemporalData(input)
	return
}

// SetBarWidth sets the width of the bars. The bars are centered around their T value of the data points
// An error is returned in w < 0
func (tbs TemporalBarSeries) SetBarWidth(w time.Duration) (err error) {
	if tbs.ser == nil {
		return
	}
	tbs.ser.SetTemporalWidthAndOffset(w, 0)
	return
}

// CategoricalBarSeries represents a bar series over a categorical c-axis
type CategoricalBarSeries struct {
	barSeries
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
// An error is returned if cat is empty
func (cbs CategoricalBarSeries) DeleteDataInRange(cat []string) (c int, err error) {
	if cbs.ser == nil {
		return
	}
	c, err = cbs.ser.DeleteCategoricalDataInRange(cat)
	return
}

// AddData adds data points to the series.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is not restricted in a cartesian chart, but Val>=0 in a polar chart
func (cbs CategoricalBarSeries) AddData(input []data.CategoricalDataPoint) (err error) {
	if cbs.ser == nil {
		return
	}
	err = cbs.ser.AddCategoricalData(input)
	return
}

// AngularBarSeries is a scatter series in a chart with angular phi-axis
type AngularBarSeries struct {
	barSeries
}

// DeleteDataInRange deletes all data points with an a-coordinate greater than min and smaller than max
// The return value gives the number of data points that have been removed
// An error is returned if min>max
func (abs AngularBarSeries) DeleteDataInRange(min float64, max float64) (c int, err error) {
	if abs.ser == nil {
		return
	}
	c, err = abs.ser.DeleteAngularDataInRange(min, max)
	return
}

// AddData adds data points to the series.
// The method does not check for duplicates (i.e. data points with same A)
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
func (abs AngularBarSeries) AddData(input []data.AngularDataPoint) (err error) {
	if abs.ser == nil {
		return
	}
	err = abs.ser.AddAngularData(input)
	return
}

// SetBarWidth sets the width of the bars. The bars are centered around their X value of the data points
// An error is returned if w < 0 or w > 2pi
func (abs AngularBarSeries) SetBarWidth(w float64) (err error) {
	if abs.ser == nil {
		return
	}
	abs.ser.SetNumericalWidthAndOffset(w, 0)
	return
}
