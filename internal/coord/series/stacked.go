package series

import (
	"errors"
	"image/color"

	"fyne.io/fyne/v2/theme"
	"github.com/s-daehling/fyne-charts/internal/interact"
	"github.com/s-daehling/fyne-charts/internal/renderer"
)

type StackedSeries struct {
	baseSeries
	stack  []*PointSeries
	valMin float64
	valMax float64
}

func EmptyStackedSeries(name string) (ser *StackedSeries) {
	ser = &StackedSeries{}
	ser.baseSeries = emptyBaseSeries(name, theme.ColorNameForeground, ser.toggleView)
	ser.legendEntry.HideBox()
	return
}

func (ser *StackedSeries) IsPolar() (b bool) {
	if ser.cont != nil {
		b = ser.cont.IsPolar()
	}
	return
}

func (ser *StackedSeries) DataChange() {
	if ser.cont != nil {
		ser.cont.DataChange()
	}
}

func (ser *StackedSeries) RasterRefresh() {
	if ser.cont != nil {
		ser.cont.RasterRefresh()
	}
}

func (ser *StackedSeries) AddLegendEntry(le *interact.LegendEntry) {
	if ser.cont != nil {
		ser.cont.AddLegendEntry(le)
	}
}

func (ser *StackedSeries) RemoveLegendEntry(name string, super string) {
	if ser.cont != nil {
		ser.cont.RemoveLegendEntry(name, super)
	}
}

func (ser *StackedSeries) CRange() (cs []string) {
	for i := range ser.stack {
		cats := ser.stack[i].CRange()
		for j := range cats {
			cExist := false
			for k := range cs {
				if cs[k] == cats[j] {
					cExist = true
					break
				}
			}
			if !cExist {
				cs = append(cs, cats[j])
			}
		}
	}
	return
}

func (ser *StackedSeries) ValRange() (isEmpty bool, min float64, max float64) {
	ser.UpdateValOffset()
	min = 0
	max = 0
	isEmpty = true
	if len(ser.stack) == 0 {
		return
	}
	for i := range ser.stack {
		sEmpty, sMin, sMax := ser.stack[i].ValRange()
		if sEmpty {
			continue
		}
		if isEmpty {
			// first non-empty series
			isEmpty = false
			min = sMin
			max = sMax
		} else {
			if sMin < min {
				min = sMin
			}
			if sMax > max {
				max = sMax
			}
		}
	}
	ser.valMin = min
	ser.valMax = max
	return
}

func (ser *StackedSeries) ConvertCtoN(cToN func(c string) (n float64)) {
	for i := range ser.stack {
		ser.stack[i].ConvertCtoN(cToN)
	}
}

func (ser *StackedSeries) CartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []renderer.CartesianRect) {
	for i := range ser.stack {
		fs = append(fs, ser.stack[i].CartesianRects(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *StackedSeries) RasterColorPolar(phi float64, r float64, x float64, y float64) (col color.Color) {
	col = ser.baseSeries.RasterColorPolar(phi, r, x, y)
	if !ser.visible || r > ser.valMax {
		return
	}
	for i := range ser.stack {
		sCol := ser.stack[i].RasterColorPolar(phi, r, x, y)
		r, g, b, _ := sCol.RGBA()
		if r > 0 || g > 0 || b > 0 {
			col = sCol
			break
		}
	}
	return
}

func (ser *StackedSeries) IsPartOfChartRaster() (b bool) {
	b = false
	if ser.cont == nil || !ser.visible {
		return
	}
	if !ser.cont.IsPolar() {
		return
	}
	b = true
	return
}

func (ser *StackedSeries) RefreshTheme() {
	ser.col = theme.Color(ser.colName)
	for i := range ser.stack {
		ser.stack[i].RefreshTheme()
	}
}

// setWidthAndOffset sets width of bars and offset from x coordinate for this series
func (ser *StackedSeries) SetNumericalBarWidthAndShift(width float64, shift float64) (err error) {
	for i := range ser.stack {
		err = ser.stack[i].SetNumericalBarWidthAndShift(width, shift)
		if err != nil {
			return
		}
	}
	return
}

func (ser *StackedSeries) UpdateValOffset() {
	valOffset := []catOffset{}
	for i := range ser.stack {
		valOffset = ser.stack[i].SetAndUpdateValBaseCategorical(valOffset)
	}
}

// Show makes the bars of the series visible
func (ser *StackedSeries) Show() {
	ser.visible = true
	for i := range ser.stack {
		ser.stack[i].Show()
	}
	ser.legendEntry.Show()
}

// Hide hides the bars of the series
func (ser *StackedSeries) Hide() {
	ser.visible = false
	for i := range ser.stack {
		ser.stack[i].Hide()
	}
	ser.legendEntry.Hide()
}

func (ser *StackedSeries) toggleView() {
	if ser.visible {
		ser.Hide()
	} else {
		ser.Show()
	}
}

func (ser *StackedSeries) BindToChart(ch container) (err error) {
	err = ser.baseSeries.BindToChart(ch)
	if err != nil {
		return
	}
	for i := range ser.stack {
		ch.AddLegendEntry(ser.stack[i].legendEntry)
	}
	return
}

func (ser *StackedSeries) Release() {
	if ser.cont == nil {
		return
	}
	for i := range ser.stack {
		ser.cont.RemoveLegendEntry(ser.stack[i].name, ser.stack[i].super)
	}
	ser.baseSeries.Release()
}

func (ser *StackedSeries) Clear() {
	ser.stack = []*PointSeries{}
	if ser.cont != nil {
		ser.cont.DataChange()
	}
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
func (ser *StackedSeries) DeleteCategoricalDataInRange(cat []string) (c int) {
	c = 0
	if len(cat) == 0 {
		return
	}
	for i := range ser.stack {
		c += ser.stack[i].DeleteCategoricalDataInRange(cat)
	}
	return
}

func (ser *StackedSeries) RemovePointSeries(name string) {
	newStack := make([]*PointSeries, 0)
	for i := range ser.stack {
		if ser.stack[i].name != name {
			newStack = append(newStack, ser.stack[i])
		} else {
			ser.stack[i].Release()
		}
	}
	ser.stack = newStack
	if ser.cont != nil {
		ser.cont.DataChange()
	}
}

func (ser *StackedSeries) AddPointSeries(ps *PointSeries) (err error) {
	if ser.seriesExist(ps.name) {
		err = errors.New("series already exists")
		return
	}
	ps.MakeBar()
	err = ps.BindToStack(ser)
	if err != nil {
		return
	}
	ser.stack = append(ser.stack, ps)
	if ser.cont != nil {
		ser.cont.DataChange()
	}
	return
}

func (ser *StackedSeries) seriesExist(name string) (exist bool) {
	exist = false
	for i := range ser.stack {
		if ser.stack[i].name == name {
			exist = true
			break
		}
	}
	return
}
