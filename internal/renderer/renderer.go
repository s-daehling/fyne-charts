package renderer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/s-daehling/fyne-charts/internal/interact"
)

type baseChart interface {
	Title() (ct *canvas.Text)
	LegendEntries() (le []LegendEntry)
	Tooltip() (tt Tooltip)
	FromAxisElements() (min float64, max float64, origin float64, label Label, ticks []Tick, arrow Arrow, show bool)
	ToAxisElements() (min float64, max float64, origin float64, label Label, ticks []Tick, arrow Arrow, show bool)
	Raster() (r *canvas.Raster)
	Overlay() (io *interact.Overlay)
	ChartSizeChange(fromSpace float32, toSpace float32)
	RefreshTheme()
}

type baseRenderer struct {
	margin     float32
	tickLength float32
	widgetSize func() fyne.Size
}

func emptyBaseRenderer(ws func() fyne.Size) (r baseRenderer) {
	r = baseRenderer{
		margin:     10.0,
		tickLength: 5.0,
		widgetSize: ws,
	}
	return
}

// Destroy has nothing to do
func (r *baseRenderer) Destroy() {}

func (r *baseRenderer) placeTitleAndLegend(size fyne.Size, ct *canvas.Text,
	les []LegendEntry) (titleWidth float32, titleHeight float32, legendWidth float32,
	legendHeight float32) {
	// place title
	if ct != nil {
		if ct.Text != "" {
			titleWidth = ct.MinSize().Width
			titleHeight = ct.MinSize().Height
			ct.Move(fyne.NewPos(size.Width/2-titleWidth/2, r.margin))
		}
	}

	// place legend
	if len(les) > 0 {
		legendWidth, legendHeight = legendSize(les)
		yLegend := titleHeight + (size.Height-titleHeight-legendHeight)/2.0
		for i := range les {
			subOffset := float32(0.0)
			if les[i].IsSub {
				subOffset = 20
			}
			labelOffset := float32(0.0)
			if les[i].ShowButton {
				labelOffset = 20
				les[i].Button.Resize(fyne.NewSize(15, 15))
				les[i].Button.Move(fyne.NewPos(size.Width-r.margin-legendWidth+5+subOffset, yLegend+20*float32(i)))
			}
			les[i].Label.Move(fyne.NewPos(size.Width-r.margin-legendWidth+5+labelOffset+subOffset, yLegend+20*float32(i)))
		}
	}
	return
}
