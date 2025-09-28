package series

import (
	"errors"
	"image/color"

	"fyne.io/fyne/v2/theme"
	"github.com/s-daehling/fyne-charts/pkg/data"
)

type StackedBarSeries struct {
	baseSeries
	stack []*BarSeries
}

func EmptyStackedBarSeries(chart chart, name string, polar bool) (ser *StackedBarSeries) {
	ser = &StackedBarSeries{}
	ser.baseSeries = emptyBaseSeries(chart, name, color.Black, polar, ser.toggleView)
	return
}

func (ser *StackedBarSeries) CRange() (cs []string) {
	ser.mutex.Lock()
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
	ser.mutex.Unlock()
	return
}

func (ser *StackedBarSeries) ValRange() (isEmpty bool, min float64, max float64) {
	min = 0
	max = 0
	isEmpty = true
	ser.mutex.Lock()
	if len(ser.stack) == 0 {
		ser.mutex.Unlock()
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
	ser.mutex.Unlock()
	return
}

func (ser *StackedBarSeries) ConvertCtoN(cToN func(c string) (n float64)) {
	ser.mutex.Lock()
	for i := range ser.stack {
		ser.stack[i].ConvertCtoN(cToN)
	}
	ser.mutex.Unlock()
}

func (ser *StackedBarSeries) CartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []CartesianRect) {
	// valOffset := []catOffset{}
	ser.mutex.Lock()
	for i := range ser.stack {
		// valOffset = ser.stack[i].SetAndUpdateValOffset(valOffset)
		fs = append(fs, ser.stack[i].CartesianRects(xMin, xMax, yMin, yMax)...)
	}
	ser.mutex.Unlock()
	return
}

func (ser *StackedBarSeries) RasterColorPolar(phi float64, r float64, x float64, y float64) (col color.Color) {
	col = ser.baseSeries.RasterColorPolar(phi, r, x, y)
	if !ser.visible {
		return
	}
	ser.mutex.Lock()
	for i := range ser.stack {
		sCol := ser.stack[i].RasterColorPolar(phi, r, x, y)
		r, g, b, _ := sCol.RGBA()
		if r > 0 || g > 0 || b > 0 {
			col = sCol
			break
		}
	}
	ser.mutex.Unlock()
	return
}

func (ser *StackedBarSeries) LegendEntries() (les []LegendEntry) {
	les = append(les, ser.baseSeries.LegendEntries()...)
	ser.mutex.Lock()
	for i := range ser.stack {
		subLes := ser.stack[i].LegendEntries()
		for j := range subLes {
			subLes[j].IsSub = true
		}
		les = append(les, subLes...)
	}
	ser.mutex.Unlock()
	return
}

func (ser *StackedBarSeries) RefreshThemeColor() {
	ser.mutex.Lock()
	ser.legendLabel.Color = theme.Color(theme.ColorNameForeground)
	for i := range ser.stack {
		ser.stack[i].RefreshThemeColor()
	}
	ser.mutex.Unlock()
}

// setWidthAndOffset sets width of bars and offset from x coordinate for this series
func (ser *StackedBarSeries) SetNumericalWidthAndOffset(width float64, offset float64) (err error) {
	ser.mutex.Lock()
	for i := range ser.stack {
		err = ser.stack[i].SetNumericalWidthAndOffset(width, offset)
		if err != nil {
			ser.mutex.Unlock()
			return
		}
	}
	ser.mutex.Unlock()
	return
}

func (ser *StackedBarSeries) UpdateValOffset() {
	valOffset := []catOffset{}
	ser.mutex.Lock()
	for i := range ser.stack {
		valOffset = ser.stack[i].SetAndUpdateValOffset(valOffset)
	}
	ser.mutex.Unlock()
}

// Show makes the bars of the series visible
func (ser *StackedBarSeries) Show() {
	ser.mutex.Lock()
	ser.visible = true
	for i := range ser.stack {
		go ser.stack[i].Show()
	}
	ser.mutex.Unlock()
}

// Hide hides the bars of the series
func (ser *StackedBarSeries) Hide() {
	ser.mutex.Lock()
	ser.visible = false
	for i := range ser.stack {
		go ser.stack[i].Hide()
	}
	ser.mutex.Unlock()
}

func (ser *StackedBarSeries) toggleView() {
	if ser.visible {
		ser.Hide()
	} else {
		ser.Show()
	}
}

func (ser *StackedBarSeries) Clear() (err error) {
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	ser.stack = []*BarSeries{}
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

// DeleteDataInRange deletes all data points with one of the given category
// The return value gives the number of data points that have been removed
func (ser *StackedBarSeries) DeleteCategoricalDataInRange(cat []string) (c int, err error) {
	c = 0
	if len(cat) == 0 {
		err = errors.New("invald range")
		return
	}
	ser.mutex.Lock()
	for i := range ser.stack {
		var cs int
		cs, err = ser.stack[i].DeleteCategoricalDataInRange(cat)
		if err != nil {
			ser.mutex.Unlock()
			return
		}
		c += cs
	}
	ser.mutex.Unlock()
	return
}

// AddData adds data points to the stacked series.
// If the single series exists, the data points will be added to it
// If the single series does not exist, nothing is done
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
func (ser *StackedBarSeries) AddCategoricalData(series string, input []data.CategoricalDataPoint) (err error) {
	err = categoricalDataPointRangeCheck(input, true)
	if err != nil {
		return
	}
	ser.mutex.Lock()
	for i := range ser.stack {
		if ser.stack[i].name == series {
			err = ser.stack[i].AddCategoricalData(input)
			break
		}
	}
	ser.mutex.Unlock()
	return
}

// AddSeries adds a new single series to the stacked bar series.
// If the single series already exists, nothing will be done.
// The method checks for duplicates (i.e. data points with same C).
// Data points with a C that already exists, will be ignored.
func (ser *StackedBarSeries) AddCategoricalSeries(series data.CategoricalDataSeries) (err error) {
	if ser.seriesExist(series.Name) {
		err = errors.New("series already exists")
		return
	}
	err = categoricalDataPointRangeCheck(series.Points, true)
	if err != nil {
		return
	}
	ser.mutex.Lock()
	bs := EmptyBarSeries(ser.chart, series.Name, series.Col, ser.polar)
	err = bs.AddCategoricalData(series.Points)
	if err != nil {
		ser.mutex.Unlock()
		return
	}
	ser.stack = append(ser.stack, bs)
	ser.mutex.Unlock()
	return
}

func (ser *StackedBarSeries) seriesExist(name string) (exist bool) {
	exist = false
	ser.mutex.Lock()
	for i := range ser.stack {
		if ser.stack[i].name == name {
			exist = true
			break
		}
	}
	ser.mutex.Unlock()
	return
}
