package coord

import "github.com/s-daehling/fyne-charts/internal/coord"

type coordChart struct {
	base *coord.BaseChart
}

func emptyCoordChart(planeType coord.PlaneType, fromType coord.FromType) (chart coordChart) {
	chart.base = coord.EmptyBaseChart(planeType, fromType)
	return
}

// DeleteSeries deletes the series with the specified name if it exists
func (chart *coordChart) DeleteSeries(name string) {
	chart.base.DeleteSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (chart *coordChart) SetTitle(l string) {
	chart.base.SetTitle(l)
}

// HideLegend hides the legend and uses the full space for the chart
func (chart *coordChart) HideLegend() {
	chart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (chart *coordChart) ShowLegend() {
	chart.base.ShowLegend()
}
