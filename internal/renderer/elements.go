package renderer

import (
	"fyne.io/fyne/v2/canvas"
	"github.com/s-daehling/fyne-charts/internal/interact"
)

type CartesianNode struct {
	X   float64
	Y   float64
	Dot *canvas.Circle
}

type CartesianEdge struct {
	X1   float64
	Y1   float64
	X2   float64
	Y2   float64
	Line *canvas.Line
}

type CartesianRect struct {
	X1   float64
	Y1   float64
	X2   float64
	Y2   float64
	Rect *canvas.Rectangle
}

type CartesianText struct {
	X    float64
	Y    float64
	Text *canvas.Text
}

type PolarNode struct {
	Phi float64
	R   float64
	Dot *canvas.Circle
}

type PolarEdge struct {
	Phi1 float64
	R1   float64
	Phi2 float64
	R2   float64
	Line *canvas.Line
}

type PolarText struct {
	Phi  float64
	R    float64
	Text *canvas.Text
}

type Label struct {
	Text  *canvas.Text
	Image *canvas.Image
}

type Tick struct {
	NLabel    float64
	Label     *canvas.Text
	NLine     float64
	Line      *canvas.Line
	SupLine   *canvas.Line
	SupCircle *canvas.Circle
}

type Arrow struct {
	Line    *canvas.Line
	Circle  *canvas.Circle
	HeadOne *canvas.Line
	HeadTwo *canvas.Line
}

func maxTickSize(ts []Tick) (maxWidth float32, maxHeight float32) {
	maxWidth = 0
	maxHeight = 0
	for i := range ts {
		if ts[i].Label.MinSize().Width > maxWidth {
			maxWidth = ts[i].Label.MinSize().Width
		}
		if ts[i].Label.MinSize().Height > maxHeight {
			maxHeight = ts[i].Label.MinSize().Height
		}
	}
	return
}

type LegendEntry struct {
	Button *interact.LegendBox
	Label  *canvas.Text
	IsSub  bool
}

func legendSize(les []LegendEntry) (w float32, h float32) {
	w = 0.0
	h = 0.0
	if len(les) == 0 {
		return
	}
	hasSubs := false
	for i := range les {
		if les[i].Label.MinSize().Width > float32(w) {
			w = les[i].Label.MinSize().Width
		}
		if les[i].IsSub {
			hasSubs = true
		}
	}
	w += 25
	if hasSubs {
		w += 20
	}
	h = float32(len(les) * 20)
	return
}

type CartesianTooltip struct {
	X       float64
	Y       float64
	Entries []*canvas.Text
}

type PolarTooltip struct {
	Phi     float64
	R       float64
	Entries []*canvas.Text
}

func tooltipSize(entries []*canvas.Text) (w float32, h float32) {
	w = 0.0
	h = 0.0
	if len(entries) == 0 {
		return
	}
	w = entries[0].MinSize().Width
	for i := range entries {
		h += entries[i].MinSize().Height + 2
		if entries[i].MinSize().Width > w {
			w = entries[i].MinSize().Width
		}
	}
	return
}
