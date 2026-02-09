package main

import (
	"math"
	"math/rand/v2"
	"time"

	"github.com/s-daehling/fyne-charts/pkg/coord"
	"github.com/s-daehling/fyne-charts/pkg/data"
	"github.com/s-daehling/fyne-charts/pkg/style"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Coordinate System Charts")
	myApp.Settings().SetTheme(style.NewColorPaletteTheme(theme.DefaultTheme()))

	var cartCharts, polCharts fyne.CanvasObject
	var err error
	cartCharts, err = cartesianCharts()
	if err != nil {
		panic(err)
	}
	polCharts, err = polarCharts()
	if err != nil {
		panic(err)
	}

	vS := container.NewVSplit(cartCharts, polCharts)

	myWindow.SetContent(vS)
	myWindow.Resize(fyne.NewSize(600, 600))
	myWindow.ShowAndRun()
}

func cartesianCharts() (obj fyne.CanvasObject, err error) {
	t, err := cartTempChart()
	if err != nil {
		return
	}
	c, err := cartCatChart()
	if err != nil {
		return
	}
	n, err := cartNumChart()
	if err != nil {
		return
	}
	hS := container.NewHSplit(t, c)
	obj = container.NewHSplit(n, hS)
	return
}

func polarCharts() (obj fyne.CanvasObject, err error) {
	t, err := polTempChart()
	if err != nil {
		return
	}
	c, err := polCatChart()
	if err != nil {
		return
	}
	a, err := polNumChart()
	if err != nil {
		return
	}
	hS := container.NewHSplit(t, c)
	obj = container.NewHSplit(a, hS)
	return
}

// Cartesian Numerical Chart
func cartNumChart() (numChart *coord.CartesianNumericalChart, err error) {
	numChart = coord.NewCartesianNumericalChart("Cartesian Numerical Chart")
	colPal := style.NewPaletteTriadic(theme.ColorNamePrimary)

	// Area Series
	data1 := make([]data.NumericalPoint, 0)
	for range 50 {
		data1 = append(data1, randomNumericalDataPoint(-100, 0, -100, 0))
	}
	for range 50 {
		data1 = append(data1, randomNumericalDataPoint(0, 100, 0, 100))
	}
	as, err := coord.NewNumericalPointSeries("area", colPal.Next(), data1)
	if err != nil {
		return
	}
	err = numChart.AddAreaSeries(as, true)
	if err != nil {
		return
	}

	// Line Series
	ls, err := coord.NewNumericalPointSeries("line", colPal.Next(), updateSineNumericalData())
	if err != nil {
		return
	}
	err = numChart.AddLineSeries(ls, true)
	if err != nil {
		return
	}
	// Change color after series creation
	ls.SetColor(theme.ColorNameWarning)
	// Update Line Series in a goroutine
	go func() {
		for {
			time.Sleep(time.Millisecond * 250)
			fyne.Do(func() {
				ls.Clear()
				ls.AddData(updateSineNumericalData())
			})
		}
	}()

	// Bar Series
	data3 := make([]data.NumericalPoint, 0)
	for range 50 {
		data3 = append(data3, randomNumericalDataPoint(-110, 110, -110, 110))
	}
	bs, err := coord.NewNumericalPointSeries("bar", colPal.Next(), data3)
	if err != nil {
		return
	}
	err = numChart.AddBarSeries(bs, 2)
	if err != nil {
		return
	}

	// Examples of methods for altering the chart appearance
	numChart.SetOrigin(0, 0)
	numChart.SetXAxisLabel("X axis")
	numChart.SetYAxisLabel("Y axis")
	numChart.HideLegend()
	return
}

// Cartesian Temporal Chart
func cartTempChart() (tempChart *coord.CartesianTemporalChart, err error) {
	tempChart = coord.NewCartesianTemporalChart("")

	// Candlestick Series
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
	tcs, err := coord.NewTemporalCandleStickSeries("candlestick", data1)
	if err != nil {
		return
	}
	err = tempChart.AddCandleStickSeries(tcs)
	if err != nil {
		return
	}

	// Lollipop Series
	data2 := make([]data.TemporalPoint, 0)
	for range 10 {
		data2 = append(data2, randomTemporalDataPoint(time.Now(), time.Now().Add(time.Hour*50), -5, 10))
	}
	tps, err := coord.NewTemporalPointSeries("lollipop", theme.ColorNamePrimary, data2)
	if err != nil {
		return
	}
	err = tempChart.AddLollipopSeries(tps)
	if err != nil {
		return
	}
	tps.SetDotSize(6)
	tps.SetLineWidth(2)

	// Examples of methods for altering the chart appearance
	tempChart.SetOrigin(time.Now().Add(time.Hour*20), 4)
	tempChart.SetTAxisLabel("T axis")
	tempChart.SetYAxisLabel("Y axis")
	tempChart.SetTitle("Cartesian Temporal Chart")
	as := style.DefaultAxisLabelStyle()
	as.SizeName = theme.SizeNameText
	as.Alignment = fyne.TextAlignLeading
	tempChart.SetTAxisStyle(as, style.DefaultAxisStyle())
	return
}

// Cartesian Categorical Chart
func cartCatChart() (catChart *coord.CartesianCategoricalChart, err error) {
	catChart = coord.NewCartesianCategoricalChart("Cartesian Categorical Chart")
	colPal := style.NewPaletteLightDark(theme.ColorNameSuccess)

	// Stacked Bar Series
	data1 := []data.CategoricalPoint{
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
	sbs1, err := coord.NewCategoricalPointSeries("Test1", colPal.Next(), data1)
	if err != nil {
		return
	}
	data2 := []data.CategoricalPoint{
		{
			C:   "One",
			Val: rand.Float64() * 30,
		},
		{
			C:   "Two",
			Val: rand.Float64() * 30,
		},
	}
	sbs2, err := coord.NewCategoricalPointSeries("Test2", colPal.Next(), data2)
	if err != nil {
		return
	}
	catSer, err := coord.NewCategoricalStackedSeries("stacked bar", []*coord.CategoricalPointSeries{sbs1, sbs2})
	if err != nil {
		return
	}
	err = catChart.AddStackedBarSeries(catSer)
	if err != nil {
		return
	}

	// Bar Series
	data3 := []data.CategoricalPoint{
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
	bs, err := coord.NewCategoricalPointSeries("bar", theme.ColorNameError, data3)
	if err != nil {
		return
	}
	err = catChart.AddBarSeries(bs)
	if err != nil {
		return
	}

	// Box Series
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
	cbs, err := coord.NewCategoricalBoxSeries("box", theme.ColorNamePrimary, data4)
	if err != nil {
		return
	}
	err = catChart.AddBoxSeries(cbs)
	if err != nil {
		return
	}

	// Examples of methods for altering the chart appearance
	catChart.SetCAxisLabel("C axis")
	catChart.SetYAxisLabel("Y axis")
	ts := style.DefaultTitleStyle()
	ts.SizeName = theme.SizeNameText
	catChart.SetTitleStyle(ts)
	catChart.SetLegendStyle(style.LegendLocationRight, style.DefaultLegendTextStyle(), true)
	catChart.SetOrientation(true)

	return
}

// Polar Numerical Chart
func polNumChart() (numChart *coord.PolarNumericalChart, err error) {
	numChart = coord.NewPolarNumericalChart("Polar Numerical Chart")
	colPal := style.ColorPalette{}
	colPal.Add(theme.ColorNameError)
	colPal.Add(theme.ColorNameSuccess)

	// Area Series
	data1 := make([]data.NumericalPoint, 0)
	for range 100 {
		data1 = append(data1, randomAngularDataPoint(63.777))
	}
	as, err := coord.NewNumericalPointSeries("area", colPal.Next(), data1)
	if err != nil {
		return
	}
	err = numChart.AddAreaSeries(as, true)
	if err != nil {
		return
	}

	// Line Series
	data2 := make([]data.NumericalPoint, 0)
	for range 150 {
		data2 = append(data2, randomSineAngularDataPoint(63.777))
	}
	ls, err := coord.NewNumericalPointSeries("line", colPal.Next(), data2)
	if err != nil {
		return
	}
	err = numChart.AddLineSeries(ls, false)
	if err != nil {
		return
	}
	ls.SetLineWidth(2)

	// Examples of methods for altering the chart appearance
	numChart.SetOrigin(0, 64)
	numChart.SetPhiAxisLabel("Phi axis")
	numChart.SetRAxisLabel("R axis")
	s := style.DefaultAxisStyle()
	s.LineColorName = theme.ColorNameSuccess
	s.TickColorName = theme.ColorNameError
	s.TickTextStyle.Italic = true
	numChart.SetPhiAxisStyle(style.DefaultAxisLabelStyle(), s)
	return
}

// Polar Temporal Chart
func polTempChart() (tempChart *coord.PolarTemporalChart, err error) {
	tempChart = coord.NewPolarTemporalChart("Polar Temporal Chart")
	colPal := style.NewPaletteComplementary(theme.ColorNamePrimary)

	// Lollipop Series
	data1 := make([]data.TemporalPoint, 0)
	for range 50 {
		data1 = append(data1, randomTemporalDataPoint(time.Now(), time.Now().Add(time.Hour*50), 0, 111))
	}
	tps1, err := coord.NewTemporalPointSeries("lollipop", colPal.Next(), data1)
	if err != nil {
		return
	}
	err = tempChart.AddLollipopSeries(tps1)
	if err != nil {
		return
	}

	// Scatter Series
	data2 := make([]data.TemporalPoint, 0)
	for range 25 {
		data2 = append(data2, randomTemporalDataPoint(time.Now(), time.Now().Add(time.Hour*50), 0, 111))
	}
	tps2, err := coord.NewTemporalPointSeries("scatter", colPal.Next(), data2)
	if err != nil {
		return
	}
	err = tempChart.AddScatterSeries(tps2)
	if err != nil {
		return
	}

	// Examples of methods for altering the chart appearance
	tempChart.SetOrigin(time.Now().Add(time.Hour*10), 120)
	tempChart.SetTAxisLabel("T axis")
	tempChart.SetRAxisLabel("R axis")
	ts := style.DefaultTitleStyle()
	ts.SizeName = theme.SizeNameHeadingText
	ts.ColorName = theme.ColorNameError
	tempChart.SetTitleStyle(ts)
	return
}

// Polar Categorical Chart
func polCatChart() (catChart *coord.PolarCategoricalChart, err error) {
	catChart = coord.NewPolarCategoricalChart("")
	colPal := style.NewPaletteLightMediumDark(theme.ColorNameError)

	// Stacked Bar Series
	data1 := []data.CategoricalPoint{
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
	sbs1, err := coord.NewCategoricalPointSeries("Test1", colPal.Next(), data1)
	if err != nil {
		return
	}
	data2 := []data.CategoricalPoint{
		{
			C:   "One",
			Val: rand.Float64() * 30,
		},
		{
			C:   "Two",
			Val: rand.Float64() * 30,
		},
	}
	sbs2, err := coord.NewCategoricalPointSeries("Test2", colPal.Next(), data2)
	if err != nil {
		return
	}
	catSer, err := coord.NewCategoricalStackedSeries("stacked bar", []*coord.CategoricalPointSeries{sbs1, sbs2})
	if err != nil {
		return
	}
	err = catChart.AddStackedBarSeries(catSer)
	if err != nil {
		return
	}

	// Bar Series
	data3 := []data.CategoricalPoint{
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
	bs, err := coord.NewCategoricalPointSeries("bar", colPal.Next(), data3)
	if err != nil {
		return
	}
	err = catChart.AddBarSeries(bs)
	if err != nil {
		return
	}

	// Examples of methods for altering the chart appearance
	catChart.SetRRange(40)
	catChart.SetCAxisLabel("C axis")
	catChart.SetRAxisLabel("R axis")
	catChart.SetTitle("Polar Categorical Chart")
	return
}

func randomNumericalDataPoint(xMin float64, xMax float64, valMin float64,
	valMax float64) (ndp data.NumericalPoint) {
	ndp.N = xMin + (rand.Float64() * (xMax - xMin))
	ndp.Val = valMin + (rand.Float64() * (valMax - valMin))
	return
}

func updateSineNumericalData() (ndp []data.NumericalPoint) {
	periodInMilliSecond := 10000
	shift := float64(time.Now().UnixMilli()%int64(periodInMilliSecond)) / float64(periodInMilliSecond) * 2 * math.Pi
	for range 50 {
		ndp = append(ndp, randomSineNumericalDataPoint(180, 50, shift))
	}
	return
}

func randomSineNumericalDataPoint(l float64, amp float64, shift float64) (ndp data.NumericalPoint) {
	ndp.N = (-l / 2) + (rand.Float64() * l)
	ndp.Val = amp * math.Sin((ndp.N/(l))*2*math.Pi-shift)
	return
}

func randomTemporalDataPoint(tMin time.Time, tMax time.Time, valMin float64,
	valMax float64) (tdp data.TemporalPoint) {
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

func randomAngularDataPoint(valMax float64) (adp data.NumericalPoint) {
	adp.N = rand.Float64() * 2 * math.Pi
	adp.Val = rand.Float64() * valMax
	return
}

func randomSineAngularDataPoint(amp float64) (adp data.NumericalPoint) {
	adp.N = rand.Float64() * 2 * math.Pi
	adp.Val = amp + (amp * math.Sin(adp.N))
	return
}
