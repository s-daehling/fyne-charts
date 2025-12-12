package prop

import (
	"errors"
	"image/color"
	"math"
	"strconv"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"github.com/s-daehling/fyne-charts/internal/legend"
	"github.com/s-daehling/fyne-charts/internal/renderer"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

func (base *BaseChart) addSeriesIfNotExist(ser *Series) (err error) {
	for i := range base.series {
		if base.series[i].Name() == ser.Name() {
			err = errors.New("series already exists")
			return
		}
	}
	base.series = append(base.series, ser)
	base.DataChange()
	return
}

func (base *BaseChart) AddProportionalSeries(name string, points []data.ProportionalDataPoint) (ser *Series, err error) {
	pSeries := EmptyProportionalSeries(base, name, base.planeType == PolarPlane)
	err = pSeries.AddData(points)
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(pSeries)
	if err != nil {
		return
	}
	ser = pSeries
	return
}

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
	legendButton *legend.LegendBox
	legendLabel  *canvas.Text
	ser          *Series
}

func emptyProportionPoint(showText bool, col color.Color) (point *proportionPoint) {
	point = &proportionPoint{
		rect:        canvas.NewRectangle(col),
		legendLabel: canvas.NewText("", theme.Color(theme.ColorNameForeground)),
		visible:     true,
	}
	point.legendButton = legend.NewLegendBox(col, point.toggleView)
	if showText {
		point.text = canvas.NewText("", theme.Color(theme.ColorNameForeground))
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
	point.legendButton.ToggleColor()
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
	point.legendButton.ToggleColor()
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

func (point *proportionPoint) legendEntry() (le renderer.LegendEntry) {
	le = renderer.LegendEntry{
		Button: point.legendButton,
		Label:  point.legendLabel,
		IsSub:  true,
	}
	return
}

func (point *proportionPoint) cartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (rs []renderer.CartesianRect) {
	if point.valOffset+point.n < xMin || point.valOffset > xMax {
		return
	}
	if point.hOffset+point.height < yMin || point.hOffset > yMax {
		return
	}
	r := renderer.CartesianRect{
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
	yMax float64) (ts []renderer.CartesianText) {
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
	t := renderer.CartesianText{
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
	rMax float64) (ts []renderer.PolarText) {
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
	t := renderer.PolarText{
		Phi:  point.valOffset + (point.n / 2),
		R:    point.hOffset + (point.height / 2),
		Text: point.text,
	}
	ts = append(ts, t)
	return
}

type Series struct {
	showText         bool
	data             []*proportionPoint
	tot              float64
	name             string
	visible          bool
	autoValTextColor bool
	legendButton     *legend.LegendBox
	legendLabel      *canvas.Text
	polar            bool
	chart            *BaseChart
}

func EmptyProportionalSeries(chart *BaseChart, name string, polar bool) (ser *Series) {
	ser = &Series{
		name:             name,
		visible:          true,
		showText:         true,
		autoValTextColor: true,
		legendLabel:      canvas.NewText(name, theme.Color(theme.ColorNameForeground)),
		polar:            polar,
		chart:            chart,
	}
	ser.legendButton = legend.NewLegendBox(theme.Color(theme.ColorNameForeground), ser.toggleView)
	ser.legendButton.UseGradient(theme.Color(theme.ColorNameForeground), theme.Color(theme.ColorNameBackground))
	return
}

// Name gives the name of the series
func (ser *Series) Name() (n string) {
	n = ser.name
	return
}

func (ser *Series) Delete() {
	ser.chart = nil
}

func (ser *Series) ConvertPtoN(pToN func(p float64) (n float64)) {
	valOffset := 0.0
	for i := range ser.data {
		ser.data[i].valOffset = valOffset
		if ser.data[i].visible {
			ser.data[i].n = pToN(ser.data[i].val / ser.tot)
			valOffset += ser.data[i].n
		} else {
			ser.data[i].n = 0
		}
	}
}

func (ser *Series) CartesianRects(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []renderer.CartesianRect) {
	for i := range ser.data {
		fs = append(fs, ser.data[i].cartesianRects(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *Series) CartesianTexts(xMin float64, xMax float64, yMin float64,
	yMax float64) (ts []renderer.CartesianText) {
	for i := range ser.data {
		ts = append(ts, ser.data[i].cartesianTexts(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *Series) RasterColorPolar(phi float64, r float64, x float64, y float64) (col color.Color) {
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	if !ser.visible {
		return
	}
	for i := range ser.data {
		pCol := ser.data[i].RasterColorPolar(phi, r, x, y)
		r, g, b, _ := pCol.RGBA()
		if r > 0 || g > 0 || b > 0 {
			col = pCol
			break
		}
	}
	return
}

func (ser *Series) PolarTexts(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ts []renderer.PolarText) {
	for i := range ser.data {
		ts = append(ts, ser.data[i].polarTexts(phiMin, phiMax, rMin, rMax)...)
	}
	return
}

func (ser *Series) RefreshTheme() {
	ser.legendLabel.Color = theme.Color(theme.ColorNameForeground)
	// ser.legendButton.SetRectColor(theme.Color(theme.ColorNameForeground))
	ser.legendButton.SetGradColor(theme.Color(theme.ColorNameForeground), theme.Color(theme.ColorNameBackground))
	for i := range ser.data {
		ser.data[i].legendLabel.Color = theme.Color(theme.ColorNameForeground)
		if ser.autoValTextColor {
			ser.data[i].text.Color = theme.Color(theme.ColorNameForeground)
		}
	}
}

func (ser *Series) SetValTextColor(col color.Color) {
	ser.autoValTextColor = false
	for i := range ser.data {
		ser.data[i].text.Color = col
	}
}

func (ser *Series) SetAutoValTextColor() {
	ser.autoValTextColor = true
	for i := range ser.data {
		ser.data[i].text.Color = theme.Color(theme.ColorNameForeground)
	}
}

// Show makes the Bars of the series visible
func (ser *Series) Show() {
	ser.visible = true
	for i := range ser.data {
		ser.data[i].show()
	}
}

// Hide hides the Barss of the series
func (ser *Series) Hide() {
	ser.visible = false
	for i := range ser.data {
		ser.data[i].hide()
	}
}

func (ser *Series) toggleView() {
	if ser.visible {
		ser.Hide()
	} else {
		ser.Show()
	}
	if ser.polar {
		ser.chart.RasterVisibilityChange()
	}
}

func (ser *Series) pointVisibilityUpdate(totChange float64) {
	ser.tot += totChange
	ser.chart.DataChange()
}

func (ser *Series) LegendEntries() (les []renderer.LegendEntry) {
	les = append(les, renderer.LegendEntry{
		Button: ser.legendButton,
		Label:  ser.legendLabel,
		IsSub:  false,
	})
	for i := range ser.data {
		les = append(les, ser.data[i].legendEntry())
	}
	return
}

func (ser *Series) SetHeightAndOffset(h float64, hOffset float64) {
	for i := range ser.data {
		ser.data[i].height = h
		ser.data[i].hOffset = hOffset
	}
}

func (ser *Series) Clear() (err error) {
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
	chart := ser.chart
	ser.data = []*proportionPoint{}
	chart.DataChange()
	return
}

func (ser *Series) DeleteDataInRange(cat []string) (c int, err error) {
	c = 0
	if len(cat) == 0 {
		err = errors.New("invald range")
		return
	}
	finalData := []*proportionPoint{}
	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
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
		return
	}
	ser.data = nil
	ser.data = finalData
	ser.tot = tot
	ser.chart.DataChange()
	return
}

func (ser *Series) AddData(input []data.ProportionalDataPoint) (err error) {
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

	if ser.chart == nil {
		err = errors.New("series is not part of any chart")
		return
	}
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
	ser.chart.DataChange()
	return
}
