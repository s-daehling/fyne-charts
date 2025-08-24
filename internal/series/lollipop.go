package series

import (
	"image/color"
)

type LollipopSeries struct {
	ScatterSeries
}

func EmptyLollipopSeries(chart chart, name string, color color.Color, polar bool) (ser *LollipopSeries) {
	ser = &LollipopSeries{
		ScatterSeries: ScatterSeries{},
	}
	ser.baseSeries = emptyBaseSeries(chart, name, color, polar, ser.toggleView)
	return
}

func (ser *LollipopSeries) CartesianEdges(xMin float64, xMax float64, yMin float64,
	yMax float64) (es []CartesianEdge) {
	ser.mutex.Lock()
	for i := range ser.data {
		es = append(es, ser.data[i].cartesianEdges(xMin, xMax, yMin, yMax)...)
	}
	ser.mutex.Unlock()
	return
}

func (ser *LollipopSeries) PolarEdges(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (es []PolarEdge) {
	ser.mutex.Lock()
	for i := range ser.data {
		es = append(es, ser.data[i].polarEdges(phiMin, phiMax, rMin, rMax)...)
	}
	ser.mutex.Unlock()
	return
}

func (ser *LollipopSeries) SetValOrigin(vo float64) {
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].setValOrigin(vo)
	}
	ser.mutex.Unlock()
}

// SetLineWidth changes the width of the line
// Standard value is 1
// The provided width must be greater than zero for this method to take effect
func (ser *LollipopSeries) SetLineWidth(lw float32) {
	if lw < 0 {
		return
	}
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].setLineWidth(lw)
	}
	ser.mutex.Unlock()
}
