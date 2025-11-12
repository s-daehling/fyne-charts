package coord

import (
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/coord"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
)

// CartesianTemporalChart implements a cartesian plane with a temporal t-axis and a numerical y-axis
type CartesianTemporalChart struct {
	coordChart
}

// NewCartesianTemporalChart returns an initialized CartesianTemporalChart
func NewCartesianTemporalChart() (tempChart *CartesianTemporalChart) {
	tempChart = &CartesianTemporalChart{
		coordChart: emptyCoordChart(coord.CartesianPlane, coord.Temporal),
	}
	return
}

// AddLineSeries adds a series of data which is visualized as line chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// data does not need to be sorted. It will be sorted by T by the method.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and Val is not restricted
func (tempChart *CartesianTemporalChart) AddLineSeries(name string, points []data.TemporalPoint,
	showDots bool, color color.Color) (tls TemporalLineSeries, err error) {
	tls.ser, err = tempChart.BaseChart.AddTemporalLineSeries(name, points, showDots, color)
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and Val is not restricted
func (tempChart *CartesianTemporalChart) AddScatterSeries(name string, points []data.TemporalPoint,
	color color.Color) (tss TemporalScatterSeries, err error) {
	tss.ser, err = tempChart.BaseChart.AddTemporalScatterSeries(name, points, color)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and Val is not restricted
func (tempChart *CartesianTemporalChart) AddLollipopSeries(name string, points []data.TemporalPoint,
	color color.Color) (tls TemporalLollipopSeries, err error) {
	tls.ser, err = tempChart.BaseChart.AddTemporalLollipopSeries(name, points, color)
	return
}

// AddCandleStickSeries adds a series of data which is visualized as canlde stick chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same T)
// The range of TStart, TEnd and values is not restricted
func (tempChart *CartesianTemporalChart) AddCandleStickSeries(name string,
	points []data.TemporalCandleStick) (tcs TemporalCandleStickSeries, err error) {
	tcs.ser, err = tempChart.BaseChart.AddTemporalCandleStickSeries(name, points)
	return
}

// AddBoxSeries adds a series of data which is visualized as box chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and values is not restricted
func (tempChart *CartesianTemporalChart) AddBoxSeries(name string,
	points []data.TemporalBox, col color.Color) (tbs TemporalBoxSeries, err error) {
	tbs.ser, err = tempChart.BaseChart.AddTemporalBoxSeries(name, points, col)
	return
}

// AddAreaSeries adds a series of data which is visualized as area chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// data does not need to be sorted. It will be sorted by T by the method.
// The method does not check for duplicates (i.e. data points with same T).
// The range of T and Val is not restricted
func (tempChart *CartesianTemporalChart) AddAreaSeries(name string, points []data.TemporalPoint,
	showDots bool, color color.Color) (tas TemporalAreaSeries, err error) {
	tas.ser, err = tempChart.BaseChart.AddTemporalAreaSeries(name, points, showDots, color)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same T)
// The range of T and Val is not restricted
// The bars are centered around their T value of the data points. barWidth is the width of the bars.
// An error is returned if barWidth < 0
func (tempChart *CartesianTemporalChart) AddBarSeries(name string, points []data.TemporalPoint,
	barWidth time.Duration, color color.Color) (tbs TemporalBarSeries, err error) {
	tbs.ser, err = tempChart.BaseChart.AddTemporalBarSeries(name, points, barWidth, color)
	return
}

// SetYAxisLabel sets the label of the y-axis, which will be displayed at the left side
func (tempChart *CartesianTemporalChart) SetYAxisLabel(l string) {
	tempChart.BaseChart.SetToAxisLabel(l)
}

// SetYRange sets a user defined range for the y-axis;
// an error is returned if min>max or if the origin has been defined by the user before and is outside the given range
func (tempChart *CartesianTemporalChart) SetYRange(min float64, max float64) (err error) {
	err = tempChart.BaseChart.SetToRange(min, max)
	return
}

// SetAutoYRange overrides a previously user defined range and lets the range be calculated automatically
func (tempChart *CartesianTemporalChart) SetAutoYRange() {
	tempChart.BaseChart.SetAutoToRange()
}

// SetYTicks sets the list of user defined ticks to be shown on the y-axis
func (tempChart *CartesianTemporalChart) SetYTicks(ts []data.NumericalTick) {
	tempChart.BaseChart.SetToTicks(ts)
}

// SetAutoYTicks overrides a previously user defined set of y-axis ticks and lets the ticks be calculated automatically
func (tempChart *CartesianTemporalChart) SetAutoYTicks(autoSupportLine bool) {
	tempChart.BaseChart.SetAutoToTicks(autoSupportLine)
}

// SetYAxisStyle changes the style of the Y-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (tempChart *CartesianTemporalChart) SetYAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	tempChart.BaseChart.SetToAxisLabelStyle(labelSize, labelColor)
	tempChart.BaseChart.SetToAxisStyle(axisColor)
}

// SetOrigin sets a user defined origin (crossing of t and y axis).
// An error is returned. if a range has been defined before and at least one coordinate is outside the range.
func (tempChart *CartesianTemporalChart) SetOrigin(t time.Time, y float64) (err error) {
	err = tempChart.BaseChart.SetTOrigin(t, y)
	return
}

// SetAutoOrigin resets a previously user defined origin and allows the chart to calculate the ideal origin automatically
func (tempChart *CartesianTemporalChart) SetAutoOrigin() {
	tempChart.BaseChart.SetAutoOrigin()
}

// SetTAxisLabel sets the label of the t-axis, which will be displayed at the bottom
func (tempChart *CartesianTemporalChart) SetTAxisLabel(l string) {
	tempChart.BaseChart.SetFromAxisLabel(l)
}

// SetTRange sets a user defined range for the t-axis.
// An error is returned, if min after max or if the origin has been defined by the user before and is outside the given range
func (tempChart *CartesianTemporalChart) SetTRange(min time.Time, max time.Time) (err error) {
	err = tempChart.BaseChart.SetFromTRange(min, max)
	return
}

// SetAutoTRange overrides a previously user defined range and lets the range be calculated automatically
func (tempChart *CartesianTemporalChart) SetAutoTRange() {
	tempChart.BaseChart.SetAutoFromRange()
}

// SetTTicks sets the list of user defined ticks to be shown on the t-axis
func (tempChart *CartesianTemporalChart) SetTTicks(ts []data.TemporalTick, format string) {
	tempChart.BaseChart.SetFromTTicks(ts, format)
}

// SetAutoTTicks overrides a previously user defined set of t-axis ticks and lets the ticks be calculated automatically
func (tempChart *CartesianTemporalChart) SetAutoTTicks(autoSupportLine bool) {
	tempChart.BaseChart.SetAutoFromTicks(autoSupportLine)
}

// SetTAxisStyle changes the style of the T-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (tempChart *CartesianTemporalChart) SetTAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	tempChart.BaseChart.SetFromAxisLabelStyle(labelSize, labelColor)
	tempChart.BaseChart.SetFromAxisStyle(axisColor)
}
