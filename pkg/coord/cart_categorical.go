package coord

import (
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
	catChart.ExtendBaseWidget(catChart)
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C and Val is not restricted
func (catChart *CartesianCategoricalChart) AddScatterSeries(cps CategoricalPointSeries) (err error) {
	err = catChart.base.AddScatterSeries(cps.ser)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C and Val is not restricted
func (catChart *CartesianCategoricalChart) AddBarSeries(cps CategoricalPointSeries) (err error) {
	err = catChart.base.AddBarSeries(cps.ser)
	return
}

// AddStackedBarSeries adds a series of data which is visualized as stacked bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0.
func (catChart *CartesianCategoricalChart) AddStackedBarSeries(css CategoricalStackedSeries) (err error) {
	err = catChart.base.AddStackedBarSeries(css.ser)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C and Val is not restricted
func (catChart *CartesianCategoricalChart) AddLollipopSeries(cps CategoricalPointSeries) (err error) {
	err = catChart.base.AddLollipopSeries(cps.ser)
	return
}

// AddBoxSeries adds a series of data which is visualized as box chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. boxes with same C).
// Boxes with a C that already exists, will be ignored.
// The range of C and values is not restricted.
func (catChart *CartesianCategoricalChart) AddBoxSeries(cbs CategoricalBoxSeries) (err error) {
	err = catChart.base.AddBoxSeries(cbs.ser)
	return
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

// SetYAxisStyle changes the style of the Y-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (catChart *CartesianCategoricalChart) SetYAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	catChart.base.SetToAxisLabelStyle(labelSize, labelColor)
	catChart.base.SetToAxisStyle(axisColor)
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

// SetCAxisStyle changes the style of the C-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (catChart *CartesianCategoricalChart) SetCAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	catChart.base.SetFromAxisLabelStyle(labelSize, labelColor)
	catChart.base.SetFromAxisStyle(axisColor)
}
