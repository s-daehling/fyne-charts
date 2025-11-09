package interact

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type Tooltip struct {
	cFrom      float64
	cTo        float64
	mouseIn    bool
	actCounter int
	entries    []*canvas.Text
}

func NewTooltip() (tt *Tooltip) {
	tt = &Tooltip{
		cFrom:   0,
		cTo:     0,
		mouseIn: false,
	}
	return
}

func (tt *Tooltip) MouseIn(from, to float64) {
	tt.mouseIn = true
	tt.cFrom = from
	tt.cTo = to

}

func (tt *Tooltip) MouseMove(from, to float64) (c int) {
	tt.cFrom = from
	tt.cTo = to
	tt.actCounter++
	c = tt.actCounter
	return
}

func (tt *Tooltip) MouseOut() {
	tt.mouseIn = false
}

func (tt *Tooltip) SetEntries(ent []string) {
	tt.entries = []*canvas.Text{}
	for i := range ent {
		tt.entries = append(tt.entries, canvas.NewText(ent[i], theme.Color(theme.ColorNameForeground)))
	}
	tt.actCounter = 0
}

func (tt *Tooltip) GetEntries() (from float64, to float64, entries []*canvas.Text) {
	from = tt.cFrom
	to = tt.cTo
	if tt.mouseIn {
		entries = append(entries, tt.entries...)
	}
	return
}
