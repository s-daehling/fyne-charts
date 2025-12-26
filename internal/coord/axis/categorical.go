package axis

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/software"
	"github.com/disintegration/imaging"
	"github.com/s-daehling/fyne-charts/pkg/data"
)

func (ax *Axis) SetCRange(cs []string) {
	ax.cs = nil
	ax.cs = append(ax.cs, cs...)
}

func (ax *Axis) CRange() (cs []string) {
	cs = append(cs, ax.cs...)
	return
}

func (ax *Axis) SetCTicks(cs []data.CategoricalTick) {
	ax.adjustNumberOfTicks(len(cs))
	spacePerCat := ax.space / float32(len(cs))
	maxTickWidth := ax.maxTickWidth()
	rot := 0.0
	if spacePerCat < maxTickWidth && ax.typ == CartesianHorAxis {
		if math.Cos(math.Pi/8)*float64(maxTickWidth) < float64(spacePerCat) {
			rot = 22.5
		} else if math.Cos(math.Pi/4)*float64(maxTickWidth) < float64(spacePerCat) {
			rot = 45
		} else if math.Cos(3*math.Pi/8)*float64(maxTickWidth) < float64(spacePerCat) {
			rot = 77.5
		} else {
			rot = 90
		}
	}
	for i := range cs {
		ax.ticks[i].c = cs[i].C
		ax.ticks[i].labelText.Text = cs[i].C
		ax.ticks[i].hasSupportLine = cs[i].SupportLine
		c := software.NewTransparentCanvas()
		c.SetPadded(false)
		c.SetContent(ax.ticks[i].labelText)
		ax.ticks[i].label.Image = c.Capture()
		ax.ticks[i].label.Resize(ax.ticks[i].labelText.MinSize())
		ax.ticks[i].label.SetMinSize(ax.ticks[i].labelText.MinSize())
		if rot > 1 {
			ax.ticks[i].label.Image = imaging.Rotate(ax.ticks[i].label.Image, rot, color.RGBA{A: 0x00})
			cos := float32(math.Cos(2 * math.Pi * rot / float64(360)))
			sin := float32(math.Sin(2 * math.Pi * rot / float64(360)))
			w := ax.ticks[i].labelText.MinSize().Width
			h := ax.ticks[i].labelText.MinSize().Height
			minSize := fyne.NewSize(w*cos+h*sin, w*sin+h*cos)
			ax.ticks[i].label.Resize(minSize)
			ax.ticks[i].label.SetMinSize(minSize)
		}
	}
}

func (ax *Axis) AutoCTicks() {
	if !ax.autoTicks {
		return
	}
	cs := make([]data.CategoricalTick, 0)
	for i := range ax.cs {
		cs = append(cs, data.CategoricalTick{C: ax.cs[i], SupportLine: ax.autoSupportLine})
	}
	ax.SetCTicks(cs)
}

func (ax *Axis) ConvertCTickstoN() {
	catSize := (ax.nMax - ax.nMin) / float64(len(ax.cs))
	for i := range ax.ticks {
		ax.ticks[i].nLabel = ax.CtoN(ax.ticks[i].c)
		ax.ticks[i].nLine = ax.CtoN(ax.ticks[i].c) - 0.5*catSize
	}
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

func (ax *Axis) NtoC(n float64) (c string) {
	numCats := len(ax.cs)
	catSize := (ax.nMax - ax.nMin) / float64(numCats)
	pos := int(n / catSize)
	if pos >= 0 && pos < numCats {
		c = ax.cs[pos]
	}
	return
}
