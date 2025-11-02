package coord

import (
	"image/color"

	"fyne.io/fyne/v2"
	"github.com/s-daehling/fyne-charts/internal/coord"
)

type coordChart struct {
	base *coord.BaseChart
}

func emptyCoordChart(planeType coord.PlaneType, fromType coord.FromType) (chart coordChart) {
	chart.base = coord.EmptyBaseChart(planeType, fromType)
	return
}

// CreateRenderer creates the renderer of the widget
func (chart *coordChart) CreateRenderer() fyne.WidgetRenderer {
	r := chart.base.GetRenderer()
	return r
}

// DeleteSeries deletes the series with the specified name if it exists
func (chart *coordChart) DeleteSeries(name string) {
	chart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (chart *coordChart) SetTitle(l string) {
	chart.base.SetTitle(l)
}

// SetTitleColor changes the color of the chart title
func (chart *coordChart) SetTitleColor(col color.Color) {
	chart.base.SetTitleColor(col)
}

// SetAutoTitleColor changes the color of the chart title back to the default (theme.ColorNameForeground)
func (chart *coordChart) SetAutoTitleColor() {
	chart.base.SetAutoTitleColor()
}

// SetTitleSize changes the size of the chart title
func (chart *coordChart) SetTitleSize(size float32) {
	chart.base.SetTitleSize(size)
}

// SetAutoTitleSize changes the size of the chart title back to the default (theme.SizeNameHeadingText)
func (chart *coordChart) SetAutoTitleSize() {
	chart.base.SetAutoTitleSize()
}

// SetAxisLabelColor changes the color of the chart axis labels
func (chart *coordChart) SetAxisLabelColor(col color.Color) {
	chart.base.SetAxisLabelColor(col)
}

// SetAutoAxisLabelColor changes the color of the chart axis labels back to the default (theme.ColorNameForeground)
func (chart *coordChart) SetAutoAxisLabelColor() {
	chart.base.SetAutoAxisLabelColor()
}

// SetAxisLabelSize changes the size of the chart axis labels
func (chart *coordChart) SetAxisLabelSize(size float32) {
	chart.base.SetAxisLabelSize(size)
}

// SetAutoAxisLabelSize changes the size of the chart axis labels back to the default (theme.SizeNameSubHeadingText)
func (chart *coordChart) SetAutoAxisLabelSize() {
	chart.base.SetAutoAxisLabelSize()
}

// HideLegend hides the legend and uses the full space for the chart
func (chart *coordChart) HideLegend() {
	chart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (chart *coordChart) ShowLegend() {
	chart.base.ShowLegend()
}
