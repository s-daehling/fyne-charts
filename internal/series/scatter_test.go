package series

import (
	"image/color"
	"math"
	"testing"
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/s-daehling/fyne-charts/pkg/data"
)

func TestAddNumericalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.NumericalDataPoint
		polar        bool
		expSuccess   bool
		expNumPoints int
		expIsEmpty   bool
		expNMin      float64
		expNMax      float64
		expValMin    float64
		expValMax    float64
	}{
		{ndpTestSetFull, false, true, len(ndpTestSetFull), false, -1000, 1000, -1000, 1000},
		{ndpTestSetFull, true, false, 0, true, 0, 0, 0, 0},
		{ndpTestSetPosVal, false, true, len(ndpTestSetPosVal), false, -1000, 1000, 0, 1000},
		{ndpTestSetPosVal, true, false, 0, true, 0, 0, 0, 0},
		{ndpTestSetPosValPolar, false, true, len(ndpTestSetPosValPolar), false, 0, 2 * math.Pi, 0, 1000},
		{ndpTestSetPosValPolar, true, true, len(ndpTestSetPosValPolar), false, 0, 2 * math.Pi, 0, 1000},
		{[]data.NumericalDataPoint{}, false, false, 0, true, 0, 0, 0, 0},
	}
	for i, tt := range tests {
		ser := EmptyScatterSeries(chartDummy{}, "test", color.Black, tt.polar)
		err := ser.AddNumericalData(tt.input)
		if err != nil && tt.expSuccess {
			t.Errorf("adding data failed incorrectly, set %d, %s", i, err.Error())
		} else if err == nil && !tt.expSuccess {
			t.Errorf("adding data succeeded incorrectly, set %d", i)
		}
		if len(ser.data) != tt.expNumPoints {
			t.Errorf("wrong number of data, set %d, exp %d, have %d", i, tt.expNumPoints, len(ser.data))
		}
		err = testNRange(ser, tt.expIsEmpty, tt.expNMin, tt.expNMax)
		if err != nil {
			t.Errorf("wrong N range, set %d, %s", i, err.Error())
		}
		err = testValRange(ser, tt.expIsEmpty, tt.expValMin, tt.expValMax)
		if err != nil {
			t.Errorf("wrong Val range, set %d, %s", i, err.Error())
		}
	}
}

func TestAddTemporalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.TemporalDataPoint
		polar        bool
		expSuccess   bool
		expNumPoints int
		expIsEmpty   bool
		expTMin      time.Time
		expTMax      time.Time
		expValMin    float64
		expValMax    float64
	}{
		{tdpTestSetFull, false, true, len(tdpTestSetFull), false, tdpTestSetFull[3].T, tdpTestSetFull[4].T, -1000, 1000},
		{tdpTestSetFull, true, false, 0, true, time.Now(), time.Now(), 0, 0},
		{tdpTestSetPosVal, false, true, len(tdpTestSetPosVal), false, tdpTestSetPosVal[1].T, tdpTestSetPosVal[2].T, 0, 1000},
		{tdpTestSetPosVal, true, true, len(tdpTestSetPosVal), false, tdpTestSetPosVal[1].T, tdpTestSetPosVal[2].T, 0, 1000},
		{[]data.TemporalDataPoint{}, false, false, 0, true, time.Now(), time.Now(), 0, 0},
	}
	for i, tt := range tests {
		ser := EmptyScatterSeries(chartDummy{}, "test", color.Black, tt.polar)
		err := ser.AddTemporalData(tt.input)
		if err != nil && tt.expSuccess {
			t.Errorf("adding data failed incorrectly, set %d, %s", i, err.Error())
		} else if err == nil && !tt.expSuccess {
			t.Errorf("adding data succeeded incorrectly, set %d", i)
		}
		if len(ser.data) != tt.expNumPoints {
			t.Errorf("wrong number of data, set %d, exp %d, have %d", i, tt.expNumPoints, len(ser.data))
		}
		err = testTRange(ser, tt.expIsEmpty, tt.expTMin, tt.expTMax)
		if err != nil {
			t.Errorf("wrong N range, set %d, %s", i, err.Error())
		}
		err = testValRange(ser, tt.expIsEmpty, tt.expValMin, tt.expValMax)
		if err != nil {
			t.Errorf("wrong Val range, set %d, %s", i, err.Error())
		}
	}
}

func TestAddCategoricalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.CategoricalDataPoint
		polar        bool
		expSuccess   bool
		expNumPoints int
		expIsEmpty   bool
		expCRange    []string
		expValMin    float64
		expValMax    float64
	}{
		{cdpTestSetFull, false, true, len(cdpTestSetFull), false, []string{"one", "two", "three", "four", "five"}, -1000, 1000},
		{cdpTestSetFull, true, false, 0, true, []string{}, 0, 0},
		{cdpTestSetPosVal, false, true, len(cdpTestSetPosVal), false, []string{"one", "two", "three"}, 0, 1000},
		{cdpTestSetPosVal, true, true, len(cdpTestSetPosVal), false, []string{"one", "two", "three"}, 0, 1000},
		{[]data.CategoricalDataPoint{}, false, false, 0, true, []string{}, 0, 0},
	}
	for i, tt := range tests {
		ser := EmptyScatterSeries(chartDummy{}, "test", color.Black, tt.polar)
		err := ser.AddCategoricalData(tt.input)
		if err != nil && tt.expSuccess {
			t.Errorf("adding data failed incorrectly, set %d, %s", i, err.Error())
		} else if err == nil && !tt.expSuccess {
			t.Errorf("adding data succeeded incorrectly, set %d", i)
		}
		if len(ser.data) != tt.expNumPoints {
			t.Errorf("wrong number of data, set %d, exp %d, have %d", i, tt.expNumPoints, len(ser.data))
		}
		err = testCRange(ser, tt.expCRange)
		if err != nil {
			t.Errorf("wrong N range, set %d, %s", i, err.Error())
		}
		err = testValRange(ser, tt.expIsEmpty, tt.expValMin, tt.expValMax)
		if err != nil {
			t.Errorf("wrong Val range, set %d, %s", i, err.Error())
		}
	}
}

func TestDeleteNumericalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input         []data.NumericalDataPoint
		delMin        float64
		delMax        float64
		expSuccess    bool
		expNumDeleted int
		expIsEmpty    bool
		expNMin       float64
		expNMax       float64
		expValMin     float64
		expValMax     float64
	}{
		{ndpTestSetFull, -1000.000001, 0, true, 10, false, 0, 1000, -1000, 1000},
		{ndpTestSetFull, -1000, 0, true, 5, false, -1000, 1000, -1000, 1000},
		{ndpTestSetFull, 0, 0, true, 0, false, -1000, 1000, -1000, 1000},
		{ndpTestSetFull, -1000.000001, 1000.000001, true, len(ndpTestSetFull), true, 0, 0, 0, 0},
		{ndpTestSetFull, 1000.000001, -1000.000001, false, 0, false, -1000, 1000, -1000, 1000},
	}
	for i, tt := range tests {
		ser := EmptyScatterSeries(chartDummy{}, "test", color.Black, false)
		ser.AddNumericalData(tt.input)
		c, err := ser.DeleteNumericalDataInRange(tt.delMin, tt.delMax)
		if err != nil && tt.expSuccess {
			t.Errorf("deleteing data failed incorrectly, set %d, %s", i, err.Error())
		} else if err == nil && !tt.expSuccess {
			t.Errorf("deleteing data succeeded incorrectly, set %d", i)
		}
		if c != tt.expNumDeleted {
			t.Errorf("wrong number of data deleted, set %d, exp %d, have %d", i, tt.expNumDeleted, c)
		}
		err = testNRange(ser, tt.expIsEmpty, tt.expNMin, tt.expNMax)
		if err != nil {
			t.Errorf("wrong N range, set %d, %s", i, err.Error())
		}
		err = testValRange(ser, tt.expIsEmpty, tt.expValMin, tt.expValMax)
		if err != nil {
			t.Errorf("wrong Val range, set %d, %s", i, err.Error())
		}
	}
}

func TestDeleteTemporalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input         []data.TemporalDataPoint
		delMin        time.Time
		delMax        time.Time
		expSuccess    bool
		expNumDeleted int
		expIsEmpty    bool
		expTMin       time.Time
		expTMax       time.Time
		expValMin     float64
		expValMax     float64
	}{
		{tdpTestSetFull, tdpTestSetFull[3].T.Add(-time.Second), tdpTestSetFull[0].T, true, 2, false, tdpTestSetFull[0].T, tdpTestSetFull[4].T, -1000, 1000},
		{tdpTestSetFull, tdpTestSetFull[3].T, tdpTestSetFull[0].T, true, 1, false, tdpTestSetFull[3].T, tdpTestSetFull[4].T, -1000, 1000},
		{tdpTestSetFull, tdpTestSetFull[0].T, tdpTestSetFull[0].T, true, 0, false, tdpTestSetFull[3].T, tdpTestSetFull[4].T, -1000, 1000},
		{tdpTestSetFull, tdpTestSetFull[3].T.Add(-time.Second), tdpTestSetFull[4].T.Add(time.Second), true, 5, true, time.Now(), time.Now(), 0, 0},
		{tdpTestSetFull, tdpTestSetFull[4].T.Add(-time.Second), tdpTestSetFull[3].T.Add(time.Second), false, 0, false, tdpTestSetFull[3].T, tdpTestSetFull[4].T, -1000, 1000},
	}
	for i, tt := range tests {
		ser := EmptyScatterSeries(chartDummy{}, "test", color.Black, false)
		ser.AddTemporalData(tt.input)
		c, err := ser.DeleteTemporalDataInRange(tt.delMin, tt.delMax)
		if err != nil && tt.expSuccess {
			t.Errorf("deleteing data failed incorrectly, set %d, %s", i, err.Error())
		} else if err == nil && !tt.expSuccess {
			t.Errorf("deleteing data succeeded incorrectly, set %d", i)
		}
		if c != tt.expNumDeleted {
			t.Errorf("wrong number of data deleted, set %d, exp %d, have %d", i, tt.expNumDeleted, c)
		}
		err = testTRange(ser, tt.expIsEmpty, tt.expTMin, tt.expTMax)
		if err != nil {
			t.Errorf("wrong T range, set %d, %s", i, err.Error())
		}
		err = testValRange(ser, tt.expIsEmpty, tt.expValMin, tt.expValMax)
		if err != nil {
			t.Errorf("wrong Val range, set %d, %s", i, err.Error())
		}
	}
}

func TestDeleteCategoricalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input         []data.CategoricalDataPoint
		del           []string
		expSuccess    bool
		expNumDeleted int
		expIsEmpty    bool
		expCRange     []string
		expValMin     float64
		expValMax     float64
	}{
		{cdpTestSetFull, []string{"one", "two"}, true, 2, false, []string{"three", "four", "five"}, 0, 1000},
		{cdpTestSetFull, []string{}, false, 0, false, []string{"one", "two", "three", "four", "five"}, -1000, 1000},
		{cdpTestSetFull, []string{"six"}, true, 0, false, []string{"one", "two", "three", "four", "five"}, -1000, 1000},
		{cdpTestSetFull, []string{"one", "two", "three", "four", "five"}, true, 5, true, []string{}, 0, 0},
	}
	for i, tt := range tests {
		ser := EmptyScatterSeries(chartDummy{}, "test", color.Black, false)
		ser.AddCategoricalData(tt.input)
		c, err := ser.DeleteCategoricalDataInRange(tt.del)
		if err != nil && tt.expSuccess {
			t.Errorf("deleteing data failed incorrectly, set %d, %s", i, err.Error())
		} else if err == nil && !tt.expSuccess {
			t.Errorf("deleteing data succeeded incorrectly, set %d", i)
		}
		if c != tt.expNumDeleted {
			t.Errorf("wrong number of data deleted, set %d, exp %d, have %d", i, tt.expNumDeleted, c)
		}
		err = testCRange(ser, tt.expCRange)
		if err != nil {
			t.Errorf("wrong C range, set %d, %s", i, err.Error())
		}
		err = testValRange(ser, tt.expIsEmpty, tt.expValMin, tt.expValMax)
		if err != nil {
			t.Errorf("wrong Val range, set %d, %s", i, err.Error())
		}
	}
}

func TestNodes(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.NumericalDataPoint
		xMin         float64
		xMax         float64
		yMin         float64
		yMax         float64
		expCartNodes int
		expPolNodes  int
	}{
		{ndpTestSetFull, -1000, 1000, -1000, 1000, len(ndpTestSetFull), len(ndpTestSetFull)},
		{ndpTestSetFull, 0, 1000, -1000, 1000, 25, 25},
		{ndpTestSetFull, -1000, 1000, 0, 1000, 21, 21},
		{ndpTestSetFull, -1000, 0, -1000, 1000, 15, 15},
		{ndpTestSetFull, -1000, 1000, -1000, 0, 21, 21},
	}
	for i, tt := range tests {
		app.New()
		ser := EmptyScatterSeries(chartDummy{}, "test", color.Black, false)
		ser.AddNumericalData(tt.input)
		cns := ser.CartesianNodes(tt.xMin, tt.xMax, tt.yMin, tt.yMax)
		if len(cns) != tt.expCartNodes {
			t.Errorf("wrong number of cartesian nodes, set %d, num %d, exp %d", i, len(cns), tt.expCartNodes)
		}
		pns := ser.PolarNodes(tt.xMin, tt.xMax, tt.yMin, tt.yMax)
		if len(pns) != tt.expCartNodes {
			t.Errorf("wrong number of polar nodes, set %d, num %d, exp %d", i, len(pns), tt.expCartNodes)
		}
	}
}
