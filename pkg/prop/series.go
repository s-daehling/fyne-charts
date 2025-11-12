package prop

import (
	"image/color"

	"github.com/s-daehling/fyne-charts/internal/prop"
	"github.com/s-daehling/fyne-charts/pkg/data"
)

// ProportionalSeries represents a proportional series over a proportional axis
type ProportionalSeries struct {
	ser *prop.Series
}

// Name returns the name of the series
func (ps ProportionalSeries) Name() (n string) {
	if ps.ser == nil {
		return
	}
	n = ps.ser.Name()
	return
}

// SetValTextColor changes the color of value labels
func (ps *ProportionalSeries) SetValTextColor(col color.Color) {
	ps.ser.SetValTextColor(col)
}

// AutoValTextColor sets the color of value labels to the default (theme.ColorNameForeground)
func (ps *ProportionalSeries) SetAutoValTextColor() {
	ps.ser.SetAutoValTextColor()
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

// Clear deletes all data
func (ps ProportionalSeries) Clear() (err error) {
	if ps.ser == nil {
		return
	}
	err = ps.ser.Clear()
	return
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
func (ps ProportionalSeries) AddData(input []data.ProportionalPoint) (err error) {
	if ps.ser == nil {
		return
	}
	err = ps.ser.AddData(input)
	if err != nil {
		return
	}
	return
}
