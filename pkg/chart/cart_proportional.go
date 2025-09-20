package chart

import (
	"github.com/s-daehling/fyne-charts/internal/chart"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// CartesianProportionalChart implements a cartesian plane with one proportional axis
type CartesianProportionalChart struct {
	base *chart.BaseChart
	widget.BaseWidget
}

// NewCartesianProportionalChart returns an initialized CartesianProportionalChart
func NewCartesianProportionalChart() (propChart *CartesianProportionalChart) {
	propChart = &CartesianProportionalChart{
		base: chart.EmptyBaseChart(chart.CartesianPlane, chart.Proportional),
	}
	propChart.ExtendBaseWidget(propChart)
	return
}

// CreateRenderer creates the renderer of the widget
func (propChart *CartesianProportionalChart) CreateRenderer() fyne.WidgetRenderer {
	r := propChart.base.GetRenderer()
	return r
}

// AddProportionalSeries adds a series of data which is visualized as proportional bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0
func (propChart *CartesianProportionalChart) AddProportionalSeries(name string,
	points []data.ProportionalDataPoint) (ps ProportionalSeries, err error) {
	ps.ser, err = propChart.base.AddProportionalSeries(name, points, nil)
	return
}

// AddProportionalSeriesWithProvider adds a series of data which is visualized as proportional bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The series data is retrieved from providerFct
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0
func (propChart *CartesianProportionalChart) AddProportionalSeriesWithProvider(name string,
	providerFct func() []data.ProportionalDataPoint) (ps ProportionalSeries, err error) {
	ps.ser, err = propChart.base.AddProportionalSeries(name, nil, providerFct)
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (propChart *CartesianProportionalChart) DeleteSeries(name string) {
	propChart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (propChart *CartesianProportionalChart) SetTitle(l string) {
	propChart.base.SetTitle(l)
}

// HideLegend hides the legend and uses the full space for the chart
func (propChart *CartesianProportionalChart) HideLegend() {
	propChart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (propChart *CartesianProportionalChart) ShowLegend() {
	propChart.base.ShowLegend()
}
