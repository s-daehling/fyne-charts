package data

import (
	"fyne.io/fyne/v2"
)

// ProportionalPoint represents one data point with a Categorical coordinate that will be scaled according to its Value
type ProportionalPoint struct {
	C       string
	Val     float64
	ColName fyne.ThemeColorName
}
