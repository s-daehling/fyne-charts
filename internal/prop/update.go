package prop

import (
	"fyne.io/fyne/v2/theme"
)

func (base *BaseChart) DataChange() {
	base.updateSeriesVariables()
	base.render.Refresh()
}

func (base *BaseChart) RasterVisibilityChange() {
	if base.rast != nil {
		base.rast.Refresh()
	}
}

func (base *BaseChart) Resize(fromSpace float32, toSpace float32) {

}

func (base *BaseChart) updateSeriesVariables() {
	nPropSeries := len(base.series)
	propHeight := base.toMax / float64(nPropSeries)
	propOffset := 0.0
	for i := range base.series {
		base.series[i].SetHeightAndOffset(propHeight*0.9, propOffset)
		propOffset += propHeight
	}

	for i := range base.series {
		base.series[i].ConvertPtoN(base.ptoN)
	}
}

func (base *BaseChart) RefreshThemeColor() {
	base.label.Color = theme.Color(theme.ColorNameForeground)
	for i := range base.series {
		base.series[i].RefreshThemeColor()
	}
}

func (base *BaseChart) ptoN(p float64) (n float64) {
	n = p * base.fromMax
	return
}
