package series

import (
	"errors"
	"math"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

func angularToNumerical(input []data.AngularDataPoint) (output []data.NumericalDataPoint) {
	for i := range input {
		output = append(output, data.NumericalDataPoint{X: input[i].A, Val: input[i].Val})
	}
	return
}

func numericalDataPointRangeCheck(input []data.NumericalDataPoint, noNegativeVal bool, isPolar bool) (err error) {
	if len(input) == 0 {
		err = errors.New("no input data")
		return
	}
	if isPolar || noNegativeVal {
		for i := range input {
			if (isPolar && (input[i].X < 0 || input[i].X > 2*math.Pi)) ||
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
