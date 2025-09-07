package chart

import (
	"github.com/s-daehling/fyne-charts/internal/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

// ProportionalSeries represents a proportional series over a proportional axis
type ProportionalSeries struct {
	ser *series.ProportionalSeries
}

// Name returns the name of the series
func (ps ProportionalSeries) Name() (n string) {
	if ps.ser == nil {
		return
	}
	n = ps.ser.Name()
	return
}

// Show makes the elements of the series visible
func (ps ProportionalSeries) Show() {
	if ps.ser == nil {
		return
	}
	ps.ser.Show()
}

// Hide makes the elements of the series invisible
func (ps ProportionalSeries) Hide() {
	if ps.ser == nil {
		return
	}
	ps.ser.Hide()
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
// An error is returned if cat is empty
func (ps ProportionalSeries) DeleteDataInRange(cat []string) (c int, err error) {
	if ps.ser == nil {
		return
	}
	c, err = ps.ser.DeleteDataInRange(cat)
	if err != nil {
		return
	}
	return
}

// AddData adds data points to the series.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0
func (ps ProportionalSeries) AddData(input []data.ProportionalDataPoint) (err error) {
	if ps.ser == nil {
		return
	}
	err = ps.ser.AddData(input)
	if err != nil {
		return
	}
	return
}
