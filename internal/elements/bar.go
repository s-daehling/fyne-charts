package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Bar struct {
	N1   float64
	Val1 float64
	N2   float64
	Val2 float64
	rect *canvas.Rectangle
	col  color.Color
	widget.BaseWidget
}

func NewBar(col color.Color) (b *Bar) {
	b = &Bar{
		col:  col,
		rect: canvas.NewRectangle(col),
	}
	b.ExtendBaseWidget(b)
	return
}

func (b *Bar) SetColor(col color.Color) {
	b.rect.FillColor = col
}

func (b *Bar) SetCornerRadius(r float32) {
	b.rect.CornerRadius = r
}

func (b *Bar) CreateRenderer() (r fyne.WidgetRenderer) {
	r = newBarRenderer(b)
	return
}

type barRenderer struct {
	bar *Bar
}

func newBarRenderer(b *Bar) (br *barRenderer) {
	br = &barRenderer{
		bar: b,
	}
	return
}

func (br *barRenderer) Layout(size fyne.Size) {
	br.bar.rect.Resize(size)
	br.bar.rect.Move(fyne.NewPos(0, 0))
}

func (br *barRenderer) MinSize() (size fyne.Size) {
	size = fyne.NewSize(0, 0)
	return
}

func (br *barRenderer) Refresh() {
	br.bar.rect.Refresh()
}

func (br *barRenderer) Objects() (canObj []fyne.CanvasObject) {
	canObj = append(canObj, br.bar.rect)
	return
}

func (br *barRenderer) Destroy() {}
