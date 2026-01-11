package renderer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/s-daehling/fyne-charts/internal/interact"
)

type baseChart interface {
	Tooltip() (tt Tooltip)
	FromAxisElements() (min float64, max float64, origin float64, ticks []Tick, arrow Arrow, show bool)
	ToAxisElements() (min float64, max float64, origin float64, ticks []Tick, arrow Arrow, show bool)
	Raster() (r *canvas.Raster)
	Overlay() (io *interact.Overlay)
	ChartSizeChange(fromSpace float32, toSpace float32)
	RefreshTheme()
	Size() (size fyne.Size)
}

type baseRenderer struct {
	margin     float32
	tickLength float32
}

func emptyBaseRenderer() (r baseRenderer) {
	r = baseRenderer{
		margin:     10.0,
		tickLength: 5.0,
	}
	return
}

// Destroy has nothing to do
func (r *baseRenderer) Destroy() {}
