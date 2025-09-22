package chart

import (
	"image/color"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/s-daehling/fyne-charts/pkg/data"

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
	mutex           *sync.Mutex // mutex to prevent concurrent access
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
		mutex:           &sync.Mutex{},
	}
	if typ == PolarPhiAxis {
		ax.nMax = 2 * math.Pi
	}
	ax.circle.StrokeColor = col
	ax.circle.StrokeWidth = 1
	// ax.circle.FillColor = color.RGBA{0x00, 0x00, 0x00, 0x00}
	return
}

func (ax *Axis) AutoTicks() (a bool) {
	a = ax.autoTicks
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
	ax.mutex.Lock()
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
	ax.mutex.Unlock()
	return
}

func (ax *Axis) Hide() {
	ax.mutex.Lock()
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
	ax.mutex.Unlock()
}

func (ax *Axis) Show() {
	ax.mutex.Lock()
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
	ax.mutex.Unlock()
}

func (ax *Axis) SetLabel(l string) {
	ax.mutex.Lock()
	ax.name = l
	ax.labelText.Text = l
	ax.mutex.Unlock()
}

func (ax *Axis) Label() (label *canvas.Image, text *canvas.Text) {
	ax.mutex.Lock()
	label = ax.label
	text = ax.labelText
	ax.mutex.Unlock()
	return
}

func (ax *Axis) SetAutoTicks(autoSupport bool) {
	ax.mutex.Lock()
	ax.autoTicks = true
	ax.autoSupportLine = autoSupport
	ax.mutex.Unlock()
}

func (ax *Axis) SetManualTicks() {
	ax.mutex.Lock()
	ax.autoTicks = false
	ax.autoSupportLine = false
	ax.mutex.Unlock()
}

func (ax *Axis) adjustNumberOfTicks(n int) {
	ax.mutex.Lock()
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
	ax.mutex.Unlock()
}

func (ax *Axis) MaxTickWidth() (maxWidth float32) {
	maxWidth = 0
	ax.mutex.Lock()
	for i := range ax.ticks {
		if ax.ticks[i].label.MinSize().Width > maxWidth {
			maxWidth = ax.ticks[i].label.MinSize().Width
		}
	}
	ax.mutex.Unlock()
	return
}

func (ax *Axis) MaxTickHeight() (maxHeight float32) {
	maxHeight = 0
	ax.mutex.Lock()
	for i := range ax.ticks {
		if ax.ticks[i].label.MinSize().Height > maxHeight {
			maxHeight = ax.ticks[i].label.MinSize().Height
		}
	}
	ax.mutex.Unlock()
	return
}

func (ax *Axis) SetCRange(cs []string) {
	ax.mutex.Lock()
	ax.cs = nil
	ax.cs = append(ax.cs, cs...)
	ax.mutex.Unlock()
}

func (ax *Axis) CRange() (cs []string) {
	ax.mutex.Lock()
	cs = append(cs, ax.cs...)
	ax.mutex.Unlock()
	return
}

func (ax *Axis) SetTOrigin(o time.Time) {
	ax.mutex.Lock()
	ax.tOrigin = o
	ax.mutex.Unlock()
}

func (ax *Axis) AutoTOrigin() {
	ax.mutex.Lock()
	ax.tOrigin = ax.tMin
	ax.mutex.Unlock()
}

func (ax *Axis) TOrigin() (o time.Time) {
	ax.mutex.Lock()
	o = ax.tOrigin
	ax.mutex.Unlock()
	return
}

func (ax *Axis) SetTRange(min time.Time, max time.Time) {
	ax.mutex.Lock()
	ax.tMin = min
	ax.tMax = max
	ax.mutex.Unlock()
}

func (ax *Axis) TRange() (min time.Time, max time.Time) {
	ax.mutex.Lock()
	min = ax.tMin
	max = ax.tMax
	ax.mutex.Unlock()
	return
}

func (ax *Axis) SetNOrigin(o float64) {
	ax.mutex.Lock()
	ax.nOrigin = o
	ax.mutex.Unlock()
}

func (ax *Axis) AutoNOrigin() {
	ax.mutex.Lock()
	switch ax.typ {
	case CartesianAxis:
		if ax.nMin < 0 && ax.nMax > 0 {
			ax.nOrigin = 0
		} else {
			ax.nOrigin = ax.nMin
		}
	case PolarPhiAxis:
		ax.nOrigin = ax.nMin
	case PolarRAxis:
		ax.nOrigin = ax.nMax
	}
	ax.mutex.Unlock()
}

func (ax *Axis) NOrigin() (o float64) {
	ax.mutex.Lock()
	o = ax.nOrigin
	ax.mutex.Unlock()
	return
}

func (ax *Axis) SetNRange(min float64, max float64) {
	ax.mutex.Lock()
	ax.nMin = min
	ax.nMax = max
	ax.mutex.Unlock()
}

func (ax *Axis) NRange() (min float64, max float64) {
	ax.mutex.Lock()
	min = ax.nMin
	max = ax.nMax
	ax.mutex.Unlock()
	return
}

func (ax *Axis) SetSpace(space float32) {
	ax.mutex.Lock()
	ax.space = space
	ax.mutex.Unlock()
}

func (ax *Axis) SetCTicks(cs []data.CategoricalTick) {
	ax.adjustNumberOfTicks(len(cs))
	ax.mutex.Lock()
	for i := range cs {
		ax.ticks[i].c = cs[i].C
		ax.ticks[i].label.Text = cs[i].C
		ax.ticks[i].hasSupportLine = cs[i].SupportLine
	}
	ax.mutex.Unlock()
}

func (ax *Axis) AutoCTicks() {
	if !ax.autoTicks {
		return
	}
	cs := make([]data.CategoricalTick, 0)
	ax.mutex.Lock()
	for i := range ax.cs {
		cs = append(cs, data.CategoricalTick{C: ax.cs[i], SupportLine: ax.autoSupportLine})
	}
	ax.mutex.Unlock()
	ax.SetCTicks(cs)
}

func (ax *Axis) ConvertCTickstoN() {
	ax.mutex.Lock()
	catSize := (ax.nMax - ax.nMin) / float64(len(ax.cs))
	for i := range ax.ticks {
		ax.ticks[i].nLabel = ax.CtoN(ax.ticks[i].c)
		ax.ticks[i].nLine = ax.CtoN(ax.ticks[i].c) - 0.5*catSize
	}
	ax.mutex.Unlock()
}

func (ax *Axis) SetTTicks(ts []data.TemporalTick, format string) {
	ax.adjustNumberOfTicks(len(ts))
	ax.mutex.Lock()
	for i := range ts {
		ax.ticks[i].t = ts[i].T
		ax.ticks[i].label.Text = ts[i].T.Format(format)
		ax.ticks[i].hasSupportLine = ts[i].SupportLine
	}
	ax.mutex.Unlock()
}

func (ax *Axis) AutoTTicks() {
	if !ax.autoTicks {
		return
	}
	ax.mutex.Lock()
	min := ax.tMin
	max := ax.tMax
	ax.mutex.Unlock()
	ts, format := calculateTTicks(ax.space, min, max, ax.autoSupportLine)
	ax.SetTTicks(ts, format)
}

func (ax *Axis) ConvertTTickstoN() {
	ax.mutex.Lock()
	for i := range ax.ticks {
		ax.ticks[i].nLabel = ax.TtoN(ax.ticks[i].t)
		ax.ticks[i].nLine = ax.TtoN(ax.ticks[i].t)
	}
	ax.nOrigin = ax.TtoN(ax.tOrigin)
	ax.mutex.Unlock()
}

func (ax *Axis) SetNTicks(ns []data.NumericalTick, orderOfMagn int) {
	ax.adjustNumberOfTicks(len(ns))
	prec := -orderOfMagn + 1
	if prec < 0 {
		prec = 0
	}
	ax.mutex.Lock()
	for i := range ns {
		ax.ticks[i].n = ns[i].N
		ax.ticks[i].nLabel = ns[i].N
		ax.ticks[i].nLine = ns[i].N
		ax.ticks[i].hasSupportLine = ns[i].SupportLine
		if ax.typ == PolarPhiAxis {
			ax.ticks[i].label.Text = strconv.FormatFloat(ns[i].N/math.Pi, 'f', 2, 64) + " pi"
		} else {
			ax.ticks[i].label.Text = strconv.FormatFloat(ns[i].N, 'f', prec, 64)
		}
	}
	ax.mutex.Unlock()
}

func (ax *Axis) AutoNTicks() {
	if !ax.autoTicks {
		return
	}
	ax.mutex.Lock()
	min := ax.nMin
	max := ax.nMax
	ax.mutex.Unlock()
	if ax.typ == PolarPhiAxis {
		ns := calculatePhiTicks(ax.autoSupportLine)
		ax.SetNTicks(ns, 1)
	} else {
		ns, orderOfMagn := calculateNTicks(ax.space, min, max, ax.autoSupportLine)
		ax.SetNTicks(ns, orderOfMagn)
	}
}

func calculateTTicks(space float32, min time.Time, max time.Time, supLine bool) (ts []data.TemporalTick, format string) {
	minSpacePerLabel := 100
	maxTickNum := int(space / float32(minSpacePerLabel))
	if maxTickNum == 0 {
		maxTickNum = 1
	}
	r := max.Sub(min)
	numYears := int(r.Hours() / (365 * 24))
	numDays := int(r.Hours() / 24)
	var coord time.Time
	var step time.Duration
	addMin := false
	if numYears > (maxTickNum / 2) {
		// #years > #ticks/2 -> use years as ticks
		inc := numYears/maxTickNum + 1
		coord = time.Date(min.Year(), time.January, 1, 0, 0, 0, 0, time.Local)
		step = time.Duration(int(time.Hour) * 24 * 365 * inc)
		format = "2006"
		if min.Month() == time.January && min.Day() <= 7 {
			addMin = true
		}
	} else if numDays/30 > maxTickNum/2 {
		// #months > #ticks/2 -> use months as ticks
		inc := (numDays/30)/maxTickNum + 1
		coord = time.Date(min.Year(), min.Month(), 1, 0, 0, 0, 0, time.Local)
		step = time.Duration(int(time.Hour) * 24 * 31 * inc)
		format = "01.2006"
		if min.Day() == 1 {
			addMin = true
		}
	} else if numDays > maxTickNum/2 {
		// #days > #ticks/2 -> days as ticks
		inc := numDays/maxTickNum + 1
		coord = time.Date(min.Year(), min.Month(), min.Day(), 0, 0, 0, 0, time.Local)
		step = time.Duration(int(time.Hour) * 24 * inc)
		format = "02.01.2006"
		if min.Hour() < 1 {
			addMin = true
		}
	} else if int(r.Hours()) > maxTickNum/2 {
		// #hours > #ticks/2 -> hours as ticks
		inc := int(r.Hours())/maxTickNum + 1
		coord = time.Date(min.Year(), min.Month(), min.Day(), min.Hour(), 0, 0, 0, time.Local)
		step = time.Duration(int(time.Hour) * inc)
		format = "15h"
		if min.Minute() < 5 {
			addMin = true
		}
	} else if int(r.Minutes()) > maxTickNum/2 {
		// #mins > #ticks/2 -> mins as ticks
		inc := int(r.Minutes())/maxTickNum + 1
		coord = time.Date(min.Year(), min.Month(), min.Day(), min.Hour(), min.Minute(), 0, 0, time.Local)
		step = time.Duration(int(time.Minute) * inc)
		format = "15:04"
		if min.Second() < 5 {
			addMin = true
		}
	} else if int(r.Seconds()) > maxTickNum/2 {
		// #secs > #ticks/2 -> secs as ticks
		inc := int(r.Seconds())/maxTickNum + 1
		coord = time.Date(min.Year(), min.Month(), min.Day(), min.Hour(), min.Minute(), min.Second(), 0, time.Local)
		step = time.Duration(int(time.Second) * inc)
		format = "15:04:05"
		if min.Sub(coord).Milliseconds() < 25 {
			addMin = true
		}
	} else {
		// #msecs as ticks
		inc := int(r.Milliseconds())/maxTickNum + 1
		coord = time.Date(min.Year(), min.Month(), min.Day(), min.Hour(), min.Minute(), min.Second(), min.Nanosecond(), time.Local)
		step = time.Duration(int(time.Millisecond) * inc)
		format = "05.000"
		if min.Sub(coord).Nanoseconds() < 100 {
			addMin = true
		}
	}
	if addMin {
		ts = append(ts, data.TemporalTick{T: min, SupportLine: supLine})
	}
	for {
		coord = coord.Add(step)
		if coord.After(max) {
			break
		}
		ts = append(ts, data.TemporalTick{T: coord, SupportLine: supLine})
	}
	return
}

func calculateNTicks(space float32, min float64, max float64,
	supLine bool) (ns []data.NumericalTick, orderOfMagn int) {
	minSpacePerLabel := 50
	maxTickNum := int(space / float32(minSpacePerLabel))
	if maxTickNum == 0 {
		maxTickNum = 1
	}
	r := max - min
	orderOfMagn = 0
	dist := 1.0
	// find upper limit for orderOfMagn
	for {
		if math.Pow10(orderOfMagn) < r {
			orderOfMagn++
		} else {
			break
		}
	}
	for {
		if r/(5.0*math.Pow10(orderOfMagn-1)) < float64(maxTickNum) {
			// more than 5*10^(magnOfOrder-1) ticks fit -> further decrease magnOfOrder
			orderOfMagn -= 1
			if orderOfMagn < -12 {
				break
			}
		} else {
			break
		}
	}
	if r/(1.0*math.Pow10(orderOfMagn)) > float64(maxTickNum) {
		// too many ticks if dist=1*10^magnOfOrder
		if r/(2.0*math.Pow10(orderOfMagn)) > float64(maxTickNum) {
			// too many ticks if dist=2*10^magnOfOrder
			dist = 5.0 * math.Pow10(orderOfMagn)
		} else {
			dist = 2.0 * math.Pow10(orderOfMagn)
		}
	} else {
		dist = 1.0 * math.Pow10(orderOfMagn)
	}

	// if min is not a multiplier of dist, calculate offset
	m := int(min / dist)
	offset := min - float64(m)*dist
	if offset > dist/1.25 {
		offset -= dist
	}

	i := 0
	for {
		coord := min + float64(i)*dist - offset
		// if i > 0 {
		// 	coord -= offset
		// }
		if coord > max+0.001*math.Pow10(orderOfMagn) {
			break
		}
		ns = append(ns, data.NumericalTick{N: coord, SupportLine: supLine})
		i += 1
	}
	return
}

func calculatePhiTicks(supLine bool) (as []data.NumericalTick) {
	// as, _ = calculateNTicks(space, min, max)
	as = []data.NumericalTick{
		{
			N:           0,
			SupportLine: supLine,
		},
		{
			N:           math.Pi / 4,
			SupportLine: supLine,
		},
		{
			N:           math.Pi / 2,
			SupportLine: supLine,
		},
		{
			N:           3 * math.Pi / 4,
			SupportLine: supLine,
		},
		{
			N:           math.Pi,
			SupportLine: supLine,
		},
		{
			N:           5 * math.Pi / 4,
			SupportLine: supLine,
		},
		{
			N:           3 * math.Pi / 2,
			SupportLine: supLine,
		},
		{
			N:           7 * math.Pi / 4,
			SupportLine: supLine,
		},
	}
	return
}

func (ax *Axis) TtoN(t time.Time) (n float64) {
	r := ax.tMax.Sub(ax.tMin).Nanoseconds()
	p := t.Sub(ax.tMin).Nanoseconds()
	n = ax.nMin + (float64(p) / float64(r) * (ax.nMax - ax.nMin))
	return
}

func (ax *Axis) CtoN(c string) (n float64) {
	numCats := len(ax.cs)
	pos := -1
	for i := range ax.cs {
		if c == ax.cs[i] {
			pos = i
			break
		}
	}
	catSize := (ax.nMax - ax.nMin) / float64(numCats)
	n = ax.nMin + (catSize * (0.5 + float64(pos)))
	return
}

func (ax *Axis) PtoN(p float64) (n float64) {
	n = p * (ax.nMax - ax.nMin)
	return
}
