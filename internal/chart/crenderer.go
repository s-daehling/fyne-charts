package chart

import (
	"github.com/s-daehling/fyne-charts/internal/axis"
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
	baseRenderer
	transposed bool
}

func EmptyCartesianRenderer(chart *BaseChart) (r *cartesianRenderer) {
	r = &cartesianRenderer{
		baseRenderer: emptyBaseRenderer(chart),
		transposed:   false,
	}
	return
}

// Layout is responsible for redrawing the chart widget; here the horizontal and vertical numerical coordinates are converted to fyne positions and objects are placed accordingly
func (r *cartesianRenderer) Layout(size fyne.Size) {
	// todo: should chart be locked from this point on?

	_, titleHeight, legendWidth, _ := r.placeTitleAndLegend(size)
	vAxisLabelWidth := float32(0.0)
	hAxisLabelHeight := float32(0.0)
	vAxisTickLabelWidth := float32(0.0)
	hAxisTickLabelHeight := float32(0.0)

	var vAxis, hAxis *axis.Axis
	if r.transposed {
		vAxis = r.chart.fromAxis()
		hAxis = r.chart.toAxis()
	} else {
		vAxis = r.chart.toAxis()
		hAxis = r.chart.fromAxis()
	}
	hAxisTickLabelHeight = hAxis.MaxTickHeight()
	vAxisTickLabelWidth = vAxis.MaxTickWidth()

	hAxLabel, hAxText := hAxis.Label()
	if hAxText.Text != "" {
		c := software.NewTransparentCanvas()
		c.SetPadded(false)
		c.SetContent(hAxText)
		hAxLabel.Image = c.Capture()
		hAxLabel.Resize(hAxText.MinSize())
		hAxLabel.SetMinSize(hAxText.MinSize())
		hAxisLabelHeight = hAxLabel.MinSize().Height
	}

	vAxLabel, vAxText := vAxis.Label()
	if vAxText.Text != "" {
		c := software.NewTransparentCanvas()
		c.SetPadded(false)
		c.SetContent(vAxText)
		img := c.Capture()
		vAxLabel.Image = imaging.Rotate90(img)
		vAxLabel.Resize(fyne.NewSize(vAxText.MinSize().Height, vAxText.MinSize().Width))
		vAxLabel.SetMinSize(fyne.NewSize(vAxText.MinSize().Height, vAxText.MinSize().Width))
		vAxisLabelWidth = vAxLabel.MinSize().Width
	}

	// determine the chart area
	hMin, hMax := hAxis.NRange()
	hOrigin := hAxis.NOrigin()
	vMin, vMax := vAxis.NRange()
	vOrigin := vAxis.NOrigin()
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

	r.chart.resize(area.maxPos.X-area.minPos.X, area.minPos.Y-area.maxPos.Y)

	// calculate conversion factors from ccordinates to positions
	area.hCoordToPos = (area.maxPos.X - area.minPos.X) / float32(hMax-hMin)
	area.vCoordToPos = (area.minPos.Y - area.maxPos.Y) / float32(vMax-vMin)

	// Place horizontal-Axis from hMin to hMax
	hLine, _, hArrowOne, hArrowTwo := hAxis.Arrow()
	if hAxText.Text != "" {
		hAxLabel.Move(fyne.NewPos(area.minPos.X+((area.maxPos.X-area.minPos.X)/2)-hAxText.MinSize().Width/2,
			size.Height-hAxText.MinSize().Height-r.margin))
	}
	hLine.Position1 = cartesianCoordinatesToPosition(hMin, vOrigin, area)
	hLine.Position2 = cartesianCoordinatesToPosition(hMax, vOrigin, area)
	hArrowOne.Position1 = fyne.NewPos(hLine.Position2.X-10, hLine.Position2.Y-5)
	hArrowOne.Position2 = hLine.Position2
	hArrowTwo.Position1 = fyne.NewPos(hLine.Position2.X-10, hLine.Position2.Y+5)
	hArrowTwo.Position2 = hLine.Position2

	// place horizontal ticks
	th := hAxis.Ticks()
	for i := range th {
		if th[i].Line != nil {
			th[i].Line.Position1 = cartesianCoordinatesToPosition(th[i].NLine, vOrigin, area)
			th[i].Line.Position2 = th[i].Line.Position1.AddXY(0, 5)
		}
		if th[i].SupLine != nil {
			th[i].SupLine.Position1 = cartesianCoordinatesToPosition(th[i].NLine, vMin, area)
			th[i].SupLine.Position2 = cartesianCoordinatesToPosition(th[i].NLine, vMax, area)
		}
		if th[i].Label != nil {
			th[i].Label.Move(cartesianCoordinatesToPosition(th[i].NLabel, vOrigin, area).AddXY(0, 5))
			th[i].Label.Alignment = fyne.TextAlignCenter
		}
	}

	// Place vertical axis from vMin to vMax
	vLine, _, vArrowOne, vArrowTwo := vAxis.Arrow()
	if vAxText.Text != "" {
		vAxLabel.Move(fyne.NewPos(r.margin, area.maxPos.Y+((area.minPos.Y-area.maxPos.Y)/2)-vAxText.MinSize().Width/2))
	}
	vLine.Position1 = cartesianCoordinatesToPosition(hOrigin, vMin, area)
	vLine.Position2 = cartesianCoordinatesToPosition(hOrigin, vMax, area)
	vArrowOne.Position1 = fyne.NewPos(vLine.Position2.X-5, vLine.Position2.Y+10)
	vArrowOne.Position2 = vLine.Position2
	vArrowTwo.Position1 = fyne.NewPos(vLine.Position2.X+5, vLine.Position2.Y+10)
	vArrowTwo.Position2 = vLine.Position2

	// place vertical ticks
	vh := vAxis.Ticks()
	for i := range vh {
		if vh[i].Line != nil {
			vh[i].Line.Position1 = cartesianCoordinatesToPosition(hOrigin,
				vh[i].NLine, area)
			vh[i].Line.Position2 = vh[i].Line.Position1.SubtractXY(5, 0)
		}
		if vh[i].SupLine != nil {
			vh[i].SupLine.Position1 = cartesianCoordinatesToPosition(hMin, vh[i].NLine, area)
			vh[i].SupLine.Position2 = cartesianCoordinatesToPosition(hMax, vh[i].NLine, area)
		}
		if vh[i].Label != nil {
			vh[i].Label.Move(cartesianCoordinatesToPosition(hOrigin,
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
}

// MinSize calculates the minimum space required to display the chart
func (r *cartesianRenderer) MinSize() fyne.Size {
	titleWidth := float32(0.0)
	titleHeight := float32(0.0)
	legendWidth := float32(0.0)
	legendHeight := float32(0.0)
	vAxisLabelWidth := float32(0.0)
	vAxisLabelHeight := float32(0)
	hAxisLabelWidth := float32(0)
	hAxisLabelHeight := float32(0.0)

	ct := r.chart.title()
	if ct.Name != "" {
		titleWidth = ct.Label.MinSize().Width
		titleHeight = ct.Label.MinSize().Height
	}

	if r.chart.legendVisibility() {
		les := r.chart.legendEntries()
		legendWidth, legendHeight = series.LegendSize(les)
	}

	var vAxis, hAxis *axis.Axis
	if r.transposed {
		vAxis = r.chart.fromAxis()
		hAxis = r.chart.toAxis()
	} else {
		vAxis = r.chart.toAxis()
		hAxis = r.chart.fromAxis()
	}
	hAxLabel, hAxText := hAxis.Label()
	if hAxText.Text != "" {
		hAxisLabelWidth = hAxLabel.MinSize().Width
		hAxisLabelHeight = hAxLabel.MinSize().Height
	}
	vAxLabel, vAxText := vAxis.Label()
	if vAxText.Text != "" {
		vAxisLabelWidth = vAxLabel.MinSize().Width
		vAxisLabelHeight = vAxLabel.MinSize().Height
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
func (r *cartesianRenderer) Objects() []fyne.CanvasObject {
	return r.chart.cartesianObjects()
}

// Refresh calls Layout if data of the chart has changes
func (r *cartesianRenderer) Refresh() {
	// if r.chart.hasChanged() {
	r.chart.refreshThemeColor()

	obj := r.Objects()
	for i := range obj {
		obj[i].Refresh()
	}
	// 	r.chart.resetHasChanged()
	// }
}

// cartesianCoordinatesToPosition converts a (h,v) coordinate to a fyne position
func cartesianCoordinatesToPosition(h float64, v float64, area cartDrawingArea) (pos fyne.Position) {
	pos.X = area.minPos.X + float32(h-area.hmin)*area.hCoordToPos
	pos.Y = area.minPos.Y - float32(v-area.vmin)*area.vCoordToPos
	return
}
