package series

import (
	"errors"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/s-daehling/fyne-charts/internal/elements"
	"github.com/s-daehling/fyne-charts/internal/interact"
)

type baseSeries struct {
	name        string
	super       string
	visible     bool
	col         color.Color
	colName     fyne.ThemeColorName
	legendEntry *interact.LegendEntry
	cont        container
}

func emptyBaseSeries(name string, colName fyne.ThemeColorName, togView func()) (ser baseSeries) {
	ser = baseSeries{
		name:        name,
		super:       "",
		visible:     true,
		colName:     colName,
		col:         theme.Color(colName),
		legendEntry: interact.NewLegendEntry(name, "", true, colName, togView),
		cont:        nil,
	}
	return
}

// Name gives the name of the series
func (ser *baseSeries) Name() (n string) {
	n = ser.name
	return
}

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

func (ser *baseSeries) CartesianDots(xMin float64, xMax float64, yMin float64,
	yMax float64) (ns []*elements.Dot) {
	return
}

func (ser *baseSeries) CartesianEdges(xMin float64, xMax float64, yMin float64,
	yMax float64) (es []elements.Edge) {
	return
}

func (ser *baseSeries) CartesianBars(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []*elements.Bar) {
	return
}

func (ser *baseSeries) CartesianBoxes(xMin float64, xMax float64, yMin float64,
	yMax float64) (bs []*elements.Box) {
	return
}

func (ser *baseSeries) CartesianCandles(xMin float64, xMax float64, yMin float64,
	yMax float64) (cs []*elements.Candle) {
	return
}

func (ser *baseSeries) CartesianTexts(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []elements.Label) {
	return
}

func (ser *baseSeries) PolarDots(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ns []*elements.Dot) {
	return
}

func (ser *baseSeries) PolarEdges(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (es []elements.Edge) {
	return
}

func (ser *baseSeries) PolarTexts(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (es []elements.Label) {
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

func (ser *baseSeries) IsPartOfChartRaster() (b bool) {
	b = false
	return
}

func (ser *baseSeries) RefreshTheme() {
	ser.col = theme.Color(ser.colName)
}

func (ser *baseSeries) Hover(n float64, val float64) (text string) {
	return
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
	CartesianDots(xMin float64, xMax float64, yMin float64, yMax float64) (ns []*elements.Dot)
	CartesianEdges(xMin float64, xMax float64, yMin float64, yMax float64) (es []elements.Edge)
	CartesianBars(xMin float64, xMax float64, yMin float64, yMax float64) (fs []*elements.Bar)
	CartesianBoxes(xMin float64, xMax float64, yMin float64, yMax float64) (bs []*elements.Box)
	CartesianCandles(xMin float64, xMax float64, yMin float64, yMax float64) (cs []*elements.Candle)
	CartesianTexts(xMin float64, xMax float64, yMin float64, yMax float64) (ts []elements.Label)
	PolarDots(phiMin float64, phiMax float64, rMin float64, rMax float64) (ns []*elements.Dot)
	PolarEdges(phiMin float64, phiMax float64, rMin float64, rMax float64) (es []elements.Edge)
	PolarTexts(phiMin float64, phiMax float64, rMin float64, rMax float64) (es []elements.Label)
	RasterColorCartesian(x float64, y float64) (col color.Color)
	RasterColorPolar(phi float64, r float64, x float64, y float64) (col color.Color)
	IsPartOfChartRaster() (b bool)
	RefreshTheme()
	Hover(n float64, val float64) (text string)
}

type container interface {
	IsPolar() (b bool)
	DataChange()
	AreaRefresh()
	AddLegendEntry(le *interact.LegendEntry)
	RemoveLegendEntry(name string, super string)
}
