package elements

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type Area struct {
	rast  *canvas.Raster
	hover func(pX, pY, w, h, absX, absY float32)
	widget.BaseWidget
}

func NewArea(pixelGen func(pX, pY, w, h int) (col color.Color), hover func(pX, pY, w, h, absX, absY float32)) (a *Area) {
	a = &Area{
		rast:  canvas.NewRasterWithPixels(pixelGen),
		hover: hover,
	}
	a.ExtendBaseWidget(a)
	return
}

func (a *Area) CreateRenderer() (r fyne.WidgetRenderer) {
	r = newAreaRenderer(a)
	return
}

func (a *Area) MouseIn(me *desktop.MouseEvent) {

}

func (a *Area) MouseMoved(me *desktop.MouseEvent) {
	if a.hover != nil {
		size := a.rast.Size()
		a.hover(me.Position.X, me.Position.Y, size.Width, size.Height, me.AbsolutePosition.X, me.AbsolutePosition.Y)
	}
}

func (a *Area) MouseOut() {

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
