package renderer

import (
	"fyne.io/fyne/v2/canvas"
)

type Tooltip struct {
	X       float32
	Y       float32
	Entries []*canvas.Text
	Box     *canvas.Rectangle
}

func tooltipSize(entries []*canvas.Text) (w float32, h float32) {
	w = 0.0
	h = 0.0
	if len(entries) == 0 {
		return
	}
	w = entries[0].MinSize().Width
	for i := range entries {
		h += entries[i].MinSize().Height + 2
		if entries[i].MinSize().Width > w {
			w = entries[i].MinSize().Width
		}
	}
	return
}
