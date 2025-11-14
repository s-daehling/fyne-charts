package coord

import (
	"time"

	"github.com/s-daehling/fyne-charts/internal/coord"
	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
)

// PolarTemporalChart implements a polar plane with one temporal t-axis and one numerical r-axis
type PolarTemporalChart struct {
	coordChart
}

// NewPolarTemporalChart returns an initialized PolarTemporalChart
func NewPolarTemporalChart() (tempChart *PolarTemporalChart) {
	tempChart = &PolarTemporalChart{
		coordChart: emptyCoordChart(coord.PolarPlane, coord.Temporal),
	}
	tempChart.ExtendBaseWidget(tempChart)
	return
}

// AddLineSeries adds a series of data which is visualized as line chart.
// If showDots is true, dots are displayed at the osition of the series points.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// An error is returned,if another series with the same name exists, if the series is already added to another chart or if Val < 0 for one or more points
func (tempChart *PolarTemporalChart) AddLineSeries(tps *TemporalPointSeries, showDots bool) (err error) {
	err = tempChart.base.AddLineSeries(tps.ser, showDots)
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// An error is returned,if another series with the same name exists, if the series is already added to another chart or if Val < 0 for one or more points
func (tempChart *PolarTemporalChart) AddScatterSeries(tps *TemporalPointSeries) (err error) {
	err = tempChart.base.AddScatterSeries(tps.ser)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// An error is returned,if another series with the same name exists, if the series is already added to another chart or if Val < 0 for one or more points
func (tempChart *PolarTemporalChart) AddLollipopSeries(tps *TemporalPointSeries) (err error) {
	err = tempChart.base.AddLollipopSeries(tps.ser)
	return
}

// AddAreaSeries adds a series of data which is visualized as area chart.
// If showDots is true, dots are displayed at the osition of the series points.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// An error is returned,if another series with the same name exists, if the series is already added to another chart or if Val < 0 for one or more points
func (tempChart *PolarTemporalChart) AddAreaSeries(tps *TemporalPointSeries, showDots bool) (err error) {
	err = tempChart.base.AddAreaSeries(tps.ser, showDots)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// The bars are centered around their T value of the data points. barWidth is the width of the bars.
// An error is returned,if another series with the same name exists, if the series is already added to another chart or if Val < 0 for one or more points
func (tempChart *PolarTemporalChart) AddBarSeries(tps *TemporalPointSeries, barWidth time.Duration) (err error) {
	err = tps.SetBarWidth(barWidth)
	if err != nil {
		return
	}
	err = tempChart.base.AddBarSeries(tps.ser)
	return
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

// SetRAxisStyle changes the style of the R-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (tempChart *PolarTemporalChart) SetRAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	tempChart.base.SetToAxisLabelStyle(labelSize, labelColor)
	tempChart.base.SetToAxisStyle(axisColor)
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

// SetTAxisStyle changes the style of the T-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (tempChart *PolarTemporalChart) SetTAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	tempChart.base.SetFromAxisLabelStyle(labelSize, labelColor)
	tempChart.base.SetFromAxisStyle(axisColor)
}
