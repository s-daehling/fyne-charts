package chart

import (
	"image/color"

	"github.com/s-daehling/fyne-charts/internal/chart"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// PolarAngularChart implements a polar plane with one angular phi-axis and one numerical r-axis
type PolarAngularChart struct {
	base *chart.BaseChart
	widget.BaseWidget
}

// NewPolarAngularChart returns an initialized PolarAngularChart
func NewPolarAngularChart() (angChart *PolarAngularChart) {
	angChart = &PolarAngularChart{
		base: chart.EmptyBaseChart(chart.PolarPlane, chart.Angular),
	}
	angChart.ExtendBaseWidget(angChart)
	return
}

// CreateRenderer creates the renderer of the widget
func (angChart *PolarAngularChart) CreateRenderer() fyne.WidgetRenderer {
	r := chart.EmptyPolarRenderer(angChart.base)
	return r
}

// AddLineSeries adds a series of data which is visualized as line chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// data does not need to be sorted. It will be sorted by A by the method.
// The method does not check for duplicates (i.e. data points with same A)
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
func (angChart *PolarAngularChart) AddLineSeries(name string, points []data.AngularDataPoint, showDots bool,
	color color.Color) (als AngularLineSeries, err error) {
	als.ser, err = angChart.base.AddAngularLineSeries(name, points, showDots, color)
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same A)
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
func (angChart *PolarAngularChart) AddScatterSeries(name string, points []data.AngularDataPoint,
	color color.Color) (ass AngularScatterSeries, err error) {
	ass.ser, err = angChart.base.AddAngularScatterSeries(name, points, color)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same A)
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
func (angChart *PolarAngularChart) AddLollipopSeries(name string, points []data.AngularDataPoint,
	color color.Color) (als AngularLollipopSeries, err error) {
	als.ser, err = angChart.base.AddAngularLollipopSeries(name, points, color)
	return
}

// AddAreaSeries adds a series of data which is visualized as area chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// data does not need to be sorted. It will be sorted by A by the method.
// The method does not check for duplicates (i.e. data points with same A).
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
func (angChart *PolarAngularChart) AddAreaSeries(name string, points []data.AngularDataPoint, showDots bool,
	color color.Color) (aas AngularAreaSeries, err error) {
	aas.ser, err = angChart.base.AddAngularAreaSeries(name, points, showDots, color)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same A)
// The range of A and Val is restricted (0<=A<=2pi; Val>0)
// The bars are centered around their A value of the data points. barWidth is the width of the bars.
// An error is returned if barWidth < 0
func (angChart *PolarAngularChart) AddBarSeries(name string, points []data.AngularDataPoint,
	barWidth float64, color color.Color) (abs AngularBarSeries, err error) {
	abs.ser, err = angChart.base.AddAngularBarSeries(name, points, barWidth, color)
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (angChart *PolarAngularChart) DeleteSeries(name string) {
	angChart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (angChart *PolarAngularChart) SetTitle(l string) {
	angChart.base.SetTitle(l)
}

// HideLegend hides the legend and uses the full space for the chart
func (angChart *PolarAngularChart) HideLegend() {
	angChart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (angChart *PolarAngularChart) ShowLegend() {
	angChart.base.ShowLegend()
}

// SetRAxisLabel sets the label of the r-axis, which will be displayed at the bottom-right
func (angChart *PolarAngularChart) SetRAxisLabel(l string) {
	angChart.base.SetToAxisLabel(l)
}

// SetRRange sets a user defined range for the r-axis;
// an error is returned if max<0 or if the origin has been defined by the user before and is outside the given range
func (angChart *PolarAngularChart) SetRRange(max float64) (err error) {
	err = angChart.base.SetToRange(0.0, max)
	return
}

// SetAutoRRange overrides a previously user defined range and lets the range be calculated automatically
func (angChart *PolarAngularChart) SetAutoRRange() {
	angChart.base.SetAutoToRange()
}

// SetRTicks sets the list of user defined ticks to be shown on the r-axis
func (angChart *PolarAngularChart) SetRTicks(ts []data.NumericalTick) {
	angChart.base.SetToTicks(ts)
}

// SetAutoRTicks overrides a previously user defined set of r-axis ticks and lets the ticks be calculated automatically
func (angChart *PolarAngularChart) SetAutoRTicks(autoSupportLine bool) {
	angChart.base.SetAutoToTicks(autoSupportLine)
}

// SetOrigin sets a user defined origin (crossing of phi and r axis).
// An error is returned, if a range has been defined before and at least one coordinate is outside the range.
func (angChart *PolarAngularChart) SetOrigin(phi float64, r float64) (err error) {
	err = angChart.base.SetNOrigin(phi, r)
	return
}

// SetAutoOrigin resets a previously user defined origin and allows the chart to calculate the ideal origin automatically
func (angChart *PolarAngularChart) SetAutoOrigin() {
	angChart.base.SetAutoOrigin()
}

// SetPhiAxisLabel sets the label of the phi-axis, which will be displayed at the left side
func (angChart *PolarAngularChart) SetPhiAxisLabel(l string) {
	angChart.base.SetFromAxisLabel(l)
}

// SetPhiTicks sets the list of user defined ticks to be shown on the phi-axis
func (angChart *PolarAngularChart) SetPhiTicks(ts []data.AngularTick) {
	angChart.base.SetFromATicks(ts)
}

// SetAutoPhiTicks overrides a previously user defined set of phi-axis ticks and lets the ticks be calculated automatically
func (angChart *PolarAngularChart) SetAutoPhiTicks(autoSupportLine bool) {
	angChart.base.SetAutoFromTicks(autoSupportLine)
}
