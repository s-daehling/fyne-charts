package chart

import (
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/chart"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// PolarTemporalChart implements a polar plane with one temporal t-axis and one numerical r-axis
type PolarTemporalChart struct {
	base *chart.BaseChart
	widget.BaseWidget
}

// NewPolarTemporalChart returns an initialized PolarTemporalChart
func NewPolarTemporalChart() (tempChart *PolarTemporalChart) {
	tempChart = &PolarTemporalChart{
		base: chart.EmptyBaseChart(chart.PolarPlane, chart.Temporal),
	}
	tempChart.ExtendBaseWidget(tempChart)
	return
}

// CreateRenderer creates the renderer of the widget
func (tempChart *PolarTemporalChart) CreateRenderer() fyne.WidgetRenderer {
	r := chart.EmptyPolarRenderer(tempChart.base)
	return r
}

// AddLineSeries adds a series of data which is visualized as line chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// data does not need to be sorted. It will be sorted by T by the method.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T is not restricted. The range of Val is restricted to Val>=0.
func (tempChart *PolarTemporalChart) AddLineSeries(name string, points []data.TemporalDataPoint,
	showDots bool, color color.Color) (tls TemporalLineSeries, err error) {
	tls.ser, err = tempChart.base.AddTemporalLineSeries(name, points, showDots, color)
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T is not restricted. The range of Val is restricted to Val>=0.
func (tempChart *PolarTemporalChart) AddScatterSeries(name string, points []data.TemporalDataPoint,
	color color.Color) (tss TemporalScatterSeries, err error) {
	tss.ser, err = tempChart.base.AddTemporalScatterSeries(name, points, color)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T is not restricted. The range of Val is restricted to Val>=0.
func (tempChart *PolarTemporalChart) AddLollipopSeries(name string, points []data.TemporalDataPoint,
	color color.Color) (tls TemporalLollipopSeries, err error) {
	tls.ser, err = tempChart.base.AddTemporalLollipopSeries(name, points, color)
	return
}

// AddAreaSeries adds a series of data which is visualized as area chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// data does not need to be sorted. It will be sorted by T by the method.
// The method does not check for duplicates (i.e. data points with same T).
// The range of T is not restricted. The range of Val is restricted to Val>=0.
func (tempChart *PolarTemporalChart) AddAreaSeries(name string, points []data.TemporalDataPoint, showDots bool,
	color color.Color) (tas TemporalAreaSeries, err error) {
	tas.ser, err = tempChart.base.AddTemporalAreaSeries(name, points, showDots, color)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T is not restricted. The range of Val is restricted to Val>=0.
// The bars are centered around their T value of the data points. barWidth is the width of the bars.
// An error is returned if barWidth < 0
func (numChart *PolarTemporalChart) AddBarSeries(name string, points []data.TemporalDataPoint,
	barWidth time.Duration, color color.Color) (tbs TemporalBarSeries, err error) {
	tbs.ser, err = numChart.base.AddTemporalBarSeries(name, points, barWidth, color)
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (tempChart *PolarTemporalChart) DeleteSeries(name string) {
	tempChart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (tempChart *PolarTemporalChart) SetTitle(l string) {
	tempChart.base.SetTitle(l)
}

// HideLegend hides the legend and uses the full space for the chart
func (tempChart *PolarTemporalChart) HideLegend() {
	tempChart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (tempChart *PolarTemporalChart) ShowLegend() {
	tempChart.base.ShowLegend()
}

// SetRAxisLabel sets the label of the r-axis, which will be displayed at the bottom-right
func (tempChart *PolarTemporalChart) SetRAxisLabel(l string) {
	tempChart.base.SetToAxisLabel(l)
}

// SetRRange sets a user defined range for the r-axis;
// an error is returned if max<0 or if the origin has been defined by the user before and is outside the given range
func (tempChart *PolarTemporalChart) SetRRange(max float64) (err error) {
	err = tempChart.base.SetToRange(0.0, max)
	return
}

// SetAutoRRange overrides a previously user defined range and lets the range be calculated automatically
func (tempChart *PolarTemporalChart) SetAutoRRange() {
	tempChart.base.SetAutoToRange()
}

// SetRTicks sets the list of user defined ticks to be shown on the r-axis
func (tempChart *PolarTemporalChart) SetRTicks(ts []data.NumericalTick) {
	tempChart.base.SetToTicks(ts)
}

// SetAutoRTicks overrides a previously user defined set of r-axis ticks and lets the ticks be calculated automatically
func (tempChart *PolarTemporalChart) SetAutoRTicks(autoSupportLine bool) {
	tempChart.base.SetAutoToTicks(autoSupportLine)
}

// SetOrigin sets a user defined origin (crossing of t and r axis).
// An error is returned, if a range has been defined before and at least one coordinate is outside the range.
func (tempChart *PolarTemporalChart) SetOrigin(t time.Time, r float64) (err error) {
	err = tempChart.base.SetTOrigin(t, r)
	return
}

// SetAutoOrigin resets a previously user defined origin and allows the chart to calculate the ideal origin automatically
func (tempChart *PolarTemporalChart) SetAutoOrigin() {
	tempChart.base.SetAutoOrigin()
}

// SetTAxisLabel sets the label of the t-axis, which will be displayed at the left side
func (tempChart *PolarTemporalChart) SetTAxisLabel(l string) {
	tempChart.base.SetFromAxisLabel(l)
}

// SetTRange sets a user defined range for the t-axis.
// An error is returned, if min after max or if the origin has been defined by the user before and is outside the given range
func (tempChart *PolarTemporalChart) SetTRange(min time.Time, max time.Time) (err error) {
	err = tempChart.base.SetFromTRange(min, max)
	return
}

// SetAutoTRange overrides a previously user defined range and lets the range be calculated automatically
func (tempChart *PolarTemporalChart) SetAutoTRange() {
	tempChart.base.SetAutoFromRange()
}

// SetTTicks sets the list of user defined ticks to be shown on the t-axis
func (tempChart *PolarTemporalChart) SetTTicks(ts []data.TemporalTick, format string) {
	tempChart.base.SetFromTTicks(ts, format)
}

// SetAutoTTicks overrides a previously user defined set of t-axis ticks and lets the ticks be calculated automatically
func (tempChart *PolarTemporalChart) SetAutoTTicks(autoSupportLine bool) {
	tempChart.base.SetAutoFromTicks(autoSupportLine)
}
