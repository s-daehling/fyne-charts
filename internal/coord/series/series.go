package series

import (
	"errors"
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/interact"
	"github.com/s-daehling/fyne-charts/internal/renderer"
)

type baseSeries struct {
	name        string
	super       string
	visible     bool
	color       color.Color
	legendEntry *interact.LegendEntry
	// legendButton *interact.LegendBox
	// legendLabel  *canvas.Text
	cont container
}

func emptyBaseSeries(name string, col color.Color, togView func()) (ser baseSeries) {
	ser = baseSeries{
		name:        name,
		super:       "",
		visible:     true,
		color:       col,
		legendEntry: interact.NewLegendEntry(name, "", true, col, togView),
		// legendButton: interact.NewLegendBox(col, togView),
		// legendLabel:  canvas.NewText(name, theme.Color(theme.ColorNameForeground)),
		cont: nil,
	}
	return
}

// Name gives the name of the series
func (ser *baseSeries) Name() (n string) {
	n = ser.name
	return
}

// func (ser *baseSeries) LegendEntries() (les []*interact.LegendEntry) {
// 	les = append(les, ser.legendEntry)
// 	return
// }

func (ser *baseSeries) BindToChart(ch container) (err error) {
	if ser.cont != nil {
		err = errors.New("series is already part of a chart")
		return
	}
	ser.cont = ch
	ch.AddLegendEntry(ser.legendEntry)
	return
}

func (ser *baseSeries) Release() {
	if ser.cont != nil {
		ser.cont.RemoveLegendEntry(ser.name, ser.super)
	}
	ser.cont = nil
}

func (ser *baseSeries) HasChart() (b bool) {
	b = false
	if ser.cont != nil {
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

}

type Series interface {
	// LegendEntries() (les []*interact.LegendEntry)
	Name() (n string)
	BindToChart(ch container) (err error)
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

type container interface {
	IsPolar() (b bool)
	DataChange()
	RasterVisibilityChange()
	AddLegendEntry(le *interact.LegendEntry)
	RemoveLegendEntry(name string, super string)
}
