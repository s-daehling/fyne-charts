package prop

import (
	"image/color"
	"math"

	"github.com/s-daehling/fyne-charts/internal/interact"
	"github.com/s-daehling/fyne-charts/internal/renderer"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type PlaneType string

const (
	CartesianPlane PlaneType = "Cartesian"
	PolarPlane     PlaneType = "Polar"
)

type BaseChart struct {
	title          *canvas.Text
	titleColorName fyne.ThemeColorName
	titleSizeName  fyne.ThemeSizeName
	series         []*Series
	changed        bool
	legendVisible  bool
	planeType      PlaneType
	rast           *canvas.Raster
	render         fyne.WidgetRenderer
	fromMin        float64
	fromMax        float64
	toMin          float64
	toMax          float64
}

func EmptyBaseChart(pType PlaneType) (base *BaseChart) {
	base = &BaseChart{
		title:         canvas.NewText("", theme.Color(theme.ColorNameForeground)),
		changed:       false,
		legendVisible: true,
		planeType:     pType,
		fromMin:       0,
		toMin:         0,
		toMax:         100,
	}
	base.SetTitleStyle(theme.SizeNameHeadingText, theme.ColorNameForeground)
	if pType == CartesianPlane {
		base.rast = nil
		base.fromMax = 100
	} else {
		base.rast = canvas.NewRasterWithPixels(base.PixelGenPolar)
		base.fromMax = 2 * math.Pi
	}
	return
}

func (base *BaseChart) CreateRenderer(ws func() fyne.Size) (r fyne.WidgetRenderer) {
	if base.planeType == CartesianPlane {
		base.render = renderer.EmptyCartesianRenderer(base, ws)
	} else {
		base.render = renderer.EmptyPolarRenderer(base, ws)
	}
	r = base.render
	return
}

func (base *BaseChart) IsPolar() (b bool) {
	b = (base.planeType == PolarPlane)
	return
}

func (base *BaseChart) SeriesExist(n string) (exist bool) {
	exist = false
	for i := range base.series {
		if base.series[i].Name() == n {
			exist = true
			break
		}
	}
	return
}

func (base *BaseChart) DeleteSeries(name string) {
	newSeries := make([]*Series, 0)
	for i := range base.series {
		if base.series[i].Name() != name {
			newSeries = append(newSeries, base.series[i])
		} else {
			base.series[i].Release()
		}
	}
	base.series = newSeries
	base.DataChange()
}

func (base *BaseChart) Title() (ct *canvas.Text) {
	ct = base.title
	return
}

func (base *BaseChart) CartesianObjects() (canObj []fyne.CanvasObject) {
	// objects will be drawn in the same order as added here

	// first get all objects from the series
	lEntries := base.LegendEntries()
	for i := range lEntries {
		canObj = append(canObj, lEntries[i].Button, lEntries[i].Label)
	}
	rects := base.CartesianRects()
	for i := range rects {
		canObj = append(canObj, rects[i].Rect)
	}
	texts := base.CartesianTexts()
	for i := range texts {
		canObj = append(canObj, texts[i].Text)
	}

	// add chart title and axis titles
	if base.title.Text != "" {
		canObj = append(canObj, base.title)
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
	lEntries := base.LegendEntries()
	for i := range lEntries {
		canObj = append(canObj, lEntries[i].Button, lEntries[i].Label)
	}
	canObj = append(canObj, base.rast)
	texts := base.PolarTexts()
	for i := range texts {
		canObj = append(canObj, texts[i].Text)
	}

	// add chart title and axis titles
	if base.title.Text != "" {
		canObj = append(canObj, base.title)
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

func (base *BaseChart) LegendEntries() (les []renderer.LegendEntry) {
	if !base.legendVisible {
		return
	}
	for i := range base.series {
		les = append(les, base.series[i].LegendEntries()...)
	}
	return
}

func (base *BaseChart) ShowLegend() {
	base.legendVisible = true
	base.DataChange()
}

func (base *BaseChart) HideLegend() {
	base.legendVisible = false
	base.DataChange()
}

func (base *BaseChart) Tooltip() (tt renderer.Tooltip) {
	return
}

func (base *BaseChart) SetTitle(l string) {
	base.title.Text = l
}

func (base *BaseChart) SetTitleStyle(sizeName fyne.ThemeSizeName, colorName fyne.ThemeColorName) {
	base.titleSizeName = sizeName
	base.title.TextSize = theme.Size(sizeName)
	base.titleColorName = colorName
	base.title.Color = theme.Color(colorName)
}

func (base *BaseChart) FromAxisElements() (min float64, max float64, origin float64,
	label renderer.Label, ticks []renderer.Tick, arrow renderer.Arrow, show bool) {
	min, max = base.fromMin, base.fromMax
	origin = 0
	label = renderer.Label{}
	ticks = []renderer.Tick{}
	arrow = renderer.Arrow{}
	show = false
	return
}

func (base *BaseChart) ToAxisElements() (min float64, max float64, origin float64,
	label renderer.Label, ticks []renderer.Tick, arrow renderer.Arrow, show bool) {
	min, max = base.toMin, base.toMax
	origin = 0
	label = renderer.Label{}
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
