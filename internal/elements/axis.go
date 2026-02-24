package elements

import "fyne.io/fyne/v2/canvas"

type Tick struct {
	NLabel    float64
	Label     *canvas.Image
	NLine     float64
	Line      *canvas.Line
	SupLine   *canvas.Line
	SupCircle *canvas.Circle
}

type Arrow struct {
	Line    *canvas.Line
	Circle  *canvas.Circle
	HeadOne *canvas.Line
	HeadTwo *canvas.Line
}

func MaxTickSize(ts []Tick) (maxWidth float32, maxHeight float32) {
	maxWidth = 0
	maxHeight = 0
	for i := range ts {
		if ts[i].Label.Size().Width > maxWidth {
			maxWidth = ts[i].Label.Size().Width
		}
		if ts[i].Label.Size().Height > maxHeight {
			maxHeight = ts[i].Label.Size().Height
		}
	}
	return
}
