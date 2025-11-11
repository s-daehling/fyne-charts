package interact

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type Tooltip struct {
	x          float32
	y          float32
	mouseIn    bool
	actCounter int
	entries    []*canvas.Text
	box        *canvas.Rectangle
}

func NewTooltip() (tt *Tooltip) {
	tt = &Tooltip{
		x:       0,
		y:       0,
		mouseIn: false,
		box:     canvas.NewRectangle(theme.Color(theme.ColorNameBackground)),
	}
	tt.box.CornerRadius = 5
	tt.box.StrokeColor = theme.Color(theme.ColorNameForeground)
	tt.box.StrokeWidth = 0.5
	return
}

func (tt *Tooltip) MouseIn(x, y float32) {
	tt.mouseIn = true
	tt.x = x
	tt.y = y

}

func (tt *Tooltip) MouseMove(x, y float32) (c int) {
	tt.x = x
	tt.y = y
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

func (tt *Tooltip) GetEntries() (x float32, y float32, entries []*canvas.Text, box *canvas.Rectangle) {
	x = tt.x
	y = tt.y
	if tt.mouseIn {
		entries = append(entries, tt.entries...)
		box = tt.box
	} else {
		box = nil
	}
	return
}
