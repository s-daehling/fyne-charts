package series

import (
	"errors"
	"image/color"
	"math"
	"time"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2/canvas"
)

type catOffset struct {
	C         string
	ValOffset float64
}

type barPoint struct {
	c         string
	n         float64
	t         time.Time
	val       float64
	nWidth    float64
	nOffset   float64
	tWidth    time.Duration
	tOffset   time.Duration
	valOffset float64
	rect      *canvas.Rectangle
}

func emptyBarPoint(color color.Color) (point *barPoint) {
	point = &barPoint{
		rect: canvas.NewRectangle(color),
	}
	return
}

func (point *barPoint) hide() {
	point.rect.Hide()
}

func (point *barPoint) show() {
	point.rect.Show()
}

func (point *barPoint) setColor(col color.Color) {
	point.rect.FillColor = col
}

func (point *barPoint) cartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []CartesianRect) {
	if point.n < xMin || point.n > xMax {
		return
	}
	f := CartesianRect{
		X1:   point.n + point.nOffset - (point.nWidth / 2),
		Y1:   point.valOffset,
		X2:   point.n + point.nOffset + (point.nWidth / 2),
		Y2:   point.valOffset + point.val,
		Rect: point.rect,
	}
	if f.Y2 < f.Y1 {
		f.Y1 = point.valOffset + point.val
		f.Y2 = point.valOffset
	}
	if yMin > f.Y1 {
		f.Y1 = yMin
	}
	if f.Y2 > yMax {
		f.Y2 = yMax
	}
	fs = append(fs, f)
	return
}

func (point *barPoint) RasterColorPolar(phi float64, r float64, x float64, y float64) (col color.Color) {
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	if phi < point.n+point.nOffset-(point.nWidth/2) ||
		phi > point.n+point.nOffset+(point.nWidth/2) ||
		r < point.valOffset || r > point.val+point.valOffset {
		return
	}
	col = point.rect.FillColor
	return
}

type BarSeries struct {
	baseSeries
	data []*barPoint
}

func EmptyBarSeries(chart chart, name string, color color.Color, polar bool) (ser *BarSeries) {
	ser = &BarSeries{}
	ser.baseSeries = emptyBaseSeries(chart, name, color, polar, ser.toggleView)
	return
}

func (ser *BarSeries) CRange() (cs []string) {
	ser.mutex.Lock()
	for i := range ser.data {
		cs = append(cs, ser.data[i].c)
	}
	ser.mutex.Unlock()
	return
}

func (ser *BarSeries) ValRange() (isEmpty bool, min float64, max float64) {
	min = 0
	max = 0
	isEmpty = false
	if len(ser.data) == 0 {
		isEmpty = true
		return
	}
	ser.mutex.Lock()
	min = ser.data[0].val
	max = ser.data[0].val
	for i := range ser.data {
		if ser.data[i].val < min {
			min = ser.data[i].val
		}
		if ser.data[i].val > max {
			max = ser.data[i].val
		}
	}
	ser.mutex.Unlock()
	return
}

func (ser *BarSeries) ConvertCtoN(cToN func(c string) (n float64)) {
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].n = cToN(ser.data[i].c)
	}
	ser.mutex.Unlock()
}

func (ser *BarSeries) ConvertTtoN(tToN func(t time.Time) (n float64)) {
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].n = tToN(ser.data[i].t)
		ser.data[i].nWidth = tToN(ser.data[i].t.Add(ser.data[i].tWidth)) - ser.data[i].n
		ser.data[i].nOffset = tToN(ser.data[i].t.Add(ser.data[i].tOffset)) - ser.data[i].n
	}
	ser.mutex.Unlock()
}

func (ser *BarSeries) CartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []CartesianRect) {
	ser.mutex.Lock()
	for i := range ser.data {
		fs = append(fs, ser.data[i].cartesianRects(xMin, xMax, yMin, yMax)...)
	}
	ser.mutex.Unlock()
	return
}

func (ser *BarSeries) RasterColorPolar(phi float64, r float64, x float64, y float64) (col color.Color) {
	col = ser.baseSeries.RasterColorPolar(phi, r, x, y)
	if !ser.visible {
		return
	}
	ser.mutex.Lock()
	for i := range ser.data {
		pCol := ser.data[i].RasterColorPolar(phi, r, x, y)
		r, g, b, _ := pCol.RGBA()
		if r > 0 || g > 0 || b > 0 {
			col = pCol
			break
		}
	}
	ser.mutex.Unlock()
	return
}

func (ser *BarSeries) SetNumericalWidthAndOffset(width float64, offset float64) (err error) {
	if width < 0 {
		err = errors.New("invalid width")
		return
	}
	if ser.polar && width > 2*math.Pi {
		err = errors.New("invalid width")
		return
	}
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].nWidth = width
		ser.data[i].nOffset = offset
	}
	ser.mutex.Unlock()
	return
}

func (ser *BarSeries) SetTemporalWidthAndOffset(width time.Duration, offset time.Duration) (err error) {
	if width < 0 {
		err = errors.New("invalid width")
		return
	}
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].tWidth = width
		ser.data[i].tOffset = offset
	}
	ser.mutex.Unlock()
	return
}

func (ser *BarSeries) SetAndUpdateValOffset(in []catOffset) (out []catOffset) {
	ser.mutex.Lock()
	for i := range in {
		for j := range ser.data {
			if in[i].C == ser.data[j].c {
				ser.data[j].valOffset = in[i].ValOffset
			}
		}
	}
	copy(out, in)
	for i := range ser.data {
		catExist := false
		for j := range out {
			if ser.data[i].c == out[j].C {
				catExist = true
				out[j].ValOffset += ser.data[i].val
				break
			}
		}
		if !catExist {
			out = append(out, catOffset{C: ser.data[i].c, ValOffset: ser.data[i].val})
		}
	}
	ser.mutex.Unlock()
	return
}

// Show makes the bars of the series visible
func (ser *BarSeries) Show() {
	ser.mutex.Lock()
	ser.visible = true
	for i := range ser.data {
		ser.data[i].show()
	}
	ser.mutex.Unlock()
}

// Hide hides the bars of the series
func (ser *BarSeries) Hide() {
	ser.mutex.Lock()
	ser.visible = false
	for i := range ser.data {
		ser.data[i].hide()
	}
	ser.mutex.Unlock()
}

func (ser *BarSeries) toggleView() {
	if ser.visible {
		ser.Hide()
	} else {
		ser.Show()
	}
	if ser.polar {
		ser.chart.RasterVisibilityChange()
	}
}

// SetColor changes the color of the bar series
func (ser *BarSeries) SetColor(col color.Color) {
	ser.mutex.Lock()
	ser.color = col
	ser.legendButton.color = col
	ser.legendButton.rect.FillColor = col
	for i := range ser.data {
		ser.data[i].setColor(col)
	}
	ser.mutex.Unlock()
}

func (ser *BarSeries) DeleteNumericalDataInRange(min float64, max float64) (c int, err error) {
	c = 0
	if min > max {
		err = errors.New("invalid range")
		return
	}
	finalData := []*barPoint{}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	for i := range ser.data {
		if ser.data[i].n > min && ser.data[i].n < max {
			c++
		} else {
			finalData = append(finalData, ser.data[i])
		}
	}
	if c == 0 {
		ser.mutex.Unlock()
		return
	}
	ser.data = nil
	ser.data = finalData
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *BarSeries) AddNumericalData(input []data.NumericalDataPoint) (err error) {
	err = numericalDataPointRangeCheck(input, ser.polar, ser.polar)
	if err != nil {
		return
	}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	for i := range input {
		lPoint := emptyBarPoint(ser.color)
		lPoint.n = input[i].X
		lPoint.val = input[i].Val
		ser.data = append(ser.data, lPoint)
	}
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *BarSeries) DeleteCategoricalDataInRange(cat []string) (c int, err error) {
	c = 0
	if len(cat) == 0 {
		err = errors.New("invald range")
		return
	}
	finalData := []*barPoint{}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	for i := range ser.data {
		del := false
		for j := range cat {
			if ser.data[i].c == cat[j] {
				del = true
				break
			}
		}
		if del {
			c++
		} else {
			finalData = append(finalData, ser.data[i])
		}
	}
	if c == 0 {
		ser.mutex.Unlock()
		return
	}
	ser.data = nil
	ser.data = finalData
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *BarSeries) AddCategoricalData(input []data.CategoricalDataPoint) (err error) {
	err = categoricalDataPointRangeCheck(input, ser.polar)
	if err != nil {
		return
	}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	for i := range input {
		catExist := false
		for j := range ser.data {
			if input[i].C == ser.data[j].c {
				catExist = true
				break
			}
		}
		if catExist {
			continue
		}
		bPoint := emptyBarPoint(ser.color)
		bPoint.c = input[i].C
		bPoint.val = input[i].Val
		ser.data = append(ser.data, bPoint)
	}
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *BarSeries) DeleteTemporalDataInRange(min time.Time, max time.Time) (c int, err error) {
	c = 0
	if min.After(max) {
		err = errors.New("invalid range")
		return
	}
	finalData := []*barPoint{}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	for i := range ser.data {
		if ser.data[i].t.After(min) && ser.data[i].t.Before(max) {
			c++
		} else {
			finalData = append(finalData, ser.data[i])
		}
	}
	if c == 0 {
		ser.mutex.Unlock()
		return
	}
	ser.data = nil
	ser.data = finalData
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *BarSeries) AddTemporalData(input []data.TemporalDataPoint) (err error) {
	err = temporalDataPointRangeCheck(input, ser.polar)
	if err != nil {
		return
	}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	for i := range input {
		bPoint := emptyBarPoint(ser.color)
		bPoint.t = input[i].T
		bPoint.val = input[i].Val
		ser.data = append(ser.data, bPoint)
	}
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *BarSeries) DeleteAngularDataInRange(min float64, max float64) (c int, err error) {
	c, err = ser.DeleteNumericalDataInRange(min, max)
	return
}

func (ser *BarSeries) AddAngularData(input []data.AngularDataPoint) (err error) {
	err = ser.AddNumericalData(angularToNumerical(input))
	return
}
