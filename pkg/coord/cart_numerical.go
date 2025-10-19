package coord

import (
	"image/color"

	"github.com/s-daehling/fyne-charts/internal/coord"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// CartesianNumericalChart implements a cartesian plane with two numerical axes x and y
type CartesianNumericalChart struct {
	base *coord.BaseChart
	widget.BaseWidget
}

// NewCartesianNumericalChart returns an initialized CartesianNumericalChart
func NewCartesianNumericalChart() (numChart *CartesianNumericalChart) {
	numChart = &CartesianNumericalChart{
		base: coord.EmptyBaseChart(coord.CartesianPlane, coord.Numerical),
	}
	numChart.ExtendBaseWidget(numChart)
	return
}

// CreateRenderer creates the renderer of the widget
func (numChart *CartesianNumericalChart) CreateRenderer() fyne.WidgetRenderer {
	r := numChart.base.GetRenderer()
	return r
}

// AddLineSeries adds a series of data which is visualized as line chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// data does not need to be sorted. It will be sorted by X by the method.
// The method does not check for duplicates (i.e. data points with same X)
// The range of X and Val is not restricted
func (numChart *CartesianNumericalChart) AddLineSeries(name string, points []data.NumericalDataPoint,
	showDots bool, color color.Color) (nls NumericalLineSeries, err error) {
	nls.ser, err = numChart.base.AddNumericalLineSeries(name, points, showDots, color)
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same X)
// The range of X and Val is not restricted
func (numChart *CartesianNumericalChart) AddScatterSeries(name string, points []data.NumericalDataPoint,
	color color.Color) (nss NumericalScatterSeries, err error) {
	nss.ser, err = numChart.base.AddNumericalScatterSeries(name, points, color)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same X)
// The range of X and Val is not restricted
func (numChart *CartesianNumericalChart) AddLollipopSeries(name string, points []data.NumericalDataPoint,
	color color.Color) (nls NumericalLollipopSeries, err error) {
	nls.ser, err = numChart.base.AddNumericalLollipopSeries(name, points, color)
	return
}

// AddCandleStickSeries adds a series of data which is visualized as canlde stick chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same X)
// The range of XStart, XEnd and values is not restricted
func (numChart *CartesianNumericalChart) AddCandleStickSeries(name string,
	points []data.NumericalCandleStick) (ncs NumericalCandleStickSeries, err error) {
	ncs.ser, err = numChart.base.AddNumericalCandleStickSeries(name, points)
	return
}

// AddBoxSeries adds a series of data which is visualized as box chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same X)
// The range of X and values is not restricted
func (numChart *CartesianNumericalChart) AddBoxSeries(name string,
	points []data.NumericalBox, col color.Color) (nbs NumericalBoxSeries, err error) {
	nbs.ser, err = numChart.base.AddNumericalBoxSeries(name, points, col)
	return
}

// AddAreaSeries adds a series of data which is visualized as area chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// data does not need to be sorted. It will be sorted by X by the method.
// The method does not check for duplicates (i.e. data points with same X).
// The range of X and Val is not restricted
func (numChart *CartesianNumericalChart) AddAreaSeries(name string, points []data.NumericalDataPoint, showDots bool,
	color color.Color) (nas NumericalAreaSeries, err error) {
	nas.ser, err = numChart.base.AddNumericalAreaSeries(name, points, showDots, color)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same X)
// The range of X and Val is not restricted
// The bars are centered around their X value of the data points. barWidth is the width of the bars.
// An error is returned if barWidth < 0
func (numChart *CartesianNumericalChart) AddBarSeries(name string, points []data.NumericalDataPoint,
	barWidth float64, color color.Color) (nbs NumericalBarSeries, err error) {
	nbs.ser, err = numChart.base.AddNumericalBarSeries(name, points, barWidth, color)
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (numChart *CartesianNumericalChart) DeleteSeries(name string) {
	numChart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (numChart *CartesianNumericalChart) SetTitle(l string) {
	numChart.base.SetTitle(l)
}

// HideLegend hides the legend and uses the full space for the chart
func (numChart *CartesianNumericalChart) HideLegend() {
	numChart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (numChart *CartesianNumericalChart) ShowLegend() {
	numChart.base.ShowLegend()
}

// SetYAxisLabel sets the label of the y-axis, which will be displayed at the left side
func (numChart *CartesianNumericalChart) SetYAxisLabel(l string) {
	numChart.base.SetToAxisLabel(l)
}

// SetYRange sets a user defined range for the y-axis;
// an error is returned if min>max or if the origin has been defined by the user before and is outside the given range
func (numChart *CartesianNumericalChart) SetYRange(min float64, max float64) (err error) {
	err = numChart.base.SetToRange(min, max)
	return
}

// SetAutoYRange overrides a previously user defined range and lets the range be calculated automatically
func (numChart *CartesianNumericalChart) SetAutoYRange() {
	numChart.base.SetAutoToRange()
}

// SetYTicks sets the list of user defined ticks to be shown on the y-axis
func (numChart *CartesianNumericalChart) SetYTicks(ts []data.NumericalTick) {
	numChart.base.SetToTicks(ts)
}

// SetAutoYTicks overrides a previously user defined set of y-axis ticks and lets the ticks be calculated automatically
func (numChart *CartesianNumericalChart) SetAutoYTicks(autoSupportLine bool) {
	numChart.base.SetAutoToTicks(autoSupportLine)
}

// SetOrigin sets a user defined origin (crossing of x and y axis).
// An error is returned, if a range has been defined before and at least one coordinate is outside the range.
func (numChart *CartesianNumericalChart) SetOrigin(x float64, y float64) (err error) {
	err = numChart.base.SetNOrigin(x, y)
	return
}

// SetAutoOrigin resets a previously user defined origin and allows the chart to calculate the ideal origin automatically
func (numChart *CartesianNumericalChart) SetAutoOrigin() {
	numChart.base.SetAutoOrigin()
}

// SetXAxisLabel sets the label of the x-axis, which will be displayed at the bottom
func (numChart *CartesianNumericalChart) SetXAxisLabel(l string) {
	numChart.base.SetFromAxisLabel(l)
}

// SetXRange sets a user defined range for the x-axis.
// An error is returned, if min>max or if the origin has been defined by the user before and is outside the given range
func (numChart *CartesianNumericalChart) SetXRange(min float64, max float64) (err error) {
	err = numChart.base.SetFromNRange(min, max)
	return
}

// SetAutoXRange overrides a previously user defined range and lets the range be calculated automatically
func (numChart *CartesianNumericalChart) SetAutoXRange() {
	numChart.base.SetAutoFromRange()
}

// SetXTicks sets the list of user defined ticks to be shown on the x-axis
func (numChart *CartesianNumericalChart) SetXTicks(ts []data.NumericalTick) {
	numChart.base.SetFromNTicks(ts)
}

// SetAutoXTicks overrides a previously user defined set of x-axis ticks and lets the ticks be calculated automatically
func (numChart *CartesianNumericalChart) SetAutoXTicks(autoSupportLine bool) {
	numChart.base.SetAutoFromTicks(autoSupportLine)
}
