package chart

import (
	"github.com/s-daehling/fyne-charts/internal/chart"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// PolarProportionalChart implements a polar plane with one proportional axis
type PolarProportionalChart struct {
	base *chart.BaseChart
	widget.BaseWidget
}

// NewPolarProportionalChart returns an initialized PolarProportionalChart
func NewPolarProportionalChart() (propChart *PolarProportionalChart) {
	propChart = &PolarProportionalChart{
		base: chart.EmptyBaseChart(chart.PolarPlane, chart.Proportional),
	}
	propChart.ExtendBaseWidget(propChart)
	return
}

// CreateRenderer creates the renderer of the widget
func (propChart *PolarProportionalChart) CreateRenderer() fyne.WidgetRenderer {
	r := chart.EmptyPolarRenderer(propChart.base)
	return r
}

// AddProportionalSeries adds a series of data which is visualized as proportional bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same C)
// The range of C is not restricted. The range of Val is restricted to Val>=0
func (propChart *PolarProportionalChart) AddProportionalSeries(name string,
	points []data.ProportionalDataPoint) (ps ProportionalSeries, err error) {
	ps.ser, err = propChart.base.AddProportionalSeries(name, points)
	ps.wid = &propChart.BaseWidget
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (propChart *PolarProportionalChart) DeleteSeries(name string) {
	propChart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (propChart *PolarProportionalChart) SetTitle(l string) {
	propChart.base.SetTitle(l)
}

// HideLegend hides the legend and uses the full space for the chart
func (propChart *PolarProportionalChart) HideLegend() {
	propChart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (propChart *PolarProportionalChart) ShowLegend() {
	propChart.base.ShowLegend()
}
