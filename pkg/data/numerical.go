package data

// NumericalDataPoint represents one data point with a numerical coordinate
type NumericalDataPoint struct {
	N   float64
	Val float64
}

// DpByNValue is used to sort slices of NumericalDataPoint by the x coordinate
type DpByNValue []NumericalDataPoint

// Len returns the length of the slice
func (m DpByNValue) Len() int { return len(m) }

// Less returns true if the x value of the ith varibale is smaller than x of the jth variable
func (m DpByNValue) Less(i, j int) bool { return m[i].N < m[j].N }

// Swap swaps the points on positions i and j
func (m DpByNValue) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

// NumericalCandleStick represents one canlde in a candlestick series over a numerical axis
type NumericalCandleStick struct {
	NStart float64
	NEnd   float64
	Open   float64
	Close  float64
	Low    float64
	High   float64
}

// NumericalBox represents one box in a box series with a numerical coordinate
type NumericalBox struct {
	N             float64
	Maximum       float64
	ThirdQuartile float64
	Median        float64
	FirstQuartile float64
	Minimum       float64
	Outlier       []float64
}

// NumericalTick represents one tick on a numerical axis
type NumericalTick struct {
	N           float64
	SupportLine bool
}
