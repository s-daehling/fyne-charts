package series

import (
	"errors"
	"image/color"
	"math"
	"strconv"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type proportionPoint struct {
	c            string
	n            float64
	val          float64
	height       float64
	hOffset      float64
	valOffset    float64
	rect         *canvas.Rectangle
	text         *canvas.Text
	visible      bool
	legendButton *LegendBox
	legendLabel  *canvas.Text
	ser          *ProportionalSeries
}

func emptyProportionPoint(showText bool, col color.Color) (point *proportionPoint) {
	point = &proportionPoint{
		rect:        canvas.NewRectangle(col),
		legendLabel: canvas.NewText("", theme.Color(theme.ColorNameForeground)),
		visible:     true,
	}
	point.legendButton = NewLegendBox(col, point.toggleView)
	if showText {
		point.text = canvas.NewText("", color.Black)
	}
	return
}

func (point *proportionPoint) toggleView() {
	if point.rect.Visible() {
		point.hide()
	} else {
		point.show()
	}
}

func (point *proportionPoint) hide() {
	if !point.visible {
		return
	}
	point.rect.Hide()
	if point.text != nil {
		point.text.Hide()
	}
	point.visible = false
	if point.ser != nil {
		point.ser.pointVisibilityUpdate(-point.val)
	}
}

func (point *proportionPoint) show() {
	if point.visible {
		return
	}
	point.rect.Show()
	if point.text != nil {
		point.text.Show()
	}
	point.visible = true
	point.ser.visible = true
	if point.ser != nil {
		point.ser.pointVisibilityUpdate(point.val)
	}
}

// func (point *proportionPoint) setColor(col color.Color) {
// 	point.rect.FillColor = col
// }

func (point *proportionPoint) legendEntry() (le LegendEntry) {
	le = LegendEntry{
		Button: point.legendButton,
		Label:  point.legendLabel,
		IsSub:  true,
	}
	return
}

func (point *proportionPoint) cartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (rs []CartesianRect) {
	if point.valOffset+point.n < xMin || point.valOffset > xMax {
		return
	}
	if point.hOffset+point.height < yMin || point.hOffset > yMax {
		return
	}
	r := CartesianRect{
		X1:   point.valOffset,
		Y1:   point.hOffset,
		X2:   point.n + point.valOffset,
		Y2:   point.hOffset + point.height,
		Rect: point.rect,
	}

	rs = append(rs, r)
	return
}

func (point *proportionPoint) cartesianTexts(xMin float64, xMax float64, yMin float64,
	yMax float64) (ts []CartesianText) {
	if point.text == nil {
		return
	}
	if point.valOffset+point.n < xMin || point.valOffset > xMax {
		return
	}
	if point.hOffset+point.height < yMin || point.hOffset > yMax {
		return
	}
	point.text.Text = strconv.FormatFloat(point.n, 'f', 0, 64) + "%"
	t := CartesianText{
		X:    point.valOffset + (point.n / 2),
		Y:    point.hOffset + (point.height / 2),
		Text: point.text,
	}
	ts = append(ts, t)
	return
}

func (point *proportionPoint) RasterColorPolar(phi float64, r float64, x float64, y float64) (col color.Color) {
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	if !point.visible {
		return
	}
	if phi < point.valOffset ||
		phi > point.valOffset+point.n ||
		r < point.hOffset || r > point.hOffset+point.height {
		return
	}
	col = point.rect.FillColor
	return
}

func (point *proportionPoint) polarTexts(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ts []PolarText) {
	if point.text == nil {
		return
	}
	if point.valOffset+point.n < phiMin || point.valOffset > phiMax {
		return
	}
	if point.hOffset+point.height < rMin || point.hOffset > rMax {
		return
	}
	point.text.Text = strconv.FormatFloat(100*(point.n/(2*math.Pi)), 'f', 0, 64) + "%"
	t := PolarText{
		Phi:  point.valOffset + (point.n / 2),
		R:    point.hOffset + (point.height / 2),
		Text: point.text,
	}
	ts = append(ts, t)
	return
}

type ProportionalSeries struct {
	baseSeries
	showText bool
	data     []*proportionPoint
	tot      float64
}

func EmptyProportionalSeries(chart chart, name string, polar bool) (ser *ProportionalSeries) {
	ser = &ProportionalSeries{showText: true}
	ser.baseSeries = emptyBaseSeries(chart, name, color.Black, polar, ser.toggleView)
	return
}

func (ser *ProportionalSeries) CRange() (cs []string) {
	ser.mutex.Lock()
	for i := range ser.data {
		cs = append(cs, ser.data[i].c)
	}
	ser.mutex.Unlock()
	return
}

func (ser *ProportionalSeries) ValRange() (isEmpty bool, min float64, max float64) {
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

func (ser *ProportionalSeries) ConvertPtoN(pToN func(p float64) (n float64)) {
	valOffset := 0.0
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].valOffset = valOffset
		if ser.data[i].visible {
			ser.data[i].n = pToN(ser.data[i].val / ser.tot)
			valOffset += ser.data[i].n
		} else {
			ser.data[i].n = 0
		}
	}
	ser.mutex.Unlock()
}

func (ser *ProportionalSeries) CartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []CartesianRect) {
	ser.mutex.Lock()
	for i := range ser.data {
		fs = append(fs, ser.data[i].cartesianRects(xMin, xMax, yMin, yMax)...)
	}
	ser.mutex.Unlock()
	return
}

func (ser *ProportionalSeries) CartesianTexts(xMin float64, xMax float64, yMin float64,
	yMax float64) (ts []CartesianText) {
	ser.mutex.Lock()
	for i := range ser.data {
		ts = append(ts, ser.data[i].cartesianTexts(xMin, xMax, yMin, yMax)...)
	}
	ser.mutex.Unlock()
	return
}

func (ser *ProportionalSeries) RasterColorPolar(phi float64, r float64, x float64, y float64) (col color.Color) {
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

func (ser *ProportionalSeries) PolarTexts(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ts []PolarText) {
	ser.mutex.Lock()
	for i := range ser.data {
		ts = append(ts, ser.data[i].polarTexts(phiMin, phiMax, rMin, rMax)...)
	}
	ser.mutex.Unlock()
	return
}

// Show makes the Bars of the series visible
func (ser *ProportionalSeries) Show() {
	ser.visible = true
	for i := range ser.data {
		ser.data[i].show()
	}
}

// Hide hides the Barss of the series
func (ser *ProportionalSeries) Hide() {
	ser.visible = false
	for i := range ser.data {
		ser.data[i].hide()
	}
}

func (ser *ProportionalSeries) toggleView() {
	if ser.visible {
		ser.Hide()
	} else {
		ser.Show()
	}
	if ser.polar {
		ser.chart.RasterVisibilityChange()
	}
}

func (ser *ProportionalSeries) pointVisibilityUpdate(totChange float64) {
	ser.mutex.Lock()
	ser.tot += totChange
	ser.mutex.Unlock()
	ser.chart.DataChange()
}

func (ser *ProportionalSeries) LegendEntries() (les []LegendEntry) {
	ser.mutex.Lock()
	les = append(les, ser.baseSeries.LegendEntries()...)
	for i := range ser.data {
		les = append(les, ser.data[i].legendEntry())
	}
	ser.mutex.Unlock()
	return
}

func (ser *ProportionalSeries) SetHeightAndOffset(h float64, hOffset float64) {
	ser.mutex.Lock()
	for i := range ser.data {
		ser.data[i].height = h
		ser.data[i].hOffset = hOffset
	}
	ser.mutex.Unlock()
}

func (ser *ProportionalSeries) Clear() (err error) {
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	ser.data = []*proportionPoint{}
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *ProportionalSeries) DeleteDataInRange(cat []string) (c int, err error) {
	c = 0
	if len(cat) == 0 {
		err = errors.New("invald range")
		return
	}
	finalData := []*proportionPoint{}
	ser.mutex.Lock()
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		ser.mutex.Unlock()
		return
	}
	chart := ser.chart
	tot := 0.0
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
			tot += ser.data[i].val
		}
	}
	if c == 0 {
		ser.mutex.Unlock()
		return
	}
	ser.data = nil
	ser.data = finalData
	ser.tot = tot
	ser.mutex.Unlock()
	chart.DataChange()
	return
}

func (ser *ProportionalSeries) AddData(input []data.ProportionalDataPoint) (err error) {
	if len(input) == 0 {
		err = errors.New("no input data")
		return
	}
	for i := range input {
		if input[i].Val < 0 {
			err = errors.New("invalid data")
			return
		}
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
		pPoint := emptyProportionPoint(ser.showText, input[i].Col)
		pPoint.c = input[i].C
		pPoint.legendLabel.Text = input[i].C
		pPoint.val = input[i].Val
		pPoint.ser = ser
		ser.data = append(ser.data, pPoint)
		ser.tot += pPoint.val
	}
	ser.mutex.Unlock()
	chart.DataChange()
	return
}
