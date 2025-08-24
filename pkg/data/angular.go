package data

// AngularDataPoint represents one data point with an angular coordinate
type AngularDataPoint struct {
	A   float64
	Val float64
}

// DpByAValue is used to sort slices of AngularDataPoint by the a coordinate
type DpByAValue []AngularDataPoint

// Len returns the length of the slice
func (m DpByAValue) Len() int { return len(m) }

// Less returns true if the x value of the ith varibale is smaller than x of the jth variable
func (m DpByAValue) Less(i, j int) bool { return m[i].A < m[j].A }

// Swap swaps the points on positions i and j
func (m DpByAValue) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

// AngularTick represents one tick on a angular axis
type AngularTick struct {
	A           float64
	SupportLine bool
}
