package coord

import (
	"image/color"

	"github.com/s-daehling/fyne-charts/internal/coord"
	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
)

// PolarCategoricalChart implements a polar plane with one categorical c-axis and one numerical r-axis
type PolarCategoricalChart struct {
	coordChart
}

// NewPolarCategoricalChart returns an initialized PolarCategoricalChart
func NewPolarCategoricalChart() (catChart *PolarCategoricalChart) {
	catChart = &PolarCategoricalChart{
		coordChart: emptyCoordChart(coord.PolarPlane, coord.Categorical),
	}
	catChart.ExtendBaseWidget(catChart)
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0.
func (catChart *PolarCategoricalChart) AddScatterSeries(name string, points []data.CategoricalPoint,
	color color.Color) (css CategoricalPointSeries, err error) {
	css.ser, err = catChart.base.AddCategoricalScatterSeries(name, points, color)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0.
func (catChart *PolarCategoricalChart) AddLollipopSeries(name string, points []data.CategoricalPoint,
	color color.Color) (cls CategoricalPointSeries, err error) {
	cls.ser, err = catChart.base.AddCategoricalLollipopSeries(name, points, color)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0.
func (catChart *PolarCategoricalChart) AddBarSeries(name string, points []data.CategoricalPoint,
	color color.Color) (cbs CategoricalPointSeries, err error) {
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

// SetRAxisStyle changes the style of the R-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (catChart *PolarCategoricalChart) SetRAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	catChart.base.SetToAxisLabelStyle(labelSize, labelColor)
	catChart.base.SetToAxisStyle(axisColor)
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

// SetCAxisStyle changes the style of the C-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (catChart *PolarCategoricalChart) SetCAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	catChart.base.SetFromAxisLabelStyle(labelSize, labelColor)
	catChart.base.SetFromAxisStyle(axisColor)
}
