package prop

import (
	"github.com/s-daehling/fyne-charts/internal/chart"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// BarChart implements a cartesian plane with one proportional axis
type BarChart struct {
	base *chart.BaseChart
	widget.BaseWidget
}

// NewBarChart returns an initialized CartesianProportionalChart
func NewBarChart() (propChart *BarChart) {
	propChart = &BarChart{
		base: chart.EmptyBaseChart(chart.CartesianPlane, chart.Proportional),
	}
	propChart.ExtendBaseWidget(propChart)
	return
}

// CreateRenderer creates the renderer of the widget
func (propChart *BarChart) CreateRenderer() fyne.WidgetRenderer {
	r := propChart.base.GetRenderer()
	return r
}

// AddSeries adds a series of data which is visualized as proportional bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// Data points with a C that already exists, will be ignored.
// The range of C is not restricted. The range of Val is restricted to Val>=0
func (propChart *BarChart) AddSeries(name string,
	points []data.ProportionalDataPoint) (ps ProportionalSeries, err error) {
	ps.ser, err = propChart.base.AddProportionalSeries(name, points)
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (propChart *BarChart) DeleteSeries(name string) {
	propChart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (propChart *BarChart) SetTitle(l string) {
	propChart.base.SetTitle(l)
}

// HideLegend hides the legend and uses the full space for the chart
func (propChart *BarChart) HideLegend() {
	propChart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (propChart *BarChart) ShowLegend() {
	propChart.base.ShowLegend()
}
