package prop

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/s-daehling/fyne-charts/internal/prop"
	"github.com/s-daehling/fyne-charts/pkg/style"
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
	if chart.base == nil {
		r = widget.NewSimpleRenderer(widget.NewLabel("not initialized"))
		return
	}
	r = widget.NewSimpleRenderer(chart.base.MainContainer())
	return
}

// Refresh chart
// chart is automatically refreshed after data changes
func (chart *propChart) Refresh() {
	if chart.base == nil {
		return
	}
	chart.base.Refresh()
}

// AddSeries adds a series of data
// The series must have a unique name throughout the chart.
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (chart *propChart) AddSeries(ps *Series) (err error) {
	if chart.base == nil || ps == nil {
		return
	}
	if ps.ser == nil {
		err = errors.New("series not initialized")
		return
	}
	err = chart.base.AddSeries(ps.ser)
	return
}

// RemoveSeries deletes the series with the specified name if it exists
func (chart *propChart) RemoveSeries(name string) {
	if chart.base == nil {
		return
	}
	chart.base.RemoveSeries(name)
}

// SetTitle sets the title of the chart, which will be displayed at the top
func (chart *propChart) SetTitle(l string) {
	if chart.base == nil {
		return
	}
	chart.base.SetTitle(l)
}

// SetTitleStyle changes the style of the chart title
// default value title size: theme.SizeNameSubHeadingText
// default value title color: theme.ColorNameForeground
func (chart *propChart) SetTitleStyle(titleSize fyne.ThemeSizeName, titleColor fyne.ThemeColorName) {
	if chart.base == nil {
		return
	}
	chart.base.SetTitleStyle(titleSize, titleColor)
}

// HideLegend hides the legend and uses the full space for the chart
func (chart *propChart) HideLegend() {
	if chart.base == nil {
		return
	}
	chart.base.HideLegend()
}

// ShowLegend shows the legend on the right side
func (chart *propChart) ShowLegend() {
	if chart.base == nil {
		return
	}
	chart.base.ShowLegend()
}

// SetLegendStyle changes the style of the chart legend
func (chart *propChart) SetLegendStyle(loc style.LegendLocation) {
	if chart.base == nil {
		return
	}
	chart.base.SetLegendStyle(loc)
}
