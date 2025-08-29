package series

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
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

type LegendBox struct {
	widget.BaseWidget
	color  color.Color
	rect   *canvas.Rectangle
	tapFct func()
}

func NewLegendBox(col color.Color, tapFct func()) *LegendBox {
	box := &LegendBox{
		rect:   canvas.NewRectangle(col),
		color:  col,
		tapFct: tapFct,
	}
	box.ExtendBaseWidget(box)
	return box
}

func (box *LegendBox) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewStack(box.rect)
	return widget.NewSimpleRenderer(c)
}

func (box *LegendBox) Tapped(_ *fyne.PointEvent) {
	box.tapFct()
}

func (box *LegendBox) MouseIn(me *desktop.MouseEvent) {
	r, g, b, _ := box.color.RGBA()
	box.rect.FillColor = color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: 0xaaaa}
	box.rect.Refresh()
}

func (box *LegendBox) MouseMoved(me *desktop.MouseEvent) {}

func (box *LegendBox) MouseOut() {
	box.rect.FillColor = box.color
	box.rect.Refresh()
}

func (box *LegendBox) SetColor(col color.Color) {
	box.color = col
	box.rect.FillColor = col
}

type LegendEntry struct {
	Button *LegendBox
	Label  *canvas.Text
	IsSub  bool
}

func LegendWidth(les []LegendEntry) (w float32) {
	w = 0.0
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
	return
}
