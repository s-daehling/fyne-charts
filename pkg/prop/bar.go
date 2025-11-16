package prop

import (
	"github.com/s-daehling/fyne-charts/internal/prop"
)

// BarChart implements a cartesian plane with one proportional axis
type BarChart struct {
	propChart
}

// NewBarChart returns an initialized BarChart
func NewBarChart(title string) (barChart *BarChart) {
	barChart = &BarChart{
		propChart: emptyPropChart(prop.CartesianPlane),
	}
	barChart.ExtendBaseWidget(barChart)
	barChart.SetTitle(title)
	return
}
