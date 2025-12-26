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

// SetOrientation defines the orientation of axes
// if transposed is false, the bar are oriented horizontally
// if transposed is true, the bars are oriented vertically
func (barChart *BarChart) SetOrientation(transposed bool) {
	barChart.base.SetCartesianOrientantion(transposed)
}
