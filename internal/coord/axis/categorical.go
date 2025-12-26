package axis

import (
	"fyne.io/fyne/v2/driver/software"
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
