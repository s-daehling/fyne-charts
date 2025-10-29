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
func (chart *coordChart) SetAutoTitleColor(col color.Color) {
	chart.base.SetAutoTitleColor()
}

// SetTitleSize changes the size of the chart title
func (chart *coordChart) SetTitleSize(size float32) {
	chart.base.SetTitleSize(size)
}

// SetAutoTitleSize changes the size of the chart title back to the default (theme.SizeNameSubHeadingText)
func (chart *coordChart) SetAutoTitleSize() {
	chart.base.SetAutoTitleSize()
}

// HideLegend hides the legend and uses the full space for the chart
func (chart *coordChart) HideLegend() {
	chart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (chart *coordChart) ShowLegend() {
	chart.base.ShowLegend()
}
