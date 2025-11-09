package renderer

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/software"
	"github.com/disintegration/imaging"
)

type CartesianChart interface {
	baseChart
	CartesianNodes() (ns []CartesianNode)
	CartesianEdges() (es []CartesianEdge)
	CartesianRects() (rs []CartesianRect)
	CartesianTexts() (ts []CartesianText)
	CartesianTooltip() (tt CartesianTooltip)
	CartesianObjects() (obj []fyne.CanvasObject)
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
	_, titleHeight, legendWidth, _ := r.placeTitleAndLegend(size, r.chart.Title(), r.chart.LegendEntries())
	vAxisLabelWidth := float32(0.0)
	hAxisLabelHeight := float32(0.0)
	vAxisTickLabelWidth := float32(0.0)
	hAxisTickLabelHeight := float32(0.0)

	var vMin, vMax, vOrigin, hMin, hMax, hOrigin float64
	var vLabel, hLabel Label
	var vTicks, hTicks []Tick
	var vArrow, hArrow Arrow
	var vShow, hShow bool
	if r.transposed {
		vMin, vMax, vOrigin, vLabel, vTicks, vArrow, vShow = r.chart.FromAxisElements()
		hMin, hMax, hOrigin, hLabel, hTicks, hArrow, hShow = r.chart.ToAxisElements()
	} else {
		vMin, vMax, vOrigin, vLabel, vTicks, vArrow, vShow = r.chart.ToAxisElements()
		hMin, hMax, hOrigin, hLabel, hTicks, hArrow, hShow = r.chart.FromAxisElements()
	}
	_, hAxisTickLabelHeight = maxTickSize(hTicks)
	vAxisTickLabelWidth, _ = maxTickSize(vTicks)

	if hShow && hLabel.Text.Text != "" {
		c := software.NewTransparentCanvas()
		c.SetPadded(false)
		c.SetContent(hLabel.Text)
		hLabel.Image.Image = c.Capture()
		hLabel.Image.Resize(hLabel.Text.MinSize())
		hLabel.Image.SetMinSize(hLabel.Text.MinSize())
		hAxisLabelHeight = hLabel.Image.MinSize().Height
	}

	if vShow && vLabel.Text.Text != "" {
		c := software.NewTransparentCanvas()
		c.SetPadded(false)
		c.SetContent(vLabel.Text)
		img := c.Capture()
		vLabel.Image.Image = imaging.Rotate90(img)
		vLabel.Image.Resize(fyne.NewSize(vLabel.Text.MinSize().Height, vLabel.Text.MinSize().Width))
		vLabel.Image.SetMinSize(fyne.NewSize(vLabel.Text.MinSize().Height, vLabel.Text.MinSize().Width))
		vAxisLabelWidth = vLabel.Image.MinSize().Width
	}

	// determine the chart area
	var area cartDrawingArea
	area.hmin = hMin
	area.vmin = vMin
	area.minPos.X = r.margin + vAxisLabelWidth + vAxisTickLabelWidth + r.tickLength -
		((size.Width - (r.tickLength + vAxisTickLabelWidth)) *
			float32((hOrigin-hMin)/(hMax-hMin)))
	if area.minPos.X < r.margin+vAxisLabelWidth {
		area.minPos.X = r.margin + vAxisLabelWidth
	}
	area.minPos.Y = size.Height - (r.margin + hAxisLabelHeight + hAxisTickLabelHeight + r.tickLength -
		((size.Height - (r.tickLength + hAxisTickLabelHeight)) *
			float32((vOrigin-vMin)/(vMax-vMin))))
	if area.minPos.Y > size.Height-r.margin-hAxisLabelHeight {
		area.minPos.Y = size.Height - r.margin - hAxisLabelHeight
	}
	area.maxPos.X = size.Width - r.margin - legendWidth
	area.maxPos.Y = r.margin + titleHeight

	// update chart with available space
	r.chart.ChartSizeChange(area.maxPos.X-area.minPos.X, area.minPos.Y-area.maxPos.Y)

	if r.transposed {
		_, _, _, _, vTicks, _, _ = r.chart.FromAxisElements()
		_, _, _, _, hTicks, _, _ = r.chart.ToAxisElements()
	} else {
		_, _, _, _, vTicks, _, _ = r.chart.ToAxisElements()
		_, _, _, _, hTicks, _, _ = r.chart.FromAxisElements()
	}

	// calculate conversion factors from ccordinates to positions
	area.hCoordToPos = (area.maxPos.X - area.minPos.X) / float32(hMax-hMin)
	area.vCoordToPos = (area.minPos.Y - area.maxPos.Y) / float32(vMax-vMin)

	// Place horizontal-Axis from hMin to hMax
	if hShow {
		if hLabel.Text.Text != "" {
			hLabel.Image.Move(fyne.NewPos(area.minPos.X+((area.maxPos.X-area.minPos.X)/2)-hLabel.Text.MinSize().Width/2,
				size.Height-hLabel.Text.MinSize().Height-r.margin))
		}
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
				hTicks[i].Label.Move(cartesianCoordinatesToPosition(hTicks[i].NLabel, vOrigin, area).AddXY(0, 5))
				hTicks[i].Label.Alignment = fyne.TextAlignCenter
			}
		}
	}

	// Place vertical axis from vMin to vMax
	if vShow {
		if vLabel.Text.Text != "" {
			vLabel.Image.Move(fyne.NewPos(r.margin, area.maxPos.Y+((area.minPos.Y-area.maxPos.Y)/2)-vLabel.Text.MinSize().Width/2))
		}
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
					vTicks[i].NLabel, area).SubtractXY(5, vTicks[i].Label.MinSize().Height/2))
				vTicks[i].Label.Alignment = fyne.TextAlignTrailing
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
	tt := r.chart.CartesianTooltip()
	for i := range tt.Entries {
		tt.Entries[i].Move(cartesianCoordinatesToPosition(tt.X, tt.Y, area))
		tt.Entries[i].Alignment = fyne.TextAlignTrailing
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
	titleWidth := float32(0.0)
	titleHeight := float32(0.0)
	legendWidth := float32(0.0)
	legendHeight := float32(0.0)
	vAxisLabelWidth := float32(0.0)
	vAxisLabelHeight := float32(0)
	hAxisLabelWidth := float32(0)
	hAxisLabelHeight := float32(0.0)

	ct := r.chart.Title()
	if ct != nil {
		if ct.Text != "" {
			titleWidth = ct.MinSize().Width
			titleHeight = ct.MinSize().Height
		}
	}

	les := r.chart.LegendEntries()
	legendWidth, legendHeight = legendSize(les)

	var vLabel, hLabel Label
	var vShow, hShow bool
	if r.transposed {
		_, _, _, vLabel, _, _, vShow = r.chart.FromAxisElements()
		_, _, _, hLabel, _, _, hShow = r.chart.ToAxisElements()
	} else {
		_, _, _, vLabel, _, _, vShow = r.chart.ToAxisElements()
		_, _, _, hLabel, _, _, hShow = r.chart.FromAxisElements()
	}

	if hShow && hLabel.Text.Text != "" {
		hAxisLabelWidth = hLabel.Image.MinSize().Width
		hAxisLabelHeight = hLabel.Image.MinSize().Height
	}
	if vShow && vLabel.Text.Text != "" {
		vAxisLabelWidth = vLabel.Image.MinSize().Width
		vAxisLabelHeight = vLabel.Image.MinSize().Height
	}

	minHeight := 2*r.margin + titleHeight
	if legendHeight > 20+vAxisLabelHeight+hAxisLabelHeight {
		minHeight += legendHeight
	} else {
		minHeight += 20 + vAxisLabelHeight + hAxisLabelHeight
	}

	minWidth := 2 * r.margin
	if titleWidth > 20+vAxisLabelWidth+hAxisLabelWidth+legendWidth {
		minWidth += titleWidth
	} else {
		minWidth += 20 + vAxisLabelWidth + hAxisLabelWidth + legendWidth
	}

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

	r.Layout(r.chart.WidgetSize())
	// 	r.chart.resetHasChanged()
	// }
}

// cartesianCoordinatesToPosition converts a (h,v) coordinate to a fyne position
func cartesianCoordinatesToPosition(h float64, v float64, area cartDrawingArea) (pos fyne.Position) {
	pos.X = area.minPos.X + float32(h-area.hmin)*area.hCoordToPos
	pos.Y = area.minPos.Y - float32(v-area.vmin)*area.vCoordToPos
	return
}
