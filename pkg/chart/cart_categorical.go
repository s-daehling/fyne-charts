package chart

import (
	"image/color"

	"github.com/s-daehling/fyne-charts/internal/chart"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// CartesianCategoricalChart implements a cartesian plane with a categorical c-axis and a numerical y-axis
type CartesianCategoricalChart struct {
	base *chart.BaseChart
	widget.BaseWidget
}

// NewCartesianCategoricalChart returns an initialized CategoricalChart
func NewCartesianCategoricalChart() (catChart *CartesianCategoricalChart) {
	catChart = &CartesianCategoricalChart{
		base: chart.EmptyBaseChart(chart.CartesianPlane, chart.Categorical),
	}
	catChart.ExtendBaseWidget(catChart)
	return
}

// CreateRenderer creates the renderer of the widget
func (catChart *CartesianCategoricalChart) CreateRenderer() fyne.WidgetRenderer {
	r := chart.EmptyCartesianRenderer(catChart.base)
	return r
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C and Val is not restricted
func (catChart *CartesianCategoricalChart) AddScatterSeries(name string, points []data.CategoricalDataPoint,
	color color.Color) (css CategoricalScatterSeries, err error) {
	css.ser, err = catChart.base.AddCategoricalScatterSeries(name, points, color)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C and Val is not restricted
func (catChart *CartesianCategoricalChart) AddBarSeries(name string, points []data.CategoricalDataPoint,
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
func (catChart *CartesianCategoricalChart) AddStackedBarSeries(name string,
	points []data.CategoricalDataSeries) (css CategoricalStackedBarSeries, err error) {
	css.ser, err = catChart.base.AddCategoricalStackedBarSeries(name, points)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C and Val is not restricted
func (catChart *CartesianCategoricalChart) AddLollipopSeries(name string, points []data.CategoricalDataPoint,
	color color.Color) (cls CategoricalLollipopSeries, err error) {
	cls.ser, err = catChart.base.AddCategoricalLollipopSeries(name, points, color)
	return
}

// AddBoxSeries adds a series of data which is visualized as canlde stick chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. boxes with same C).
// Boxes with a C that already exists, will be ignored.
// The range of C and values is not restricted.
func (catChart *CartesianCategoricalChart) AddBoxSeries(name string,
	points []data.CategoricalBox, col color.Color) (cbs CategoricalBoxSeries, err error) {
	cbs.ser, err = catChart.base.AddCategoricalBoxSeries(name, points, col)
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (catChart *CartesianCategoricalChart) DeleteSeries(name string) {
	catChart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (catChart *CartesianCategoricalChart) SetTitle(l string) {
	catChart.base.SetTitle(l)
}

// HideLegend hides the legend and uses the full space for the chart
func (catChart *CartesianCategoricalChart) HideLegend() {
	catChart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (catChart *CartesianCategoricalChart) ShowLegend() {
	catChart.base.ShowLegend()
}

// SetYAxisLabel sets the label of the y-axis, which will be displayed at the left side
func (catChart *CartesianCategoricalChart) SetYAxisLabel(l string) {
	catChart.base.SetToAxisLabel(l)
}

// SetYRange sets a user defined range for the y-axis;
// an error is returned if min>max or if the origin has been defined by the user before and is outside the given range
func (catChart *CartesianCategoricalChart) SetYRange(min float64, max float64) (err error) {
	err = catChart.base.SetToRange(min, max)
	return
}

// SetAutoYRange overrides a previously user defined range and lets the range be calculated automatically
func (catChart *CartesianCategoricalChart) SetAutoYRange() {
	catChart.base.SetAutoToRange()
}

// SetYTicks sets the list of user defined ticks to be shown on the y-axis
func (catChart *CartesianCategoricalChart) SetYTicks(ts []data.NumericalTick) {
	catChart.base.SetToTicks(ts)
}

// SetAutoYTicks overrides a previously user defined set of y-axis ticks and lets the ticks be calculated automatically
func (catChart *CartesianCategoricalChart) SetAutoYTicks(autoSupportLine bool) {
	catChart.base.SetAutoToTicks(autoSupportLine)
}

// SetCAxisLabel sets the label of the c-axis, which will be displayed at the bottom
func (catChart *CartesianCategoricalChart) SetCAxisLabel(l string) {
	catChart.base.SetFromAxisLabel(l)
}

// SetCRange sets a user defined range for the c-axis.
// This will also determine the ticks and their order on the c-axis.
// An error is returned if cs is empty.
func (catChart *CartesianCategoricalChart) SetCRange(cs []string) (err error) {
	err = catChart.base.SetFromCRange(cs)
	return
}

// SetAutoCRange overrides a previously user defined set of categories to be shown and lets the set be calculated automatically
func (catChart *CartesianCategoricalChart) SetAutoCRange() {
	catChart.base.SetAutoFromRange()
}
