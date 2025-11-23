package data

// CategoricalPoint represents one data point with a categorical coordinate
type CategoricalPoint struct {
	C   string
	Val float64
}

// CategoricalBox represents one box in a box series with a categorical coordinate
type CategoricalBox struct {
	C             string
	Maximum       float64
	ThirdQuartile float64
	Median        float64
	FirstQuartile float64
	Minimum       float64
	Outlier       []float64
}

// CategoricalTick represents one tick on a categorical axis
type CategoricalTick struct {
	C           string
	SupportLine bool
}
