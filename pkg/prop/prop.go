package prop

import (
	"fyne.io/fyne/v2"
	"github.com/s-daehling/fyne-charts/internal/prop"
	"github.com/s-daehling/fyne-charts/pkg/data"
)

type propChart struct {
	*prop.BaseChart
}

func emptyPropChart(planeType prop.PlaneType) (chart propChart) {
	chart.BaseChart = prop.EmptyBaseChart(planeType)
	return
}

// AddSeries adds a series of data which is visualized as proportional bar chart.
// The series can be accessed via the name later, it must be unique throughout the chart.
// An error is returned,if another series with the same name exists.
// The method does not check for duplicates (i.e. data points with same C)
// The range of C is not restricted. The range of Val is restricted to Val>=0
func (chart *propChart) AddSeries(name string,
	points []data.ProportionalPoint) (ps ProportionalSeries, err error) {
	ps.ser, err = chart.BaseChart.AddProportionalSeries(name, points)
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (chart *propChart) DeleteSeries(name string) {
	chart.BaseChart.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (chart *propChart) SetTitle(l string) {
	chart.BaseChart.SetTitle(l)
}

// SetTitleStyle changes the style of the chart title
// default value title size: theme.SizeNameSubHeadingText
// default value title color: theme.ColorNameForeground
func (chart *propChart) SetTitleStyle(titleSize fyne.ThemeSizeName, titleColor fyne.ThemeColorName) {
	chart.BaseChart.SetTitleStyle(titleSize, titleColor)
}

// HideLegend hides the legend and uses the full space for the chart
func (chart *propChart) HideLegend() {
	chart.BaseChart.HideLegend()
}

// ShowLegend shows the legend on the right side
func (chart *propChart) ShowLegend() {
	chart.BaseChart.ShowLegend()
}
