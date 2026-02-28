package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type Dot struct {
	N           float64
	Val         float64
	size        float32
	circle      *canvas.Circle
	highlight   func()
	unhighlight func()
	widget.BaseWidget
}

func NewDot(col color.Color, size float32, highlight func(), unhighlight func()) (d *Dot) {
	d = &Dot{
		size:        size,
		circle:      canvas.NewCircle(col),
		highlight:   highlight,
		unhighlight: unhighlight,
	}
	d.ExtendBaseWidget(d)
	return
}

func (d *Dot) SetMinSize(size float32) {
	d.size = size
}

func (d *Dot) SetColor(col color.Color) {
	d.circle.FillColor = col
}

func (d *Dot) MouseIn(me *desktop.MouseEvent) {
	if d.highlight == nil {
		return
	}
	d.highlight()
}

func (d *Dot) MouseMoved(me *desktop.MouseEvent) {

}

func (d *Dot) MouseOut() {
	if d.unhighlight == nil {
		return
	}
	d.unhighlight()
}

func (d *Dot) CreateRenderer() (r fyne.WidgetRenderer) {
	r = newDotRenderer(d)
	return
}

type dotRenderer struct {
	dot *Dot
}

func newDotRenderer(d *Dot) (dr *dotRenderer) {
	dr = &dotRenderer{
		dot: d,
	}
	return
}

func (dr *dotRenderer) Layout(size fyne.Size) {
	dr.dot.circle.Resize(size)
	dr.dot.circle.Move(fyne.NewPos(0, 0))
}

func (dr *dotRenderer) MinSize() (size fyne.Size) {
	size = fyne.NewSize(dr.dot.size, dr.dot.size)
	return
}

func (dr *dotRenderer) Refresh() {
	dr.dot.circle.Refresh()
}

func (dr *dotRenderer) Objects() (canObj []fyne.CanvasObject) {
	canObj = append(canObj, dr.dot.circle)
	return
}

func (dr *dotRenderer) Destroy() {}
