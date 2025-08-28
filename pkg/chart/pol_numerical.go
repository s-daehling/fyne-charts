package chart

import (
	"image/color"

	"github.com/s-daehling/fyne-charts/internal/chart"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// PolarNumericalChart implements a polar plane with one numerical phi-axis and one numerical r-axis
type PolarNumericalChart struct {
	base *chart.BaseChart
	widget.BaseWidget
}

// NewPolarNumericalChart returns an initialized PolarNumericalChart
func NewPolarNumericalChart() (numChart *PolarNumericalChart) {
	numChart = &PolarNumericalChart{
		base: chart.EmptyBaseChart(chart.PolarPlane, chart.Numerical),
	}
	numChart.ExtendBaseWidget(numChart)
	return
}

// CreateRenderer creates the renderer of the widget
func (numChart *PolarNumericalChart) CreateRenderer() fyne.WidgetRenderer {
	r := chart.EmptyPolarRenderer(numChart.base)
	return r
}

// AddLineSeries adds a series of data which is visualized as line chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// data does not need to be sorted. It will be sorted by A by the method.
// The method does not check for duplicates (i.e. data points with same A)
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
func (numChart *PolarNumericalChart) AddLineSeries(name string, points []data.NumericalDataPoint, showDots bool,
	color color.Color) (nls NumericalLineSeries, err error) {
	nls.ser, err = numChart.base.AddNumericalLineSeries(name, points, showDots, color)
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same A)
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
func (numChart *PolarNumericalChart) AddScatterSeries(name string, points []data.NumericalDataPoint,
	color color.Color) (nss NumericalScatterSeries, err error) {
	nss.ser, err = numChart.base.AddNumericalScatterSeries(name, points, color)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same A)
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
func (numChart *PolarNumericalChart) AddLollipopSeries(name string, points []data.NumericalDataPoint,
	color color.Color) (nls NumericalLollipopSeries, err error) {
	nls.ser, err = numChart.base.AddNumericalLollipopSeries(name, points, color)
	return
}

// AddAreaSeries adds a series of data which is visualized as area chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// data does not need to be sorted. It will be sorted by A by the method.
// The method does not check for duplicates (i.e. data points with same A).
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
func (numChart *PolarNumericalChart) AddAreaSeries(name string, points []data.NumericalDataPoint, showDots bool,
	color color.Color) (nas NumericalAreaSeries, err error) {
	nas.ser, err = numChart.base.AddNumericalAreaSeries(name, points, showDots, color)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same A)
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
// The bars are centered around their A value of the data points. barWidth is the width of the bars.
// An error is returned if barWidth < 0
func (numChart *PolarNumericalChart) AddBarSeries(name string, points []data.NumericalDataPoint,
	barWidth float64, color color.Color) (nbs NumericalBarSeries, err error) {
	nbs.ser, err = numChart.base.AddNumericalBarSeries(name, points, barWidth, color)
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (numChart *PolarNumericalChart) DeleteSeries(name string) {
	numChart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (numChart *PolarNumericalChart) SetTitle(l string) {
	numChart.base.SetTitle(l)
}

// HideLegend hides the legend and uses the full space for the chart
func (numChart *PolarNumericalChart) HideLegend() {
	numChart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (numChart *PolarNumericalChart) ShowLegend() {
	numChart.base.ShowLegend()
}

// SetRAxisLabel sets the label of the r-axis, which will be displayed at the bottom-right
func (numChart *PolarNumericalChart) SetRAxisLabel(l string) {
	numChart.base.SetToAxisLabel(l)
}

// SetRRange sets a user defined range for the r-axis;
// an error is returned if max<0 or if the origin has been defined by the user before and is outside the given range
func (numChart *PolarNumericalChart) SetRRange(max float64) (err error) {
	err = numChart.base.SetToRange(0.0, max)
	return
}

// SetAutoRRange overrides a previously user defined range and lets the range be calculated automatically
func (numChart *PolarNumericalChart) SetAutoRRange() {
	numChart.base.SetAutoToRange()
}

// SetRTicks sets the list of user defined ticks to be shown on the r-axis
func (numChart *PolarNumericalChart) SetRTicks(ts []data.NumericalTick) {
	numChart.base.SetToTicks(ts)
}

// SetAutoRTicks overrides a previously user defined set of r-axis ticks and lets the ticks be calculated automatically
func (numChart *PolarNumericalChart) SetAutoRTicks(autoSupportLine bool) {
	numChart.base.SetAutoToTicks(autoSupportLine)
}

// SetOrigin sets a user defined origin (crossing of phi and r axis).
// An error is returned, if a range has been defined before and at least one coordinate is outside the range.
func (numChart *PolarNumericalChart) SetOrigin(phi float64, r float64) (err error) {
	err = numChart.base.SetNOrigin(phi, r)
	return
}

// SetAutoOrigin resets a previously user defined origin and allows the chart to calculate the ideal origin automatically
func (numChart *PolarNumericalChart) SetAutoOrigin() {
	numChart.base.SetAutoOrigin()
}

// SetPhiAxisLabel sets the label of the phi-axis, which will be displayed at the left side
func (numChart *PolarNumericalChart) SetPhiAxisLabel(l string) {
	numChart.base.SetFromAxisLabel(l)
}

// SetPhiTicks sets the list of user defined ticks to be shown on the phi-axis
func (numChart *PolarNumericalChart) SetPhiTicks(ts []data.NumericalTick) {
	numChart.base.SetFromNTicks(ts)
}

// SetAutoPhiTicks overrides a previously user defined set of phi-axis ticks and lets the ticks be calculated automatically
func (numChart *PolarNumericalChart) SetAutoPhiTicks(autoSupportLine bool) {
	numChart.base.SetAutoFromTicks(autoSupportLine)
}
