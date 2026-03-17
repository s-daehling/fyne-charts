package prop

import (
	"image/color"
	"math"

	"github.com/s-daehling/fyne-charts/internal/elements"
	"github.com/s-daehling/fyne-charts/internal/interact"
	"github.com/s-daehling/fyne-charts/internal/renderer"
	"github.com/s-daehling/fyne-charts/pkg/style"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type PlaneType string

const (
	CartesianPlane PlaneType = "Cartesian"
	PolarPlane     PlaneType = "Polar"
)

type BaseChart struct {
	widget.BaseWidget
	title             *canvas.Text
	titleStyle        style.ChartTextStyle
	series            []*Series
	changed           bool
	legend            *interact.Legend
	legendVisible     bool
	planeType         PlaneType
	transposed        bool
	onHoverHighlight  bool
	onHoverValueLabel bool
	area              *elements.Area
	render            fyne.WidgetRenderer
	fromMin           float64
	fromMax           float64
	toMin             float64
	toMax             float64
	mainCont          *fyne.Container
	rLegendCont       *fyne.Container
	lLegendCont       *fyne.Container
	bLegendCont       *fyne.Container
	tLegendCont       *fyne.Container
}

func EmptyBaseChart(pType PlaneType) (base *BaseChart) {
	base = &BaseChart{
		title:             canvas.NewText("", theme.Color(theme.ColorNameForeground)),
		changed:           false,
		legend:            interact.NewLegend(),
		legendVisible:     true,
		planeType:         pType,
		transposed:        false,
		onHoverHighlight:  true,
		onHoverValueLabel: true,
		fromMin:           0,
		toMin:             0,
		toMax:             100,
		rLegendCont:       container.NewCenter(),
		lLegendCont:       container.NewCenter(),
		bLegendCont:       container.NewStack(),
		tLegendCont:       container.NewStack(),
	}
	base.mainCont = container.NewBorder(
		container.NewVBox(
			base.title,
			base.tLegendCont),
		base.bLegendCont,
		base.lLegendCont,
		base.rLegendCont,
		base)
	base.SetTitleStyle(style.DefaultTitleStyle())
	base.SetLegendStyle(style.LegendLocationRight, style.DefaultLegendTextStyle(), true)
	if pType == CartesianPlane {
		base.area = nil
		base.fromMax = 100
	} else {
		base.area = elements.NewArea(base.PixelGenPolar, base.MouseMove)
		base.fromMax = 2 * math.Pi
	}
	base.ExtendBaseWidget(base)
	return
}

func (base *BaseChart) CreateRenderer() (r fyne.WidgetRenderer) {
	if base.planeType == CartesianPlane {
		base.render = renderer.EmptyCartesianRenderer(base)
	} else {
		base.render = renderer.EmptyPolarRenderer(base)
	}
	r = base.render
	return
}

func (base *BaseChart) MainContainer() (cont *fyne.Container) {
	cont = base.mainCont
	return
}

func (base *BaseChart) IsPolar() (b bool) {
	b = (base.planeType == PolarPlane)
	return
}

func (base *BaseChart) SetCartesianOrientantion(transposed bool) {
	if base.transposed != transposed {
		base.transposed = transposed
		base.DataChange()
	}
}

func (base *BaseChart) CartesianOrientation() (transposed bool) {
	transposed = base.transposed
	return
}

func (base *BaseChart) Objects() (canObj []fyne.CanvasObject) {
	// objects will be drawn in the same order as added here

	// first get all objects from the series
	if base.area != nil {
		canObj = append(canObj, base.area)
	}
	bars := base.Bars()
	for i := range bars {
		canObj = append(canObj, bars[i])
	}
	labels := base.Labels()
	for i := range labels {
		canObj = append(canObj, labels[i])
	}
	return
}

func (base *BaseChart) Dots() (ns []*elements.Dot) {
	return
}

func (base *BaseChart) Edges() (es []elements.Edge) {
	return
}

func (base *BaseChart) Bars() (as []*elements.Bar) {
	for i := range base.series {
		as = append(as, base.series[i].Bars(base.fromMin, base.fromMax, base.toMin, base.toMax)...)
	}
	return
}

func (base *BaseChart) Boxes() (ns []*elements.Box) {
	return
}

func (base *BaseChart) Candles() (ns []*elements.Candle) {
	return
}

func (base *BaseChart) Labels() (ls []*elements.Label) {
	for i := range base.series {
		ls = append(ls, base.series[i].Labels(base.fromMin, base.fromMax, base.toMin, base.toMax, base.onHoverValueLabel)...)
	}
	return
}

func (base *BaseChart) Area() (rs *elements.Area) {
	rs = base.area
	return
}

func (base *BaseChart) SetLegendStyle(loc style.LegendLocation, ls style.ChartTextStyle, interactive bool) {
	base.legend.SetStyle(loc, ls, interactive)
	base.lLegendCont.RemoveAll()
	base.rLegendCont.RemoveAll()
	base.tLegendCont.RemoveAll()
	base.bLegendCont.RemoveAll()
	switch loc {
	case style.LegendLocationBottom:
		base.bLegendCont.Add(base.legend)
	case style.LegendLocationLeft:
		base.lLegendCont.Add(base.legend)
	case style.LegendLocationRight:
		base.rLegendCont.Add(base.legend)
	case style.LegendLocationTop:
		base.tLegendCont.Add(base.legend)
	}
}

func (base *BaseChart) ShowLegend() {
	base.legend.Show()
}

func (base *BaseChart) HideLegend() {
	base.legend.Hide()
}

func (base *BaseChart) SetTitle(l string) {
	base.title.Text = l
	if l == "" && !base.title.Hidden {
		base.title.Hide()
	} else if l != "" && base.title.Hidden {
		base.title.Show()
	}
	base.title.Refresh()
}

func (base *BaseChart) SetTitleStyle(ts style.ChartTextStyle) {
	base.titleStyle = ts
	base.title.Alignment = ts.Alignment
	base.title.TextSize = theme.Size(ts.SizeName)
	base.title.Color = theme.Color(ts.ColorName)
	base.title.TextStyle = ts.TextStyle
	base.title.Refresh()
}

func (base *BaseChart) FromAxisElements() (min float64, max float64, origin float64,
	ticks []elements.Tick, arrow elements.Arrow, show bool) {
	min, max = base.fromMin, base.fromMax
	origin = 0
	ticks = []elements.Tick{}
	arrow = elements.Arrow{}
	show = false
	return
}

func (base *BaseChart) ToAxisElements() (min float64, max float64, origin float64,
	ticks []elements.Tick, arrow elements.Arrow, show bool) {
	min, max = base.toMin, base.toMax
	origin = 0
	ticks = []elements.Tick{}
	arrow = elements.Arrow{}
	show = false
	return
}

func (base *BaseChart) MouseMove(pX, pY, w, h, absX, absY float32) {
	if base.planeType == CartesianPlane {

	} else {
		phi, r, _, _ := base.PositionToPolarCoordinates(pX, pY, w, h)
		for i := range base.series {
			base.series[i].Hover(phi, r)
		}
	}
}

func (base *BaseChart) PixelGenPolar(pX, pY, w, h int) (col color.Color) {
	phi, r, _, _ := base.PositionToPolarCoordinates(float32(pX), float32(pY), float32(w), float32(h))
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	if r > base.toMax {
		return
	}
	for i := range base.series {
		serCol, useColor := base.series[i].RasterColorPolar(phi, r)
		if useColor {
			col = serCol
			break
		}
	}
	return
}

func (base *BaseChart) PositionToPolarCoordinates(pX float32, pY float32, w float32, h float32) (phi float64,
	r float64, x float64, y float64) {
	rot := 0.0
	mathPos := true
	posToCoord := base.toMax / (float64(w) / 2.0)
	x = (float64(pX) - (float64(w) / 2.0)) * posToCoord
	y = ((float64(h) / 2.0) - float64(pY)) * posToCoord
	r = math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	phi = math.Acos(x / r)
	if y < 0 {
		phi = -phi + (2 * math.Pi)
	}
	dirCor := 0.0
	if !mathPos {
		dirCor = 1.0
	}
	phi -= (rot + (dirCor * 2 * math.Pi))
	if phi < 0 {
		phi += 2 * math.Pi
	} else if phi > 2*math.Pi {
		phi -= 2 * math.Pi
	}
	return
}
