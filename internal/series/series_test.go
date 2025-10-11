package series

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
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
