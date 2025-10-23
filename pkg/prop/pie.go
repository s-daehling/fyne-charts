package prop

import (
	"github.com/s-daehling/fyne-charts/internal/prop"
	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// PieChart implements a polar plane with one proportional axis
type PieChart struct {
	base *prop.BaseChart
	widget.BaseWidget
}

// NewPieChart returns an initialized PolarProportionalChart
func NewPieChart() (propChart *PieChart) {
	propChart = &PieChart{
		base: prop.EmptyBaseChart(prop.PolarPlane),
	}
	propChart.ExtendBaseWidget(propChart)
	return
}

// CreateRenderer creates the renderer of the widget
func (propChart *PieChart) CreateRenderer() fyne.WidgetRenderer {
	r := propChart.base.GetRenderer()
	return r
}

// AddSeries adds a series of data which is visualized as proportional bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same C)
// The range of C is not restricted. The range of Val is restricted to Val>=0
func (propChart *PieChart) AddSeries(name string,
	points []data.ProportionalDataPoint) (ps ProportionalSeries, err error) {
	ps.ser, err = propChart.base.AddProportionalSeries(name, points)
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (propChart *PieChart) DeleteSeries(name string) {
	propChart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (propChart *PieChart) SetTitle(l string) {
	propChart.base.SetTitle(l)
}

// HideLegend hides the legend and uses the full space for the chart
func (propChart *PieChart) HideLegend() {
	propChart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (propChart *PieChart) ShowLegend() {
	propChart.base.ShowLegend()
}
