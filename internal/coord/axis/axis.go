package axis

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/software"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"github.com/disintegration/imaging"
	"github.com/s-daehling/fyne-charts/internal/elements"
	"github.com/s-daehling/fyne-charts/pkg/style"
)

type AxisType string

const (
	CartesianHorAxis  AxisType = "CartesianHor"
	CartesianVertAxis AxisType = "CartesianVert"
	PolarPhiAxis      AxisType = "PolarPhi"
	PolarRAxis        AxisType = "PolarR"
)

type axisTick struct {
	c              string
	t              time.Time
	n              float64
	nLabel         float64
	nLine          float64
	label          *canvas.Image
	labelText      *canvas.Text // the text label
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
	labelStyle      style.ChartTextStyle
	space           float32
	style           style.AxisStyle
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
	}
	ax.SetAxisLabelStyle(style.DefaultAxisLabelStyle())
	ax.SetAxisStyle(style.DefaultAxisStyle())
	if typ == PolarPhiAxis {
		ax.nMax = 2 * math.Pi
	}
	ax.circle.StrokeColor = col
	ax.circle.StrokeWidth = 1
	ax.SetLabel("")
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
	return
}

func (ax *Axis) Arrow() (ar elements.Arrow) {
	ar.Line = ax.line
	ar.Circle = ax.circle
	ar.HeadOne = ax.arrowOne
	ar.HeadTwo = ax.arrowTwo
	return
}

func (ax *Axis) Ticks() (ts []elements.Tick) {
	for i := range ax.ticks {
		if ax.ticks[i].n < ax.nMin || ax.ticks[i].n > ax.nMax {
			continue
		}
		t := elements.Tick{
			NLabel:  ax.ticks[i].nLabel,
			NLine:   ax.ticks[i].nLine,
			Label:   nil,
			Line:    nil,
			SupLine: nil,
		}
		if t.NLabel > ax.nMin || t.NLabel < ax.nMax {
			// t.Label.Text = ax.ticks[i].labelText
			t.Label = ax.ticks[i].label
		}
		if t.NLine > ax.nMin || t.NLine < ax.nMax {
			t.Line = ax.ticks[i].line
			if ax.ticks[i].hasSupportLine {
				if ax.typ == CartesianHorAxis || ax.typ == CartesianVertAxis || ax.typ == PolarPhiAxis {
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

func (ax *Axis) maxTickWidth() (maxWidth float32) {
	maxWidth = 0
	for i := range ax.ticks {
		if ax.ticks[i].labelText.MinSize().Width > maxWidth {
			maxWidth = ax.ticks[i].labelText.MinSize().Width
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
		ax.ticks[i].labelText.Hide()
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
		ax.ticks[i].labelText.Show()
		ax.ticks[i].line.Show()
		ax.ticks[i].supportCircle.Show()
		ax.ticks[i].supportLine.Show()
	}
}

func (ax *Axis) Visible() (b bool) {
	b = ax.visible
	return
}

func (ax *Axis) RefreshTheme() {
	ax.arrowOne.StrokeColor = theme.Color(ax.style.LineColorName)
	ax.arrowTwo.StrokeColor = theme.Color(ax.style.LineColorName)
	ax.line.StrokeColor = theme.Color(ax.style.LineColorName)
	ax.circle.StrokeColor = theme.Color(ax.style.LineColorName)
	for i := range ax.ticks {
		ax.ticks[i].labelText.Color = theme.Color(ax.style.TickColorName)
		ax.ticks[i].labelText.TextSize = theme.Size(ax.style.TickSizeName)
		ax.ticks[i].line.StrokeColor = theme.Color(ax.style.LineColorName)
		ax.ticks[i].supportCircle.StrokeColor = theme.Color(ax.style.SupportLineColorName)
		ax.ticks[i].supportLine.StrokeColor = theme.Color(ax.style.SupportLineColorName)
	}
}

func (ax *Axis) SetAxisStyle(s style.AxisStyle) {
	ax.style = s
	ax.arrowOne.StrokeColor = theme.Color(s.LineColorName)
	ax.arrowOne.StrokeWidth = s.LineWidth
	ax.arrowTwo.StrokeColor = theme.Color(s.LineColorName)
	ax.arrowTwo.StrokeWidth = s.LineWidth
	if s.LineShowArrow {
		ax.arrowOne.Show()
		ax.arrowTwo.Show()
	} else {
		ax.arrowOne.Hide()
		ax.arrowTwo.Hide()
	}
	ax.line.StrokeColor = theme.Color(s.LineColorName)
	ax.line.StrokeWidth = s.LineWidth
	ax.circle.StrokeColor = theme.Color(s.LineColorName)
	ax.circle.StrokeWidth = s.LineWidth
	for i := range ax.ticks {
		ax.ticks[i].labelText.Color = theme.Color(s.TickColorName)
		ax.ticks[i].labelText.TextSize = theme.Size(s.TickSizeName)
		ax.ticks[i].labelText.TextStyle = s.TickTextStyle
		ax.ticks[i].line.StrokeColor = theme.Color(s.LineColorName)
		ax.ticks[i].line.StrokeWidth = s.LineWidth
		ax.ticks[i].supportCircle.StrokeColor = theme.Color(s.SupportLineColorName)
		ax.ticks[i].supportCircle.StrokeWidth = s.SupportLineWidth
		ax.ticks[i].supportLine.StrokeColor = theme.Color(s.SupportLineColorName)
		ax.ticks[i].supportLine.StrokeWidth = s.SupportLineWidth
	}
}

func (ax *Axis) SetLabel(l string) {
	ax.name = l
	ax.labelText.Text = l
	c := software.NewTransparentCanvas()
	c.SetPadded(false)
	c.SetContent(ax.labelText)
	ax.label.Image = c.Capture()
	minSize := ax.labelText.MinSize()
	// if l == "" {
	// 	minSize.Height = 0
	// }
	ax.label.Resize(minSize)
	ax.label.SetMinSize(minSize)
	if ax.typ == CartesianVertAxis || ax.typ == PolarPhiAxis {
		ax.label.Image = imaging.Rotate90(ax.label.Image)
		ax.label.Resize(fyne.NewSize(minSize.Height, minSize.Width))
		ax.label.SetMinSize(fyne.NewSize(minSize.Height, minSize.Width))
	}
	if l == "" && !ax.label.Hidden {
		ax.label.Hide()
	} else if l != "" && ax.label.Hidden {
		ax.label.Show()
	}
}

func (ax *Axis) SetAxisLabelStyle(ls style.ChartTextStyle) {
	ax.labelStyle = ls
	ax.labelText.TextSize = theme.Size(ls.SizeName)
	ax.labelText.Color = theme.Color(ls.ColorName)
	ax.labelText.Alignment = ls.Alignment
	ax.labelText.TextStyle = ls.TextStyle
	ax.SetLabel(ax.name)
}

func (ax *Axis) Label() (l *canvas.Image) {
	l = ax.label
	return
}

func (ax *Axis) AddLabelToContainer(cont *fyne.Container) {
	if ax.labelStyle.Alignment != fyne.TextAlignLeading {
		cont.Add(layout.NewSpacer())
	}
	cont.Add(ax.label)
	if ax.labelStyle.Alignment != fyne.TextAlignTrailing {
		cont.Add(layout.NewSpacer())
	}
}

func (ax *Axis) CartesianTranspose() {
	switch ax.typ {
	case CartesianVertAxis:
		ax.typ = CartesianHorAxis
	case CartesianHorAxis:
		ax.typ = CartesianVertAxis
	}
	ax.SetLabel(ax.name)
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
				labelText:      canvas.NewText("", theme.Color(ax.style.TickColorName)),
				label:          canvas.NewImageFromImage(software.NewTransparentCanvas().Capture()),
				line:           canvas.NewLine(theme.Color(ax.style.LineColorName)),
				hasSupportLine: false,
				supportLine:    canvas.NewLine(theme.Color(ax.style.SupportLineColorName)),
				supportCircle:  canvas.NewCircle(color.RGBA{0x00, 0x00, 0x00, 0x00}),
			}
			// tick.supportLine.StrokeWidth = 0.5
			tick.labelText.TextSize = theme.Size(ax.style.TickSizeName)
			tick.labelText.TextStyle = ax.style.TickTextStyle
			tick.line.StrokeWidth = ax.style.LineWidth
			tick.supportCircle.StrokeColor = theme.Color(ax.style.SupportLineColorName)
			tick.supportCircle.StrokeWidth = ax.style.SupportLineWidth
			if !ax.visible {
				tick.labelText.Hide()
				tick.line.Hide()
				tick.supportCircle.Hide()
				tick.supportLine.Hide()
			}
			ax.ticks = append(ax.ticks, tick)
		}
	}
}
