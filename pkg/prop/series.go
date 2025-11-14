package prop

import (
	"image/color"

	"github.com/s-daehling/fyne-charts/internal/prop"
	"github.com/s-daehling/fyne-charts/pkg/data"
)

// Series represents a proportional series over a proportional axis
type Series struct {
	ser *prop.Series
}

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

// SetValTextColor changes the color of value labels
func (ps *Series) SetValTextColor(col color.Color) {
	ps.ser.SetValTextColor(col)
}

// AutoValTextColor sets the color of value labels to the default (theme.ColorNameForeground)
func (ps *Series) SetAutoValTextColor() {
	ps.ser.SetAutoValTextColor()
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
// The range of C is not restricted. The range of Val is restricted to Val>=0
func (ps *Series) AddData(input []data.ProportionalPoint) (err error) {
	if ps.ser == nil {
		return
	}
	err = ps.ser.AddData(input)
	return
}
