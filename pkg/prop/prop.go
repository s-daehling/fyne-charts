package prop

import (
	"fyne.io/fyne/v2"
	"github.com/s-daehling/fyne-charts/internal/prop"
	"github.com/s-daehling/fyne-charts/pkg/data"
)

type propChart struct {
	base *prop.BaseChart
}

func emptyPropChart(planeType prop.PlaneType) (chart propChart) {
	chart.base = prop.EmptyBaseChart(planeType)
	return
}

// CreateRenderer creates the renderer of the widget
func (chart *propChart) CreateRenderer() fyne.WidgetRenderer {
	r := chart.base.GetRenderer()
	return r
}

// AddSeries adds a series of data which is visualized as proportional bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same C)
// The range of C is not restricted. The range of Val is restricted to Val>=0
func (chart *propChart) AddSeries(name string,
	points []data.ProportionalDataPoint) (ps ProportionalSeries, err error) {
	ps.ser, err = chart.base.AddProportionalSeries(name, points)
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (chart *propChart) DeleteSeries(name string) {
	chart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (chart *propChart) SetTitle(l string) {
	chart.base.SetTitle(l)
}

// HideLegend hides the legend and uses the full space for the chart
func (chart *propChart) HideLegend() {
	chart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (chart *propChart) ShowLegend() {
	chart.base.ShowLegend()
}
