package coord

import (
	"github.com/s-daehling/fyne-charts/internal/coord/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

type stackedBarSeries struct {
	ser *series.StackedBarSeries
}

// Name returns the name of the series
func (ss stackedBarSeries) Name() (n string) {
	if ss.ser == nil {
		return
	}
	n = ss.ser.Name()
	return
}

// Show makes the elements of the series visible
func (ss stackedBarSeries) Show() {
	if ss.ser == nil {
		return
	}
	ss.ser.Show()
}

// Hide makes the elements of the series invisible
func (ss stackedBarSeries) Hide() {
	if ss.ser == nil {
		return
	}
	ss.ser.Hide()
}

// Clear deletes all data
func (ss stackedBarSeries) Clear() {
	if ss.ser == nil {
		return
	}
	ss.ser.Clear()
}

// CategoricalStackedBarSeries represents a stacked bar series over a categorical c-axis
type CategoricalStackedBarSeries struct {
	stackedBarSeries
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
func (css CategoricalStackedBarSeries) DeleteDataInRange(cat []string) (c int) {
	if css.ser == nil {
		return
	}
	c = css.ser.DeleteCategoricalDataInRange(cat)
	return
}

// AddData adds data points to the stacked series.
// If the single series exists, the data points will be added to it
// If the single series does not exist, nothing is done
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0
func (css CategoricalStackedBarSeries) AddData(series string, input []data.CategoricalPoint) (err error) {
	if css.ser == nil {
		return
	}
	err = css.ser.AddCategoricalData(series, input)
	if err != nil {
		return
	}
	return
}

// AddSeries adds a new single series to the stacked bar series.
// If the single series already exists, nothing will be done.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0
func (css CategoricalStackedBarSeries) AddSeries(series data.CategoricalDataSeries) (err error) {
	if css.ser == nil {
		return
	}
	err = css.ser.AddCategoricalSeries(series)
	if err != nil {
		return
	}
	return
}
