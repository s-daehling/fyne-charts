package chart

import (
	"image/color"

	"github.com/s-daehling/fyne-charts/internal/chart"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// PolarCategoricalChart implements a polar plane with one categorical c-axis and one numerical r-axis
type PolarCategoricalChart struct {
	base *chart.BaseChart
	widget.BaseWidget
}

// NewPolarCategoricalChart returns an initialized PolarCategoricalChart
func NewPolarCategoricalChart() (catChart *PolarCategoricalChart) {
	catChart = &PolarCategoricalChart{
		base: chart.EmptyBaseChart(chart.PolarPlane, chart.Categorical),
	}
	catChart.ExtendBaseWidget(catChart)
	return
}

// CreateRenderer creates the renderer of the widget
func (catChart *PolarCategoricalChart) CreateRenderer() fyne.WidgetRenderer {
	r := catChart.base.GetRenderer()
	return r
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0.
func (catChart *PolarCategoricalChart) AddScatterSeries(name string, points []data.CategoricalDataPoint,
	color color.Color) (css CategoricalScatterSeries, err error) {
	css.ser, err = catChart.base.AddCategoricalScatterSeries(name, points, color)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0.
func (catChart *PolarCategoricalChart) AddLollipopSeries(name string, points []data.CategoricalDataPoint,
	color color.Color) (cls CategoricalLollipopSeries, err error) {
	cls.ser, err = catChart.base.AddCategoricalLollipopSeries(name, points, color)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0.
func (catChart *PolarCategoricalChart) AddBarSeries(name string, points []data.CategoricalDataPoint,
	color color.Color) (cbs CategoricalBarSeries, err error) {
	cbs.ser, err = catChart.base.AddCategoricalBarSeries(name, points, color)
	return
}

// AddStackedBarSeries adds a series of data which is visualized as stacked bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0.
func (catChart *PolarCategoricalChart) AddStackedBarSeries(name string,
	points []data.CategoricalDataSeries) (css CategoricalStackedBarSeries, err error) {
	css.ser, err = catChart.base.AddCategoricalStackedBarSeries(name, points)
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (catChart *PolarCategoricalChart) DeleteSeries(name string) {
	catChart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (catChart *PolarCategoricalChart) SetTitle(l string) {
	catChart.base.SetTitle(l)
}

// HideLegend hides the legend and uses the full space for the chart
func (catChart *PolarCategoricalChart) HideLegend() {
	catChart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (catChart *PolarCategoricalChart) ShowLegend() {
	catChart.base.ShowLegend()
}

// SetYAxisLabel sets the label of the y-axis, which will be displayed at the left side
func (catChart *PolarCategoricalChart) SetRAxisLabel(l string) {
	catChart.base.SetToAxisLabel(l)
}

// SetRRange sets a user defined range for the r-axis;
// an error is returned if max<0 or if the origin has been defined by the user before and is outside the given range
func (catChart *PolarCategoricalChart) SetRRange(max float64) (err error) {
	err = catChart.base.SetToRange(0.0, max)
	return
}

// SetAutoRRange overrides a previously user defined range and lets the range be calculated automatically
func (catChart *PolarCategoricalChart) SetAutoRRange() {
	catChart.base.SetAutoToRange()
}

// SetRTicks sets the list of user defined ticks to be shown on the r-axis
func (catChart *PolarCategoricalChart) SetRTicks(ts []data.NumericalTick) {
	catChart.base.SetToTicks(ts)
}

// SetAutoRTicks overrides a previously user defined set of r-axis ticks and lets the ticks be calculated automatically
func (catChart *PolarCategoricalChart) SetAutoRTicks(autoSupportLine bool) {
	catChart.base.SetAutoToTicks(autoSupportLine)
}

// SetCAxisLabel sets the label of the c-axis, which will be displayed at the left side
func (catChart *PolarCategoricalChart) SetCAxisLabel(l string) {
	catChart.base.SetFromAxisLabel(l)
}

// SetCRange sets a user defined range for the c-axis.
// This will also determine the ticks and their order on the c-axis.
// An error is returned if cs is empty.
func (catChart *PolarCategoricalChart) SetCRange(cs []string) (err error) {
	err = catChart.base.SetFromCRange(cs)
	return
}

// SetAutoCRange overrides a previously user defined range and lets the range be calculated automatically
func (catChart *PolarCategoricalChart) SetAutoCRange() {
	catChart.base.SetAutoFromRange()
}
