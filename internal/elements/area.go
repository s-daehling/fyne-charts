package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Area struct {
	rast *canvas.Raster
	widget.BaseWidget
}

func NewArea(pixelGen func(pX, pY, w, h int) (col color.Color)) (a *Area) {
	a = &Area{
		rast: canvas.NewRasterWithPixels(pixelGen),
	}
	a.ExtendBaseWidget(a)
	return
}

func (a *Area) CreateRenderer() (r fyne.WidgetRenderer) {
	r = newAreaRenderer(a)
	return
}

type areaRenderer struct {
	area *Area
}

func newAreaRenderer(a *Area) (ar *areaRenderer) {
	ar = &areaRenderer{
		area: a,
	}
	return
}

func (ar *areaRenderer) Layout(size fyne.Size) {
	ar.area.rast.Resize(size)
	ar.area.rast.Move(fyne.NewPos(0, 0))
}

func (ar *areaRenderer) MinSize() (size fyne.Size) {
	size = fyne.NewSize(0, 0)
	return
}

func (ar *areaRenderer) Refresh() {
	ar.area.rast.Refresh()
}

func (ar *areaRenderer) Objects() (canObj []fyne.CanvasObject) {
	canObj = append(canObj, ar.area.rast)
	return
}

func (ar *areaRenderer) Destroy() {}
