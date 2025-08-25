package data

import (
	"image/color"
	"time"
)

// TemporalDataPoint represents one data point with a temporal coordinate
type TemporalDataPoint struct {
	T   time.Time
	Val float64
}

// DpByTValue is used to sort slices of TemporalDataPoint by the t coordinate
type DpByTValue []TemporalDataPoint

// Len returns the length of the slice
func (m DpByTValue) Len() int { return len(m) }

// Less returns true if the t value of the ith varibale is before the t of the jth variable
func (m DpByTValue) Less(i, j int) bool { return m[i].T.Before(m[i].T) }

// Swap swaps the points on positions i and j
func (m DpByTValue) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

// TemporalDataSeries represents a series of data points with a temporal coordinate
type TemporalDataSeries struct {
	Name   string
	Col    color.Color
	Points []TemporalDataPoint
}

// TemporalCandleStick represents one canlde in a candlestick series over a temoral axis
type TemporalCandleStick struct {
	TStart time.Time
	TEnd   time.Time
	Open   float64
	Close  float64
	Low    float64
	High   float64
}

// TemporalBox represents one box in a box series with a temporal coordinate
type TemporalBox struct {
	T             time.Time
	Maximum       float64
	ThirdQuartile float64
	Median        float64
	FirstQuartile float64
	Minimum       float64
	Outlier       []float64
}

// TemporalTick represents one tick on a temporal axis
type TemporalTick struct {
	T           time.Time
	SupportLine bool
}
