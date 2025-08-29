package chart

import (
	"github.com/s-daehling/fyne-charts/internal/series"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/software"
	"github.com/disintegration/imaging"
)

// cartDrawingArea represents the area of the widget that can be used for the chart
type cartDrawingArea struct {
	minPos      fyne.Position // fyne position of the (hor,vert) coordinate (min,min)
	hmin        float64       // min coordinate of the horizontal axis
	vmin        float64       // min coordinate of the vertical axis
	maxPos      fyne.Position // fyne position of the (hor,vert) coordinates (max,max)
	hCoordToPos float32       // conversion factor from horizontal coordinate to fyne position X
	vCoordToPos float32       // conversion factor from vertical coordinate to fyne position Y
}

// cartesianRenderer is the renderer for all cartesian plane widgets
type cartesianRenderer struct {
	chart      *BaseChart // reference to the chart widget
	margin     float32    // free space to the border
	tickLength float32    // length of ticks
	transposed bool
}

func EmptyCartesianRenderer(chart *BaseChart) (r *cartesianRenderer) {
	r = &cartesianRenderer{
		chart:      chart,
		margin:     10.0,
		tickLength: 5.0,
		transposed: false,
	}
	return
}

// Destroy has nothing to do
func (r *cartesianRenderer) Destroy() {}

// Layout is responsible for redrawing the chart widget; here the horizontal and vertical numerical coordinates are converted to fyne positions and objects are placed accordingly
func (r *cartesianRenderer) Layout(size fyne.Size) {
	r.chart.resize(size.Width, size.Height)
	// todo: should chart be locked from this point on?
	ct := r.chart.title()
	if ct.Name != "" {
		ct.Label.Text = ct.Name
		ct.Label.Move(fyne.NewPos(size.Width/2-ct.Label.MinSize().Width/2, r.margin))
	}
	var vAxis, hAxis *Axis
	if r.transposed {
		vAxis = r.chart.fromAxis()
		hAxis = r.chart.toAxis()
	} else {
		vAxis = r.chart.toAxis()
		hAxis = r.chart.fromAxis()
	}
	hAxLabel, hAxText := hAxis.Label()
	if hAxText.Text != "" {
		// l := canvas.NewText(hAxis.name, color.Black)
		c := software.NewTransparentCanvas()
		c.SetPadded(false)
		c.SetContent(hAxText)
		hAxLabel.Image = c.Capture()
		hAxLabel.Resize(hAxText.MinSize())
		hAxLabel.SetMinSize(hAxText.MinSize())
	}

	vAxLabel, vAxText := vAxis.Label()
	if vAxText.Text != "" {
		// l := canvas.NewText(vAxis.name, color.Black)
		c := software.NewTransparentCanvas()
		c.SetPadded(false)
		c.SetContent(vAxText)
		img := c.Capture()
		vAxLabel.Image = imaging.Rotate90(img)
		vAxLabel.Resize(fyne.NewSize(vAxText.MinSize().Height, vAxText.MinSize().Width))
		vAxLabel.SetMinSize(fyne.NewSize(vAxText.MinSize().Height, vAxText.MinSize().Width))
	}

	legendVisible := r.chart.legendVisibility()
	les := r.chart.legendEntries()
	legendWidth := float32(0.0)
	if legendVisible {
		legendWidth = series.LegendWidth(les)
	}

	// determine the chart area
	var area cartDrawingArea
	area.hmin = hAxis.nMin
	area.vmin = vAxis.nMin

	vLabelWidth := r.maxVertLabelWidth()
	area.minPos.X = r.margin + r.tickLength + vLabelWidth -
		((size.Width - (r.tickLength + vLabelWidth)) *
			float32((hAxis.nOrigin-hAxis.nMin)/(hAxis.nMax-hAxis.nMin)))
	if area.minPos.X < r.margin {
		area.minPos.X = r.margin
	}
	if vAxis.name != "" {
		area.minPos.X += vAxis.label.MinSize().Width
	}
	hLabelHeight := r.maxHorLabelHeight()
	area.minPos.Y = size.Height - (r.margin + r.tickLength + hLabelHeight -
		((size.Height - (r.tickLength + hLabelHeight)) *
			float32((vAxis.nOrigin-vAxis.nMin)/(vAxis.nMax-vAxis.nMin))))
	if area.minPos.Y > size.Height-r.margin {
		area.minPos.Y = size.Height - r.margin
	}
	if hAxis.name != "" {
		area.minPos.Y -= hAxis.label.MinSize().Height
	}
	area.maxPos.X = size.Width - r.margin - legendWidth
	area.maxPos.Y = r.margin
	if ct.Name != "" {
		area.maxPos.Y += ct.Label.MinSize().Height
	}

	// calculate conversion factors from ccordinates to positions
	area.hCoordToPos = (area.maxPos.X - area.minPos.X) / float32(hAxis.nMax-hAxis.nMin)
	area.vCoordToPos = (area.minPos.Y - area.maxPos.Y) / float32(vAxis.nMax-vAxis.nMin)

	// Place horizontal-Axis from hMin to hMax
	if hAxText.Text != "" {
		hAxLabel.Move(fyne.NewPos(area.minPos.X+((area.maxPos.X-area.minPos.X)/2)-hAxText.MinSize().Width/2,
			size.Height-hAxText.MinSize().Height-r.margin))
	}
	hAxis.line.Position1 = cartesianCoordinatesToPosition(hAxis.nMin, vAxis.nOrigin, area)
	hAxis.line.Position2 = cartesianCoordinatesToPosition(hAxis.nMax, vAxis.nOrigin, area)
	hAxis.arrowOne.Position1 = fyne.NewPos(hAxis.line.Position2.X-10, hAxis.line.Position2.Y-5)
	hAxis.arrowOne.Position2 = hAxis.line.Position2
	hAxis.arrowTwo.Position1 = fyne.NewPos(hAxis.line.Position2.X-10, hAxis.line.Position2.Y+5)
	hAxis.arrowTwo.Position2 = hAxis.line.Position2

	// place horizontal ticks
	th := hAxis.Ticks()
	for i := range th {
		if th[i].Line != nil {
			th[i].Line.Position1 = cartesianCoordinatesToPosition(th[i].NLine, vAxis.nOrigin, area)
			th[i].Line.Position2 = th[i].Line.Position1.AddXY(0, 5)
		}
		if th[i].SupLine != nil {
			th[i].SupLine.Position1 = cartesianCoordinatesToPosition(th[i].NLine, vAxis.nMin, area)
			th[i].SupLine.Position2 = cartesianCoordinatesToPosition(th[i].NLine, vAxis.nMax, area)
		}
		if th[i].Label != nil {
			th[i].Label.Move(cartesianCoordinatesToPosition(th[i].NLabel, vAxis.nOrigin, area).AddXY(0, 5))
			th[i].Label.Alignment = fyne.TextAlignCenter
		}
	}

	// Place vertical axis from vMin to vMax
	if vAxText.Text != "" {
		vAxLabel.Move(fyne.NewPos(r.margin, area.maxPos.Y+((area.minPos.Y-area.maxPos.Y)/2)-vAxText.MinSize().Width/2))
	}
	vAxis.line.Position1 = cartesianCoordinatesToPosition(hAxis.nOrigin, vAxis.nMin, area)
	vAxis.line.Position2 = cartesianCoordinatesToPosition(hAxis.nOrigin, vAxis.nMax, area)
	vAxis.arrowOne.Position1 = fyne.NewPos(vAxis.line.Position2.X-5, vAxis.line.Position2.Y+10)
	vAxis.arrowOne.Position2 = vAxis.line.Position2
	vAxis.arrowTwo.Position1 = fyne.NewPos(vAxis.line.Position2.X+5, vAxis.line.Position2.Y+10)
	vAxis.arrowTwo.Position2 = vAxis.line.Position2

	// place vertical ticks
	vh := vAxis.Ticks()
	for i := range vh {
		if vh[i].Line != nil {
			vh[i].Line.Position1 = cartesianCoordinatesToPosition(hAxis.nOrigin,
				vh[i].NLine, area)
			vh[i].Line.Position2 = vh[i].Line.Position1.SubtractXY(5, 0)
		}
		if vh[i].SupLine != nil {
			vh[i].SupLine.Position1 = cartesianCoordinatesToPosition(hAxis.nMin, vh[i].NLine, area)
			vh[i].SupLine.Position2 = cartesianCoordinatesToPosition(hAxis.nMax, vh[i].NLine, area)
		}
		if vh[i].Label != nil {
			vh[i].Label.Move(cartesianCoordinatesToPosition(hAxis.nOrigin,
				vh[i].NLabel, area).SubtractXY(5, vh[i].Label.MinSize().Height/2))
			vh[i].Label.Alignment = fyne.TextAlignTrailing
		}
	}

	// place nodes
	ns := r.chart.cartesianNodes()
	for i := range ns {
		var dotPos fyne.Position
		if r.transposed {
			dotPos = cartesianCoordinatesToPosition(ns[i].Y, ns[i].X, area)
		} else {
			dotPos = cartesianCoordinatesToPosition(ns[i].X, ns[i].Y, area)
		}
		dotPos = dotPos.SubtractXY(ns[i].Dot.Size().Width/2.0, ns[i].Dot.Size().Height/2.0)
		ns[i].Dot.Move(dotPos)
	}

	// place edges
	es := r.chart.cartesianEdges()
	for i := range es {
		if r.transposed {
			es[i].Line.Position1 = cartesianCoordinatesToPosition(es[i].Y1, es[i].X1, area)
			es[i].Line.Position2 = cartesianCoordinatesToPosition(es[i].Y2, es[i].X2, area)
		} else {
			es[i].Line.Position1 = cartesianCoordinatesToPosition(es[i].X1, es[i].Y1, area)
			es[i].Line.Position2 = cartesianCoordinatesToPosition(es[i].X2, es[i].Y2, area)
		}
	}

	// place rects
	fs := r.chart.cartesianRects()
	for i := range fs {
		if r.transposed {
			p1 := cartesianCoordinatesToPosition(fs[i].Y1, fs[i].X2, area)
			p2 := cartesianCoordinatesToPosition(fs[i].Y2, fs[i].X1, area)
			fs[i].Rect.Move(p1)
			fs[i].Rect.Resize(fyne.NewSize(p2.X-p1.X, p2.Y-p1.Y))
		} else {
			p1 := cartesianCoordinatesToPosition(fs[i].X1, fs[i].Y2, area)
			p2 := cartesianCoordinatesToPosition(fs[i].X2, fs[i].Y1, area)
			fs[i].Rect.Move(p1)
			fs[i].Rect.Resize(fyne.NewSize(p2.X-p1.X, p2.Y-p1.Y))
		}
	}

	// place texts
	ts := r.chart.cartesianTexts()
	for i := range ts {
		if r.transposed {

		} else {
			tPos := cartesianCoordinatesToPosition(ts[i].X, ts[i].Y, area)
			tPos = tPos.SubtractXY(0, ts[i].Text.MinSize().Height/2)
			ts[i].Text.Move(tPos)
			ts[i].Text.Alignment = fyne.TextAlignCenter
		}
	}

	// place raster
	rs := r.chart.chartRaster()
	rs.Move(fyne.NewPos(area.minPos.X, area.maxPos.Y))
	rs.Resize(fyne.NewSize(area.maxPos.X-area.minPos.X, area.minPos.Y-area.maxPos.Y))

	// place legend entries
	if legendVisible {
		yLegend := (size.Height - float32(len(les)*20)) / 2.0
		for i := range les {
			subOffset := float32(0.0)
			if les[i].IsSub {
				subOffset = 20
			}
			les[i].Button.Resize(fyne.NewSize(15, 15))
			les[i].Button.Move(fyne.NewPos(area.maxPos.X+5+subOffset, yLegend+20*float32(i)))
			les[i].Label.Move(fyne.NewPos(area.maxPos.X+25+subOffset, yLegend+20*float32(i)))
		}
	}
}

// MinSize calculates the minimum space required to display the chart
func (r *cartesianRenderer) MinSize() fyne.Size {
	legendWidth := float32(0)
	if r.chart.legendVisibility() {
		legendWidth = series.LegendWidth(r.chart.legendEntries())
	}
	minWidth := 2*r.margin + r.maxVertLabelWidth() + legendWidth

	if !r.transposed && r.chart.fixedFromTicks() {
		ts := r.chart.fromAxis().Ticks()
		for i := range ts {
			minWidth += ts[i].Label.MinSize().Width
		}
	} else if r.transposed && r.chart.fixedToTicks() {
		ts := r.chart.toAxis().Ticks()
		for i := range ts {
			minWidth += ts[i].Label.MinSize().Width
		}
	}
	minHeight := 2*r.margin + r.maxHorLabelHeight()
	if !r.transposed && r.chart.fixedToTicks() {
		ts := r.chart.toAxis().Ticks()
		for i := range ts {
			minHeight += ts[i].Label.MinSize().Height
		}
	} else if r.transposed && r.chart.fixedFromTicks() {
		ts := r.chart.fromAxis().Ticks()
		for i := range ts {
			minHeight += ts[i].Label.MinSize().Height
		}
	}
	return fyne.NewSize(minWidth, minHeight)
}

// Objects returns a list of all objects to be drawn
func (r *cartesianRenderer) Objects() []fyne.CanvasObject {
	return r.chart.cartesianObjects()
}

// Refresh calls Layout if data of the chart has changes
func (r *cartesianRenderer) Refresh() {
	// if r.chart.hasChanged() {

	obj := r.Objects()
	for i := range obj {
		obj[i].Refresh()
	}
	// 	r.chart.resetHasChanged()
	// }
}

// maxVertLabelWidth returns the maximum width of the vertical axis labels
func (r *cartesianRenderer) maxVertLabelWidth() (maxWidth float32) {
	var vAxis *Axis
	if r.transposed {
		vAxis = r.chart.fromAxis()
	} else {
		vAxis = r.chart.toAxis()
	}
	maxWidth = vAxis.MaxTickWidth()
	return
}

// maxHorLabelHeight returns the maximum height of the horizontal axis labels
func (r *cartesianRenderer) maxHorLabelHeight() (maxHeight float32) {
	var hAxis *Axis
	if r.transposed {
		hAxis = r.chart.toAxis()
	} else {
		hAxis = r.chart.fromAxis()
	}
	maxHeight = hAxis.MaxTickHeight()
	return
}

// cartesianCoordinatesToPosition converts a (h,v) coordinate to a fyne position
func cartesianCoordinatesToPosition(h float64, v float64, area cartDrawingArea) (pos fyne.Position) {
	pos.X = area.minPos.X + float32(h-area.hmin)*area.hCoordToPos
	pos.Y = area.minPos.Y - float32(v-area.vmin)*area.vCoordToPos
	return
}

// func positionToCartesianCoordinates(pX int, pY int, w int, h int, xMin float64, xMax float64,
// 	yMin float64, yMax float64, trans bool) (x float64, y float64) {
// 	if trans {
// 		x = xMin + ((float64(h-pY) / float64(h)) * (xMax - xMin))
// 		y = yMin + ((float64(pX) / float64(w)) * (yMax - yMin))
// 	} else {
// 		x = xMin + ((float64(pX) / float64(w)) * (xMax - xMin))
// 		y = yMin + ((float64(h-pY) / float64(h)) * (yMax - yMin))
// 	}
// 	return
// }
