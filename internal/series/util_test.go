package series

import (
	"math"
	"testing"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

func TestNumericalDataPointRangeCheck(t *testing.T) {
	for _, tt := range ndpTestSetFull {
		err := numericalDataPointRangeCheck([]data.NumericalDataPoint{tt}, false, false)
		if err != nil {
			t.Errorf("point incorrectly rejected; N: %f, Val: %f, noNegative: %t, polar: %t", tt.N, tt.Val, false, false)
		}

		err = numericalDataPointRangeCheck([]data.NumericalDataPoint{tt}, true, false)
		if tt.Val < 0 {
			if err == nil {
				t.Errorf("point incorrectly accpeted; N: %f, Val: %f, noNegative: %t, polar: %t", tt.N, tt.Val, true, false)
			}
		} else if err != nil {
			t.Errorf("point incorrectly rejected; N: %f, Val: %f, noNegative: %t, polar: %t", tt.N, tt.Val, true, false)
		}

		err = numericalDataPointRangeCheck([]data.NumericalDataPoint{tt}, false, true)
		if tt.N < 0 || tt.N > 2*math.Pi {
			if err == nil {
				t.Errorf("point incorrectly accepted; N: %f, Val: %f, noNegative: %t, polar: %t", tt.N, tt.Val, false, true)
			}
		} else if err != nil {
			t.Errorf("point incorrectly rejected; N: %f, Val: %f, noNegative: %t, polar: %t", tt.N, tt.Val, false, true)
		}

		err = numericalDataPointRangeCheck([]data.NumericalDataPoint{tt}, true, true)
		if (tt.N < 0 || tt.N > 2*math.Pi) || tt.Val < 0 {
			if err == nil {
				t.Errorf("point incorrectly accepted; N: %f, Val: %f, noNegative: %t, polar: %t", tt.N, tt.Val, true, true)
			}
		} else if err != nil {
			t.Errorf("point incorrectly rejected; N: %f, Val: %f, noNegative: %t, polar: %t", tt.N, tt.Val, true, true)
		}
	}
}

func TestTemporalDataPointRangeCheck(t *testing.T) {
	for _, tt := range tdpTestSetFull {
		err := temporalDataPointRangeCheck([]data.TemporalDataPoint{tt}, false)
		if err != nil {
			t.Errorf("point incorrectly rejected; T: %s, Val: %f, noNegative: %t", tt.T.String(), tt.Val, false)
		}
		err = temporalDataPointRangeCheck([]data.TemporalDataPoint{tt}, true)
		if tt.Val < 0 {
			if err == nil {
				t.Errorf("point incorrectly accpeted; T: %s, Val: %f, noNegative: %t", tt.T.String(), tt.Val, true)
			}
		} else if err != nil {
			t.Errorf("point incorrectly rejected; T: %s, Val: %f, noNegative: %t", tt.T.String(), tt.Val, true)
		}
	}
}

func TestCategoricalDataPointRangeCheck(t *testing.T) {
	for _, tt := range cdpTestSetFull {
		err := categoricalDataPointRangeCheck([]data.CategoricalDataPoint{tt}, false)
		if err != nil {
			t.Errorf("point incorrectly rejected; C: %s, Val: %f, noNegative: %t", tt.C, tt.Val, false)
		}
		err = categoricalDataPointRangeCheck([]data.CategoricalDataPoint{tt}, true)
		if tt.Val < 0 {
			if err == nil {
				t.Errorf("point incorrectly accpeted; C: %s, Val: %f, noNegative: %t", tt.C, tt.Val, true)
			}
		} else if err != nil {
			t.Errorf("point incorrectly rejected; C: %s, Val: %f, noNegative: %t", tt.C, tt.Val, true)
		}
	}
}
