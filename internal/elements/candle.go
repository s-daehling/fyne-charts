package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Candle struct {
	N1         float64
	N2         float64
	Open       float64
	Close      float64
	High       float64
	Low        float64
	rect       *canvas.Rectangle
	upperLine  *canvas.Line
	lowerLine  *canvas.Line
	col        color.Color
	transposed bool
	widget.BaseWidget
}

func NewCandle(col color.Color) (c *Candle) {
	c = &Candle{
		col:       col,
		rect:      canvas.NewRectangle(col),
		upperLine: canvas.NewLine(theme.Color(theme.ColorNameForeground)),
		lowerLine: canvas.NewLine(theme.Color(theme.ColorNameForeground)),
	}
	c.rect.CornerRadius = 2
	c.ExtendBaseWidget(c)
	return
}

func (c *Candle) SetCandleColor(col color.Color) {
	c.rect.FillColor = col
}

func (c *Candle) SetLineColor(col color.Color) {
	c.upperLine.StrokeColor = col
}

func (c *Candle) SetLineWidth(lw float32) {
	c.upperLine.StrokeWidth = lw
	c.lowerLine.StrokeWidth = lw
}

func (c *Candle) SetOrientantion(transposed bool) {
	c.transposed = transposed
}

func (c *Candle) CreateRenderer() (r fyne.WidgetRenderer) {
	r = newCandleRenderer(c)
	return
}

type candleRenderer struct {
	candle *Candle
}

func newCandleRenderer(c *Candle) (cr *candleRenderer) {
	cr = &candleRenderer{
		candle: c,
	}
	return
}

func (cr *candleRenderer) Layout(size fyne.Size) {
	cMax := cr.candle.Open
	cMin := cr.candle.Close
	if cr.candle.Open < cr.candle.Close {
		cMax = cr.candle.Close
		cMin = cr.candle.Open
	}
	if cr.candle.transposed {
		cr.candle.upperLine.Position1 = fyne.NewPos(size.Width, size.Height/2)
		cr.candle.upperLine.Position2 = fyne.NewPos(size.Width*float32((cMax-cr.candle.Low)/(cr.candle.High-cr.candle.Low)), size.Height/2)
		cr.candle.lowerLine.Position1 = fyne.NewPos(0, size.Height/2)
		cr.candle.lowerLine.Position2 = fyne.NewPos(size.Width*float32((cMin-cr.candle.Low)/(cr.candle.High-cr.candle.Low)), size.Height/2)
		cr.candle.rect.Resize(fyne.NewSize(size.Width*float32((cMax-cMin)/(cr.candle.High-cr.candle.Low)), size.Height))
		cr.candle.rect.Move(fyne.NewPos(size.Width*float32((cMin-cr.candle.Low)/(cr.candle.High-cr.candle.Low)), 0))
	} else {
		cr.candle.upperLine.Position1 = fyne.NewPos(size.Width/2, 0)
		cr.candle.upperLine.Position2 = fyne.NewPos(size.Width/2, size.Height*float32((cr.candle.High-cMax)/(cr.candle.High-cr.candle.Low)))
		cr.candle.lowerLine.Position1 = fyne.NewPos(size.Width/2, size.Height*float32((cr.candle.High-cMin)/(cr.candle.High-cr.candle.Low)))
		cr.candle.lowerLine.Position2 = fyne.NewPos(size.Width/2, size.Height)
		cr.candle.rect.Resize(fyne.NewSize(size.Width, size.Height*float32((cMax-cMin)/(cr.candle.High-cr.candle.Low))))
		cr.candle.rect.Move(fyne.NewPos(0, size.Height*float32((cr.candle.High-cMax)/(cr.candle.High-cr.candle.Low))))
	}
}

func (cr *candleRenderer) MinSize() (size fyne.Size) {
	size = fyne.NewSize(0, 0)
	return
}

func (cr *candleRenderer) Refresh() {
	cr.candle.rect.Refresh()
	cr.candle.upperLine.Refresh()
	cr.candle.lowerLine.Refresh()
}

func (cr *candleRenderer) Objects() (canObj []fyne.CanvasObject) {
	canObj = append(canObj, cr.candle.rect)
	canObj = append(canObj, cr.candle.upperLine)
	canObj = append(canObj, cr.candle.lowerLine)
	return
}

func (cr *candleRenderer) Destroy() {}
