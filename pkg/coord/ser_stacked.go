package coord

import (
	"github.com/s-daehling/fyne-charts/internal/coord/series"
)

type stackedSeries struct {
	ser *series.StackedSeries
}

// Name returns the name of the series
func (ss stackedSeries) Name() (n string) {
	if ss.ser == nil {
		return
	}
	n = ss.ser.Name()
	return
}

// Show makes the elements of the series visible
func (ss stackedSeries) Show() {
	if ss.ser == nil {
		return
	}
	ss.ser.Show()
}

// Hide makes the elements of the series invisible
func (ss stackedSeries) Hide() {
	if ss.ser == nil {
		return
	}
	ss.ser.Hide()
}

// Clear deletes all data
func (ss stackedSeries) Clear() {
	if ss.ser == nil {
		return
	}
	ss.ser.Clear()
}

// CategoricalStackedSeries represents a stacked bar series over a categorical c-axis
type CategoricalStackedSeries struct {
	stackedSeries
}

func NewCategoricalStackedSeries(name string) (cps CategoricalStackedSeries) {
	cps.ser = series.EmptyStackedSeries(name)
	return
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
func (css CategoricalStackedSeries) DeleteDataInRange(cat []string) (c int) {
	if css.ser == nil {
		return
	}
	c = css.ser.DeleteCategoricalDataInRange(cat)
	return
}

func (css CategoricalStackedSeries) RemoveSeries(name string) {
	if css.ser == nil {
		return
	}
	css.ser.RemovePointSeries(name)
}

func (css CategoricalStackedSeries) AddSeries(cps CategoricalPointSeries) (err error) {
	if css.ser == nil {
		return
	}
	err = css.ser.AddPointSeries(cps.ser)
	return
}
