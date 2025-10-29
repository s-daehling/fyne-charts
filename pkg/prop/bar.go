package prop

import (
	"github.com/s-daehling/fyne-charts/internal/prop"

	"fyne.io/fyne/v2/widget"
)

// BarChart implements a cartesian plane with one proportional axis
type BarChart struct {
	propChart
	widget.BaseWidget
}

// NewBarChart returns an initialized CartesianProportionalChart
func NewBarChart() (barChart *BarChart) {
	barChart = &BarChart{
		propChart: emptyPropChart(prop.CartesianPlane),
	}
	barChart.ExtendBaseWidget(barChart)
	return
}
