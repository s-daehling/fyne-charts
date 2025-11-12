package coord

import (
	"image/color"

	"github.com/s-daehling/fyne-charts/internal/coord"
	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
)

// CartesianCategoricalChart implements a cartesian plane with a categorical c-axis and a numerical y-axis
type CartesianCategoricalChart struct {
	coordChart
}

// NewCartesianCategoricalChart returns an initialized CategoricalChart
func NewCartesianCategoricalChart() (catChart *CartesianCategoricalChart) {
	catChart = &CartesianCategoricalChart{
		coordChart: emptyCoordChart(coord.CartesianPlane, coord.Categorical),
	}
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C and Val is not restricted
func (catChart *CartesianCategoricalChart) AddScatterSeries(name string, points []data.CategoricalPoint,
	color color.Color) (css CategoricalPointSeries, err error) {
	css.ser, err = catChart.BaseChart.AddCategoricalScatterSeries(name, points, color)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C and Val is not restricted
func (catChart *CartesianCategoricalChart) AddBarSeries(name string, points []data.CategoricalPoint,
	color color.Color) (cbs CategoricalPointSeries, err error) {
	cbs.ser, err = catChart.BaseChart.AddCategoricalBarSeries(name, points, color)
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
	css.ser, err = catChart.BaseChart.AddCategoricalStackedBarSeries(name, points)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C and Val is not restricted
func (catChart *CartesianCategoricalChart) AddLollipopSeries(name string, points []data.CategoricalPoint,
	color color.Color) (cls CategoricalPointSeries, err error) {
	cls.ser, err = catChart.BaseChart.AddCategoricalLollipopSeries(name, points, color)
	return
}

// AddBoxSeries adds a series of data which is visualized as box chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. boxes with same C).
// Boxes with a C that already exists, will be ignored.
// The range of C and values is not restricted.
func (catChart *CartesianCategoricalChart) AddBoxSeries(name string,
	points []data.CategoricalBox, col color.Color) (cbs CategoricalBoxSeries, err error) {
	cbs.ser, err = catChart.BaseChart.AddCategoricalBoxSeries(name, points, col)
	return
}

// SetYAxisLabel sets the label of the y-axis, which will be displayed at the left side
func (catChart *CartesianCategoricalChart) SetYAxisLabel(l string) {
	catChart.BaseChart.SetToAxisLabel(l)
}

// SetYRange sets a user defined range for the y-axis;
// an error is returned if min>max or if the origin has been defined by the user before and is outside the given range
func (catChart *CartesianCategoricalChart) SetYRange(min float64, max float64) (err error) {
	err = catChart.BaseChart.SetToRange(min, max)
	return
}

// SetAutoYRange overrides a previously user defined range and lets the range be calculated automatically
func (catChart *CartesianCategoricalChart) SetAutoYRange() {
	catChart.BaseChart.SetAutoToRange()
}

// SetYTicks sets the list of user defined ticks to be shown on the y-axis
func (catChart *CartesianCategoricalChart) SetYTicks(ts []data.NumericalTick) {
	catChart.BaseChart.SetToTicks(ts)
}

// SetAutoYTicks overrides a previously user defined set of y-axis ticks and lets the ticks be calculated automatically
func (catChart *CartesianCategoricalChart) SetAutoYTicks(autoSupportLine bool) {
	catChart.BaseChart.SetAutoToTicks(autoSupportLine)
}

// SetYAxisStyle changes the style of the Y-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (catChart *CartesianCategoricalChart) SetYAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	catChart.BaseChart.SetToAxisLabelStyle(labelSize, labelColor)
	catChart.BaseChart.SetToAxisStyle(axisColor)
}

// SetCAxisLabel sets the label of the c-axis, which will be displayed at the bottom
func (catChart *CartesianCategoricalChart) SetCAxisLabel(l string) {
	catChart.BaseChart.SetFromAxisLabel(l)
}

// SetCRange sets a user defined range for the c-axis.
// This will also determine the ticks and their order on the c-axis.
// An error is returned if cs is empty.
func (catChart *CartesianCategoricalChart) SetCRange(cs []string) (err error) {
	err = catChart.BaseChart.SetFromCRange(cs)
	return
}

// SetAutoCRange overrides a previously user defined set of categories to be shown and lets the set be calculated automatically
func (catChart *CartesianCategoricalChart) SetAutoCRange() {
	catChart.BaseChart.SetAutoFromRange()
}

// SetCAxisStyle changes the style of the C-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (catChart *CartesianCategoricalChart) SetCAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	catChart.BaseChart.SetFromAxisLabelStyle(labelSize, labelColor)
	catChart.BaseChart.SetFromAxisStyle(axisColor)
}
