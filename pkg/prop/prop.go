package prop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/s-daehling/fyne-charts/internal/prop"
)

type propChart struct {
	base *prop.BaseChart
	widget.BaseWidget
}

func emptyPropChart(planeType prop.PlaneType) (chart propChart) {
	chart.base = prop.EmptyBaseChart(planeType)
	return
}

func (chart *propChart) CreateRenderer() (r fyne.WidgetRenderer) {
	r = chart.base.CreateRenderer(chart.Size)
	return
}

// Refresh chart
// chart is automatically refreshed after data changes
func (chart *propChart) Refresh() {
	chart.base.Refresh()
}

// AddSeries adds a series of data
// The series must have a unique name throughout the chart.
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (chart *propChart) AddSeries(ps *Series) (err error) {
	err = chart.base.AddSeries(ps.ser)
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

// SetTitleStyle changes the style of the chart title
// default value title size: theme.SizeNameSubHeadingText
// default value title color: theme.ColorNameForeground
func (chart *propChart) SetTitleStyle(titleSize fyne.ThemeSizeName, titleColor fyne.ThemeColorName) {
	chart.base.SetTitleStyle(titleSize, titleColor)
}

// HideLegend hides the legend and uses the full space for the chart
func (chart *propChart) HideLegend() {
	chart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (chart *propChart) ShowLegend() {
	chart.base.ShowLegend()
}
