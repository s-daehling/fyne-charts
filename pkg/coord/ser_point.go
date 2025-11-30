package coord

import (
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/coord/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

type pointSeries struct {
	ser *series.PointSeries
}

// Name returns the name of the series
func (ps *pointSeries) Name() (n string) {
	if ps.ser == nil {
		return
	}
	n = ps.ser.Name()
	return
}

// Show makes the elements of the series visible
func (ps *pointSeries) Show() {
	if ps.ser == nil {
		return
	}
	ps.ser.Show()
}

// Hide makes the elements of the series invisible
func (ps *pointSeries) Hide() {
	if ps.ser == nil {
		return
	}
	ps.ser.Hide()
}

// SetColor changes the color of series elements
func (ps *pointSeries) SetColor(col color.Color) {
	if ps.ser == nil {
		return
	}
	ps.ser.SetColor(col)
}

// SetLineWidth sets the width of the line
func (ps *pointSeries) SetLineWidth(lw float32) {
	if ps.ser == nil {
		return
	}
	ps.ser.SetLineWidth(lw)
}

// SetDotSize sets the size of the dots at series data points
func (ps *pointSeries) SetDotSize(ds float32) {
	if ps.ser == nil {
		return
	}
	ps.ser.SetDotSize(ds)
}

// Clear deletes all data
func (ps *pointSeries) Clear() {
	if ps.ser == nil {
		return
	}
	ps.ser.Clear()
}

// NumericalPointSeries represents an area series over a numerical x-axis
type NumericalPointSeries struct {
	pointSeries
}

// NewNumericalPointSeries creates a new NumericalPointSeries and populates it with input data
// A series can only be added to a polar chart is Val >= 0 for all points
// In a polar chart only points with 0 <= N <= 2pi are displayed
// An error is returned if the input data is invalid
func NewNumericalPointSeries(name string, col color.Color, input []data.NumericalPoint) (nps *NumericalPointSeries, err error) {
	nps = &NumericalPointSeries{
		pointSeries: pointSeries{
			ser: series.EmptyPointSeries(name, col),
		},
	}
	err = nps.AddData(input)
	if err != nil {
		nps = nil
	}
	return
}

// DeleteDataInRange deletes all data points with a x-coordinate greater than min and smaller than max
// The return value gives the number of data points that have been removed
func (nps *NumericalPointSeries) DeleteDataInRange(min float64, max float64) (c int) {
	if nps.ser == nil {
		return
	}
	c = nps.ser.DeleteNumericalDataInRange(min, max)
	return
}

// AddData adds data points to the series.
// If the series has been added to a polar chart only points with Val >= 0 are allowed
// In a polar chart only points with 0 <= N <= 2pi are displayed
// An error is returned if the input data is invalid
func (nps *NumericalPointSeries) AddData(input []data.NumericalPoint) (err error) {
	if nps.ser == nil {
		return
	}
	err = nps.ser.AddNumericalData(input)
	return
}

// SetBarWidth sets the width of the bars. The bars are centered around the N value of the data points
// An error is returned in w < 0
// only effective if series is displayed as bar series
func (nps *NumericalPointSeries) SetBarWidth(w float64) (err error) {
	if nps.ser == nil {
		return
	}
	err = nps.ser.SetNumericalBarWidth(w)
	if err != nil {
		return
	}
	return
}

// TemporalPointSeries represents an area series over a temporal t-axis
type TemporalPointSeries struct {
	pointSeries
}

// NewTemporalPointSeries creates a new TemporalPointSeries and populates it with input data
// A series can only be added to a polar chart is Val >= 0 for all points
// An error is returned if the input data is invalid
func NewTemporalPointSeries(name string, col color.Color, input []data.TemporalPoint) (tps *TemporalPointSeries, err error) {
	tps = &TemporalPointSeries{
		pointSeries: pointSeries{
			ser: series.EmptyPointSeries(name, col),
		},
	}
	err = tps.AddData(input)
	if err != nil {
		tps = nil
	}
	return
}

// DeleteDataInRange deletes all data points with a t-coordinate after min and before max.
// The return value gives the number of data points that have been removed
func (tps *TemporalPointSeries) DeleteDataInRange(min time.Time, max time.Time) (c int) {
	if tps.ser == nil {
		return
	}
	c = tps.ser.DeleteTemporalDataInRange(min, max)
	return
}

// AddData adds data points to the series.
// If the series has been added to a polar chart only points with Val >= 0 are allowed
// An error is returned if the input data is invalid
func (tps *TemporalPointSeries) AddData(input []data.TemporalPoint) (err error) {
	if tps.ser == nil {
		return
	}
	err = tps.ser.AddTemporalData(input)
	return
}

// SetBarWidth sets the width of the bars. The bars are centered around the T value of the data points
// only effective if series is displayed as bar series
func (tps *TemporalPointSeries) SetBarWidth(w time.Duration) (err error) {
	if tps.ser == nil {
		return
	}
	err = tps.ser.SetTemporalBarWidth(w)
	if err != nil {
		return
	}
	return
}

// CategoricalPointSeries represents a bar series over a categorical c-axis
type CategoricalPointSeries struct {
	pointSeries
}

// NewCategoricalPointSeries creates a new CategoricalPointSeries and populates it with input data
// The method checks for duplicates (i.e. data points with same C).
// If multiple entries with the same C exist only the first is added to the series
// A series can only be added to a polar chart is Val >= 0 for all points
// An error is returned if the input data is invalid
func NewCategoricalPointSeries(name string, col color.Color, input []data.CategoricalPoint) (cps *CategoricalPointSeries, err error) {
	cps = &CategoricalPointSeries{
		pointSeries: pointSeries{
			ser: series.EmptyPointSeries(name, col),
		},
	}
	err = cps.AddData(input)
	if err != nil {
		cps = nil
	}
	return
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
func (cps *CategoricalPointSeries) DeleteDataInRange(cat []string) (c int) {
	if cps.ser == nil {
		return
	}
	c = cps.ser.DeleteCategoricalDataInRange(cat)
	return
}

// AddData adds data points to the series.
// The method checks for duplicates (i.e. data points with same C).
// If multiple entries with the same C exist only the first is added to the series
// If the series has been added to a polar chart only points with Val >= 0 are allowed
// An error is returned if the input data is invalid
func (cps *CategoricalPointSeries) AddData(input []data.CategoricalPoint) (err error) {
	if cps.ser == nil {
		return
	}
	err = cps.ser.AddCategoricalData(input)
	return
}
