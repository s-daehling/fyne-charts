package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Test")

	l := widget.NewLabel("long long long text")
	l.Wrapping = fyne.TextWrapWord

	myWindow.SetContent(l)
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}
