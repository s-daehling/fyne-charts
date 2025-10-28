package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand/v2"
	"time"

	"github.com/s-daehling/fyne-charts/pkg/coord"
	"github.com/s-daehling/fyne-charts/pkg/data"
	"github.com/s-daehling/fyne-charts/pkg/prop"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Chart")

	vS := container.NewVSplit(cartesianCharts(), polarCharts())

	myWindow.SetContent(vS)
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}

func cartesianCharts() (obj fyne.CanvasObject) {
	t := cartTempChart()
	c := cartCatChart()
	n := cartNumChart()
	p := cartPropChart()
	hS1 := container.NewHSplit(c, p)
	hS2 := container.NewHSplit(t, hS1)
	obj = container.NewHSplit(n, hS2)
	return
}

func polarCharts() (obj fyne.CanvasObject) {
	t := polTempChart()
	c := polCatChart()
	a := polAngChart()
	p := polPropChart()
	hS1 := container.NewHSplit(c, p)
	hS2 := container.NewHSplit(t, hS1)
	obj = container.NewHSplit(a, hS2)
	return
}

func updateSineNumericalData() (ndp []data.NumericalDataPoint) {
	periodInMilliSecond := 10000
	shift := float64(time.Now().UnixMilli()%int64(periodInMilliSecond)) / float64(periodInMilliSecond) * 2 * math.Pi
	for range 50 {
		ndp = append(ndp, randomSineNumericalDataPoint(180, 50, shift))
	}
	return
}

func cartNumChart() (numChart *coord.CartesianNumericalChart) {
	numChart = coord.NewCartesianNumericalChart()
	data1 := make([]data.NumericalDataPoint, 0)
	for range 50 {
		data1 = append(data1, randomNumericalDataPoint(-100, 0, -100, 0))
	}
	for range 50 {
		data1 = append(data1, randomNumericalDataPoint(0, 100, 0, 100))
	}
	numChart.AddAreaSeries("area", data1, true, color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff})

	// data2 := make([]data.NumericalDataPoint, 0)
	// for range 50 {
	// 	data2 = append(data2, randomSineNumericalDataPoint(180, 50))
	// }
	ls, err := numChart.AddLineSeries("line", updateSineNumericalData(), true, color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff})
	if err != nil {
		fmt.Println(err)
		return
	}
	ls.SetColor(color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff})

	go func() {
		// data2a := data.NumericalDataPoint{
		// 	N:   150,
		// 	Val: 150,
		// }
		// time.Sleep(time.Second * 5)
		// ls.AddData([]data.NumericalDataPoint{data2a})
		for {
			time.Sleep(time.Millisecond * 250)
			fyne.Do(func() {
				ls.Clear()
				ls.AddData(updateSineNumericalData())
			})
		}
	}()

	data3 := make([]data.NumericalDataPoint, 0)
	for range 50 {
		data3 = append(data3, randomNumericalDataPoint(-110, 110, -110, 110))
	}
	_, err = numChart.AddBarSeries("scatter", data3, 2, color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff})
	if err != nil {
		fmt.Println(err)
	}

	numChart.SetOrigin(0, 0)
	numChart.SetXAxisLabel("X axis")
	numChart.SetYAxisLabel("Y axis")
	numChart.SetTitle("Cartesian Numerical Chart")
	numChart.HideLegend()
	return
}

func cartTempChart() (tempChart *coord.CartesianTemporalChart) {
	tempChart = coord.NewCartesianTemporalChart()
	data1 := make([]data.TemporalCandleStick, 0)
	tStart := time.Now()
	close := 3.14
	span := time.Hour
	for range 50 {
		cs := randomTemporalCandleStick(tStart, span, close, 2.1)
		data1 = append(data1, cs)
		tStart = tStart.Add(span)
		close = cs.Close
	}
	tempChart.AddCandleStickSeries("candlestick", data1)

	data2 := make([]data.TemporalDataPoint, 0)
	for range 10 {
		data2 = append(data2, randomTemporalDataPoint(time.Now(), time.Now().Add(time.Hour*50), -5, 10))
	}
	tempChart.AddLollipopSeries("lollipop", data2, color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff})
	// tls.SetDotSize(6)

	tempChart.SetOrigin(time.Now().Add(time.Hour*20), 4)
	tempChart.SetTAxisLabel("T axis")
	tempChart.SetYAxisLabel("Y axis")
	tempChart.SetTitle("Cartesian Temporal Chart")
	return
}

func cartCatChart() (catChart *coord.CartesianCategoricalChart) {
	catChart = coord.NewCartesianCategoricalChart()
	data1 := []data.CategoricalDataPoint{
		{
			C:   "One",
			Val: rand.Float64() * 30,
		},
		{
			C:   "Two",
			Val: rand.Float64() * 30,
		},
		{
			C:   "Three",
			Val: rand.Float64() * 30,
		},
	}
	data2 := []data.CategoricalDataPoint{
		{
			C:   "One",
			Val: rand.Float64() * 30,
		},
		{
			C:   "Two",
			Val: rand.Float64() * 30,
		},
	}
	catSer := []data.CategoricalDataSeries{
		{
			Name:   "Test1",
			Col:    color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff},
			Points: data1,
		},
		{
			Name:   "Test2",
			Col:    color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff},
			Points: data2,
		},
	}
	catChart.AddStackedBarSeries("stacked bar", catSer)
	data3 := []data.CategoricalDataPoint{
		{
			C:   "One",
			Val: -10 + rand.Float64()*20,
		},
		{
			C:   "Two",
			Val: -10 + rand.Float64()*20,
		},
		{
			C:   "Three",
			Val: -10 + rand.Float64()*20,
		},
	}
	catChart.AddBarSeries("bar", data3, color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff})
	data4 := []data.CategoricalBox{
		{
			C:             "One",
			Maximum:       25.12,
			ThirdQuartile: 23.4,
			Median:        21.5,
			FirstQuartile: 19.3,
			Minimum:       18,
			Outlier:       []float64{26.3, 25.3, 17.45, 12.8},
		},

		{
			C:             "Three",
			Maximum:       20.12,
			ThirdQuartile: 18.4,
			Median:        16.5,
			FirstQuartile: 14.3,
			Minimum:       13,
			Outlier:       []float64{21.3, 20.3, 12.45, 7.8},
		},
	}
	_, err := catChart.AddBoxSeries("box", data4, color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff})
	if err != nil {
		fmt.Print(err)
	}
	// catChart.SetYRange(-5, 35)
	// chart.SetCRange([]string{"One", "Three"})
	catChart.SetCAxisLabel("C axis")
	catChart.SetYAxisLabel("Y axis")
	catChart.SetTitle("Cartesian Categorical Chart")
	return
}

func cartPropChart() (propChart *prop.BarChart) {
	propChart = prop.NewBarChart()
	data1 := []data.ProportionalDataPoint{
		{
			Val: rand.Float64() * 222,
			C:   "One",
			Col: color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff},
		},
		{
			Val: rand.Float64() * 222,
			C:   "Two",
			Col: color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff},
		},
	}
	propChart.AddSeries("proportion", data1)
	data2 := []data.ProportionalDataPoint{
		{
			Val: rand.Float64() * 222,
			C:   "One",
			Col: color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff},
		},
		{
			Val: rand.Float64() * 222,
			C:   "Two",
			Col: color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff},
		},
		{
			Val: rand.Float64() * 222,
			C:   "Three",
			Col: color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff},
		},
	}
	propChart.AddSeries("proportion 2", data2)
	propChart.SetTitle("Proportional Bar Chart")
	return
}

func polAngChart() (angChart *coord.PolarNumericalChart) {
	angChart = coord.NewPolarNumericalChart()
	data1 := make([]data.NumericalDataPoint, 0)
	for range 100 {
		data1 = append(data1, randomAngularDataPoint(63.777))
	}
	angChart.AddAreaSeries("area", data1, true, color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff})

	data2 := make([]data.NumericalDataPoint, 0)
	for range 150 {
		data2 = append(data2, randomSineAngularDataPoint(63.777))
	}
	ls, _ := angChart.AddLineSeries("line", data2, false, color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff})
	ls.SetLineWidth(2)

	angChart.SetOrigin(0, 64)
	angChart.SetPhiAxisLabel("Phi axis")
	angChart.SetRAxisLabel("R axis")
	angChart.SetTitle("Polar Numerical Chart")
	return
}

func polTempChart() (tempChart *coord.PolarTemporalChart) {
	tempChart = coord.NewPolarTemporalChart()
	data1 := make([]data.TemporalDataPoint, 0)
	for range 50 {
		data1 = append(data1, randomTemporalDataPoint(time.Now(), time.Now().Add(time.Hour*50), 0, 111))
	}
	tempChart.AddLollipopSeries("lollipop", data1, color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff})

	data2 := make([]data.TemporalDataPoint, 0)
	for range 25 {
		data2 = append(data2, randomTemporalDataPoint(time.Now(), time.Now().Add(time.Hour*50), 0, 111))
	}
	tempChart.AddScatterSeries("scatter", data2, color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff})

	tempChart.SetOrigin(time.Now().Add(time.Hour*10), 120)
	tempChart.SetTAxisLabel("T axis")
	tempChart.SetRAxisLabel("R axis")
	tempChart.SetTitle("Polar Temporal Chart")
	return
}

func polCatChart() (catChart *coord.PolarCategoricalChart) {
	catChart = coord.NewPolarCategoricalChart()
	data1 := []data.CategoricalDataPoint{
		{
			C:   "One",
			Val: rand.Float64() * 30,
		},
		{
			C:   "Two",
			Val: rand.Float64() * 30,
		},
		{
			C:   "Three",
			Val: rand.Float64() * 30,
		},
	}

	data2 := []data.CategoricalDataPoint{
		{
			C:   "One",
			Val: rand.Float64() * 30,
		},
		{
			C:   "Two",
			Val: rand.Float64() * 30,
		},
	}
	catSer := []data.CategoricalDataSeries{
		{
			Name:   "Test1",
			Col:    color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff},
			Points: data1,
		},
		{
			Name:   "Test2",
			Col:    color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff},
			Points: data2,
		},
	}
	catChart.AddStackedBarSeries("stacked bar", catSer)
	data3 := []data.CategoricalDataPoint{
		{
			C:   "One",
			Val: rand.Float64() * 30,
		},
		{
			C:   "Two",
			Val: rand.Float64() * 30,
		},
		{
			C:   "Three",
			Val: rand.Float64() * 30,
		},
	}
	catChart.AddBarSeries("bar", data3, color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff})

	catChart.SetRRange(40)
	// chart.SetCRange([]string{"One", "Three"})
	catChart.SetCAxisLabel("C axis")
	catChart.SetRAxisLabel("R axis")
	catChart.SetTitle("Polar Categorical Chart")
	return
}

func polPropChart() (propChart *prop.PieChart) {
	propChart = prop.NewPieChart()
	data1 := []data.ProportionalDataPoint{
		{
			Val: rand.Float64() * 222,
			C:   "One",
			Col: color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff},
		},
		{
			Val: rand.Float64() * 222,
			C:   "Two",
			Col: color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff},
		},
	}
	propChart.AddSeries("proportion", data1)
	data2 := []data.ProportionalDataPoint{
		{
			Val: rand.Float64() * 222,
			C:   "One",
			Col: color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff},
		},
		{
			Val: rand.Float64() * 222,
			C:   "Two",
			Col: color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff},
		},
		{
			Val: rand.Float64() * 222,
			C:   "Three",
			Col: color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff},
		},
	}
	ser, _ := propChart.AddSeries("proportion2", data2)
	ser.SetValTextColor(color.Black)
	propChart.SetTitle("Proportional Pie/Doughnut Chart")
	return
}

func randomNumericalDataPoint(xMin float64, xMax float64, valMin float64,
	valMax float64) (ndp data.NumericalDataPoint) {
	ndp.N = xMin + (rand.Float64() * (xMax - xMin))
	ndp.Val = valMin + (rand.Float64() * (valMax - valMin))
	return
}

func randomSineNumericalDataPoint(l float64, amp float64, shift float64) (ndp data.NumericalDataPoint) {
	ndp.N = (-l / 2) + (rand.Float64() * l)
	ndp.Val = amp * math.Sin((ndp.N/(l))*2*math.Pi-shift)
	return
}

func randomTemporalDataPoint(tMin time.Time, tMax time.Time, valMin float64,
	valMax float64) (tdp data.TemporalDataPoint) {
	tdp.T = time.Unix(rand.Int64N(tMax.Unix()-tMin.Unix())+tMin.Unix(), 0)
	tdp.Val = valMin + (rand.Float64() * (valMax - valMin))
	return
}

func randomTemporalCandleStick(tStart time.Time, span time.Duration,
	prevClose float64, vol float64) (tcs data.TemporalCandleStick) {
	tcs.TStart = tStart
	tcs.TEnd = tStart.Add(span)
	tcs.Open = prevClose + ((-0.5 + rand.Float64()) * vol * 0.1)
	tcs.High = tcs.Open + (rand.Float64() * vol)
	tcs.Low = tcs.Open - (rand.Float64() * vol)
	tcs.Close = tcs.Low + (rand.Float64() * (tcs.High - tcs.Low))

	return
}

func randomAngularDataPoint(valMax float64) (adp data.NumericalDataPoint) {
	adp.N = rand.Float64() * 2 * math.Pi
	adp.Val = rand.Float64() * valMax
	return
}

func randomSineAngularDataPoint(amp float64) (adp data.NumericalDataPoint) {
	adp.N = rand.Float64() * 2 * math.Pi
	adp.Val = amp + (amp * math.Sin(adp.N))
	return
}
