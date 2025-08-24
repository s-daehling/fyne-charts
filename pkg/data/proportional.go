package data

import (
	"image/color"
)

// ProportionalDataPoint represents one data point with a Categorical coordinate that will be scaled according to its Value
type ProportionalDataPoint struct {
	C   string
	Val float64
	Col color.Color
}
