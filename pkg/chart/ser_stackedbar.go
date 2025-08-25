package chart

import (
	"time"

	"github.com/s-daehling/fyne-charts/internal/series"

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

// NumericalStackedBarSeries represents a stacked bar series over a numerical x-axis
type NumericalStackedBarSeries struct {
	stackedBarSeries
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
// An error is returned if min>max
func (nss NumericalStackedBarSeries) DeleteDataInRange(min float64, max float64) (c int, err error) {
	if nss.ser == nil {
		return
	}
	c, err = nss.ser.DeleteNumericalDataInRange(min, max)
	return
}

// AddData adds data points to the stacked series.
// If the single series exists, the data points will be added to it
// If the single series does not exist, nothing is done
// The method does not check for duplicates (i.e. data points with same X).
// The range of X is not restricted. The range of Val is restricted to Val>=0
func (nss NumericalStackedBarSeries) AddData(series string, input []data.NumericalDataPoint) (err error) {
	if nss.ser == nil {
		return
	}
	err = nss.ser.AddNumericalData(series, input)
	return
}

// AddSeries adds a new single series to the stacked bar series.
// If the single series already exists, nothing will be done.
// The method does not check for duplicates (i.e. data points with same X).
// The range of X is not restricted. The range of Val is restricted to Val>=0
func (nss NumericalStackedBarSeries) AddSeries(series data.NumericalDataSeries) (err error) {
	if nss.ser == nil {
		return
	}
	err = nss.ser.AddNumericalSeries(series)
	return
}

// SetWidth sets the width of the bars. The bars are centered around their X value of the data points
// An error is returned in w < 0
func (nss NumericalStackedBarSeries) SetWidth(w float64) (err error) {
	if nss.ser == nil {
		return
	}
	nss.ser.SetNumericalWidthAndOffset(w, 0)
	return
}

// TemporalStackedBarSeries represents a stacked bar series over a temporal t-axis
type TemporalStackedBarSeries struct {
	stackedBarSeries
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
// An error is returned if min after max
func (tss TemporalStackedBarSeries) DeleteDataInRange(min time.Time, max time.Time) (c int, err error) {
	if tss.ser == nil {
		return
	}
	c, err = tss.ser.DeleteTemporalDataInRange(min, max)
	return
}

// AddData adds data points to the stacked series.
// If the single series exists, the data points will be added to it
// If the single series does not exist, nothing is done
// The method does not check for duplicates (i.e. data points with same T).
// The range of T is not restricted. The range of Val is restricted to Val>=0
func (tss TemporalStackedBarSeries) AddData(series string, input []data.TemporalDataPoint) (err error) {
	if tss.ser == nil {
		return
	}
	err = tss.ser.AddTemporalData(series, input)
	return
}

// AddSeries adds a new single series to the stacked bar series.
// If the single series already exists, nothing will be done.
// The method does not check for duplicates (i.e. data points with same T).
// The range of T is not restricted. The range of Val is restricted to Val>=0
func (tss TemporalStackedBarSeries) AddSeries(series data.TemporalDataSeries) (err error) {
	if tss.ser == nil {
		return
	}
	err = tss.ser.AddTemporalSeries(series)
	return
}

// SetWidth sets the width of the bars. The bars are centered around their X value of the data points
// An error is returned in w < 0
func (tss TemporalStackedBarSeries) SetWidth(w time.Duration) (err error) {
	if tss.ser == nil {
		return
	}
	tss.ser.SetTemporalWidthAndOffset(w, 0)
	return
}

// CategoricalStackedBarSeries represents a stacked bar series over a categorical c-axis
type CategoricalStackedBarSeries struct {
	stackedBarSeries
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
// An error iis returned if cat is empty
func (css CategoricalStackedBarSeries) DeleteDataInRange(cat []string) (c int, err error) {
	if css.ser == nil {
		return
	}
	c, err = css.ser.DeleteCategoricalDataInRange(cat)
	return
}

// AddData adds data points to the stacked series.
// If the single series exists, the data points will be added to it
// If the single series does not exist, nothing is done
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0
func (css CategoricalStackedBarSeries) AddData(series string, input []data.CategoricalDataPoint) (err error) {
	if css.ser == nil {
		return
	}
	err = css.ser.AddCategoricalData(series, input)
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
	return
}

// AngularStackedBarSeries represents a stacked bar series over an angular phi-axis
type AngularStackedBarSeries struct {
	stackedBarSeries
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
// An error is returned if min>max
func (ass AngularStackedBarSeries) DeleteDataInRange(min float64, max float64) (c int, err error) {
	if ass.ser == nil {
		return
	}
	c, err = ass.ser.DeleteAngularDataInRange(min, max)
	return
}

// AddData adds data points to the stacked series.
// If the single series exists, the data points will be added to it
// If the single series does not exist, nothing is done
// The method does not check for duplicates (i.e. data points with same A).
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
func (ass AngularStackedBarSeries) AddData(series string, input []data.AngularDataPoint) (err error) {
	if ass.ser == nil {
		return
	}
	err = ass.ser.AddAngularData(series, input)
	return
}

// AddSeries adds a new single series to the stacked bar series.
// If the single series already exists, nothing will be done.
// The method does not check for duplicates (i.e. data points with same A).
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
func (ass AngularStackedBarSeries) AddSeries(series data.AngularDataSeries) (err error) {
	if ass.ser == nil {
		return
	}
	err = ass.ser.AddAngularSeries(series)
	return
}

// SetWidth sets the width of the bars. The bars are centered around their X value of the data points
// An error is returned if w < 0 or w > 2pi
func (ass AngularStackedBarSeries) SetWidth(w float64) (err error) {
	if ass.ser == nil {
		return
	}
	ass.ser.SetNumericalWidthAndOffset(w, 0)
	return
}
