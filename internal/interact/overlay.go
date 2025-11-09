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

type chart interface {
	MouseIn(pX, pY, w, h float32)
	MouseMove(pX, pY, w, h float32)
	MouseOut()
}

type Overlay struct {
	widget.BaseWidget
	chart chart
	rect  *canvas.Rectangle
	text  *canvas.Text
}

func NewOverlay(chart chart) (io *Overlay) {
	io = &Overlay{
		chart: chart,
		rect:  canvas.NewRectangle(color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00}),
		text:  canvas.NewText("", theme.Color(theme.ColorNameForeground)),
	}
	return
}

func (ol *Overlay) CreateRenderer() (r fyne.WidgetRenderer) {
	c := container.NewStack(ol.rect, ol.text)
	r = widget.NewSimpleRenderer(c)
	return
}

func (ol *Overlay) Tapped(_ *fyne.PointEvent) {

}

func (ol *Overlay) MouseIn(me *desktop.MouseEvent) {
	size := ol.rect.Size()
	ol.chart.MouseIn(me.Position.X, me.Position.Y, size.Width, size.Height)
}

func (ol *Overlay) MouseMoved(me *desktop.MouseEvent) {
	size := ol.rect.Size()
	ol.chart.MouseMove(me.Position.X, me.Position.Y, size.Width, size.Height)
	// size := ol.rect.Size()
	// rx := me.Position.X / size.Width
	// ry := (size.Height - me.Position.Y) / size.Height
	// ol.text.Text = fmt.Sprintf("x: %f, y: %f", rx, ry)
}

func (ol *Overlay) MouseOut() {
	ol.chart.MouseOut()
}
