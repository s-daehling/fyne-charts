package legend

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

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
