package axis

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/software"
	"fyne.io/fyne/v2/theme"
)

type AxisType string

const (
	CartesianAxis AxisType = "Cartesian"
	PolarPhiAxis  AxisType = "PolarPhi"
	PolarRAxis    AxisType = "PolarR"
)

type Tick struct {
	NLabel    float64
	Label     *canvas.Text
	NLine     float64
	Line      *canvas.Line
	SupLine   *canvas.Line
	SupCircle *canvas.Circle
}

type axisTick struct {
	c              string
	t              time.Time
	n              float64
	nLabel         float64
	nLine          float64
	label          *canvas.Text // the text label
	line           *canvas.Line // the tick line
	hasSupportLine bool         // if true, a orthogonal support line is drawn at the coordLine coordinate, ranging from min to max value of the opposite axis
	supportLine    *canvas.Line // the support line
	supportCircle  *canvas.Circle
}

type Axis struct {
	typ             AxisType
	visible         bool
	ticks           []axisTick
	autoTicks       bool
	autoSupportLine bool
	tOrigin         time.Time
	nOrigin         float64
	cs              []string
	tMin            time.Time
	tMax            time.Time
	nMin            float64
	nMax            float64
	line            *canvas.Line // the line representing the axis
	circle          *canvas.Circle
	arrowOne        *canvas.Line  // first part of the arrow at the end of the axis line
	arrowTwo        *canvas.Line  // second part of the arrow at the end of the axis line
	name            string        // name/title of the axis
	label           *canvas.Image // name/title of the axis; rotated if the axis is vertical
	labelText       *canvas.Text
	space           float32
	col             color.Color
	supCol          color.Color
}

func EmptyAxis(name string, typ AxisType) (ax *Axis) {
	col := theme.Color(theme.ColorNameForeground)
	ax = &Axis{
		typ:             typ,
		visible:         true,
		ticks:           []axisTick{},
		autoTicks:       true,
		autoSupportLine: true,
		nOrigin:         0.0,
		nMin:            0.0,
		nMax:            100.0,
		line:            canvas.NewLine(col),
		circle:          canvas.NewCircle(color.RGBA{0x00, 0x00, 0x00, 0x00}),
		arrowOne:        canvas.NewLine(col),
		arrowTwo:        canvas.NewLine(col),
		name:            name,
		label:           canvas.NewImageFromImage(software.NewTransparentCanvas().Capture()),
		labelText:       canvas.NewText(name, col),
		col:             col,
		supCol:          theme.Color(theme.ColorNameShadow),
	}
	if typ == PolarPhiAxis {
		ax.nMax = 2 * math.Pi
	}
	ax.circle.StrokeColor = col
	ax.circle.StrokeWidth = 1
	return
}

func (ax *Axis) Objects() (canObj []fyne.CanvasObject) {
	if ax.typ == PolarPhiAxis {
		canObj = append(canObj, ax.circle)
	} else {
		canObj = append(canObj, ax.line)
	}
	canObj = append(canObj, ax.arrowOne)
	canObj = append(canObj, ax.arrowTwo)

	ts := ax.Ticks()
	for i := range ts {
		if ts[i].Label != nil {
			canObj = append(canObj, ts[i].Label)
		}
		if ts[i].Line != nil {
			canObj = append(canObj, ts[i].Line)
		}
		if ts[i].SupLine != nil {
			canObj = append(canObj, ts[i].SupLine)
		}
		if ts[i].SupCircle != nil {
			canObj = append(canObj, ts[i].SupCircle)
		}
	}

	if ax.name != "" {
		canObj = append(canObj, ax.label)
	}
	return
}

func (ax *Axis) Arrow() (line *canvas.Line, circle *canvas.Circle, arrowOne *canvas.Line,
	arrowTwo *canvas.Line) {
	line = ax.line
	circle = ax.circle
	arrowOne = ax.arrowOne
	arrowTwo = ax.arrowTwo
	return
}

func (ax *Axis) Ticks() (ts []Tick) {
	for i := range ax.ticks {
		if ax.ticks[i].n < ax.nMin || ax.ticks[i].n > ax.nMax {
			continue
		}
		t := Tick{
			NLabel:  ax.ticks[i].nLabel,
			NLine:   ax.ticks[i].nLine,
			Label:   nil,
			Line:    nil,
			SupLine: nil,
		}
		if t.NLabel > ax.nMin || t.NLabel < ax.nMax {
			t.Label = ax.ticks[i].label
		}
		if t.NLine > ax.nMin || t.NLine < ax.nMax {
			t.Line = ax.ticks[i].line
			if ax.ticks[i].hasSupportLine {
				if ax.typ == CartesianAxis || ax.typ == PolarPhiAxis {
					t.SupLine = ax.ticks[i].supportLine
				} else {
					t.SupCircle = ax.ticks[i].supportCircle
				}
			}
		}
		if t.Label != nil || t.Line != nil {
			ts = append(ts, t)
		}
	}
	return
}

func (ax *Axis) MaxTickWidth() (maxWidth float32) {
	maxWidth = 0
	for i := range ax.ticks {
		if ax.ticks[i].label.MinSize().Width > maxWidth {
			maxWidth = ax.ticks[i].label.MinSize().Width
		}
	}
	return
}

func (ax *Axis) MaxTickHeight() (maxHeight float32) {
	maxHeight = 0
	for i := range ax.ticks {
		if ax.ticks[i].label.MinSize().Height > maxHeight {
			maxHeight = ax.ticks[i].label.MinSize().Height
		}
	}
	return
}

func (ax *Axis) SetSpace(space float32) {
	ax.space = space
}

func (ax *Axis) Hide() {
	ax.visible = false
	ax.arrowOne.Hide()
	ax.arrowTwo.Hide()
	ax.line.Hide()
	ax.label.Hide()
	ax.circle.Hide()
	for i := range ax.ticks {
		ax.ticks[i].label.Hide()
		ax.ticks[i].line.Hide()
		ax.ticks[i].supportCircle.Hide()
		ax.ticks[i].supportLine.Hide()
	}
}

func (ax *Axis) Show() {
	ax.visible = true
	ax.arrowOne.Show()
	ax.arrowTwo.Show()
	ax.line.Show()
	ax.label.Show()
	ax.circle.Show()
	for i := range ax.ticks {
		ax.ticks[i].label.Show()
		ax.ticks[i].line.Show()
		ax.ticks[i].supportCircle.Show()
		ax.ticks[i].supportLine.Show()
	}
}

func (ax *Axis) RefreshThemeColor() {
	ax.col = theme.Color(theme.ColorNameForeground)
	ax.supCol = theme.Color(theme.ColorNameShadow)
	ax.labelText.Color = ax.col
	ax.arrowOne.StrokeColor = ax.col
	ax.arrowTwo.StrokeColor = ax.col
	ax.line.StrokeColor = ax.col
	ax.circle.StrokeColor = ax.col
	for i := range ax.ticks {
		ax.ticks[i].label.Color = ax.col
		ax.ticks[i].line.StrokeColor = ax.col
		ax.ticks[i].supportCircle.StrokeColor = ax.supCol
		ax.ticks[i].supportLine.StrokeColor = ax.supCol
	}
}

func (ax *Axis) SetLabel(l string) {
	ax.name = l
	ax.labelText.Text = l
}

func (ax *Axis) Label() (label *canvas.Image, text *canvas.Text) {
	label = ax.label
	text = ax.labelText
	return
}

func (ax *Axis) SetAutoTicks(autoSupport bool) {
	ax.autoTicks = true
	ax.autoSupportLine = autoSupport
}

func (ax *Axis) SetManualTicks() {
	ax.autoTicks = false
	ax.autoSupportLine = false
}

func (ax *Axis) AutoTicks() (a bool) {
	a = ax.autoTicks
	return
}

func (ax *Axis) adjustNumberOfTicks(n int) {
	//adjust size of ticks
	if n < len(ax.ticks) {
		ax.ticks = ax.ticks[:n]
	} else {
		n = n - len(ax.ticks)
		for range n {
			tick := axisTick{
				label:          canvas.NewText("", ax.col),
				line:           canvas.NewLine(ax.col),
				hasSupportLine: false,
				supportLine:    canvas.NewLine(ax.supCol),
				supportCircle:  canvas.NewCircle(color.RGBA{0x00, 0x00, 0x00, 0x00}),
			}
			// tick.supportLine.StrokeWidth = 0.5
			tick.supportCircle.StrokeWidth = 1
			tick.supportCircle.StrokeColor = ax.supCol
			if !ax.visible {
				tick.label.Hide()
				tick.line.Hide()
				tick.supportCircle.Hide()
				tick.supportLine.Hide()
			}
			ax.ticks = append(ax.ticks, tick)
		}
	}
}
