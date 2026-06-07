package data

import (
	"sort"
	"time"
)

// TemporalPoint represents one data point with a temporal coordinate
type TemporalPoint struct {
	T   time.Time
	Val float64
}

// DpByTValue is used to sort slices of TemporalPoint by the t coordinate
type DpByTValue []TemporalPoint

// Len returns the length of the slice
func (m DpByTValue) Len() int { return len(m) }

// Less returns true if the t value of the ith varibale is before the t of the jth variable
func (m DpByTValue) Less(i, j int) bool { return m[i].T.Before(m[i].T) }

// Swap swaps the points on positions i and j
func (m DpByTValue) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

type TemporalPointList struct {
	boundList[TemporalPoint]
}

func NewTemporalPointList() (l ConstrainedList[TemporalPoint]) {
	l = &TemporalPointList{boundList: newList[TemporalPoint]()}
	return
}

func (tpl *TemporalPointList) Add(val TemporalPoint) (err error) {
	tpl.lock.Lock()
	v := *tpl.val
	tpl.lock.Unlock()
	v = append(v, val)
	err = tpl.Set(v)
	return
}

func (tpl *TemporalPointList) Set(l []TemporalPoint) (err error) {
	sort.Sort(DpByTValue(l))
	err = tpl.boundList.Set(l)
	return
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
