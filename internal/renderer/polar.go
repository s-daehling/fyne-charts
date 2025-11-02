package renderer

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/software"
	"github.com/disintegration/imaging"
)

type PolarChart interface {
	baseChart
	PolarNodes() (ns []PolarNode)
	PolarEdges() (es []PolarEdge)
	PolarTexts() (ts []PolarText)
	PolarObjects() (obj []fyne.CanvasObject)
}

// polDrawingArea represents the area of the widget that can be used for the chart
type polDrawingArea struct {
	zeroPos    fyne.Position
	radius     float32
	coordToPos float32
	rot        float64
	mathPos    bool
}

// Polar is the renderer for all cartesian plane widgets
type Polar struct {
	baseRenderer
	chart                     PolarChart
	rot                       float64
	mathPos                   bool
	prevPhiAxisTickLabelWidth float32
}

func EmptyPolarRenderer(chart PolarChart) (r *Polar) {
	r = &Polar{
		baseRenderer: emptyBaseRenderer(),
		rot:          0.0,
		mathPos:      true,
		chart:        chart,
	}
	return
}

// Layout is responsible for redrawing the chart widget
func (r *Polar) Layout(size fyne.Size) {
	_, titleHeight, legendWidth, _ := r.placeTitleAndLegend(size, r.chart.Title(), r.chart.LegendEntries())
	rAxisLabelHeight := float32(0.0)
	phiAxisLabelWidth := float32(0.0)
	phiAxisTickLabelWidth := float32(0.0)
	phiAxisTickLabelHeight := float32(0.0)

	var phiOrigin, rMax, rOrigin float64
	var phiLabel, rLabel Label
	var phiTicks, rTicks []Tick
	var phiArrow, rArrow Arrow
	var phiShow, rShow bool
	_, _, phiOrigin, phiLabel, phiTicks, phiArrow, phiShow = r.chart.FromAxisElements()
	_, rMax, rOrigin, rLabel, _, rArrow, rShow = r.chart.ToAxisElements()

	phiOriginAbs := absAngle(phiOrigin, r.mathPos, r.rot)
	phiAxisTickLabelWidth, phiAxisTickLabelHeight = maxTickSize(phiTicks)
	if phiAxisTickLabelWidth < 0.001 {
		phiAxisTickLabelWidth = r.prevPhiAxisTickLabelWidth
	} else {
		r.prevPhiAxisTickLabelWidth = phiAxisTickLabelWidth
	}

	if phiShow && phiLabel.Text.Text != "" {
		c := software.NewTransparentCanvas()
		c.SetPadded(false)
		c.SetContent(phiLabel.Text)
		img := c.Capture()
		phiLabel.Image.Image = imaging.Rotate90(img)
		phiLabel.Image.Resize(fyne.NewSize(phiLabel.Text.MinSize().Height, phiLabel.Text.MinSize().Width))
		phiLabel.Image.SetMinSize(fyne.NewSize(phiLabel.Text.MinSize().Height, phiLabel.Text.MinSize().Width))
		phiAxisLabelWidth = phiLabel.Image.MinSize().Width
	}

	if rShow && rLabel.Text.Text != "" {
		c := software.NewTransparentCanvas()
		c.SetPadded(false)
		c.SetContent(rLabel.Text)
		rLabel.Image.Image = c.Capture()
		rLabel.Image.Resize(rLabel.Text.MinSize())
		rLabel.Image.SetMinSize(rLabel.Text.MinSize())
		rAxisLabelHeight = rLabel.Image.MinSize().Height
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

	r.chart.Resize(2*area.radius*math.Pi, area.radius)

	_, _, _, _, phiTicks, _, _ = r.chart.FromAxisElements()
	_, _, _, _, rTicks, _, _ = r.chart.ToAxisElements()

	// place phi axis
	if phiShow {
		if phiLabel.Text.Text != "" {
			phiLabel.Image.Move(fyne.NewPos(r.margin, area.zeroPos.Y-phiLabel.Text.MinSize().Width/2))
		}
		phiRadius := area.coordToPos * float32(rOrigin)
		phiArrow.Circle.Resize(fyne.NewSize(2*phiRadius, 2*phiRadius))
		phiArrow.Circle.Move(area.zeroPos.SubtractXY(phiRadius, phiRadius))
		ri, ai := arrowCoordinates(rOrigin, -5, 10, area)
		phiArrow.HeadOne.Position1 = polarCoordinatesToPosition(0, rOrigin, area)
		phiArrow.HeadOne.Position2 = polarCoordinatesToPosition(-ai, ri, area)
		ra, aa := arrowCoordinates(rOrigin, 5, 10, area)
		phiArrow.HeadTwo.Position1 = polarCoordinatesToPosition(0, rOrigin, area)
		phiArrow.HeadTwo.Position2 = polarCoordinatesToPosition(-aa, ra, area)

		// place phi ticks
		for i := range phiTicks {
			if phiTicks[i].Line != nil {
				phiTicks[i].Line.Position1 = polarCoordinatesToPosition(phiTicks[i].NLine, rOrigin, area)
				phiTicks[i].Line.Position2 = polarCoordinatesToPosition(phiTicks[i].NLine, rOrigin+5.0/float64(area.coordToPos), area)
			}
			if phiTicks[i].SupLine != nil {
				phiTicks[i].SupLine.Position1 = area.zeroPos
				phiTicks[i].SupLine.Position2 = polarCoordinatesToPosition(phiTicks[i].NLine, rOrigin, area)
			}
			if phiTicks[i].Label != nil {
				lPos := polarCoordinatesToPosition(phiTicks[i].NLabel, rOrigin+5.0/float64(area.coordToPos), area)
				aLabelAbs := absAngle(phiTicks[i].NLabel, r.mathPos, r.rot)
				if aLabelAbs > math.Pi/8 && aLabelAbs < 7*math.Pi/8 {
					lPos = lPos.AddXY(0, -phiTicks[i].Label.MinSize().Height)
				} else if aLabelAbs < math.Pi/8 || aLabelAbs > 15*math.Pi/8 || (aLabelAbs > 7*math.Pi/8 && aLabelAbs < 9*math.Pi/8) {
					lPos = lPos.AddXY(0, -phiTicks[i].Label.MinSize().Height/2)
				}
				phiTicks[i].Label.Move(lPos)
				if aLabelAbs < 3*math.Pi/8 || aLabelAbs > 13*math.Pi/8 {
					phiTicks[i].Label.Alignment = fyne.TextAlignLeading
				} else if aLabelAbs > 5*math.Pi/8 && aLabelAbs < 11*math.Pi/8 {
					phiTicks[i].Label.Alignment = fyne.TextAlignTrailing
				} else {
					phiTicks[i].Label.Alignment = fyne.TextAlignCenter
				}
			}
		}
	}

	// place r axis
	if rShow {
		if rLabel.Text.Text != "" {
			rLabel.Image.Move(fyne.NewPos(area.zeroPos.X+(area.radius/2)-rLabel.Text.MinSize().Width/2,
				size.Height-rLabel.Text.MinSize().Height-r.margin))
		}
		rArrow.Line.Position1 = area.zeroPos
		rArrow.Line.Position2 = polarCoordinatesToPosition(phiOrigin, rMax, area)
		ri, ai := arrowCoordinates(rMax, -10, 5, area)
		rArrow.HeadOne.Position1 = polarCoordinatesToPosition(phiOrigin, rMax, area)
		rArrow.HeadOne.Position2 = polarCoordinatesToPosition(phiOrigin+ai, ri, area)
		rArrow.HeadTwo.Position1 = polarCoordinatesToPosition(phiOrigin, rMax, area)
		rArrow.HeadTwo.Position2 = polarCoordinatesToPosition(phiOrigin-ai, ri, area)

		// place r ticks
		for i := range rTicks {
			if rTicks[i].Line != nil {
				rTicks[i].Line.Position1 = polarCoordinatesToPosition(phiOrigin, rTicks[i].NLine, area)
				rt, at := arrowCoordinates(rTicks[i].NLine, 0, 5, area)
				if !r.mathPos {
					at = -at
				}
				if phiOriginAbs < math.Pi/2 {
					rTicks[i].Line.Position2 = polarCoordinatesToPosition(phiOrigin-at, rt, area)
				} else if phiOriginAbs < 3*math.Pi/2 {
					rTicks[i].Line.Position2 = polarCoordinatesToPosition(phiOrigin+at, rt, area)
				} else {
					rTicks[i].Line.Position2 = polarCoordinatesToPosition(phiOrigin-at, rt, area)
				}
			}
			if rTicks[i].SupCircle != nil {
				supRadius := area.coordToPos * float32(rTicks[i].NLine)
				rTicks[i].SupCircle.Resize(fyne.NewSize(2*supRadius, 2*supRadius))
				rTicks[i].SupCircle.Move(area.zeroPos.SubtractXY(supRadius, supRadius))
			}
			if rTicks[i].Label != nil {
				rl, al := arrowCoordinates(rTicks[i].NLabel, math.Sin(phiOriginAbs)*float64(rTicks[i].Label.MinSize().Height/2), 5, area)
				if !r.mathPos {
					al = -al
				}

				if phiOriginAbs < math.Pi/8 {
					rTicks[i].Label.Move(polarCoordinatesToPosition(phiOrigin-al, rl, area))
					rTicks[i].Label.Alignment = fyne.TextAlignCenter
				} else if phiOriginAbs < math.Pi/2 {
					rTicks[i].Label.Move(polarCoordinatesToPosition(phiOrigin-al, rl, area))
					rTicks[i].Label.Alignment = fyne.TextAlignLeading
				} else if phiOriginAbs < 7*math.Pi/8 {
					rTicks[i].Label.Move(polarCoordinatesToPosition(phiOrigin+al, rl, area))
					rTicks[i].Label.Alignment = fyne.TextAlignTrailing
				} else if phiOriginAbs < 9*math.Pi/8 {
					rTicks[i].Label.Move(polarCoordinatesToPosition(phiOrigin+al, rl, area))
					rTicks[i].Label.Alignment = fyne.TextAlignCenter
				} else if phiOriginAbs < 3*math.Pi/2 {
					rTicks[i].Label.Move(polarCoordinatesToPosition(phiOrigin+al, rl, area))
					rTicks[i].Label.Alignment = fyne.TextAlignLeading
				} else if phiOriginAbs < 15*math.Pi/8 {
					rTicks[i].Label.Move(polarCoordinatesToPosition(phiOrigin-al, rl, area))
					rTicks[i].Label.Alignment = fyne.TextAlignTrailing
				} else {
					rTicks[i].Label.Move(polarCoordinatesToPosition(phiOrigin-al, rl, area))
					rTicks[i].Label.Alignment = fyne.TextAlignCenter
				}
			}
		}
	}

	// place nodes
	ns := r.chart.PolarNodes()
	for i := range ns {
		var dotPos fyne.Position
		dotPos = polarCoordinatesToPosition(ns[i].Phi, ns[i].R, area)
		dotSize := ns[i].Dot.Size().Width
		dotPos = dotPos.SubtractXY(dotSize/2, dotSize/2)
		ns[i].Dot.Move(dotPos)
	}

	// place edges
	es := r.chart.PolarEdges()
	for i := range es {
		es[i].Line.Position1 = polarCoordinatesToPosition(es[i].Phi1, es[i].R1, area)
		es[i].Line.Position2 = polarCoordinatesToPosition(es[i].Phi2, es[i].R2, area)
	}

	// place texts
	ts := r.chart.PolarTexts()
	for i := range ts {
		tPos := polarCoordinatesToPosition(ts[i].Phi, ts[i].R, area)
		tPos = tPos.SubtractXY(0, ts[i].Text.MinSize().Height/2)
		ts[i].Text.Move(tPos)
		ts[i].Text.Alignment = fyne.TextAlignCenter
	}

	// place raster
	rs := r.chart.Raster()
	if rs != nil {
		rs.Move(fyne.NewPos(area.zeroPos.X-area.radius, area.zeroPos.Y-area.radius))
		rs.Resize(fyne.NewSize(2*area.radius, 2*area.radius))
	}
}

// MinSize calculates the minimum space required to display the chart
func (r *Polar) MinSize() fyne.Size {
	titleWidth := float32(0.0)
	titleHeight := float32(0.0)
	legendWidth := float32(0.0)
	legendHeight := float32(0.0)
	rAxisLabelWidth := float32(0.0)
	rAxisLabelHeight := float32(0)
	phiAxisLabelWidth := float32(0)
	phiAxisLabelHeight := float32(0.0)

	ct := r.chart.Title()
	if ct != nil {
		if ct.Text != "" {
			titleWidth = ct.MinSize().Width
			titleHeight = ct.MinSize().Height
		}
	}

	les := r.chart.LegendEntries()
	legendWidth, legendHeight = legendSize(les)

	var phiLabel, rLabel Label
	var phiShow, rShow bool
	_, _, _, phiLabel, _, _, phiShow = r.chart.FromAxisElements()
	_, _, _, rLabel, _, _, rShow = r.chart.ToAxisElements()

	if phiShow && phiLabel.Text.Text != "" {
		phiAxisLabelWidth = phiLabel.Image.MinSize().Width
		phiAxisLabelHeight = phiLabel.Image.MinSize().Height
	}
	if rShow && rLabel.Text.Text != "" {
		rAxisLabelWidth = rLabel.Image.MinSize().Width
		rAxisLabelHeight = rLabel.Image.MinSize().Height
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
func (r *Polar) Objects() []fyne.CanvasObject {
	return r.chart.PolarObjects()
}

// Refresh calls Layout if data of the chart has changes
func (r *Polar) Refresh() {
	// if r.chart.hasChanged() {
	r.chart.RefreshTheme()

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
