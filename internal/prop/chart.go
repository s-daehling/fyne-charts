package prop

import (
	"image/color"
	"math"

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
	title          *canvas.Text
	titleColorName fyne.ThemeColorName
	titleSizeName  fyne.ThemeSizeName
	series         []*Series
	changed        bool
	legend         *interact.Legend
	legendVisible  bool
	planeType      PlaneType
	transposed     bool
	rast           *canvas.Raster
	render         fyne.WidgetRenderer
	fromMin        float64
	fromMax        float64
	toMin          float64
	toMax          float64
	mainCont       *fyne.Container
	rLegendCont    *fyne.Container
	lLegendCont    *fyne.Container
	bLegendCont    *fyne.Container
	tLegendCont    *fyne.Container
}

func EmptyBaseChart(pType PlaneType) (base *BaseChart) {
	base = &BaseChart{
		title:         canvas.NewText("", theme.Color(theme.ColorNameForeground)),
		changed:       false,
		legend:        interact.NewLegend(),
		legendVisible: true,
		planeType:     pType,
		transposed:    false,
		fromMin:       0,
		toMin:         0,
		toMax:         100,
		rLegendCont:   container.NewCenter(),
		lLegendCont:   container.NewCenter(),
		bLegendCont:   container.NewStack(),
		tLegendCont:   container.NewStack(),
	}
	base.mainCont = container.NewBorder(
		container.NewVBox(
			base.title,
			base.tLegendCont),
		base.bLegendCont,
		base.lLegendCont,
		base.rLegendCont,
		base)
	base.SetTitleStyle(theme.SizeNameHeadingText, theme.ColorNameForeground)
	base.SetLegendStyle(style.LegendLocationRight)
	if pType == CartesianPlane {
		base.rast = nil
		base.fromMax = 100
	} else {
		base.rast = canvas.NewRasterWithPixels(base.PixelGenPolar)
		base.fromMax = 2 * math.Pi
	}
	base.ExtendBaseWidget(base)
	return
}

func (base *BaseChart) CreateRenderer() (r fyne.WidgetRenderer) {
	if base.planeType == CartesianPlane {
		base.render = renderer.EmptyCartesianRenderer(base, base.Size)
	} else {
		base.render = renderer.EmptyPolarRenderer(base, base.Size)
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

func (base *BaseChart) CartesianObjects() (canObj []fyne.CanvasObject) {
	// objects will be drawn in the same order as added here

	// first get all objects from the series
	rects := base.CartesianRects()
	for i := range rects {
		canObj = append(canObj, rects[i].Rect)
	}
	texts := base.CartesianTexts()
	for i := range texts {
		canObj = append(canObj, texts[i].Text)
	}

	return
}

func (base *BaseChart) CartesianNodes() (ns []renderer.CartesianNode) {
	return
}

func (base *BaseChart) CartesianEdges() (es []renderer.CartesianEdge) {
	return
}

func (base *BaseChart) CartesianRects() (as []renderer.CartesianRect) {
	for i := range base.series {
		as = append(as, base.series[i].CartesianRects(base.fromMin, base.fromMax, base.toMin, base.toMax)...)
	}
	return
}

func (base *BaseChart) CartesianTexts() (ts []renderer.CartesianText) {
	for i := range base.series {
		ts = append(ts, base.series[i].CartesianTexts(base.fromMin, base.fromMax, base.toMin, base.toMax)...)
	}
	return
}

func (base *BaseChart) PolarObjects() (canObj []fyne.CanvasObject) {
	// objects will be drawn in the same order as added here

	// first get all objects from the series
	canObj = append(canObj, base.rast)
	texts := base.PolarTexts()
	for i := range texts {
		canObj = append(canObj, texts[i].Text)
	}

	return
}

func (base *BaseChart) PolarNodes() (ns []renderer.PolarNode) {
	return
}

func (base *BaseChart) PolarEdges() (es []renderer.PolarEdge) {
	return
}

func (base *BaseChart) PolarTexts() (ts []renderer.PolarText) {
	for i := range base.series {
		ts = append(ts, base.series[i].PolarTexts(base.fromMin, base.fromMax, base.toMin, base.toMax)...)
	}
	return
}

func (base *BaseChart) Raster() (rs *canvas.Raster) {
	rs = base.rast
	return
}

func (base *BaseChart) Overlay() (io *interact.Overlay) {
	io = nil
	return
}

func (base *BaseChart) SetLegendStyle(loc style.LegendLocation) {
	base.legend.SetLocation(loc)
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
	base.mainCont.Refresh()
}

func (base *BaseChart) ShowLegend() {
	base.legend.Show()
	base.mainCont.Refresh()
}

func (base *BaseChart) HideLegend() {
	base.legend.Hide()
	base.mainCont.Refresh()
}

func (base *BaseChart) Tooltip() (tt renderer.Tooltip) {
	return
}

func (base *BaseChart) SetTitle(l string) {
	base.title.Text = l
	if l == "" {
		base.title.Hide()
	} else {
		base.title.Show()
	}
	base.title.Refresh()
}

func (base *BaseChart) SetTitleStyle(sizeName fyne.ThemeSizeName, colorName fyne.ThemeColorName) {
	base.title.Alignment = fyne.TextAlignCenter
	base.titleSizeName = sizeName
	base.title.TextSize = theme.Size(sizeName)
	base.titleColorName = colorName
	base.title.Color = theme.Color(colorName)
	base.title.Refresh()
}

func (base *BaseChart) FromAxisElements() (min float64, max float64, origin float64,
	ticks []renderer.Tick, arrow renderer.Arrow, show bool) {
	min, max = base.fromMin, base.fromMax
	origin = 0
	ticks = []renderer.Tick{}
	arrow = renderer.Arrow{}
	show = false
	return
}

func (base *BaseChart) ToAxisElements() (min float64, max float64, origin float64,
	ticks []renderer.Tick, arrow renderer.Arrow, show bool) {
	min, max = base.toMin, base.toMax
	origin = 0
	ticks = []renderer.Tick{}
	arrow = renderer.Arrow{}
	show = false
	return
}

func (base *BaseChart) PixelGenPolar(pX, pY, w, h int) (col color.Color) {
	phi, r, x, y := base.PositionToPolarCoordinates(pX, pY, w, h)
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	if r > base.toMax {
		return
	}
	for i := range base.series {
		serCol := base.series[i].RasterColorPolar(phi, r, x, y)
		r, g, b, _ := serCol.RGBA()
		if r > 0 || g > 0 || b > 0 {
			col = serCol
			break
		}
	}
	return
}

func (base *BaseChart) PositionToPolarCoordinates(pX int, pY int, w int, h int) (phi float64,
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
