package prop

import (
	"github.com/s-daehling/fyne-charts/internal/prop"
)

// PieChart implements a polar plane with one proportional axis
type PieChart struct {
	propChart
}

// NewPieChart returns an initialized PieChart
func NewPieChart() (pieChart *PieChart) {
	pieChart = &PieChart{
		propChart: emptyPropChart(prop.PolarPlane),
	}
	pieChart.ExtendBaseWidget(pieChart)
	return
}
