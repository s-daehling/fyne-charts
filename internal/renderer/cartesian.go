package renderer

import (
	"fyne.io/fyne/v2"
)

type CartesianChart interface {
	baseChart
	CartesianNodes() (ns []CartesianNode)
	CartesianEdges() (es []CartesianEdge)
	CartesianRects() (rs []CartesianRect)
	CartesianTexts() (ts []CartesianText)
	CartesianObjects() (obj []fyne.CanvasObject)
	CartesianOrientation() (trans bool)
}

// cartDrawingArea represents the area of the widget that can be used for the chart
type cartDrawingArea struct {
	minPos      fyne.Position // fyne position of the (hor,vert) coordinate (min,min)
	hmin        float64       // min coordinate of the horizontal axis
	vmin        float64       // min coordinate of the vertical axis
	maxPos      fyne.Position // fyne position of the (hor,vert) coordinates (max,max)
	hCoordToPos float32       // conversion factor from horizontal coordinate to fyne position X
	vCoordToPos float32       // conversion factor from vertical coordinate to fyne position Y
}

// Cartesian is the renderer for all cartesian plane widgets
type Cartesian struct {
	baseRenderer
	chart      CartesianChart
	transposed bool
}

func EmptyCartesianRenderer(chart CartesianChart) (r *Cartesian) {
	r = &Cartesian{
		baseRenderer: emptyBaseRenderer(),
		transposed:   false,
		chart:        chart,
	}
	return
}

// Layout is responsible for redrawing the chart widget; here the horizontal and vertical numerical coordinates are converted to fyne positions and objects are placed accordingly
func (r *Cartesian) Layout(size fyne.Size) {
	r.transposed = r.chart.CartesianOrientation()

	vAxisTickLabelWidth := float32(0.0)
	hAxisTickLabelHeight := float32(0.0)

	var vMin, vMax, vOrigin, hMin, hMax, hOrigin float64
	var vTicks, hTicks []Tick
	var vArrow, hArrow Arrow
	var vShow, hShow bool
	if r.transposed {
		vMin, vMax, vOrigin, vTicks, vArrow, vShow = r.chart.FromAxisElements()
		hMin, hMax, hOrigin, hTicks, hArrow, hShow = r.chart.ToAxisElements()
	} else {
		vMin, vMax, vOrigin, vTicks, vArrow, vShow = r.chart.ToAxisElements()
		hMin, hMax, hOrigin, hTicks, hArrow, hShow = r.chart.FromAxisElements()
	}
	_, hAxisTickLabelHeight = maxTickSize(hTicks)
	vAxisTickLabelWidth, _ = maxTickSize(vTicks)

	// determine the chart area
	var area cartDrawingArea
	area.hmin = hMin
	area.vmin = vMin
	area.minPos.X = r.margin + vAxisTickLabelWidth + r.tickLength -
		((size.Width - (r.tickLength + vAxisTickLabelWidth)) *
			float32((hOrigin-hMin)/(hMax-hMin)))
	if area.minPos.X < r.margin {
		area.minPos.X = r.margin
	}
	area.minPos.Y = size.Height - (r.margin + hAxisTickLabelHeight + r.tickLength -
		((size.Height - (r.tickLength + hAxisTickLabelHeight)) *
			float32((vOrigin-vMin)/(vMax-vMin))))
	if area.minPos.Y > size.Height-r.margin {
		area.minPos.Y = size.Height - r.margin
	}
	area.maxPos.X = size.Width - r.margin
	area.maxPos.Y = r.margin

	// update chart with available space
	r.chart.ChartSizeChange(area.maxPos.X-area.minPos.X, area.minPos.Y-area.maxPos.Y)

	if r.transposed {
		_, _, _, vTicks, _, _ = r.chart.FromAxisElements()
		_, _, _, hTicks, _, _ = r.chart.ToAxisElements()
	} else {
		_, _, _, vTicks, _, _ = r.chart.ToAxisElements()
		_, _, _, hTicks, _, _ = r.chart.FromAxisElements()
	}

	// calculate conversion factors from ccordinates to positions
	area.hCoordToPos = (area.maxPos.X - area.minPos.X) / float32(hMax-hMin)
	area.vCoordToPos = (area.minPos.Y - area.maxPos.Y) / float32(vMax-vMin)

	// Place horizontal-Axis from hMin to hMax
	if hShow {
		hArrow.Line.Position1 = cartesianCoordinatesToPosition(hMin, vOrigin, area)
		hArrow.Line.Position2 = cartesianCoordinatesToPosition(hMax, vOrigin, area)
		hArrow.HeadOne.Position1 = fyne.NewPos(hArrow.Line.Position2.X-10, hArrow.Line.Position2.Y-5)
		hArrow.HeadOne.Position2 = hArrow.Line.Position2
		hArrow.HeadTwo.Position1 = fyne.NewPos(hArrow.Line.Position2.X-10, hArrow.Line.Position2.Y+5)
		hArrow.HeadTwo.Position2 = hArrow.Line.Position2

		// place horizontal ticks
		for i := range hTicks {
			if hTicks[i].Line != nil {
				hTicks[i].Line.Position1 = cartesianCoordinatesToPosition(hTicks[i].NLine, vOrigin, area)
				hTicks[i].Line.Position2 = hTicks[i].Line.Position1.AddXY(0, 5)
			}
			if hTicks[i].SupLine != nil {
				hTicks[i].SupLine.Position1 = cartesianCoordinatesToPosition(hTicks[i].NLine, vMin, area)
				hTicks[i].SupLine.Position2 = cartesianCoordinatesToPosition(hTicks[i].NLine, vMax, area)
			}
			if hTicks[i].Label != nil {
				hTicks[i].Label.Move(cartesianCoordinatesToPosition(hTicks[i].NLabel, vOrigin, area).AddXY(-hTicks[i].Label.Size().Width/2, 5))
			}
		}
	}

	// Place vertical axis from vMin to vMax
	if vShow {
		vArrow.Line.Position1 = cartesianCoordinatesToPosition(hOrigin, vMin, area)
		vArrow.Line.Position2 = cartesianCoordinatesToPosition(hOrigin, vMax, area)
		vArrow.HeadOne.Position1 = fyne.NewPos(vArrow.Line.Position2.X-5, vArrow.Line.Position2.Y+10)
		vArrow.HeadOne.Position2 = vArrow.Line.Position2
		vArrow.HeadTwo.Position1 = fyne.NewPos(vArrow.Line.Position2.X+5, vArrow.Line.Position2.Y+10)
		vArrow.HeadTwo.Position2 = vArrow.Line.Position2

		// place vertical ticks
		for i := range vTicks {
			if vTicks[i].Line != nil {
				vTicks[i].Line.Position1 = cartesianCoordinatesToPosition(hOrigin,
					vTicks[i].NLine, area)
				vTicks[i].Line.Position2 = vTicks[i].Line.Position1.SubtractXY(5, 0)
			}
			if vTicks[i].SupLine != nil {
				vTicks[i].SupLine.Position1 = cartesianCoordinatesToPosition(hMin, vTicks[i].NLine, area)
				vTicks[i].SupLine.Position2 = cartesianCoordinatesToPosition(hMax, vTicks[i].NLine, area)
			}
			if vTicks[i].Label != nil {
				vTicks[i].Label.Move(cartesianCoordinatesToPosition(hOrigin,
					vTicks[i].NLabel, area).SubtractXY(5+vTicks[i].Label.Size().Width, vTicks[i].Label.Size().Height/2))
			}
		}
	}

	// place nodes
	ns := r.chart.CartesianNodes()
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
	es := r.chart.CartesianEdges()
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
	fs := r.chart.CartesianRects()
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
	ts := r.chart.CartesianTexts()
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
	rs := r.chart.Raster()
	if rs != nil {
		rs.Move(fyne.NewPos(area.minPos.X, area.maxPos.Y))
		rs.Resize(fyne.NewSize(area.maxPos.X-area.minPos.X, area.minPos.Y-area.maxPos.Y))
	}

	// place tooltip
	tt := r.chart.Tooltip()
	ttWidth, ttHeigth := tooltipSize(tt.Entries)
	ttPos := fyne.NewPos(area.minPos.X, area.maxPos.Y).AddXY(tt.X, tt.Y).SubtractXY(ttWidth+5, ttHeigth)
	if tt.Box != nil {
		tt.Box.Move(ttPos.SubtractXY(5, 0))
		tt.Box.Resize(fyne.NewSize(ttWidth+10, ttHeigth))
	}
	for i := range tt.Entries {
		tt.Entries[i].Move(ttPos)
		tt.Entries[i].Alignment = fyne.TextAlignLeading
		ttPos = ttPos.AddXY(0, tt.Entries[i].MinSize().Height)
	}

	// place overlay
	ov := r.chart.Overlay()
	if ov != nil {
		ov.Move(fyne.NewPos(area.minPos.X, area.maxPos.Y))
		ov.Resize(fyne.NewSize(area.maxPos.X-area.minPos.X, area.minPos.Y-area.maxPos.Y))
	}
}

// MinSize calculates the minimum space required to display the chart
func (r *Cartesian) MinSize() fyne.Size {
	minHeight := 2*r.margin + 20
	minWidth := 2*r.margin + 20
	return fyne.NewSize(minWidth, minHeight)
}

// Objects returns a list of all objects to be drawn
func (r *Cartesian) Objects() []fyne.CanvasObject {
	return r.chart.CartesianObjects()
}

// Refresh calls Layout if data of the chart has changes
func (r *Cartesian) Refresh() {
	// if r.chart.hasChanged() {
	r.chart.RefreshTheme()

	obj := r.Objects()
	for i := range obj {
		obj[i].Refresh()
	}

	r.Layout(r.chart.Size())
	// 	r.chart.resetHasChanged()
	// }
}

// cartesianCoordinatesToPosition converts a (h,v) coordinate to a fyne position
func cartesianCoordinatesToPosition(h float64, v float64, area cartDrawingArea) (pos fyne.Position) {
	pos.X = area.minPos.X + float32(h-area.hmin)*area.hCoordToPos
	pos.Y = area.minPos.Y - float32(v-area.vmin)*area.vCoordToPos
	return
}
