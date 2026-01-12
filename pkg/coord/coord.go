package coord

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/s-daehling/fyne-charts/internal/coord"
	"github.com/s-daehling/fyne-charts/pkg/style"
)

type coordChart struct {
	base *coord.BaseChart
	widget.BaseWidget
}

func emptyCoordChart(planeType coord.PlaneType, fromType coord.FromType) (chart coordChart) {
	chart.base = coord.EmptyBaseChart(planeType, fromType)
	return
}

func (chart *coordChart) CreateRenderer() (r fyne.WidgetRenderer) {
	if chart.base == nil {
		r = widget.NewSimpleRenderer(widget.NewLabel("not initialized"))
		return
	}
	r = widget.NewSimpleRenderer(chart.base.MainContainer())
	return
}

// Refresh chart
// chart is automatically refreshed after data changes
func (chart *coordChart) Refresh() {
	if chart.base == nil {
		return
	}
	chart.base.Refresh()
}

// RemoveSeries deletes the series with the specified name if it exists
func (chart *coordChart) RemoveSeries(name string) {
	if chart.base == nil {
		return
	}
	chart.base.RemoveSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (chart *coordChart) SetTitle(l string) {
	if chart.base == nil {
		return
	}
	chart.base.SetTitle(l)
}

// SetTitleStyle changes the style of the chart title
func (chart *coordChart) SetTitleStyle(ts style.ChartTextStyle) {
	if chart.base == nil {
		return
	}
	chart.base.SetTitleStyle(ts)
}

// HideLegend hides the legend and uses the full space for the chart
func (chart *coordChart) HideLegend() {
	if chart.base == nil {
		return
	}
	chart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (chart *coordChart) ShowLegend() {
	if chart.base == nil {
		return
	}
	chart.base.ShowLegend()
}

// SetLegendStyle changes the style of the chart legend
func (chart *coordChart) SetLegendStyle(loc style.LegendLocation, labelStyle style.ChartTextStyle, interactive bool) {
	if chart.base == nil {
		return
	}
	chart.base.SetLegendStyle(loc, labelStyle, interactive)
}
