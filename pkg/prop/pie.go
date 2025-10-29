package prop

import (
	"github.com/s-daehling/fyne-charts/internal/prop"

	"fyne.io/fyne/v2/widget"
)

// PieChart implements a polar plane with one proportional axis
type PieChart struct {
	propChart
	widget.BaseWidget
}

// NewPieChart returns an initialized PolarProportionalChart
func NewPieChart() (pieChart *PieChart) {
	pieChart = &PieChart{
		propChart: emptyPropChart(prop.PolarPlane),
	}
	pieChart.ExtendBaseWidget(pieChart)
	return
}
