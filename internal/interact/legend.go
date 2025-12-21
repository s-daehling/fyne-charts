package interact

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type LegendBox struct {
	widget.BaseWidget
	rectColor      color.Color
	rect           *canvas.Rectangle
	grad           *canvas.RadialGradient
	gradStartColor color.Color
	gradEndColor   color.Color
	showGrad       bool
	tapFct         func()
}

func NewLegendBox(rectColor color.Color, tapFct func()) *LegendBox {
	box := &LegendBox{
		rect:           canvas.NewRectangle(rectColor),
		grad:           canvas.NewRadialGradient(theme.Color(theme.ColorNameForeground), theme.Color(theme.ColorNameBackground)),
		showGrad:       false,
		rectColor:      rectColor,
		gradStartColor: theme.Color(theme.ColorNameForeground),
		gradEndColor:   theme.Color(theme.ColorNameBackground),
		tapFct:         tapFct,
	}
	box.grad.Hide()
	box.ExtendBaseWidget(box)
	return box
}

func (box *LegendBox) UseGradient(startColor color.Color, endColor color.Color) {
	box.rect.Hide()
	box.grad.Show()
	box.gradStartColor = startColor
	box.grad.StartColor = startColor
	box.gradEndColor = endColor
	box.grad.EndColor = endColor
	box.grad.Refresh()
}

func (box *LegendBox) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewStack(box.rect, box.grad)
	return widget.NewSimpleRenderer(c)
}

func (box *LegendBox) Tapped(_ *fyne.PointEvent) {
	box.tapFct()
}

func (box *LegendBox) MouseIn(me *desktop.MouseEvent) {
	r, g, b, a := box.rectColor.RGBA()
	rb, gb, bb, _ := theme.Color(theme.ColorNameBackground).RGBA()
	// box.rect.FillColor = color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: 0xaaaa}
	box.rect.FillColor = color.RGBA64{R: uint16(float32(r+rb) * 0.5), G: uint16(float32(g+gb) * 0.5), B: uint16(float32(b+bb) * 0.5), A: uint16(a)}
	box.rect.Refresh()
	box.grad.StartColor = box.gradEndColor
	box.grad.EndColor = box.gradStartColor
	box.grad.Refresh()
}

func (box *LegendBox) MouseMoved(me *desktop.MouseEvent) {}

func (box *LegendBox) MouseOut() {
	box.rect.FillColor = box.rectColor
	box.rect.Refresh()
	box.grad.StartColor = box.gradStartColor
	box.grad.EndColor = box.gradEndColor
	box.grad.Refresh()
}

func (box *LegendBox) SetRectColor(col color.Color) {
	box.rectColor = col
	box.rect.FillColor = col
}

func (box *LegendBox) SetGradColor(startColor color.Color, endColor color.Color) {
	box.gradStartColor = startColor
	box.grad.StartColor = startColor
	box.gradEndColor = endColor
	box.grad.EndColor = endColor
}
