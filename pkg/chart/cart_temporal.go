package chart

import (
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/chart"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// CartesianTemporalChart implements a cartesian plane with a temporal t-axis and a numerical y-axis
type CartesianTemporalChart struct {
	base *chart.BaseChart
	widget.BaseWidget
}

// NewCartesianTemporalChart returns an initialized CartesianTemporalChart
func NewCartesianTemporalChart() (tempChart *CartesianTemporalChart) {
	tempChart = &CartesianTemporalChart{
		base: chart.EmptyBaseChart(chart.CartesianPlane, chart.Temporal),
	}
	tempChart.ExtendBaseWidget(tempChart)
	return
}

// CreateRenderer creates the renderer of the widget
func (tempChart *CartesianTemporalChart) CreateRenderer() fyne.WidgetRenderer {
	r := tempChart.base.GetRenderer()
	return r
}

// AddLineSeries adds a series of data which is visualized as line chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// data does not need to be sorted. It will be sorted by T by the method.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and Val is not restricted
func (tempChart *CartesianTemporalChart) AddLineSeries(name string, points []data.TemporalDataPoint,
	showDots bool, color color.Color) (tls TemporalLineSeries, err error) {
	tls.ser, err = tempChart.base.AddTemporalLineSeries(name, points, showDots, nil, color)
	return
}

// AddLineSeriesWithProvider adds a series of data which is visualized as line chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The series data is retrieved from providerFct
// data does not need to be sorted. It will be sorted by T by the method.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and Val is not restricted
func (tempChart *CartesianTemporalChart) AddLineSeriesWithProvider(name string, providerFct func() []data.TemporalDataPoint,
	showDots bool, color color.Color) (tls TemporalLineSeries, err error) {
	tls.ser, err = tempChart.base.AddTemporalLineSeries(name, nil, showDots, providerFct, color)
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and Val is not restricted
func (tempChart *CartesianTemporalChart) AddScatterSeries(name string, points []data.TemporalDataPoint,
	color color.Color) (tss TemporalScatterSeries, err error) {
	tss.ser, err = tempChart.base.AddTemporalScatterSeries(name, points, nil, color)
	return
}

// AddScatterSeriesWithProvider adds a series of data which is visualized as scatter chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The series data is retrieved from providerFct
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and Val is not restricted
func (tempChart *CartesianTemporalChart) AddScatterSeriesWithProvider(name string, providerFct func() []data.TemporalDataPoint,
	color color.Color) (tss TemporalScatterSeries, err error) {
	tss.ser, err = tempChart.base.AddTemporalScatterSeries(name, nil, providerFct, color)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and Val is not restricted
func (tempChart *CartesianTemporalChart) AddLollipopSeries(name string, points []data.TemporalDataPoint,
	color color.Color) (tls TemporalLollipopSeries, err error) {
	tls.ser, err = tempChart.base.AddTemporalLollipopSeries(name, points, nil, color)
	return
}

// AddLollipopSeriesWithProvider adds a series of data which is visualized as lollipop chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The series data is retrieved from providerFct
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and Val is not restricted
func (tempChart *CartesianTemporalChart) AddLollipopSeriesWithProvider(name string, providerFct func() []data.TemporalDataPoint,
	color color.Color) (tls TemporalLollipopSeries, err error) {
	tls.ser, err = tempChart.base.AddTemporalLollipopSeries(name, nil, providerFct, color)
	return
}

// AddCandleStickSeries adds a series of data which is visualized as canlde stick chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same T)
// The range of TStart, TEnd and values is not restricted
func (tempChart *CartesianTemporalChart) AddCandleStickSeries(name string,
	points []data.TemporalCandleStick) (tcs TemporalCandleStickSeries, err error) {
	tcs.ser, err = tempChart.base.AddTemporalCandleStickSeries(name, points, nil)
	return
}

// AddCandleStickSeriesWithProvider adds a series of data which is visualized as canlde stick chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The series data is retrieved from providerFct
// The method does not check for duplicates (i.e. data points with same T)
// The range of TStart, TEnd and values is not restricted
func (tempChart *CartesianTemporalChart) AddCandleStickSeriesWithProvider(name string,
	providerFct func() []data.TemporalCandleStick) (tcs TemporalCandleStickSeries, err error) {
	tcs.ser, err = tempChart.base.AddTemporalCandleStickSeries(name, nil, providerFct)
	return
}

// AddBoxSeries adds a series of data which is visualized as box chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and values is not restricted
func (tempChart *CartesianTemporalChart) AddBoxSeries(name string,
	points []data.TemporalBox, col color.Color) (tbs TemporalBoxSeries, err error) {
	tbs.ser, err = tempChart.base.AddTemporalBoxSeries(name, points, nil, col)
	return
}

// AddBoxSeriesWithProvider adds a series of data which is visualized as box chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The series data is retrieved from providerFct
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and values is not restricted
func (tempChart *CartesianTemporalChart) AddBoxSeriesWithProvider(name string,
	providerFct func() []data.TemporalBox, col color.Color) (tbs TemporalBoxSeries, err error) {
	tbs.ser, err = tempChart.base.AddTemporalBoxSeries(name, nil, providerFct, col)
	return
}

// AddAreaSeries adds a series of data which is visualized as area chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// data does not need to be sorted. It will be sorted by T by the method.
// The method does not check for duplicates (i.e. data points with same T).
// The range of T and Val is not restricted
func (tempChart *CartesianTemporalChart) AddAreaSeries(name string, points []data.TemporalDataPoint,
	showDots bool, color color.Color) (tas TemporalAreaSeries, err error) {
	tas.ser, err = tempChart.base.AddTemporalAreaSeries(name, points, showDots, nil, color)
	return
}

// AddAreaSeriesWithProvider adds a series of data which is visualized as area chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The series data is retrieved from providerFct
// data does not need to be sorted. It will be sorted by T by the method.
// The method does not check for duplicates (i.e. data points with same T).
// The range of T and Val is not restricted
func (tempChart *CartesianTemporalChart) AddAreaSeriesWithProvider(name string, providerFct func() []data.TemporalDataPoint,
	showDots bool, color color.Color) (tas TemporalAreaSeries, err error) {
	tas.ser, err = tempChart.base.AddTemporalAreaSeries(name, nil, showDots, providerFct, color)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and Val is not restricted
// The bars are centered around their T value of the data points. barWidth is the width of the bars.
// An error is returned if barWidth < 0
func (tempChart *CartesianTemporalChart) AddBarSeries(name string, points []data.TemporalDataPoint,
	barWidth time.Duration, color color.Color) (tbs TemporalBarSeries, err error) {
	tbs.ser, err = tempChart.base.AddTemporalBarSeries(name, points, barWidth, nil, color)
	return
}

// AddBarSeriesWithProvider adds a series of data which is visualized as bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The series data is retrieved from providerFct
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and Val is not restricted
// The bars are centered around their T value of the data points. barWidth is the width of the bars.
// An error is returned if barWidth < 0
func (tempChart *CartesianTemporalChart) AddBarSeriesWithProvider(name string, providerFct func() []data.TemporalDataPoint,
	barWidth time.Duration, color color.Color) (tbs TemporalBarSeries, err error) {
	tbs.ser, err = tempChart.base.AddTemporalBarSeries(name, nil, barWidth, providerFct, color)
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (tempChart *CartesianTemporalChart) DeleteSeries(name string) {
	tempChart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (tempChart *CartesianTemporalChart) SetTitle(l string) {
	tempChart.base.SetTitle(l)
}

// HideLegend hides the legend and uses the full space for the chart
func (tempChart *CartesianTemporalChart) HideLegend() {
	tempChart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (tempChart *CartesianTemporalChart) ShowLegend() {
	tempChart.base.ShowLegend()
}

// SetYAxisLabel sets the label of the y-axis, which will be displayed at the left side
func (tempChart *CartesianTemporalChart) SetYAxisLabel(l string) {
	tempChart.base.SetToAxisLabel(l)
}

// SetYRange sets a user defined range for the y-axis;
// an error is returned if min>max or if the origin has been defined by the user before and is outside the given range
func (tempChart *CartesianTemporalChart) SetYRange(min float64, max float64) (err error) {
	err = tempChart.base.SetToRange(min, max)
	return
}

// SetAutoYRange overrides a previously user defined range and lets the range be calculated automatically
func (tempChart *CartesianTemporalChart) SetAutoYRange() {
	tempChart.base.SetAutoToRange()
}

// SetYTicks sets the list of user defined ticks to be shown on the y-axis
func (tempChart *CartesianTemporalChart) SetYTicks(ts []data.NumericalTick) {
	tempChart.base.SetToTicks(ts)
}

// SetAutoYTicks overrides a previously user defined set of y-axis ticks and lets the ticks be calculated automatically
func (tempChart *CartesianTemporalChart) SetAutoYTicks(autoSupportLine bool) {
	tempChart.base.SetAutoToTicks(autoSupportLine)
}

// SetOrigin sets a user defined origin (crossing of t and y axis).
// An error is returned. if a range has been defined before and at least one coordinate is outside the range.
func (tempChart *CartesianTemporalChart) SetOrigin(t time.Time, y float64) (err error) {
	err = tempChart.base.SetTOrigin(t, y)
	return
}

// SetAutoOrigin resets a previously user defined origin and allows the chart to calculate the ideal origin automatically
func (tempChart *CartesianTemporalChart) SetAutoOrigin() {
	tempChart.base.SetAutoOrigin()
}

// SetTAxisLabel sets the label of the t-axis, which will be displayed at the bottom
func (tempChart *CartesianTemporalChart) SetTAxisLabel(l string) {
	tempChart.base.SetFromAxisLabel(l)
}

// SetTRange sets a user defined range for the t-axis.
// An error is returned, if min after max or if the origin has been defined by the user before and is outside the given range
func (tempChart *CartesianTemporalChart) SetTRange(min time.Time, max time.Time) (err error) {
	err = tempChart.base.SetFromTRange(min, max)
	return
}

// SetAutoTRange overrides a previously user defined range and lets the range be calculated automatically
func (tempChart *CartesianTemporalChart) SetAutoTRange() {
	tempChart.base.SetAutoFromRange()
}

// SetTTicks sets the list of user defined ticks to be shown on the t-axis
func (tempChart *CartesianTemporalChart) SetTTicks(ts []data.TemporalTick, format string) {
	tempChart.base.SetFromTTicks(ts, format)
}

// SetAutoTTicks overrides a previously user defined set of t-axis ticks and lets the ticks be calculated automatically
func (tempChart *CartesianTemporalChart) SetAutoTTicks(autoSupportLine bool) {
	tempChart.base.SetAutoFromTicks(autoSupportLine)
}
