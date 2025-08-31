package chart

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/s-daehling/fyne-charts/internal/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

type barSeries struct {
	ser *series.BarSeries
	wid *widget.BaseWidget
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
	if nbs.ser == nil || nbs.wid == nil {
		return
	}
	c, err = nbs.ser.DeleteNumericalDataInRange(min, max)
	if err != nil {
		return
	}
	nbs.wid.Refresh()
	return
}

// AddData adds data points to the series.
// The method does not check for duplicates (i.e. data points with same X)
// The range of X and Val is not restricted
func (nbs NumericalBarSeries) AddData(input []data.NumericalDataPoint) (err error) {
	if nbs.ser == nil || nbs.wid == nil {
		return
	}
	err = nbs.ser.AddNumericalData(input)
	if err != nil {
		return
	}
	nbs.wid.Refresh()
	return
}

// SetBarWidth sets the width of the bars. The bars are centered around their X value of the data points
// An error is returned in w < 0
func (nbs NumericalBarSeries) SetBarWidth(w float64) (err error) {
	if nbs.ser == nil || nbs.wid == nil {
		return
	}
	err = nbs.ser.SetNumericalWidthAndOffset(w, 0)
	if err != nil {
		return
	}
	nbs.wid.Refresh()
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
	if tbs.ser == nil || tbs.wid == nil {
		return
	}
	c, err = tbs.ser.DeleteTemporalDataInRange(min, max)
	if err != nil {
		return
	}
	tbs.wid.Refresh()
	return
}

// AddData adds data points to the series.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T is not restricted. The range of Val is not restricted in a cartesian chart, but Val>=0 in a polar chart
func (tbs TemporalBarSeries) AddData(input []data.TemporalDataPoint) (err error) {
	if tbs.ser == nil || tbs.wid == nil {
		return
	}
	err = tbs.ser.AddTemporalData(input)
	if err != nil {
		return
	}
	tbs.wid.Refresh()
	return
}

// SetBarWidth sets the width of the bars. The bars are centered around their T value of the data points
// An error is returned in w < 0
func (tbs TemporalBarSeries) SetBarWidth(w time.Duration) (err error) {
	if tbs.ser == nil || tbs.wid == nil {
		return
	}
	err = tbs.ser.SetTemporalWidthAndOffset(w, 0)
	if err != nil {
		return
	}
	tbs.wid.Refresh()
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
	if cbs.ser == nil || cbs.wid == nil {
		return
	}
	c, err = cbs.ser.DeleteCategoricalDataInRange(cat)
	if err != nil {
		return
	}
	cbs.wid.Refresh()
	return
}

// AddData adds data points to the series.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is not restricted in a cartesian chart, but Val>=0 in a polar chart
func (cbs CategoricalBarSeries) AddData(input []data.CategoricalDataPoint) (err error) {
	if cbs.ser == nil || cbs.wid == nil {
		return
	}
	err = cbs.ser.AddCategoricalData(input)
	if err != nil {
		return
	}
	cbs.wid.Refresh()
	return
}
