package series

import (
	"image/color"
	"testing"
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/s-daehling/fyne-charts/pkg/data"
)

var nBoxTestSet = []data.NumericalBox{
	{N: -1000, Maximum: -999, ThirdQuartile: -999.5, Median: -1000, FirstQuartile: -1000.5, Minimum: -1001, Outlier: []float64{}},
	{N: 0, Maximum: 1, ThirdQuartile: 0.5, Median: 0, FirstQuartile: -0.5, Minimum: -1, Outlier: []float64{-2, 2}},
	{N: 0, Maximum: 1, ThirdQuartile: 0.5, Median: 0, FirstQuartile: -0.5, Minimum: -1, Outlier: []float64{-1, 1}},
	{N: 0, Maximum: 1, ThirdQuartile: 0.5, Median: 0, FirstQuartile: -0.5, Minimum: -1, Outlier: []float64{-0.5, 0.5}},
	{N: 0.000001, Maximum: 1, ThirdQuartile: 0.5, Median: 0, FirstQuartile: -0.5, Minimum: -1, Outlier: []float64{}},
	{N: 1000, Maximum: 1001, ThirdQuartile: 1000.5, Median: 1000, FirstQuartile: 999.5, Minimum: 999, Outlier: []float64{1005}},
}

var tBoxTestSet = []data.TemporalBox{
	{T: time.Now(), Maximum: -999, ThirdQuartile: -999.5, Median: -1000, FirstQuartile: -1000.5, Minimum: -1001, Outlier: []float64{}},
	{T: time.Now().Add(time.Second), Maximum: 1, ThirdQuartile: 0.5, Median: 0, FirstQuartile: -0.5, Minimum: -1, Outlier: []float64{-2, 2}},
	{T: time.Now().Add(time.Second), Maximum: 1, ThirdQuartile: 0.5, Median: 0, FirstQuartile: -0.5, Minimum: -1, Outlier: []float64{-1, 1}},
	{T: time.Now().Add(time.Second), Maximum: 1, ThirdQuartile: 0.5, Median: 0, FirstQuartile: -0.5, Minimum: -1, Outlier: []float64{-0.5, 0.5}},
	{T: time.Now().Add(5 * time.Second), Maximum: 1001, ThirdQuartile: 1000.5, Median: 1000, FirstQuartile: 999.5, Minimum: 999, Outlier: []float64{1005}},
}

var cBoxTestSet = []data.CategoricalBox{
	{C: "one", Maximum: -999, ThirdQuartile: -999.5, Median: -1000, FirstQuartile: -1000.5, Minimum: -1001, Outlier: []float64{}},
	{C: "two", Maximum: 1, ThirdQuartile: 0.5, Median: 0, FirstQuartile: -0.5, Minimum: -1, Outlier: []float64{-2, 2}},
	{C: "three", Maximum: 1, ThirdQuartile: 0.5, Median: 0, FirstQuartile: -0.5, Minimum: -1, Outlier: []float64{-1, 1}},
	{C: "four", Maximum: 1, ThirdQuartile: 0.5, Median: 0, FirstQuartile: -0.5, Minimum: -1, Outlier: []float64{-0.5, 0.5}},
	{C: "five", Maximum: 1001, ThirdQuartile: 1000.5, Median: 1000, FirstQuartile: 999.5, Minimum: 999, Outlier: []float64{1005}},
}

func TestBoxAddNumericalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.NumericalBox
		expSuccess   bool
		expNumPoints int
		expIsEmpty   bool
		expNMin      float64
		expNMax      float64
		expValMin    float64
		expValMax    float64
	}{
		{nBoxTestSet, true, len(nBoxTestSet), false, -1000, 1000, -1001, 1005},
	}
	for i, tt := range tests {
		ser := EmptyBoxSeries(chartDummy{polar: false}, "test", color.Black)
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

func TestBoxAddTemporalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.TemporalBox
		expSuccess   bool
		expNumPoints int
		expIsEmpty   bool
		expTMin      time.Time
		expTMax      time.Time
		expValMin    float64
		expValMax    float64
	}{
		{tBoxTestSet, true, len(tBoxTestSet), false, tBoxTestSet[0].T, tBoxTestSet[4].T, -1001, 1005},
	}
	for i, tt := range tests {
		ser := EmptyBoxSeries(chartDummy{polar: false}, "test", color.Black)
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
			t.Errorf("wrong T range, set %d, %s", i, err.Error())
		}
		err = testValRange(ser, tt.expIsEmpty, tt.expValMin, tt.expValMax)
		if err != nil {
			t.Errorf("wrong Val range, set %d, %s", i, err.Error())
		}
	}
}

func TestBoxAddCategoricalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.CategoricalBox
		expSuccess   bool
		expNumPoints int
		expIsEmpty   bool
		expCRange    []string
		expValMin    float64
		expValMax    float64
	}{
		{cBoxTestSet, true, len(cBoxTestSet), false, []string{"one", "two", "three", "four", "five"}, -1001, 1005},
	}
	for i, tt := range tests {
		ser := EmptyBoxSeries(chartDummy{polar: false}, "test", color.Black)
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
			t.Errorf("wrong C range, set %d, %s", i, err.Error())
		}
		err = testValRange(ser, tt.expIsEmpty, tt.expValMin, tt.expValMax)
		if err != nil {
			t.Errorf("wrong Val range, set %d, %s", i, err.Error())
		}
	}
}

func TestBoxDeleteNumericalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input         []data.NumericalBox
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
		{nBoxTestSet, -1000.000001, 0, true, 1, false, 0, 1000, -2, 1005},
	}
	for i, tt := range tests {
		ser := EmptyBoxSeries(chartDummy{polar: false}, "test", color.Black)
		ser.AddNumericalData(tt.input)
		c := ser.DeleteNumericalDataInRange(tt.delMin, tt.delMax)
		if c != tt.expNumDeleted {
			t.Errorf("wrong number of data deleted, set %d, exp %d, have %d", i, tt.expNumDeleted, c)
		}
		err := testNRange(ser, tt.expIsEmpty, tt.expNMin, tt.expNMax)
		if err != nil {
			t.Errorf("wrong N range, set %d, %s", i, err.Error())
		}
		err = testValRange(ser, tt.expIsEmpty, tt.expValMin, tt.expValMax)
		if err != nil {
			t.Errorf("wrong Val range, set %d, %s", i, err.Error())
		}
	}
}

func TestBoxDeleteTemporalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input         []data.TemporalBox
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
		{tBoxTestSet, tBoxTestSet[0].T.Add(-time.Second), tBoxTestSet[0].T.Add(time.Second), true, 1, false, tBoxTestSet[1].T, tBoxTestSet[4].T, -2, 1005},
	}
	for i, tt := range tests {
		ser := EmptyBoxSeries(chartDummy{polar: false}, "test", color.Black)
		ser.AddTemporalData(tt.input)
		c := ser.DeleteTemporalDataInRange(tt.delMin, tt.delMax)
		if c != tt.expNumDeleted {
			t.Errorf("wrong number of data deleted, set %d, exp %d, have %d", i, tt.expNumDeleted, c)
		}
		err := testTRange(ser, tt.expIsEmpty, tt.expTMin, tt.expTMax)
		if err != nil {
			t.Errorf("wrong T range, set %d, %s", i, err.Error())
		}
		err = testValRange(ser, tt.expIsEmpty, tt.expValMin, tt.expValMax)
		if err != nil {
			t.Errorf("wrong Val range, set %d, %s", i, err.Error())
		}
	}
}

func TestBoxDeleteCategoricalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input         []data.CategoricalBox
		del           []string
		expSuccess    bool
		expNumDeleted int
		expIsEmpty    bool
		expCRange     []string
		expValMin     float64
		expValMax     float64
	}{
		{cBoxTestSet, []string{"one"}, true, 1, false, []string{"two", "three", "four", "five"}, -2, 1005},
	}
	for i, tt := range tests {
		ser := EmptyBoxSeries(chartDummy{polar: false}, "test", color.Black)
		ser.AddCategoricalData(tt.input)
		c := ser.DeleteCategoricalDataInRange(tt.del)
		if c != tt.expNumDeleted {
			t.Errorf("wrong number of data deleted, set %d, exp %d, have %d", i, tt.expNumDeleted, c)
		}
		err := testCRange(ser, tt.expCRange)
		if err != nil {
			t.Errorf("wrong C range, set %d, %s", i, err.Error())
		}
		err = testValRange(ser, tt.expIsEmpty, tt.expValMin, tt.expValMax)
		if err != nil {
			t.Errorf("wrong Val range, set %d, %s", i, err.Error())
		}
	}
}

func TestBoxNodes(t *testing.T) {
	app.New()
	var tests = []struct {
		input    []data.NumericalBox
		xMin     float64
		xMax     float64
		yMin     float64
		yMax     float64
		expEdges int
	}{
		{nBoxTestSet, -1001, -999, -1001, -999, 0},
		{nBoxTestSet, -1000, 1000, -1000, 1000, 6},
		{nBoxTestSet, -1000, 1000, -1000, 1005, 7},
	}
	for i, tt := range tests {
		app.New()
		ser := EmptyBoxSeries(chartDummy{polar: false}, "test", color.Black)
		ser.AddNumericalData(tt.input)
		cns := ser.CartesianNodes(tt.xMin, tt.xMax, tt.yMin, tt.yMax)
		if len(cns) != tt.expEdges {
			t.Errorf("wrong number of nodes, set %d, num %d, exp %d", i, len(cns), tt.expEdges)
		}
	}
}

func TestBoxEdges(t *testing.T) {
	app.New()
	var tests = []struct {
		input    []data.NumericalBox
		xMin     float64
		xMax     float64
		yMin     float64
		yMax     float64
		expEdges int
	}{
		{nBoxTestSet, -1001, -999, -1001, -999, 5},
		{nBoxTestSet, -1000, 1000, -1000, 1000, 20},
		{nBoxTestSet, -1000, 1000, -1000, 1005, 25},
	}
	for i, tt := range tests {
		app.New()
		ser := EmptyBoxSeries(chartDummy{polar: false}, "test", color.Black)
		ser.AddNumericalData(tt.input)
		cns := ser.CartesianEdges(tt.xMin, tt.xMax, tt.yMin, tt.yMax)
		if len(cns) != tt.expEdges {
			t.Errorf("wrong number of edges, set %d, num %d, exp %d", i, len(cns), tt.expEdges)
		}
	}
}

func TestBoxRects(t *testing.T) {
	app.New()
	var tests = []struct {
		input    []data.NumericalBox
		xMin     float64
		xMax     float64
		yMin     float64
		yMax     float64
		expRects int
	}{
		{nBoxTestSet, -1001, -999, -1001, -999, 1},
		{nBoxTestSet, -1000, 1000, -1000, 1000, 4},
		{nBoxTestSet, -1000, 1000, -1000, 1005, 5},
	}
	for i, tt := range tests {
		app.New()
		ser := EmptyBoxSeries(chartDummy{polar: false}, "test", color.Black)
		ser.AddNumericalData(tt.input)
		cns := ser.CartesianRects(tt.xMin, tt.xMax, tt.yMin, tt.yMax)
		if len(cns) != tt.expRects {
			t.Errorf("wrong number of rects, set %d, num %d, exp %d", i, len(cns), tt.expRects)
		}
	}
}
