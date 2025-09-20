package series

import (
	"errors"
	"image/color"
	"sync"
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type baseSeries struct {
	name         string
	visible      bool
	mutex        *sync.Mutex
	color        color.Color
	legendButton *LegendBox
	legendLabel  *canvas.Text
	polar        bool
	chart        chart
	updateFct    func() error
}

func emptyBaseSeries(chart chart, name string, col color.Color, polar bool, togView func()) (ser baseSeries) {
	ser = baseSeries{
		name:         name,
		visible:      true,
		mutex:        &sync.Mutex{},
		color:        col,
		legendButton: NewLegendBox(col, togView),
		legendLabel:  canvas.NewText(name, theme.Color(theme.ColorNameForeground)),
		polar:        polar,
		chart:        chart,
		updateFct:    nil,
	}
	return
}

// Name gives the name of the series
func (ser *baseSeries) Name() (n string) {
	n = ser.name
	return
}

func (ser *baseSeries) LegendEntries() (les []LegendEntry) {
	le := LegendEntry{
		Button: ser.legendButton,
		Label:  ser.legendLabel,
		IsSub:  false,
	}
	les = append(les, le)
	return
}

func (ser *baseSeries) Delete() {
	ser.mutex.Lock()
	ser.chart = nil
	ser.mutex.Unlock()
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

func (ser *baseSeries) CartesianNodes(xMin float64, xMax float64, yMin float64, yMax float64) (ns []CartesianNode) {
	return
}

func (ser *baseSeries) CartesianEdges(xMin float64, xMax float64, yMin float64, yMax float64) (es []CartesianEdge) {
	return
}

func (ser *baseSeries) CartesianRects(xMin float64, xMax float64, yMin float64, yMax float64) (fs []CartesianRect) {
	return
}

func (ser *baseSeries) CartesianTexts(xMin float64, xMax float64, yMin float64, yMax float64) (fs []CartesianText) {
	return
}

func (ser *baseSeries) PolarNodes(phiMin float64, phiMax float64, rMin float64, rMax float64) (ns []PolarNode) {
	return
}

func (ser *baseSeries) PolarEdges(phiMin float64, phiMax float64, rMin float64, rMax float64) (es []PolarEdge) {
	return
}

func (ser *baseSeries) PolarTexts(phiMin float64, phiMax float64, rMin float64, rMax float64) (es []PolarText) {
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

func (ser *baseSeries) UpdateData() (err error) {
	if ser.updateFct == nil {
		err = errors.New("series has no provider function")
		return
	}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	ser.mutex.Unlock()
	err = ser.updateFct()
	return
}

type Series interface {
	LegendEntries() (les []LegendEntry)
	Name() (n string)
	Delete()
	Hide()
	Show()
	CRange() (cs []string)
	TRange() (isEmpty bool, min time.Time, max time.Time)
	NRange() (isEmpty bool, min float64, max float64)
	ValRange() (isEmpty bool, min float64, max float64)
	ConvertPtoN(pToN func(p float64) (n float64))
	ConvertCtoN(cToN func(c string) (n float64))
	ConvertTtoN(tToN func(t time.Time) (n float64))
	CartesianNodes(xMin float64, xMax float64, yMin float64, yMax float64) (ns []CartesianNode)
	CartesianEdges(xMin float64, xMax float64, yMin float64, yMax float64) (es []CartesianEdge)
	CartesianRects(xMin float64, xMax float64, yMin float64, yMax float64) (fs []CartesianRect)
	CartesianTexts(xMin float64, xMax float64, yMin float64, yMax float64) (ts []CartesianText)
	PolarNodes(phiMin float64, phiMax float64, rMin float64, rMax float64) (ns []PolarNode)
	PolarEdges(phiMin float64, phiMax float64, rMin float64, rMax float64) (es []PolarEdge)
	PolarTexts(phiMin float64, phiMax float64, rMin float64, rMax float64) (es []PolarText)
	RasterColorCartesian(x float64, y float64) (col color.Color)
	RasterColorPolar(phi float64, r float64, x float64, y float64) (col color.Color)
	UpdateData() (err error)
}

type chart interface {
	DataChange()
	RasterVisibilityChange()
	PositionToCartesianCoordinates(pX int, pY int, w int, h int) (x float64, y float64)
	PositionToPolarCoordinates(pX int, pY int, w int, h int) (phi float64, r float64, x float64, y float64)
}
