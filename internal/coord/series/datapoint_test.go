package series

import (
	"math"
	"testing"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/s-daehling/fyne-charts/pkg/data"
)

var ndpTestSetFull = []data.NumericalPoint{
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

var ndpTestSetPosVal = []data.NumericalPoint{
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

var ndpTestSetPosValPolar = []data.NumericalPoint{
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

var tdpTestSetFull = []data.TemporalPoint{
	{T: time.Now(), Val: -1000},
	{T: time.Now().Add(-time.Hour), Val: -0.000001},
	{T: time.Now().Add(time.Hour), Val: 0},
	{T: time.Now().Add(-2 * time.Hour), Val: 0.000001},
	{T: time.Now().Add(2 * time.Hour), Val: 1000},
}

var tdpTestSetPosVal = []data.TemporalPoint{
	{T: time.Now().Add(time.Hour), Val: 0},
	{T: time.Now().Add(-2 * time.Hour), Val: 0.000001},
	{T: time.Now().Add(2 * time.Hour), Val: 1000},
}

var cdpTestSetFull = []data.CategoricalPoint{
	{C: "one", Val: -1000},
	{C: "two", Val: -0.000001},
	{C: "three", Val: 0},
	{C: "four", Val: 0.000001},
	{C: "five", Val: 1000},
}

var cdpTestSetPosVal = []data.CategoricalPoint{
	{C: "one", Val: 0},
	{C: "two", Val: 0.000001},
	{C: "three", Val: 1000},
}

func TestDataPointAddNumericalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.NumericalPoint
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
		{ndpTestSetPosVal, true, true, len(ndpTestSetPosVal), false, -1000, 1000, 0, 1000},
		{ndpTestSetPosValPolar, false, true, len(ndpTestSetPosValPolar), false, 0, 2 * math.Pi, 0, 1000},
		{ndpTestSetPosValPolar, true, true, len(ndpTestSetPosValPolar), false, 0, 2 * math.Pi, 0, 1000},
		{[]data.NumericalPoint{}, false, false, 0, true, 0, 0, 0, 0},
	}
	for i, tt := range tests {
		ser := EmptyPointSeries("test", theme.ColorNameBackground)
		ch := chartDummy{polar: tt.polar}
		ser.BindToChart(ch)
		ser.AddNumericalData(tt.input)
		if len(ser.data) != tt.expNumPoints {
			t.Errorf("wrong number of data, set %d, exp %d, have %d", i, tt.expNumPoints, len(ser.data))
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

func TestDataPointAddTemporalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.TemporalPoint
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
		{[]data.TemporalPoint{}, false, false, 0, true, time.Now(), time.Now(), 0, 0},
	}
	for i, tt := range tests {
		ser := EmptyPointSeries("test", theme.ColorNameBackground)
		ch := chartDummy{polar: tt.polar}
		ser.BindToChart(ch)
		ser.AddTemporalData(tt.input)
		if len(ser.data) != tt.expNumPoints {
			t.Errorf("wrong number of data, set %d, exp %d, have %d", i, tt.expNumPoints, len(ser.data))
		}
		err := testTRange(ser, tt.expIsEmpty, tt.expTMin, tt.expTMax)
		if err != nil {
			t.Errorf("wrong N range, set %d, %s", i, err.Error())
		}
		err = testValRange(ser, tt.expIsEmpty, tt.expValMin, tt.expValMax)
		if err != nil {
			t.Errorf("wrong Val range, set %d, %s", i, err.Error())
		}
	}
}

func TestDataPointAddCategoricalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.CategoricalPoint
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
		{[]data.CategoricalPoint{}, false, false, 0, true, []string{}, 0, 0},
	}
	for i, tt := range tests {
		ser := EmptyPointSeries("test", theme.ColorNameBackground)
		ch := chartDummy{polar: tt.polar}
		ser.BindToChart(ch)
		ser.AddCategoricalData(tt.input)
		if len(ser.data) != tt.expNumPoints {
			t.Errorf("wrong number of data, set %d, exp %d, have %d", i, tt.expNumPoints, len(ser.data))
		}
		err := testCRange(ser, tt.expCRange)
		if err != nil {
			t.Errorf("wrong N range, set %d, %s", i, err.Error())
		}
		err = testValRange(ser, tt.expIsEmpty, tt.expValMin, tt.expValMax)
		if err != nil {
			t.Errorf("wrong Val range, set %d, %s", i, err.Error())
		}
	}
}

func TestDataPointDeleteNumericalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input         []data.NumericalPoint
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
		ser := EmptyPointSeries("test", theme.ColorNameBackground)
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

func TestDataPointDeleteTemporalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input         []data.TemporalPoint
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
		ser := EmptyPointSeries("test", theme.ColorNameBackground)
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

func TestDataPointDeleteCategoricalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input         []data.CategoricalPoint
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
		ser := EmptyPointSeries("test", theme.ColorNameBackground)
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

func TestDataPointNodes(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.NumericalPoint
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
		ser := EmptyPointSeries("test", theme.ColorNameBackground)
		ser.showDot = true
		ser.AddNumericalData(tt.input)
		cns := ser.CartesianDots(tt.xMin, tt.xMax, tt.yMin, tt.yMax)
		if len(cns) != tt.expCartNodes {
			t.Errorf("wrong number of cartesian nodes, set %d, num %d, exp %d", i, len(cns), tt.expCartNodes)
		}
		pns := ser.PolarDots(tt.xMin, tt.xMax, tt.yMin, tt.yMax)
		if len(pns) != tt.expPolNodes {
			t.Errorf("wrong number of polar nodes, set %d, num %d, exp %d", i, len(pns), tt.expPolNodes)
		}
	}
}

func TestDataPointEdges(t *testing.T) {
	app.New()
	var tests = []struct {
		input               []data.NumericalPoint
		xMin                float64
		xMax                float64
		yMin                float64
		yMax                float64
		showFromValBaseLine bool
		showFromPrevLine    bool
		expCartEdges        int
		expPolEdges         int
	}{
		{ndpTestSetFull, -1000, 1000, -1000, 1000, true, true, 69, 69},
		{ndpTestSetFull, 0, 1000, -1000, 1000, true, true, 50, 49},
		{ndpTestSetFull, -1000, 1000, 0, 1000, true, true, 62, 35},
		{ndpTestSetFull, -1000, 0, -1000, 1000, true, true, 30, 29},
		{ndpTestSetFull, -1000, 1000, -1000, 0, true, true, 62, 49},
	}
	for i, tt := range tests {
		app.New()
		ser := EmptyPointSeries("test", theme.ColorNameBackground)
		ser.showFromPrevLine = tt.showFromPrevLine
		ser.showFromValBaseLine = tt.showFromValBaseLine
		if ser.showFromPrevLine {
			ser.sortPoints = true
		}
		ser.AddNumericalData(tt.input)
		ces := ser.CartesianEdges(tt.xMin, tt.xMax, tt.yMin, tt.yMax)
		if len(ces) != tt.expCartEdges {
			t.Errorf("wrong number of cartesian edges, set %d, num %d, exp %d", i, len(ces), tt.expCartEdges)
		}
		pes := ser.PolarEdges(tt.xMin, tt.xMax, tt.yMin, tt.yMax)
		if len(pes) != tt.expPolEdges {
			t.Errorf("wrong number of polar edges, set %d, num %d, exp %d", i, len(pes), tt.expPolEdges)
		}
	}
}

func TestDataPointRects(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.NumericalPoint
		xMin         float64
		xMax         float64
		yMin         float64
		yMax         float64
		expCartRects int
	}{
		{ndpTestSetFull, -1000, 1000, -1000, 1000, len(ndpTestSetFull)},
		{ndpTestSetFull, 0, 1000, -1000, 1000, 25},
		{ndpTestSetFull, -1000, 1000, 0, 1000, 35},
		{ndpTestSetFull, -1000, 0, -1000, 1000, 15},
		{ndpTestSetFull, -1000, 1000, -1000, 0, 35},
	}
	for i, tt := range tests {
		app.New()
		ser := EmptyPointSeries("test", theme.ColorNameBackground)
		ser.showBar = true
		ser.AddNumericalData(tt.input)
		crs := ser.CartesianBars(tt.xMin, tt.xMax, tt.yMin, tt.yMax)
		if len(crs) != tt.expCartRects {
			t.Errorf("wrong number of cartesian rects, set %d, num %d, exp %d", i, len(crs), tt.expCartRects)
		}
	}
}
