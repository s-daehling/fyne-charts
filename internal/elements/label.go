package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/s-daehling/fyne-charts/pkg/style"
)

type Label struct {
	N     float64
	Val   float64
	Shift bool
	frame *canvas.Rectangle
	text  *canvas.Text
	widget.BaseWidget
}

func NewLabel(shift bool, labelStyle style.ValueLabelStyle) (vl *Label) {
	vl = &Label{
		frame: canvas.NewRectangle(theme.Color(theme.ColorNameBackground)),
		text:  canvas.NewText("", theme.Color(theme.ColorNameForeground)),
		Shift: shift,
	}
	vl.SetStyle(labelStyle)
	vl.frame.CornerRadius = 5
	vl.text.Alignment = fyne.TextAlignCenter
	vl.ExtendBaseWidget(vl)
	return
}

func (vl *Label) CreateRenderer() (r fyne.WidgetRenderer) {
	r = newValueLabelRenderer(vl)
	return
}

func (vl *Label) SetText(t string) {
	vl.text.Text = t
}

func (vl *Label) SetStyle(labelStyle style.ValueLabelStyle) {
	vl.text.Color = theme.Color(labelStyle.ValueTextStyle.ColorName)
	vl.text.TextSize = theme.Size(labelStyle.ValueTextStyle.SizeName)
	vl.text.TextStyle = labelStyle.ValueTextStyle.TextStyle
	vl.frame.FillColor = theme.Color(labelStyle.BackgroundColorName)
	vl.frame.StrokeColor = theme.Color(labelStyle.StrokeColorName)
	vl.frame.StrokeWidth = labelStyle.StrokeWidth
}

type labelRenderer struct {
	vl *Label
}

func newValueLabelRenderer(vl *Label) (vlr *labelRenderer) {
	vlr = &labelRenderer{
		vl: vl,
	}
	return
}

func (vlr *labelRenderer) Layout(size fyne.Size) {
	vlr.vl.frame.Resize(size)
	vlr.vl.frame.Move(fyne.NewPos(0, 0))
	vlr.vl.text.Resize(vlr.vl.text.MinSize())
	vlr.vl.text.Move(fyne.NewPos(3, 0))
}

func (vlr *labelRenderer) MinSize() (size fyne.Size) {
	size = vlr.vl.text.MinSize()
	size.Width += 6
	return
}

func (vlr *labelRenderer) Refresh() {
	vlr.vl.frame.Refresh()
	vlr.vl.text.Refresh()
}

func (vlr *labelRenderer) Objects() (canObj []fyne.CanvasObject) {
	canObj = append(canObj, vlr.vl.frame)
	canObj = append(canObj, vlr.vl.text)
	return
}

func (vlr *labelRenderer) Destroy() {}
