package prop

import (
	"github.com/s-daehling/fyne-charts/internal/prop"
	"github.com/s-daehling/fyne-charts/pkg/data"
	"github.com/s-daehling/fyne-charts/pkg/style"
)

// Series represents a proportional series over a proportional axis
type Series struct {
	ser *prop.Series
}

// NewSeries creates a new Series and populates it with input data
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// Val is restricted to Val>=0
// An error is returned if the input data is invalid
func NewSeries(name string, input []data.ProportionalPoint) (ps *Series, err error) {
	ps = &Series{
		ser: prop.EmptyProportionalSeries(name),
	}
	err = ps.AddData(input)
	if err != nil {
		ps = nil
	}
	return
}

// Name returns the name of the series
func (ps *Series) Name() (n string) {
	if ps.ser == nil {
		return
	}
	n = ps.ser.Name()
	return
}

// SetValTextColor changes the style of value labels
func (ps *Series) SetValueTextStyle(textStyle style.ChartTextStyle) {
	ps.ser.SetValTextStyle(textStyle)
}

// Show makes the elements of the series visible
func (ps *Series) Show() {
	if ps.ser == nil {
		return
	}
	ps.ser.Show()
}

// Hide makes the elements of the series invisible
func (ps *Series) Hide() {
	if ps.ser == nil {
		return
	}
	ps.ser.Hide()
}

// Clear deletes all data
func (ps *Series) Clear() {
	if ps.ser == nil {
		return
	}
	ps.ser.Clear()
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
func (ps *Series) DeleteDataInRange(cat []string) (c int) {
	if ps.ser == nil {
		return
	}
	c = ps.ser.DeleteDataInRange(cat)
	return
}

// AddData adds data points to the series.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// Val is restricted to Val>=0
// An error is returned if the input data is invalid
func (ps *Series) AddData(input []data.ProportionalPoint) (err error) {
	if ps.ser == nil {
		return
	}
	err = ps.ser.AddData(input)
	return
}
