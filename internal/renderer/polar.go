package renderer

import (
	"math"

	"fyne.io/fyne/v2"
	"github.com/s-daehling/fyne-charts/internal/elements"
)

type PolarChart interface {
	baseChart
	PolarDots() (ns []*elements.Dot)
	PolarEdges() (es []elements.Edge)
	PolarTexts() (ts []elements.Label)
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
	phiAxisTickLabelWidth := float32(0.0)
	phiAxisTickLabelHeight := float32(0.0)

	var phiOrigin, rMax, rOrigin float64
	var phiTicks, rTicks []elements.Tick
	var phiArrow, rArrow elements.Arrow
	var phiShow, rShow bool
	_, _, phiOrigin, phiTicks, phiArrow, phiShow = r.chart.FromAxisElements()
	_, rMax, rOrigin, _, rArrow, rShow = r.chart.ToAxisElements()

	phiOriginAbs := absAngle(phiOrigin, r.mathPos, r.rot)
	phiAxisTickLabelWidth, phiAxisTickLabelHeight = elements.MaxTickSize(phiTicks)
	if phiAxisTickLabelWidth < 0.001 {
		phiAxisTickLabelWidth = r.prevPhiAxisTickLabelWidth
	} else {
		r.prevPhiAxisTickLabelWidth = phiAxisTickLabelWidth
	}

	// determine the chart area
	area := polDrawingArea{
		rot:     r.rot,
		mathPos: r.mathPos,
	}
	availWidth := size.Width - (2 * r.margin) - (2 * phiAxisTickLabelWidth)
	area.zeroPos.X = r.margin + phiAxisTickLabelWidth + (availWidth / 2)

	availHeight := size.Height - (2 * r.margin) - (2 * phiAxisTickLabelHeight)
	area.zeroPos.Y = r.margin + phiAxisTickLabelHeight + (availHeight / 2)

	area.radius = availHeight / 2
	if availWidth < availHeight {
		area.radius = availWidth / 2
	}
	if area.radius < 0 {
		area.radius = 0
	}
	area.coordToPos = area.radius / float32(rMax)

	r.chart.ChartSizeChange(2*area.radius*math.Pi, area.radius)

	_, _, _, phiTicks, _, _ = r.chart.FromAxisElements()
	_, _, _, rTicks, _, _ = r.chart.ToAxisElements()

	// place phi axis
	if phiShow {
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
					lPos = lPos.AddXY(0, -phiTicks[i].Label.Size().Height)
				} else if aLabelAbs < math.Pi/8 || aLabelAbs > 15*math.Pi/8 || (aLabelAbs > 7*math.Pi/8 && aLabelAbs < 9*math.Pi/8) {
					lPos = lPos.AddXY(0, -phiTicks[i].Label.Size().Height/2)
				}
				if aLabelAbs < 3*math.Pi/8 || aLabelAbs > 13*math.Pi/8 {
				} else if aLabelAbs > 5*math.Pi/8 && aLabelAbs < 11*math.Pi/8 {
					lPos = lPos.AddXY(-phiTicks[i].Label.Size().Width, 0)
				} else {
					lPos = lPos.AddXY(-phiTicks[i].Label.Size().Width/2, 0)
				}
				phiTicks[i].Label.Move(lPos)
			}
		}
	}

	// place r axis
	if rShow {
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
				var lPos fyne.Position
				rl, al := arrowCoordinates(rTicks[i].NLabel, math.Sin(phiOriginAbs)*float64(rTicks[i].Label.Size().Height/2), 5, area)
				if !r.mathPos {
					al = -al
				}

				if phiOriginAbs < math.Pi/8 {
					lPos = polarCoordinatesToPosition(phiOrigin-al, rl, area)
					lPos = lPos.AddXY(-rTicks[i].Label.Size().Width/2, 0)
				} else if phiOriginAbs < math.Pi/2 {
					lPos = polarCoordinatesToPosition(phiOrigin-al, rl, area)
				} else if phiOriginAbs < 7*math.Pi/8 {
					lPos = polarCoordinatesToPosition(phiOrigin+al, rl, area)
					lPos = lPos.AddXY(-rTicks[i].Label.Size().Width, 0)
				} else if phiOriginAbs < 9*math.Pi/8 {
					lPos = polarCoordinatesToPosition(phiOrigin+al, rl, area)
					lPos = lPos.AddXY(-rTicks[i].Label.Size().Width/2, 0)
				} else if phiOriginAbs < 3*math.Pi/2 {
					lPos = polarCoordinatesToPosition(phiOrigin+al, rl, area)
				} else if phiOriginAbs < 15*math.Pi/8 {
					lPos = polarCoordinatesToPosition(phiOrigin-al, rl, area)
					lPos = lPos.AddXY(-rTicks[i].Label.Size().Width, 0)
				} else {
					lPos = polarCoordinatesToPosition(phiOrigin-al, rl, area)
					lPos = lPos.AddXY(-rTicks[i].Label.Size().Width/2, 0)
				}
				rTicks[i].Label.Move(lPos)
			}
		}
	}

	// place nodes
	ns := r.chart.PolarDots()
	for i := range ns {
		var dotPos fyne.Position
		dotPos = polarCoordinatesToPosition(ns[i].N, ns[i].Val, area)
		dotSize := ns[i].Size().Width
		dotPos = dotPos.SubtractXY(dotSize/2, dotSize/2)
		ns[i].Move(dotPos)
	}

	// place edges
	es := r.chart.PolarEdges()
	for i := range es {
		es[i].Line.Position1 = polarCoordinatesToPosition(es[i].N1, es[i].Val1, area)
		es[i].Line.Position2 = polarCoordinatesToPosition(es[i].N2, es[i].Val2, area)
	}

	// place texts
	ts := r.chart.PolarTexts()
	for i := range ts {
		tPos := polarCoordinatesToPosition(ts[i].N, ts[i].Val, area)
		tPos = tPos.SubtractXY(0, ts[i].Text.MinSize().Height/2)
		ts[i].Text.Move(tPos)
		ts[i].Text.Alignment = fyne.TextAlignCenter
	}

	// place area
	rs := r.chart.Area()
	if rs != nil {
		rs.Move(fyne.NewPos(area.zeroPos.X-area.radius, area.zeroPos.Y-area.radius))
		rs.Resize(fyne.NewSize(2*area.radius, 2*area.radius))
	}

	// place tooltip
	tt := r.chart.Tooltip()
	ttWidth, ttHeigth := tooltipSize(tt.Entries)
	ttPos := fyne.NewPos(area.zeroPos.X-area.radius, area.zeroPos.Y-area.radius).AddXY(tt.X, tt.Y).SubtractXY(ttWidth+5, ttHeigth)
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
		ov.Move(fyne.NewPos(area.zeroPos.X-area.radius, area.zeroPos.Y-area.radius))
		ov.Resize(fyne.NewSize(2*area.radius, 2*area.radius))
	}
}

// MinSize calculates the minimum space required to display the chart
func (r *Polar) MinSize() fyne.Size {
	minHeight := 2*r.margin + 20
	minWidth := 2*r.margin + 20
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

	r.Layout(r.chart.Size())
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
