package main

import (
	"image/color"
	"math/rand/v2"
	"time"

	"github.com/s-daehling/fyne-charts/pkg/data"
	"github.com/s-daehling/fyne-charts/pkg/prop"
	"github.com/s-daehling/fyne-charts/pkg/style"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Proportional Charts")

	var barCh, pieCh fyne.CanvasObject
	var err error
	barCh, err = barChart()
	if err != nil {
		panic(err)
	}
	pieCh, err = pieChart()
	if err != nil {
		panic(err)
	}

	vS := container.NewHSplit(barCh, pieCh)

	myWindow.SetContent(vS)
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}

// Stacked Bar Chart
func barChart() (propChart *prop.BarChart, err error) {
	propChart = prop.NewBarChart("Proportional Bar Chart")

	// Series 1
	data1 := []data.ProportionalPoint{
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
	ps, err := prop.NewSeries("proportion", data1)
	if err != nil {
		return
	}
	err = propChart.AddSeries(ps)
	if err != nil {
		return
	}

	// Series 2
	data2 := []data.ProportionalPoint{
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
	ps2, err := prop.NewSeries("proportion 2", data2)
	if err != nil {
		return
	}
	err = propChart.AddSeries(ps2)
	if err != nil {
		return
	}
	return
}

func pieChart() (propChart *prop.PieChart, err error) {
	propChart = prop.NewPieChart("")

	// Series 1
	data1 := []data.ProportionalPoint{
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
	ps, err := prop.NewSeries("proportion", data1)
	if err != nil {
		return
	}
	err = propChart.AddSeries(ps)
	if err != nil {
		return
	}

	// Series 2
	data2 := []data.ProportionalPoint{
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
	ps2, err := prop.NewSeries("proportion2", data2)
	if err != nil {
		return
	}
	err = propChart.AddSeries(ps2)
	if err != nil {
		return
	}
	ps.SetValTextColor(color.Black)

	propChart.SetTitle("Proportional Pie/Doughnut Chart")

	go func() {
		time.Sleep(time.Second * 2)
		fyne.Do(func() {
			propChart.SetLegendStyle(style.LegendLocationTop)
		})
	}()
	return
}
