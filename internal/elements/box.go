package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type Box struct {
	N1           float64
	N2           float64
	Max          float64
	ThirdQuart   float64
	Median       float64
	FirstQuart   float64
	Min          float64
	maxLine      *canvas.Line
	upperWhisker *canvas.Line
	medianLine   *canvas.Line
	lowerWhisker *canvas.Line
	minLine      *canvas.Line
	rect         *canvas.Rectangle
	col          color.Color
	transposed   bool
	widget.BaseWidget
}

func NewBox(col color.Color) (b *Box) {
	b = &Box{
		col:          col,
		maxLine:      canvas.NewLine(col),
		upperWhisker: canvas.NewLine(col),
		medianLine:   canvas.NewLine(col),
		lowerWhisker: canvas.NewLine(col),
		minLine:      canvas.NewLine(col),
		transposed:   false,
	}
	b.rect = canvas.NewRectangle(color.RGBA{0x00, 0x00, 0x00, 0x00})
	b.rect.StrokeColor = col
	b.rect.StrokeWidth = 1
	b.ExtendBaseWidget(b)
	return
}

func (b *Box) SetColor(col color.Color) {
	b.maxLine.StrokeColor = col
	b.upperWhisker.StrokeColor = col
	b.medianLine.StrokeColor = col
	b.lowerWhisker.StrokeColor = col
	b.minLine.StrokeColor = col
	b.rect.StrokeColor = col
}

func (b *Box) SetLineWidth(lw float32) {
	b.maxLine.StrokeWidth = lw
	b.upperWhisker.StrokeWidth = lw
	b.medianLine.StrokeWidth = lw
	b.lowerWhisker.StrokeWidth = lw
	b.minLine.StrokeWidth = lw
	b.rect.StrokeWidth = lw
}

func (b *Box) SetOrientantion(transposed bool) {
	b.transposed = transposed
}

func (b *Box) MouseIn(me *desktop.MouseEvent) {
}

func (b *Box) MouseMoved(me *desktop.MouseEvent) {

}

func (b *Box) MouseOut() {
}

func (b *Box) CreateRenderer() (r fyne.WidgetRenderer) {
	r = newBoxRenderer(b)
	return
}

type boxRenderer struct {
	box *Box
}

func newBoxRenderer(b *Box) (br *boxRenderer) {
	br = &boxRenderer{
		box: b,
	}
	return
}

func (br *boxRenderer) Layout(size fyne.Size) {
	if !br.box.transposed {
		br.box.maxLine.Position1 = fyne.NewPos(0, 0)
		br.box.maxLine.Position2 = fyne.NewPos(size.Width, 0)
		br.box.minLine.Position1 = fyne.NewPos(0, size.Height)
		br.box.minLine.Position2 = fyne.NewPos(size.Width, size.Height)
		yMedian := float32((br.box.Max-br.box.Median)/(br.box.Max-br.box.Min)) * size.Height
		br.box.medianLine.Position1 = fyne.NewPos(0, yMedian)
		br.box.medianLine.Position2 = fyne.NewPos(size.Width, yMedian)
		br.box.rect.Resize(fyne.NewSize(size.Width,
			size.Height*float32((br.box.ThirdQuart-br.box.FirstQuart)/(br.box.Max-br.box.Min))))
		br.box.rect.Move(fyne.NewPos(0, size.Height*float32((br.box.Max-br.box.ThirdQuart)/(br.box.Max-br.box.Min))))
		br.box.upperWhisker.Position1 = fyne.NewPos(size.Width/2, 0)
		br.box.upperWhisker.Position2 = fyne.NewPos(size.Width/2, size.Height*float32((br.box.Max-br.box.ThirdQuart)/(br.box.Max-br.box.Min)))
		br.box.lowerWhisker.Position1 = fyne.NewPos(size.Width/2, size.Height*float32((br.box.Max-br.box.FirstQuart)/(br.box.Max-br.box.Min)))
		br.box.lowerWhisker.Position2 = fyne.NewPos(size.Width/2, size.Height)
	} else {
		br.box.maxLine.Position1 = fyne.NewPos(size.Width, 0)
		br.box.maxLine.Position2 = fyne.NewPos(size.Width, size.Height)
		br.box.minLine.Position1 = fyne.NewPos(0, 0)
		br.box.minLine.Position2 = fyne.NewPos(0, size.Height)
		xMedian := float32((br.box.Median-br.box.Min)/(br.box.Max-br.box.Min)) * size.Width
		br.box.medianLine.Position1 = fyne.NewPos(xMedian, 0)
		br.box.medianLine.Position2 = fyne.NewPos(xMedian, size.Height)
		br.box.rect.Resize(fyne.NewSize(size.Width*float32((br.box.ThirdQuart-br.box.FirstQuart)/(br.box.Max-br.box.Min)), size.Height))
		br.box.rect.Move(fyne.NewPos(size.Width*float32((br.box.FirstQuart-br.box.Min)/(br.box.Max-br.box.Min)), 0))
		br.box.upperWhisker.Position1 = fyne.NewPos(size.Width*float32((br.box.ThirdQuart-br.box.Min)/(br.box.Max-br.box.Min)), size.Height/2)
		br.box.upperWhisker.Position2 = fyne.NewPos(size.Width, size.Height/2)
		br.box.lowerWhisker.Position1 = fyne.NewPos(size.Width*float32((br.box.FirstQuart-br.box.Min)/(br.box.Max-br.box.Min)), size.Height/2)
		br.box.lowerWhisker.Position2 = fyne.NewPos(0, size.Height/2)
	}
}

func (br *boxRenderer) MinSize() (size fyne.Size) {
	size = fyne.NewSize(0, 0)
	return
}

func (br *boxRenderer) Refresh() {
	br.box.maxLine.Refresh()
	br.box.upperWhisker.Refresh()
	br.box.medianLine.Refresh()
	br.box.lowerWhisker.Refresh()
	br.box.minLine.Refresh()
	br.box.rect.Refresh()
}

func (br *boxRenderer) Objects() (canObj []fyne.CanvasObject) {
	canObj = append(canObj, br.box.maxLine)
	canObj = append(canObj, br.box.upperWhisker)
	canObj = append(canObj, br.box.medianLine)
	canObj = append(canObj, br.box.lowerWhisker)
	canObj = append(canObj, br.box.minLine)
	canObj = append(canObj, br.box.rect)
	return
}

func (br *boxRenderer) Destroy() {}
