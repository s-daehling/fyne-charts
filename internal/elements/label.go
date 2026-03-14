package elements

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/s-daehling/fyne-charts/pkg/style"
)

type Label struct {
	N    float64
	Val  float64
	Text *canvas.Text
}

type ValueLabel struct {
	N         float64
	Val       float64
	frame     *canvas.Rectangle
	text      *canvas.Text
	textStyle style.ChartTextStyle
	widget.BaseWidget
}

func NewValueLabel() (vl *ValueLabel) {
	vl = &ValueLabel{
		frame: canvas.NewRectangle(theme.Color(theme.ColorNameBackground)),
		text:  canvas.NewText("", theme.Color(theme.ColorNameForeground)),
	}
	vl.frame.CornerRadius = 5
	vl.frame.StrokeColor = theme.Color(theme.ColorNameForeground)
	vl.frame.StrokeWidth = 0.5
	vl.text.Alignment = fyne.TextAlignCenter
	vl.ExtendBaseWidget(vl)
	return
}

func (vl *ValueLabel) CreateRenderer() (r fyne.WidgetRenderer) {
	r = newValueLabelRenderer(vl)
	return
}

func (vl *ValueLabel) SetText(t string) {
	vl.text.Text = t
}

type valueLabelRenderer struct {
	vl *ValueLabel
}

func newValueLabelRenderer(vl *ValueLabel) (vlr *valueLabelRenderer) {
	vlr = &valueLabelRenderer{
		vl: vl,
	}
	return
}

func (vlr *valueLabelRenderer) Layout(size fyne.Size) {
	vlr.vl.frame.Resize(size)
	vlr.vl.frame.Move(fyne.NewPos(0, 0))
	vlr.vl.text.Resize(vlr.vl.text.MinSize())
	vlr.vl.text.Move(fyne.NewPos(3, 0))
}

func (vlr *valueLabelRenderer) MinSize() (size fyne.Size) {
	size = vlr.vl.text.MinSize()
	size.Width += 6
	return
}

func (vlr *valueLabelRenderer) Refresh() {
	vlr.vl.frame.Refresh()
	vlr.vl.text.Refresh()
}

func (vlr *valueLabelRenderer) Objects() (canObj []fyne.CanvasObject) {
	canObj = append(canObj, vlr.vl.frame)
	canObj = append(canObj, vlr.vl.text)
	return
}

func (vlr *valueLabelRenderer) Destroy() {}
