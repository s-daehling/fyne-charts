package series

import (
	"testing"
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/s-daehling/fyne-charts/pkg/data"
)

var nCandleStickTestSet = []data.NumericalCandleStick{
	{NStart: -1001, NEnd: -999, Open: -1000, Low: -1000, High: -1000, Close: -1000},
}

var tCandleStickTestSet = []data.TemporalCandleStick{
	{TStart: time.Now(), TEnd: time.Now().Add(time.Minute), Open: -1000, Low: -1000, High: -1000, Close: -1000},
}

func TestCandleStickAddNumericalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.NumericalCandleStick
		expSuccess   bool
		expNumPoints int
		expIsEmpty   bool
		expNMin      float64
		expNMax      float64
		expValMin    float64
		expValMax    float64
	}{
		{nCandleStickTestSet, true, len(nCandleStickTestSet), false, -1001, -999, -1000, -1000},
	}
	for i, tt := range tests {
		ser := EmptyCandleStickSeries(chartDummy{polar: false}, "test")
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

func TestCandleStickAddTemporalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input        []data.TemporalCandleStick
		expSuccess   bool
		expNumPoints int
		expIsEmpty   bool
		expTMin      time.Time
		expTMax      time.Time
		expValMin    float64
		expValMax    float64
	}{
		{tCandleStickTestSet, true, len(tCandleStickTestSet), false, tCandleStickTestSet[0].TStart, tCandleStickTestSet[0].TEnd, -1000, -1000},
	}
	for i, tt := range tests {
		ser := EmptyCandleStickSeries(chartDummy{polar: false}, "test")
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

func TestCandleStickDeleteNumericalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input         []data.NumericalCandleStick
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
		{nCandleStickTestSet, -1000, 0, true, 0, false, -1001, -999, -1000, -1000},
	}
	for i, tt := range tests {
		ser := EmptyCandleStickSeries(chartDummy{polar: false}, "test")
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

func TestCandleStickDeleteTemporalData(t *testing.T) {
	app.New()
	var tests = []struct {
		input         []data.TemporalCandleStick
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
		{tCandleStickTestSet, tCandleStickTestSet[0].TEnd.Add(-time.Second), tCandleStickTestSet[0].TEnd.Add(time.Second), true, 0, false, tCandleStickTestSet[0].TStart, tCandleStickTestSet[0].TEnd, -1000, -1000},
	}
	for i, tt := range tests {
		ser := EmptyCandleStickSeries(chartDummy{polar: false}, "test")
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

func TestCandleStickEdges(t *testing.T) {
	app.New()
	var tests = []struct {
		input    []data.NumericalCandleStick
		xMin     float64
		xMax     float64
		yMin     float64
		yMax     float64
		expEdges int
	}{
		{nCandleStickTestSet, -1001, -999, -1001, -999, 2},
		{nCandleStickTestSet, -1000, 1000, -1000, 1000, 0},
	}
	for i, tt := range tests {
		app.New()
		ser := EmptyCandleStickSeries(chartDummy{polar: false}, "test")
		ser.AddNumericalData(tt.input)
		cns := ser.CartesianEdges(tt.xMin, tt.xMax, tt.yMin, tt.yMax)
		if len(cns) != tt.expEdges {
			t.Errorf("wrong number of edges, set %d, num %d, exp %d", i, len(cns), tt.expEdges)
		}
	}
}

func TestCandleStickRects(t *testing.T) {
	app.New()
	var tests = []struct {
		input    []data.NumericalCandleStick
		xMin     float64
		xMax     float64
		yMin     float64
		yMax     float64
		expRects int
	}{
		{nCandleStickTestSet, -1001, -999, -1001, -999, 1},
		{nCandleStickTestSet, -1000, 1000, -1000, 1000, 0},
	}
	for i, tt := range tests {
		app.New()
		ser := EmptyCandleStickSeries(chartDummy{polar: false}, "test")
		ser.AddNumericalData(tt.input)
		cns := ser.CartesianRects(tt.xMin, tt.xMax, tt.yMin, tt.yMax)
		if len(cns) != tt.expRects {
			t.Errorf("wrong number of rects, set %d, num %d, exp %d", i, len(cns), tt.expRects)
		}
	}
}
