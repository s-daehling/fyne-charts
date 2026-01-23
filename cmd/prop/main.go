package main

import (
	"math/rand/v2"
	"time"

	"github.com/s-daehling/fyne-charts/pkg/data"
	"github.com/s-daehling/fyne-charts/pkg/prop"
	"github.com/s-daehling/fyne-charts/pkg/style"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
			Val:     rand.Float64() * 222,
			C:       "One",
			ColName: theme.ColorNameError,
		},
		{
			Val:     rand.Float64() * 222,
			C:       "Two",
			ColName: theme.ColorNameSuccess,
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
			Val:     rand.Float64() * 222,
			C:       "One",
			ColName: theme.ColorNamePrimary,
		},
		{
			Val:     rand.Float64() * 222,
			C:       "Two",
			ColName: theme.ColorNameError,
		},
		{
			Val:     rand.Float64() * 222,
			C:       "Three",
			ColName: theme.ColorNameSuccess,
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
			Val:     rand.Float64() * 222,
			C:       "One",
			ColName: theme.ColorNameError,
		},
		{
			Val:     rand.Float64() * 222,
			C:       "Two",
			ColName: theme.ColorNameSuccess,
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
			Val:     rand.Float64() * 222,
			C:       "One",
			ColName: theme.ColorNameForeground,
		},
		{
			Val:     rand.Float64() * 222,
			C:       "Two",
			ColName: theme.ColorNameError,
		},
		{
			Val:     rand.Float64() * 222,
			C:       "Three",
			ColName: theme.ColorNameSuccess,
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
	ts := style.DefaultValueTextStyle()
	ts.ColorName = theme.ColorNameBackground
	ts.TextStyle.Italic = true
	ts.TextStyle.Bold = true
	ps.SetValueTextStyle(ts)

	propChart.SetTitle("Proportional Pie/Doughnut Chart")

	go func() {
		time.Sleep(time.Second * 2)
		fyne.Do(func() {
			ts := style.DefaultLegendTextStyle()
			ts.SizeName = theme.SizeNameCaptionText
			propChart.SetLegendStyle(style.LegendLocationTop, ts, true)
		})
	}()
	return
}
