package series

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

type chartDummy struct{}

func (cd chartDummy) DataChange()             {}
func (cd chartDummy) RasterVisibilityChange() {}
func (cd chartDummy) PositionToCartesianCoordinates(pX int, pY int, w int, h int) (x float64, y float64) {
	return
}
func (cd chartDummy) PositionToPolarCoordinates(pX int, pY int, w int, h int) (phi float64, r float64, x float64, y float64) {
	return
}

var ndpTestSetFull = []data.NumericalDataPoint{
	{N: -1000, Val: -1000},
	{N: -1000, Val: -0.000001},
	{N: -1000, Val: 0},
	{N: -1000, Val: 0.000001},
	{N: -1000, Val: 1000},
	{N: -0.000001, Val: -1000},
	{N: -0.000001, Val: -0.000001},
	{N: -0.000001, Val: 0},
	{N: -0.000001, Val: 0.000001},
	{N: -0.000001, Val: 1000},
	{N: 0, Val: -1000},
	{N: 0, Val: -0.000001},
	{N: 0, Val: 0},
	{N: 0, Val: 0.000001},
	{N: 0, Val: 1000},
	{N: 0.000001, Val: -1000},
	{N: 0.000001, Val: -0.000001},
	{N: 0.000001, Val: 0},
	{N: 0.000001, Val: 0.000001},
	{N: 0.000001, Val: 1000},
	{N: math.Pi, Val: -1000},
	{N: math.Pi, Val: -0.000001},
	{N: math.Pi, Val: 0},
	{N: math.Pi, Val: 0.000001},
	{N: math.Pi, Val: 1000},
	{N: 2 * math.Pi, Val: -1000},
	{N: 2 * math.Pi, Val: -0.000001},
	{N: 2 * math.Pi, Val: 0},
	{N: 2 * math.Pi, Val: 0.000001},
	{N: 2 * math.Pi, Val: 1000},
	{N: 1000, Val: -1000},
	{N: 1000, Val: -0.000001},
	{N: 1000, Val: 0},
	{N: 1000, Val: 0.000001},
	{N: 1000, Val: 1000},
}

var ndpTestSetPosVal = []data.NumericalDataPoint{
	{N: -1000, Val: 0},
	{N: -1000, Val: 0.000001},
	{N: -1000, Val: 1000},
	{N: -0.000001, Val: 0},
	{N: -0.000001, Val: 0.000001},
	{N: -0.000001, Val: 1000},
	{N: 0, Val: 0},
	{N: 0, Val: 0.000001},
	{N: 0, Val: 1000},
	{N: 0.000001, Val: 0},
	{N: 0.000001, Val: 0.000001},
	{N: 0.000001, Val: 1000},
	{N: math.Pi, Val: 0},
	{N: math.Pi, Val: 0.000001},
	{N: math.Pi, Val: 1000},
	{N: 2 * math.Pi, Val: 0},
	{N: 2 * math.Pi, Val: 0.000001},
	{N: 2 * math.Pi, Val: 1000},
	{N: 1000, Val: 0},
	{N: 1000, Val: 0.000001},
	{N: 1000, Val: 1000},
}

var ndpTestSetPosValPolar = []data.NumericalDataPoint{
	{N: 0, Val: 0},
	{N: 0, Val: 0.000001},
	{N: 0, Val: 1000},
	{N: 0.000001, Val: 0},
	{N: 0.000001, Val: 0.000001},
	{N: 0.000001, Val: 1000},
	{N: math.Pi, Val: 0},
	{N: math.Pi, Val: 0.000001},
	{N: math.Pi, Val: 1000},
	{N: 2 * math.Pi, Val: 0},
	{N: 2 * math.Pi, Val: 0.000001},
	{N: 2 * math.Pi, Val: 1000},
}

var tdpTestSetFull = []data.TemporalDataPoint{
	{T: time.Now(), Val: -1000},
	{T: time.Now().Add(-time.Hour), Val: -0.000001},
	{T: time.Now().Add(time.Hour), Val: 0},
	{T: time.Now().Add(-2 * time.Hour), Val: 0.000001},
	{T: time.Now().Add(2 * time.Hour), Val: 1000},
}

var tdpTestSetPosVal = []data.TemporalDataPoint{
	{T: time.Now().Add(time.Hour), Val: 0},
	{T: time.Now().Add(-2 * time.Hour), Val: 0.000001},
	{T: time.Now().Add(2 * time.Hour), Val: 1000},
}

var cdpTestSetFull = []data.CategoricalDataPoint{
	{C: "one", Val: -1000},
	{C: "two", Val: -0.000001},
	{C: "three", Val: 0},
	{C: "four", Val: 0.000001},
	{C: "five", Val: 1000},
}

var cdpTestSetPosVal = []data.CategoricalDataPoint{
	{C: "one", Val: 0},
	{C: "two", Val: 0.000001},
	{C: "three", Val: 1000},
}

func testNRange(ser Series, expIsEmpty bool, expMin float64, expMax float64) (err error) {
	isEmpty, min, max := ser.NRange()
	if isEmpty && !expIsEmpty {
		err = errors.New("NRange incorrect, is empty")
		return
	}
	if !isEmpty && expIsEmpty {
		err = errors.New("NRange incorrect, is not empty")
		return
	}
	if isEmpty {
		return
	}
	if min < expMin-0.000001 || min > expMin+0.000001 {
		err = fmt.Errorf("NRange incorrect, min: %f, exp. min: %f", min, expMin)
		return
	}
	if max < expMax-0.000001 || max > expMax+0.000001 {
		err = fmt.Errorf("NRange incorrect, max: %f, exp. max: %f", max, expMax)
		return
	}
	return
}

func testTRange(ser Series, expIsEmpty bool, expMin time.Time, expMax time.Time) (err error) {
	isEmpty, min, max := ser.TRange()
	if isEmpty && !expIsEmpty {
		err = errors.New("NRange incorrect, is empty")
		return
	}
	if !isEmpty && expIsEmpty {
		err = errors.New("NRange incorrect, is not empty")
		return
	}
	if isEmpty {
		return
	}
	if min != expMin {
		err = fmt.Errorf("TRange incorrect, min: %s, exp: %s", min.String(), expMin.String())
		return
	}
	if max != expMax {
		err = fmt.Errorf("TRange incorrect, max: %s, exp: %s", max.String(), expMax.String())
		return
	}
	return
}

func testCRange(ser Series, expCRange []string) (err error) {
	cRange := ser.CRange()
	if !slices.Equal(expCRange, cRange) {
		err = fmt.Errorf("CRange incorrect, range: %s, exp: %s", strings.Join(cRange, " "), strings.Join(expCRange, " "))
		return
	}
	return
}

func testValRange(ser Series, expIsEmpty bool, expMin float64, expMax float64) (err error) {
	isEmpty, min, max := ser.ValRange()
	if isEmpty && !expIsEmpty {
		err = errors.New("ValRange incorrect, is empty")
		return
	}
	if !isEmpty && expIsEmpty {
		err = errors.New("ValRange incorrect, is not empty")
		return
	}
	if isEmpty {
		return
	}
	if min < expMin-0.000001 || min > expMin+0.000001 {
		err = fmt.Errorf("ValRange incorrect, min: %f, exp. min: %f", min, expMin)
		return
	}
	if max < expMax-0.000001 || max > expMax+0.000001 {
		err = fmt.Errorf("ValRange incorrect, max: %f, exp. max: %f", max, expMax)
		return
	}
	return
}

func testCartesianNodes(ser Series, xMin float64, xMax float64, yMin float64, yMax float64, expNodes int) (err error) {
	ns := ser.CartesianNodes(xMin, xMax, yMin, yMax)
	if len(ns) != expNodes {
		err = fmt.Errorf("CartesianNodes incorrect, nodes %d, exp. nodes: %d", len(ns), expNodes)
	}
	return
}
