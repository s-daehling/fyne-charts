package series

import (
	"errors"
	"image/color"
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"github.com/s-daehling/fyne-charts/internal/interact"
	"github.com/s-daehling/fyne-charts/internal/renderer"
)

type baseSeries struct {
	name         string
	visible      bool
	color        color.Color
	legendButton *interact.LegendBox
	legendLabel  *canvas.Text
	chart        chart
}

func emptyBaseSeries(name string, col color.Color, togView func()) (ser baseSeries) {
	ser = baseSeries{
		name:         name,
		visible:      true,
		color:        col,
		legendButton: interact.NewLegendBox(col, togView),
		legendLabel:  canvas.NewText(name, theme.Color(theme.ColorNameForeground)),
		chart:        nil,
	}
	return
}

// Name gives the name of the series
func (ser *baseSeries) Name() (n string) {
	n = ser.name
	return
}

func (ser *baseSeries) LegendEntries() (les []renderer.LegendEntry) {
	le := renderer.LegendEntry{
		Button: ser.legendButton,
		Label:  ser.legendLabel,
		IsSub:  false,
	}
	les = append(les, le)
	return
}

func (ser *baseSeries) Bind(ch chart) (err error) {
	if ser.chart != nil {
		err = errors.New("series is already part of a chart")
		return
	}
	ser.chart = ch
	return
}

func (ser *baseSeries) Release() {
	ser.chart = nil
}

func (ser *baseSeries) HasChart() (b bool) {
	b = false
	if ser.chart != nil {
		b = true
	}
	return
}

func (ser *baseSeries) CRange() (cs []string) { return }

func (ser *baseSeries) TRange() (isEmpty bool, min time.Time, max time.Time) {
	isEmpty = true
	return
}

func (ser *baseSeries) NRange() (isEmpty bool, min float64, max float64) {
	isEmpty = true
	return
}

func (ser *baseSeries) ValRange() (isEmpty bool, min float64, max float64) {
	isEmpty = true
	return
}

func (ser *baseSeries) ConvertPtoN(pToN func(p float64) (n float64)) {}

func (ser *baseSeries) ConvertCtoN(cToN func(c string) (n float64)) {}

func (ser *baseSeries) ConvertTtoN(tToN func(t time.Time) (n float64)) {}

func (ser *baseSeries) CartesianNodes(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []renderer.CartesianNode) {
	return
}

func (ser *baseSeries) CartesianEdges(xMin float64, xMax float64, yMin float64,
	yMax float64) (es []renderer.CartesianEdge) {
	return
}

func (ser *baseSeries) CartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []renderer.CartesianRect) {
	return
}

func (ser *baseSeries) CartesianTexts(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []renderer.CartesianText) {
	return
}

func (ser *baseSeries) PolarNodes(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ns []renderer.PolarNode) {
	return
}

func (ser *baseSeries) PolarEdges(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (es []renderer.PolarEdge) {
	return
}

func (ser *baseSeries) PolarTexts(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (es []renderer.PolarText) {
	return
}

func (ser *baseSeries) RasterColorCartesian(x float64, y float64) (col color.Color) {
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	return
}

func (ser *baseSeries) RasterColorPolar(phi float64, r float64, x float64, y float64) (col color.Color) {
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	return
}

func (ser *baseSeries) RefreshTheme() {
	ser.legendLabel.Color = theme.Color(theme.ColorNameForeground)
}

type Series interface {
	LegendEntries() (les []renderer.LegendEntry)
	Name() (n string)
	Bind(ch chart) (err error)
	Release()
	Hide()
	Show()
	CRange() (cs []string)
	TRange() (isEmpty bool, min time.Time, max time.Time)
	NRange() (isEmpty bool, min float64, max float64)
	ValRange() (isEmpty bool, min float64, max float64)
	ConvertPtoN(pToN func(p float64) (n float64))
	ConvertCtoN(cToN func(c string) (n float64))
	ConvertTtoN(tToN func(t time.Time) (n float64))
	CartesianNodes(xMin float64, xMax float64, yMin float64, yMax float64) (ns []renderer.CartesianNode)
	CartesianEdges(xMin float64, xMax float64, yMin float64, yMax float64) (es []renderer.CartesianEdge)
	CartesianRects(xMin float64, xMax float64, yMin float64, yMax float64) (fs []renderer.CartesianRect)
	CartesianTexts(xMin float64, xMax float64, yMin float64, yMax float64) (ts []renderer.CartesianText)
	PolarNodes(phiMin float64, phiMax float64, rMin float64, rMax float64) (ns []renderer.PolarNode)
	PolarEdges(phiMin float64, phiMax float64, rMin float64, rMax float64) (es []renderer.PolarEdge)
	PolarTexts(phiMin float64, phiMax float64, rMin float64, rMax float64) (es []renderer.PolarText)
	RasterColorCartesian(x float64, y float64) (col color.Color)
	RasterColorPolar(phi float64, r float64, x float64, y float64) (col color.Color)
	RefreshTheme()
}

type chart interface {
	IsPolar() (b bool)
	DataChange()
	RasterVisibilityChange()
}
