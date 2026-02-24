package prop

import (
	"errors"
	"image/color"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"github.com/s-daehling/fyne-charts/internal/elements"
	"github.com/s-daehling/fyne-charts/internal/interact"

	"github.com/s-daehling/fyne-charts/pkg/data"
	"github.com/s-daehling/fyne-charts/pkg/style"
)

func (base *BaseChart) addSeriesIfNotExist(ser *Series) (err error) {
	for i := range base.series {
		if base.series[i].Name() == ser.Name() {
			err = errors.New("series already exists")
			return
		}
	}
	err = ser.BindToChart(base)
	if err != nil {
		return
	}
	base.series = append(base.series, ser)
	base.DataChange()
	return
}

func (base *BaseChart) AddSeries(ps *Series) (err error) {
	err = base.addSeriesIfNotExist(ps)
	return
}

func (base *BaseChart) AddLegendEntry(le *interact.LegendEntry) {
	base.legend.AddEntry(le)
}

func (base *BaseChart) RemoveLegendEntry(name string, super string) {
	base.legend.RemoveEntry(name, super)
}

func (base *BaseChart) RemoveSeries(name string) {
	newSeries := make([]*Series, 0)
	for i := range base.series {
		if base.series[i].Name() != name {
			newSeries = append(newSeries, base.series[i])
		} else {
			base.series[i].Release()
		}
	}
	base.series = newSeries
	base.DataChange()
}

type proportionPoint struct {
	c           string
	n           float64
	val         float64
	height      float64
	hOffset     float64
	valOffset   float64
	bar         *elements.Bar
	text        *canvas.Text
	textStyle   style.ChartTextStyle
	visible     bool
	colName     fyne.ThemeColorName
	col         color.Color
	legendEntry *interact.LegendEntry
	ser         *Series
}

func emptyProportionPoint(c string, colName fyne.ThemeColorName, ser *Series) (point *proportionPoint) {
	point = &proportionPoint{
		c:       c,
		bar:     elements.NewBar(theme.Color(colName)),
		visible: true,
		ser:     ser,
		colName: colName,
		col:     theme.Color(colName),
	}
	point.legendEntry = interact.NewLegendEntry(c, ser.name, true, colName, point.toggleView)
	if ser.showText {
		point.text = canvas.NewText("", theme.Color(theme.ColorNameForeground))
	}
	return
}

func (point *proportionPoint) toggleView() {
	if point.bar.Visible() {
		point.hide()
	} else {
		point.show()
	}
}

func (point *proportionPoint) hide() {
	if !point.visible {
		return
	}
	point.bar.Hide()
	if point.text != nil {
		point.text.Hide()
	}
	point.visible = false
	point.legendEntry.Hide()
	if point.ser != nil {
		point.ser.pointVisibilityUpdate(-point.val)
	}
}

func (point *proportionPoint) show() {
	if point.visible {
		return
	}
	point.bar.Show()
	if point.text != nil {
		point.text.Show()
	}
	point.visible = true
	point.ser.visible = true
	point.legendEntry.Show()
	if point.ser != nil {
		point.ser.pointVisibilityUpdate(point.val)
	}
}

func (point *proportionPoint) setTextStyle(ts style.ChartTextStyle) {
	point.textStyle = ts
	point.text.TextSize = theme.Size(ts.SizeName)
	point.text.Color = theme.Color(ts.ColorName)
	point.text.TextStyle = ts.TextStyle
	point.text.Refresh()
}

func (point *proportionPoint) refreshTheme() {
	point.col = theme.Color(point.colName)
	point.text.Color = theme.Color(point.textStyle.ColorName)
	point.text.TextSize = theme.Size(point.textStyle.SizeName)
	point.bar.SetColor(point.col)
}

func (point *proportionPoint) cartesianBars(xMin float64, xMax float64, yMin float64,
	yMax float64) (rs []*elements.Bar) {
	if point.valOffset+point.n < xMin || point.valOffset > xMax {
		return
	}
	if point.hOffset+point.height < yMin || point.hOffset > yMax {
		return
	}
	point.bar.N1 = point.valOffset
	point.bar.Val1 = point.hOffset
	point.bar.N2 = point.n + point.valOffset
	point.bar.Val2 = point.hOffset + point.height
	rs = append(rs, point.bar)
	return
}

func (point *proportionPoint) cartesianTexts(xMin float64, xMax float64, yMin float64,
	yMax float64) (ts []elements.Label) {
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
	t := elements.Label{
		N:    point.valOffset + (point.n / 2),
		Val:  point.hOffset + (point.height / 2),
		Text: point.text,
	}
	ts = append(ts, t)
	return
}

func (point *proportionPoint) RasterColorPolar(phi float64, r float64) (col color.Color, useColor bool) {
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	useColor = false
	if !point.visible {
		return
	}
	if phi < point.valOffset ||
		phi > point.valOffset+point.n ||
		r < point.hOffset || r > point.hOffset+point.height {
		return
	}
	useColor = true
	col = point.col
	return
}

func (point *proportionPoint) polarTexts(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ts []elements.Label) {
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
	t := elements.Label{
		N:    point.valOffset + (point.n / 2),
		Val:  point.hOffset + (point.height / 2),
		Text: point.text,
	}
	ts = append(ts, t)
	return
}

type Series struct {
	showText    bool
	data        []*proportionPoint
	tot         float64
	name        string
	visible     bool
	legendEntry *interact.LegendEntry
	textStyle   style.ChartTextStyle
	chart       *BaseChart
	height      float64
	hOffset     float64
}

func EmptyProportionalSeries(name string) (ser *Series) {
	ser = &Series{
		name:     name,
		visible:  true,
		showText: true,
	}
	ser.SetValTextStyle(style.DefaultValueTextStyle())
	ser.legendEntry = interact.NewLegendEntry(name, "", false, theme.ColorNameForeground, ser.toggleView)
	return
}

// Name gives the name of the series
func (ser *Series) Name() (n string) {
	n = ser.name
	return
}

func (ser *Series) BindToChart(ch *BaseChart) (err error) {
	if ser.chart != nil {
		err = errors.New("series is already part of a chart")
		return
	}
	ser.chart = ch
	ch.AddLegendEntry(ser.legendEntry)
	for i := range ser.data {
		ch.AddLegendEntry(ser.data[i].legendEntry)
	}
	return
}

func (ser *Series) Release() {
	if ser.chart != nil {
		for i := range ser.data {
			ser.chart.RemoveLegendEntry(ser.data[i].c, ser.name)
		}
		ser.chart.RemoveLegendEntry(ser.name, "")
	}
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

func (ser *Series) CartesianBars(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []*elements.Bar) {
	for i := range ser.data {
		fs = append(fs, ser.data[i].cartesianBars(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *Series) CartesianTexts(xMin float64, xMax float64, yMin float64,
	yMax float64) (ts []elements.Label) {
	for i := range ser.data {
		ts = append(ts, ser.data[i].cartesianTexts(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *Series) RasterColorPolar(phi float64, r float64) (col color.Color, useColor bool) {
	col = color.RGBA{0x00, 0x00, 0x00, 0x00}
	useColor = false
	if !ser.visible || r < ser.hOffset || r > ser.hOffset+ser.height {
		return
	}
	pCol := col
	for i := range ser.data {
		pCol, useColor = ser.data[i].RasterColorPolar(phi, r)
		if useColor {
			col = pCol
			break
		}
	}
	return
}

func (ser *Series) PolarTexts(phiMin float64, phiMax float64, rMin float64,
	rMax float64) (ts []elements.Label) {
	for i := range ser.data {
		ts = append(ts, ser.data[i].polarTexts(phiMin, phiMax, rMin, rMax)...)
	}
	return
}

func (ser *Series) RefreshTheme() {
	for i := range ser.data {
		ser.data[i].refreshTheme()
	}
}

func (ser *Series) SetValTextStyle(ts style.ChartTextStyle) {
	ser.textStyle = ts
	for i := range ser.data {
		ser.data[i].setTextStyle(ts)
	}
}

// Show makes the Bars of the series visible
func (ser *Series) Show() {
	ser.visible = true
	for i := range ser.data {
		ser.data[i].show()
	}
	ser.legendEntry.Show()
}

// Hide hides the Barss of the series
func (ser *Series) Hide() {
	ser.visible = false
	for i := range ser.data {
		ser.data[i].hide()
	}
	ser.legendEntry.Hide()
}

func (ser *Series) toggleView() {
	if ser.visible {
		ser.Hide()
	} else {
		ser.Show()
	}
	if ser.chart != nil {
		if ser.chart.IsPolar() {
			ser.chart.RasterVisibilityChange()
		}
	}
}

func (ser *Series) pointVisibilityUpdate(totChange float64) {
	ser.tot += totChange
	if ser.chart != nil {
		ser.chart.DataChange()
	}
}

func (ser *Series) SetHeightAndOffset(h float64, hOffset float64) {
	ser.height = h
	ser.hOffset = hOffset
	for i := range ser.data {
		ser.data[i].height = h
		ser.data[i].hOffset = hOffset
	}
}

func (ser *Series) Clear() {
	if ser.chart != nil {
		for i := range ser.data {
			ser.chart.RemoveLegendEntry(ser.data[i].c, ser.name)
		}
	}
	ser.data = []*proportionPoint{}
	if ser.chart != nil {
		ser.chart.DataChange()
	}
}

func (ser *Series) DeleteDataInRange(cat []string) (c int) {
	c = 0
	if len(cat) == 0 {
		return
	}
	finalData := []*proportionPoint{}
	tot := 0.0
	for i := range ser.data {
		del := false
		for j := range cat {
			if ser.data[i].c == cat[j] {
				del = true
				if ser.chart != nil {
					ser.chart.RemoveLegendEntry(ser.data[i].c, ser.name)
				}
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
	if ser.chart != nil {
		ser.chart.DataChange()
	}
	return
}

func (ser *Series) AddData(input []data.ProportionalPoint) (err error) {
	if len(input) == 0 {
		return
	}
	for i := range input {
		if input[i].Val < 0 {
			err = errors.New("invalid data")
			return
		}
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
		pPoint := emptyProportionPoint(input[i].C, input[i].ColName, ser)
		pPoint.setTextStyle(ser.textStyle)
		pPoint.val = input[i].Val
		ser.data = append(ser.data, pPoint)
		ser.tot += pPoint.val
		if ser.chart != nil {
			ser.chart.AddLegendEntry(pPoint.legendEntry)
		}
	}
	if ser.chart != nil {
		ser.chart.DataChange()
	}
	return
}
