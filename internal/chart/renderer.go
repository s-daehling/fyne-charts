package chart

import (
	"fyne.io/fyne/v2"
	"github.com/s-daehling/fyne-charts/internal/series"
)

type baseRenderer struct {
	chart      *BaseChart
	margin     float32
	tickLength float32
}

func emptyBaseRenderer(chart *BaseChart) (r baseRenderer) {
	r = baseRenderer{
		chart:      chart,
		margin:     10.0,
		tickLength: 5.0,
	}
	return
}

// Destroy has nothing to do
func (r *baseRenderer) Destroy() {}

func (r *baseRenderer) placeTitleAndLegend(size fyne.Size) (titleWidth float32, titleHeight float32, legendWidth float32, legendHeight float32) {
	// place title
	ct := r.chart.title()
	if ct.Name != "" {
		ct.Label.Text = ct.Name
		titleWidth = ct.Label.MinSize().Width
		titleHeight = ct.Label.MinSize().Height
		ct.Label.Move(fyne.NewPos(size.Width/2-titleWidth/2, r.margin))
	}

	// place legend
	legendVisible := r.chart.legendVisibility()
	if legendVisible {
		les := r.chart.legendEntries()
		legendWidth, legendHeight = series.LegendSize(les)
		yLegend := (size.Height - legendHeight) / 2.0
		for i := range les {
			subOffset := float32(0.0)
			if les[i].IsSub {
				subOffset = 20
			}
			les[i].Button.Resize(fyne.NewSize(15, 15))
			les[i].Button.Move(fyne.NewPos(size.Width-r.margin-legendWidth+5+subOffset, yLegend+20*float32(i)))
			les[i].Label.Move(fyne.NewPos(size.Width-r.margin-legendWidth+25+subOffset, yLegend+20*float32(i)))
		}
	}
	return
}
