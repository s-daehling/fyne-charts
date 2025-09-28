package chart

import (
	"math"

	"github.com/s-daehling/fyne-charts/internal/series"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/software"
	"github.com/disintegration/imaging"
)

// polDrawingArea represents the area of the widget that can be used for the chart
type polDrawingArea struct {
	zeroPos    fyne.Position
	radius     float32
	coordToPos float32
	rot        float64
	mathPos    bool
}

// polarRenderer is the renderer for all cartesian plane widgets
type polarRenderer struct {
	baseRenderer
	rot     float64
	mathPos bool
}

func EmptyPolarRenderer(chart *BaseChart) (r *polarRenderer) {
	r = &polarRenderer{
		baseRenderer: emptyBaseRenderer(chart),
		rot:          0.0,
		mathPos:      true,
	}
	return
}

// Layout is responsible for redrawing the chart widget
func (r *polarRenderer) Layout(size fyne.Size) {
	_, titleHeight, legendWidth, _ := r.placeTitleAndLegend(size)
	rAxisLabelHeight := float32(0.0)
	phiAxisLabelWidth := float32(0.0)
	phiAxisTickLabelWidth := float32(0.0)
	phiAxisTickLabelHeight := float32(0.0)

	phiAx := r.chart.fromAxis()
	phiOrigin := phiAx.NOrigin()
	phiOriginAbs := absAngle(phiOrigin, r.mathPos, r.rot)
	rAx := r.chart.toAxis()
	_, rMax := rAx.NRange()
	rOrigin := rAx.NOrigin()
	phiAxisTickLabelWidth = phiAx.MaxTickWidth()
	phiAxisTickLabelHeight = phiAx.MaxTickHeight()

	phiAxLabel, phiAxText := phiAx.Label()
	if phiAxText.Text != "" {
		c := software.NewTransparentCanvas()
		c.SetPadded(false)
		c.SetContent(phiAxText)
		img := c.Capture()
		phiAxLabel.Image = imaging.Rotate90(img)
		phiAxLabel.Resize(fyne.NewSize(phiAxText.MinSize().Height, phiAxText.MinSize().Width))
		phiAxLabel.SetMinSize(fyne.NewSize(phiAxText.MinSize().Height, phiAxText.MinSize().Width))
		phiAxisLabelWidth = phiAxLabel.MinSize().Width
	}

	rAxLabel, rAxText := rAx.Label()
	if rAxText.Text != "" {
		c := software.NewTransparentCanvas()
		c.SetPadded(false)
		c.SetContent(rAxText)
		rAxLabel.Image = c.Capture()
		rAxLabel.Resize(rAxText.MinSize())
		rAxLabel.SetMinSize(rAxText.MinSize())
		rAxisLabelHeight = rAxLabel.MinSize().Height
	}

	// determine the chart area
	area := polDrawingArea{
		rot:     r.rot,
		mathPos: r.mathPos,
	}
	availWidth := size.Width - (2 * r.margin) - legendWidth - (2 * phiAxisTickLabelWidth) - phiAxisLabelWidth
	area.zeroPos.X = r.margin + phiAxisLabelWidth + phiAxisTickLabelWidth + (availWidth / 2)

	availHeight := size.Height - (2 * r.margin) - titleHeight - (2 * phiAxisTickLabelHeight) - rAxisLabelHeight
	area.zeroPos.Y = r.margin + titleHeight + phiAxisTickLabelHeight + (availHeight / 2)

	area.radius = availHeight / 2
	if availWidth < availHeight {
		area.radius = availWidth / 2
	}
	if area.radius < 0 {
		area.radius = 0
	}
	area.coordToPos = area.radius / float32(rMax)

	r.chart.resize(2*area.radius*math.Pi, area.radius)

	// place phi axis
	if phiAxText.Text != "" {
		phiAxLabel.Move(fyne.NewPos(r.margin, area.zeroPos.Y-phiAxText.MinSize().Width/2))
	}
	_, circle, arrowOne, arrowTwo := phiAx.Arrow()
	phiRadius := area.coordToPos * float32(rOrigin)
	circle.Resize(fyne.NewSize(2*phiRadius, 2*phiRadius))
	circle.Move(area.zeroPos.SubtractXY(phiRadius, phiRadius))
	ri, ai := arrowCoordinates(rOrigin, -5, 10, area)
	arrowOne.Position1 = polarCoordinatesToPosition(0, rOrigin, area)
	arrowOne.Position2 = polarCoordinatesToPosition(-ai, ri, area)
	ra, aa := arrowCoordinates(rOrigin, 5, 10, area)
	arrowTwo.Position1 = polarCoordinatesToPosition(0, rOrigin, area)
	arrowTwo.Position2 = polarCoordinatesToPosition(-aa, ra, area)

	// place phi ticks
	tPhi := phiAx.Ticks()
	for i := range tPhi {
		if tPhi[i].Line != nil {
			tPhi[i].Line.Position1 = polarCoordinatesToPosition(tPhi[i].NLine, rOrigin, area)
			tPhi[i].Line.Position2 = polarCoordinatesToPosition(tPhi[i].NLine, rOrigin+5.0/float64(area.coordToPos), area)
		}
		if tPhi[i].SupLine != nil {
			tPhi[i].SupLine.Position1 = area.zeroPos
			tPhi[i].SupLine.Position2 = polarCoordinatesToPosition(tPhi[i].NLine, rOrigin, area)
		}
		if tPhi[i].Label != nil {
			lPos := polarCoordinatesToPosition(tPhi[i].NLabel, rOrigin+5.0/float64(area.coordToPos), area)
			aLabelAbs := absAngle(tPhi[i].NLabel, r.mathPos, r.rot)
			if aLabelAbs > math.Pi/8 && aLabelAbs < 7*math.Pi/8 {
				lPos = lPos.AddXY(0, -tPhi[i].Label.MinSize().Height)
			} else if aLabelAbs < math.Pi/8 || aLabelAbs > 15*math.Pi/8 || (aLabelAbs > 7*math.Pi/8 && aLabelAbs < 9*math.Pi/8) {
				lPos = lPos.AddXY(0, -tPhi[i].Label.MinSize().Height/2)
			}
			tPhi[i].Label.Move(lPos)
			if aLabelAbs < 3*math.Pi/8 || aLabelAbs > 13*math.Pi/8 {
				tPhi[i].Label.Alignment = fyne.TextAlignLeading
			} else if aLabelAbs > 5*math.Pi/8 && aLabelAbs < 11*math.Pi/8 {
				tPhi[i].Label.Alignment = fyne.TextAlignTrailing
			} else {
				tPhi[i].Label.Alignment = fyne.TextAlignCenter
			}
		}
	}

	// place r axis
	if rAxText.Text != "" {
		rAxLabel.Move(fyne.NewPos(area.zeroPos.X+(area.radius/2)-rAxText.MinSize().Width/2,
			size.Height-rAxText.MinSize().Height-r.margin))
	}
	line, _, arrowOne, arrowTwo := rAx.Arrow()
	line.Position1 = area.zeroPos
	line.Position2 = polarCoordinatesToPosition(phiOrigin, rMax, area)
	ri, ai = arrowCoordinates(rMax, -10, 5, area)
	arrowOne.Position1 = polarCoordinatesToPosition(phiOrigin, rMax, area)
	arrowOne.Position2 = polarCoordinatesToPosition(phiOrigin+ai, ri, area)
	arrowTwo.Position1 = polarCoordinatesToPosition(phiOrigin, rMax, area)
	arrowTwo.Position2 = polarCoordinatesToPosition(phiOrigin-ai, ri, area)

	// place r ticks
	tR := rAx.Ticks()
	for i := range tR {
		if tR[i].Line != nil {
			tR[i].Line.Position1 = polarCoordinatesToPosition(phiOrigin, tR[i].NLine, area)
			rt, at := arrowCoordinates(tR[i].NLine, 0, 5, area)
			if !r.mathPos {
				at = -at
			}
			if phiOriginAbs < math.Pi/2 {
				tR[i].Line.Position2 = polarCoordinatesToPosition(phiOrigin-at, rt, area)
			} else if phiOriginAbs < 3*math.Pi/2 {
				tR[i].Line.Position2 = polarCoordinatesToPosition(phiOrigin+at, rt, area)
			} else {
				tR[i].Line.Position2 = polarCoordinatesToPosition(phiOrigin-at, rt, area)
			}
		}
		if tR[i].SupCircle != nil {
			supRadius := area.coordToPos * float32(tR[i].NLine)
			tR[i].SupCircle.Resize(fyne.NewSize(2*supRadius, 2*supRadius))
			tR[i].SupCircle.Move(area.zeroPos.SubtractXY(supRadius, supRadius))
		}
		if tR[i].Label != nil {
			rl, al := arrowCoordinates(tR[i].NLabel, math.Sin(phiOriginAbs)*float64(tR[i].Label.MinSize().Height/2), 5, area)
			if !r.mathPos {
				al = -al
			}

			if phiOriginAbs < math.Pi/8 {
				tR[i].Label.Move(polarCoordinatesToPosition(phiOrigin-al, rl, area))
				tR[i].Label.Alignment = fyne.TextAlignCenter
			} else if phiOriginAbs < math.Pi/2 {
				tR[i].Label.Move(polarCoordinatesToPosition(phiOrigin-al, rl, area))
				tR[i].Label.Alignment = fyne.TextAlignLeading
			} else if phiOriginAbs < 7*math.Pi/8 {
				tR[i].Label.Move(polarCoordinatesToPosition(phiOrigin+al, rl, area))
				tR[i].Label.Alignment = fyne.TextAlignTrailing
			} else if phiOriginAbs < 9*math.Pi/8 {
				tR[i].Label.Move(polarCoordinatesToPosition(phiOrigin+al, rl, area))
				tR[i].Label.Alignment = fyne.TextAlignCenter
			} else if phiOriginAbs < 3*math.Pi/2 {
				tR[i].Label.Move(polarCoordinatesToPosition(phiOrigin+al, rl, area))
				tR[i].Label.Alignment = fyne.TextAlignLeading
			} else if phiOriginAbs < 15*math.Pi/8 {
				tR[i].Label.Move(polarCoordinatesToPosition(phiOrigin-al, rl, area))
				tR[i].Label.Alignment = fyne.TextAlignTrailing
			} else {
				tR[i].Label.Move(polarCoordinatesToPosition(phiOrigin-al, rl, area))
				tR[i].Label.Alignment = fyne.TextAlignCenter
			}
		}
	}

	// place nodes
	ns := r.chart.polarNodes()
	for i := range ns {
		var dotPos fyne.Position
		dotPos = polarCoordinatesToPosition(ns[i].Phi, ns[i].R, area)
		dotSize := ns[i].Dot.Size().Width
		dotPos = dotPos.SubtractXY(dotSize/2, dotSize/2)
		ns[i].Dot.Move(dotPos)
	}

	// place edges
	es := r.chart.polarEdges()
	for i := range es {
		es[i].Line.Position1 = polarCoordinatesToPosition(es[i].Phi1, es[i].R1, area)
		es[i].Line.Position2 = polarCoordinatesToPosition(es[i].Phi2, es[i].R2, area)
	}

	// place texts
	ts := r.chart.polarTexts()
	for i := range ts {
		tPos := polarCoordinatesToPosition(ts[i].Phi, ts[i].R, area)
		tPos = tPos.SubtractXY(0, ts[i].Text.MinSize().Height/2)
		ts[i].Text.Move(tPos)
		ts[i].Text.Alignment = fyne.TextAlignCenter
	}

	// place raster
	rs := r.chart.chartRaster()
	rs.Move(fyne.NewPos(area.zeroPos.X-area.radius, area.zeroPos.Y-area.radius))
	rs.Resize(fyne.NewSize(2*area.radius, 2*area.radius))
}

// MinSize calculates the minimum space required to display the chart
func (r *polarRenderer) MinSize() fyne.Size {
	titleWidth := float32(0.0)
	titleHeight := float32(0.0)
	legendWidth := float32(0.0)
	legendHeight := float32(0.0)
	rAxisLabelWidth := float32(0.0)
	rAxisLabelHeight := float32(0)
	phiAxisLabelWidth := float32(0)
	phiAxisLabelHeight := float32(0.0)

	ct := r.chart.title()
	if ct.Name != "" {
		titleWidth = ct.Label.MinSize().Width
		titleHeight = ct.Label.MinSize().Height
	}

	if r.chart.legendVisibility() {
		les := r.chart.legendEntries()
		legendWidth, legendHeight = series.LegendSize(les)
	}

	phiAxis := r.chart.fromAxis()
	rAxis := r.chart.toAxis()
	phiAxLabel, phiAxText := phiAxis.Label()
	if phiAxText.Text != "" {
		phiAxisLabelWidth = phiAxLabel.MinSize().Width
		phiAxisLabelHeight = phiAxLabel.MinSize().Height
	}
	rAxLabel, rAxText := rAxis.Label()
	if rAxText.Text != "" {
		rAxisLabelWidth = rAxLabel.MinSize().Width
		rAxisLabelHeight = rAxLabel.MinSize().Height
	}

	minHeight := 2*r.margin + titleHeight
	if legendHeight > 20+rAxisLabelHeight+phiAxisLabelHeight {
		minHeight += legendHeight
	} else {
		minHeight += 20 + rAxisLabelHeight + phiAxisLabelHeight
	}

	minWidth := 2 * r.margin
	if titleWidth > 20+rAxisLabelWidth+phiAxisLabelWidth+legendWidth {
		minWidth += titleWidth
	} else {
		minWidth += 20 + rAxisLabelWidth + phiAxisLabelWidth + legendWidth
	}
	return fyne.NewSize(minWidth, minHeight)
}

// Objects returns a list of all objects to be drawn
func (r *polarRenderer) Objects() []fyne.CanvasObject {
	return r.chart.polarObjects()
}

// Refresh calls Layout if data of the chart has changes
func (r *polarRenderer) Refresh() {
	// if r.chart.hasChanged() {
	r.chart.refreshThemeColor()

	obj := r.Objects()
	for i := range obj {
		obj[i].Refresh()
	}
	// 	r.chart.resetHasChanged()
	// }
}

func absAngle(a float64, mathPos bool, rot float64) (aCor float64) {
	dir := 1.0
	if !mathPos {
		dir = -1.0
	}
	aCor = (a * dir) + rot
	if aCor < 0 {
		for {
			aCor += 2 * math.Pi
			if aCor > -0.000001 {
				break
			}
		}
	} else if aCor > 2*math.Pi {
		for {
			aCor -= 2 * math.Pi
			if aCor < 2*math.Pi+0.000001 {
				break
			}
		}
	}
	return
}

func arrowCoordinates(rTip float64, radialInPos float64, tangentialInPos float64, area polDrawingArea) (rBack float64, aArrow float64) {
	rTipInPos := rTip * float64(area.coordToPos)
	rBackInPos := math.Sqrt(math.Pow(rTipInPos+radialInPos, 2) + math.Pow(tangentialInPos, 2))
	aArrow = math.Asin(tangentialInPos / rBackInPos)
	rBack = (rBackInPos / float64(area.coordToPos)) //* ((rTipInPos + radialInPos) / math.Abs(rTipInPos+radialInPos))
	return
}

// polarCoordinatesToPosition converts a (h,v) coordinate to a fyne position
func polarCoordinatesToPosition(phi float64, r float64, area polDrawingArea) (pos fyne.Position) {
	if r < 0 {
		phi -= math.Pi
		r = -r
	}
	phi = absAngle(phi, area.mathPos, area.rot)
	pos.X = area.zeroPos.X + (float32(math.Cos(phi)*r) * area.coordToPos)
	pos.Y = area.zeroPos.Y - (float32(math.Sin(phi)*r) * area.coordToPos)
	return
}
