package series

import (
	"errors"
	"math"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

func numericalDataPointRangeCheck(input []data.NumericalDataPoint, noNegativeVal bool, isPolar bool) (err error) {
	if len(input) == 0 {
		err = errors.New("no input data")
		return
	}
	if isPolar || noNegativeVal {
		for i := range input {
			if (isPolar && (input[i].N < 0 || input[i].N > 2*math.Pi)) ||
				(noNegativeVal && input[i].Val < 0) {
				err = errors.New("invalid data")
				return
			}
		}
	}
	return
}

func temporalDataPointRangeCheck(input []data.TemporalDataPoint, noNegativeVal bool) (err error) {
	if len(input) == 0 {
		err = errors.New("no input data")
		return
	}
	if noNegativeVal {
		for i := range input {
			if input[i].Val < 0 {
				err = errors.New("invalid data")
				return
			}
		}
	}
	return
}

func categoricalDataPointRangeCheck(input []data.CategoricalDataPoint, noNegativeVal bool) (err error) {
	if len(input) == 0 {
		err = errors.New("no input data")
		return
	}
	if noNegativeVal {
		for i := range input {
			if input[i].Val < 0 {
				err = errors.New("invalid data")
				return
			}
		}
	}
	return
}
