package data

import (
	"image/color"
)

// ProportionalPoint represents one data point with a Categorical coordinate that will be scaled according to its Value
type ProportionalPoint struct {
	C   string
	Val float64
	Col color.Color
}
