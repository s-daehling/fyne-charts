package elements

import "fyne.io/fyne/v2/canvas"

type Edge struct {
	N1   float64
	Val1 float64
	N2   float64
	Val2 float64
	Line *canvas.Line
}
