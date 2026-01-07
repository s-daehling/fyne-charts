package coord

import (
	"fmt"
	"image/color"
	"math"
	"strconv"

	"github.com/s-daehling/fyne-charts/internal/coord/axis"
	"github.com/s-daehling/fyne-charts/internal/coord/series"
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

type FromType string

const (
	Numerical   FromType = "Numerical"
	Temporal    FromType = "Temporal"
	Categorical FromType = "Categorical"
)

type BaseChart struct {
	widget.BaseWidget
	title             *canvas.Text
	titleStyle        style.LabelStyle
	fromAx            *axis.Axis
	toAx              *axis.Axis
	series            []series.Series
	overlay           *interact.Overlay
	tooltip           *interact.Tooltip
	changed           bool
	autoFromRange     bool
	autoToRange       bool
	autoOrigin        bool
	legend            *interact.Legend
	tooltipVisible    bool
	planeType         PlaneType
	transposed        bool
	fromType          FromType
	rast              *canvas.Raster
	render            fyne.WidgetRenderer
	mainCont          *fyne.Container
	hLabelCont        *fyne.Container
	hLabelLeftSpacer  *canvas.Rectangle
	hLabelRightSpacer *canvas.Rectangle
	vLabelCont        *fyne.Container
	rLegendCont       *fyne.Container
	lLegendCont       *fyne.Container
	bLegendCont       *fyne.Container
	tLegendCont       *fyne.Container
}

func EmptyBaseChart(pType PlaneType, fType FromType) (base *BaseChart) {
	base = &BaseChart{
		title:             canvas.NewText("", theme.Color(theme.ColorNameForeground)),
		tooltip:           interact.NewTooltip(),
		changed:           false,
		autoFromRange:     true,
		autoToRange:       true,
		autoOrigin:        true,
		legend:            interact.NewLegend(),
		tooltipVisible:    false,
		planeType:         pType,
		transposed:        false,
		fromType:          fType,
		hLabelCont:        container.NewHBox(),
		hLabelLeftSpacer:  canvas.NewRectangle(color.Alpha16{}),
		hLabelRightSpacer: canvas.NewRectangle(color.Alpha16{}),
		vLabelCont:        container.NewVBox(),
		rLegendCont:       container.NewCenter(),
		lLegendCont:       container.NewCenter(),
		bLegendCont:       container.NewStack(),
		tLegendCont:       container.NewStack(),
	}
	base.mainCont = container.NewBorder(
		container.NewVBox(
			base.title,
			base.tLegendCont),
		container.NewVBox(
			base.hLabelCont,
			base.bLegendCont),
		container.NewHBox(
			base.lLegendCont,
			base.vLabelCont),
		base.rLegendCont,
		base)
	base.overlay = interact.NewOverlay(base)
	base.hLabelLeftSpacer.SetMinSize(fyne.NewSize(0, 0))
	base.hLabelRightSpacer.SetMinSize(fyne.NewSize(0, 0))
	if pType == CartesianPlane {
		base.fromAx = axis.EmptyAxis("", axis.CartesianHorAxis)
		base.toAx = axis.EmptyAxis("", axis.CartesianVertAxis)
		base.rast = canvas.NewRasterWithPixels(base.PixelGenCartesian)
		base.vLabelCont.Add(base.toAx.Label())
		base.hLabelCont.Add(base.fromAx.Label())
	} else {
		base.fromAx = axis.EmptyAxis("", axis.PolarPhiAxis)
		base.toAx = axis.EmptyAxis("", axis.PolarRAxis)
		base.rast = canvas.NewRasterWithPixels(base.PixelGenPolar)
		base.vLabelCont.Add(base.fromAx.Label())
		base.hLabelCont.Add(base.toAx.Label())
	}
	base.SetTitleStyle(style.DefaultTitleStyle())
	base.SetFromAxisStyle(style.DefaultAxisStyle())
	base.SetFromAxisLabelStyle(style.DefaultAxisLabelStyle())
	base.SetToAxisStyle(style.DefaultAxisStyle())
	base.SetToAxisLabelStyle(style.DefaultAxisLabelStyle())
	base.SetLegendStyle(style.LegendLocationRight)
	base.updateRangeAndOrigin()
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
		base.fromAx.CartesianTranspose()
		base.toAx.CartesianTranspose()
		base.hLabelCont.RemoveAll()
		base.vLabelCont.RemoveAll()
		base.hLabelCont.Add(base.hLabelLeftSpacer)
		if transposed {
			base.fromAx.AddLabelToContainer(base.vLabelCont)
			base.toAx.AddLabelToContainer(base.hLabelCont)
		} else {
			base.fromAx.AddLabelToContainer(base.hLabelCont)
			base.toAx.AddLabelToContainer(base.vLabelCont)
		}
		base.hLabelCont.Add(base.hLabelRightSpacer)
		base.updateHLabelSpacer()
		base.mainCont.Refresh()
	}
}

func (base *BaseChart) CartesianOrientation() (transposed bool) {
	transposed = base.transposed
	return
}

func (base *BaseChart) PolarOrientation() (rot float64, mathPos bool) {
	rot = 0
	mathPos = true
	return
}

func (base *BaseChart) CartesianObjects() (canObj []fyne.CanvasObject) {
	// objects will be drawn in the same order as added here

	// first get all objects from the series
	canObj = append(canObj, base.rast)
	rects := base.CartesianRects()
	for i := range rects {
		canObj = append(canObj, rects[i].Rect)
	}
	edges := base.CartesianEdges()
	for i := range edges {
		canObj = append(canObj, edges[i].Line)
	}
	nodes := base.CartesianNodes()
	for i := range nodes {
		canObj = append(canObj, nodes[i].Dot)
	}
	texts := base.CartesianTexts()
	for i := range texts {
		canObj = append(canObj, texts[i].Text)
	}

	// add axis elements
	canObj = append(canObj, base.fromAx.Objects()...)
	canObj = append(canObj, base.toAx.Objects()...)

	if base.tooltipVisible {
		// add tooltip
		tt := base.Tooltip()
		if tt.Box != nil {
			canObj = append(canObj, tt.Box)
		}
		for i := range tt.Entries {
			canObj = append(canObj, tt.Entries[i])
		}

		// add overlay
		canObj = append(canObj, base.overlay)
	}
	return
}

func (base *BaseChart) CartesianNodes() (ns []renderer.CartesianNode) {
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	for i := range base.series {
		ns = append(ns, base.series[i].CartesianNodes(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (base *BaseChart) CartesianEdges() (es []renderer.CartesianEdge) {
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	for i := range base.series {
		es = append(es, base.series[i].CartesianEdges(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (base *BaseChart) CartesianRects() (as []renderer.CartesianRect) {
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	for i := range base.series {
		as = append(as, base.series[i].CartesianRects(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (base *BaseChart) CartesianTexts() (ts []renderer.CartesianText) {
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	for i := range base.series {
		ts = append(ts, base.series[i].CartesianTexts(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (base *BaseChart) PolarObjects() (canObj []fyne.CanvasObject) {
	// objects will be drawn in the same order as added here

	// first get all objects from the series
	canObj = append(canObj, base.rast)
	edges := base.PolarEdges()
	for i := range edges {
		canObj = append(canObj, edges[i].Line)
	}
	nodes := base.PolarNodes()
	for i := range nodes {
		canObj = append(canObj, nodes[i].Dot)
	}
	texts := base.PolarTexts()
	for i := range texts {
		canObj = append(canObj, texts[i].Text)
	}

	// add axis elements
	canObj = append(canObj, base.fromAx.Objects()...)
	canObj = append(canObj, base.toAx.Objects()...)

	if base.tooltipVisible {
		// add tooltip
		tt := base.Tooltip()
		if tt.Box != nil {
			canObj = append(canObj, tt.Box)
		}
		for i := range tt.Entries {
			canObj = append(canObj, tt.Entries[i])
		}

		// add overlay
		canObj = append(canObj, base.overlay)
	}
	return
}

func (base *BaseChart) PolarNodes() (ns []renderer.PolarNode) {
	phiMin, phiMax := base.fromAx.NRange()
	rMin, rMax := base.toAx.NRange()
	for i := range base.series {
		ns = append(ns, base.series[i].PolarNodes(phiMin, phiMax, rMin, rMax)...)
	}
	return
}

func (base *BaseChart) PolarEdges() (es []renderer.PolarEdge) {
	phiMin, phiMax := base.fromAx.NRange()
	rMin, rMax := base.toAx.NRange()
	for i := range base.series {
		es = append(es, base.series[i].PolarEdges(phiMin, phiMax, rMin, rMax)...)
	}
	return
}

func (base *BaseChart) PolarTexts() (ts []renderer.PolarText) {
	phiMin, phiMax := base.fromAx.NRange()
	rMin, rMax := base.toAx.NRange()
	for i := range base.series {
		ts = append(ts, base.series[i].PolarTexts(phiMin, phiMax, rMin, rMax)...)
	}
	return
}

func (base *BaseChart) Raster() (rs *canvas.Raster) {
	rs = base.rast
	return
}

func (base *BaseChart) Overlay() (io *interact.Overlay) {
	io = base.overlay
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
	base.updateHLabelSpacer()
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
	tt.X, tt.Y, tt.Entries, tt.Box = base.tooltip.GetEntries()
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

func (base *BaseChart) SetTitleStyle(ts style.LabelStyle) {
	base.titleStyle = ts
	base.title.Alignment = ts.Alignment
	base.title.TextSize = theme.Size(ts.SizeName)
	base.title.Color = theme.Color(ts.ColorName)
	base.title.TextStyle = ts.TextStyle
	base.title.Refresh()
}

func (base *BaseChart) MouseIn(pX, pY, w, h, absX, absY float32) {
	if base.planeType == CartesianPlane {
		x, y := base.PositionToCartesianCoordinates(pX, pY, w, h)
		base.tooltip.MouseIn(pX, pY)
		text := ""
		switch base.fromType {
		case Numerical:
			text = fmt.Sprintf("x: %s, y: %s", strconv.FormatFloat(x, 'f', base.fromAx.NTipPrecision(), 64), strconv.FormatFloat(y, 'f', base.toAx.NTipPrecision(), 64))
		case Temporal:
			text = fmt.Sprintf("t: %s, y: %s", base.fromAx.NtoT(x).Format(base.fromAx.TTipFormat()), strconv.FormatFloat(y, 'f', base.toAx.NTipPrecision(), 64))
		case Categorical:
			text = fmt.Sprintf("c: %s, y: %s", base.fromAx.NtoC(x), strconv.FormatFloat(y, 'f', base.toAx.NTipPrecision(), 64))
		}
		base.tooltip.SetEntries([]string{text})
	} else {
		phi, r, _, _ := base.PositionToPolarCoordinates(pX, pY, w, h)
		base.tooltip.MouseIn(pX, pY)
		text := ""
		switch base.fromType {
		case Numerical:
			text = fmt.Sprintf("phi: %s, r: %s", strconv.FormatFloat(phi, 'f', base.fromAx.NTipPrecision(), 64), strconv.FormatFloat(r, 'f', base.toAx.NTipPrecision(), 64))
		case Temporal:
			text = fmt.Sprintf("t: %s, r: %s", base.fromAx.NtoT(phi).Format(base.fromAx.TTipFormat()), strconv.FormatFloat(r, 'f', base.toAx.NTipPrecision(), 64))
		case Categorical:
			text = fmt.Sprintf("c: %s, r: %s", base.fromAx.NtoC(phi), strconv.FormatFloat(r, 'f', base.toAx.NTipPrecision(), 64))
		}
		base.tooltip.SetEntries([]string{text})
	}
	base.Refresh()
}

func (base *BaseChart) MouseMove(pX, pY, w, h, absX, absY float32) {
	if base.planeType == CartesianPlane {
		x, y := base.PositionToCartesianCoordinates(pX, pY, w, h)
		c := base.tooltip.MouseMove(pX, pY)
		if c > 3 {
			text := ""
			switch base.fromType {
			case Numerical:
				text = fmt.Sprintf("x: %s, y: %s", strconv.FormatFloat(x, 'f', base.fromAx.NTipPrecision(), 64), strconv.FormatFloat(y, 'f', base.toAx.NTipPrecision(), 64))
			case Temporal:
				text = fmt.Sprintf("t: %s, y: %s", base.fromAx.NtoT(x).Format(base.fromAx.TTipFormat()), strconv.FormatFloat(y, 'f', base.toAx.NTipPrecision(), 64))
			case Categorical:
				text = fmt.Sprintf("c: %s, y: %s", base.fromAx.NtoC(x), strconv.FormatFloat(y, 'f', base.toAx.NTipPrecision(), 64))
			}
			base.tooltip.SetEntries([]string{text})
			base.Refresh()
		}
	} else {
		phi, r, _, _ := base.PositionToPolarCoordinates(pX, pY, w, h)
		c := base.tooltip.MouseMove(pX, pY)
		if c > 3 {
			text := ""
			switch base.fromType {
			case Numerical:
				text = fmt.Sprintf("phi: %s, r: %s", strconv.FormatFloat(phi, 'f', base.fromAx.NTipPrecision(), 64), strconv.FormatFloat(r, 'f', base.toAx.NTipPrecision(), 64))
			case Temporal:
				text = fmt.Sprintf("t: %s, r: %s", base.fromAx.NtoT(phi).Format(base.fromAx.TTipFormat()), strconv.FormatFloat(r, 'f', base.toAx.NTipPrecision(), 64))
			case Categorical:
				text = fmt.Sprintf("c: %s, r: %s", base.fromAx.NtoC(phi), strconv.FormatFloat(r, 'f', base.toAx.NTipPrecision(), 64))
			}
			base.tooltip.SetEntries([]string{text})
			base.Refresh()
		}
	}
}

func (base *BaseChart) MouseOut() {
	base.tooltip.MouseOut()
	base.Refresh()
}

func (base *BaseChart) PixelGenCartesian(pX, pY, w, h int) (col color.Color) {
	x, y := base.PositionToCartesianCoordinates(float32(pX), float32(pY), float32(w), float32(h))
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	for i := range base.series {
		serCol := base.series[i].RasterColorCartesian(x, y)
		r, g, b, _ := serCol.RGBA()
		if r > 0 || g > 0 || b > 0 {
			col = serCol
			break
		}
	}
	return
}

func (base *BaseChart) PixelGenPolar(pX, pY, w, h int) (col color.Color) {
	phi, r, x, y := base.PositionToPolarCoordinates(float32(pX), float32(pY), float32(w), float32(h))
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	_, rMax := base.toAx.NRange()
	if r > rMax {
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

func (base *BaseChart) PositionToCartesianCoordinates(pX float32, pY float32, w float32, h float32) (x float64, y float64) {
	xMin, xMax := base.fromAx.NRange()
	yMin, yMax := base.toAx.NRange()
	if base.transposed {
		x = xMin + ((float64(h-pY) / float64(h)) * (xMax - xMin))
		y = yMin + ((float64(pX) / float64(w)) * (yMax - yMin))
	} else {
		x = xMin + ((float64(pX) / float64(w)) * (xMax - xMin))
		y = yMin + ((float64(h-pY) / float64(h)) * (yMax - yMin))
	}
	return
}

func (base *BaseChart) PositionToPolarCoordinates(pX float32, pY float32, w float32, h float32) (phi float64,
	r float64, x float64, y float64) {
	_, rMax := base.toAx.NRange()
	rot := 0.0
	mathPos := true
	posToCoord := rMax / (float64(w) / 2.0)
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
