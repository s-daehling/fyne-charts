package prop

import (
	"errors"
	"image/color"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/s-daehling/fyne-charts/internal/elements"
	"github.com/s-daehling/fyne-charts/internal/interact"

	intstyle "github.com/s-daehling/fyne-charts/internal/style"
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
	label       *elements.Label
	visible     bool
	colName     fyne.ThemeColorName
	col         color.Color
	colFaded    color.Color
	legendEntry *interact.LegendEntry
	highlighted bool
	isFaded     bool
	ser         *Series
}

func emptyProportionPoint(c string, colName fyne.ThemeColorName, labelStyle style.ValueLabelStyle, ser *Series) (point *proportionPoint) {
	point = &proportionPoint{
		label:       elements.NewLabel(labelStyle),
		c:           c,
		visible:     true,
		ser:         ser,
		colName:     colName,
		col:         theme.Color(colName),
		highlighted: false,
		isFaded:     false,
	}
	point.colFaded = intstyle.MakeFaded(point.col)
	point.bar = elements.NewBar(theme.Color(colName), point.highlight, point.unhighlight)
	point.legendEntry = interact.NewLegendEntry(c, ser.name, true, point.col, point.toggleView, point.highlight, point.unhighlight)
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
	point.visible = true
	point.ser.visible = true
	point.legendEntry.Show()
	if point.ser != nil {
		point.ser.pointVisibilityUpdate(point.val)
	}
}

func (point *proportionPoint) refreshTheme() {
	point.col = theme.Color(point.colName)
	point.colFaded = intstyle.MakeFaded(point.col)
	col := point.col
	if point.isFaded {
		col = point.colFaded
	}
	point.bar.SetColor(col)
	point.legendEntry.SetColor(col)
}

func (point *proportionPoint) setLabelStyle(labelStyle style.ValueLabelStyle) {
	point.label.SetStyle(labelStyle)
}

func (point *proportionPoint) checkIfHoveringPolarBar(phi float64, r float64) {
	if phi > point.valOffset &&
		phi < point.valOffset+point.n &&
		r > point.hOffset && r < point.hOffset+point.height {
		if !point.highlighted && !point.isFaded {
			point.highlight()
		}
	} else if point.highlighted {
		point.unhighlight()
	}
}

func (point *proportionPoint) highlight() {
	point.highlighted = true
	point.ser.highlight()
}

func (point *proportionPoint) unhighlight() {
	point.highlighted = false
	point.ser.unhighlight()
}

func (point *proportionPoint) fadeUnhighlighted() {
	if point.highlighted {
		return
	}
	point.isFaded = true
}

func (point *proportionPoint) unFade() {
	point.isFaded = false
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

func (point *proportionPoint) cartesianLabels(xMin float64, xMax float64, yMin float64,
	yMax float64, labelIfHighlighted bool, permLabel bool) (ls []*elements.Label) {
	if !point.visible || ((!point.highlighted || !labelIfHighlighted) && !permLabel) {
		return
	}
	if point.valOffset+point.n < xMin || point.valOffset > xMax {
		return
	}
	if point.hOffset+point.height < yMin || point.hOffset > yMax {
		return
	}
	point.label.SetText(strconv.FormatFloat(point.n, 'f', 0, 64) + "%")
	point.label.N = point.valOffset + (point.n / 2)
	point.label.Val = point.hOffset + (point.height / 2)
	ls = append(ls, point.label)
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
	if point.isFaded {
		col = point.colFaded
	} else {
		col = point.col
	}
	return
}

func (point *proportionPoint) polarLabels(phiMin float64, phiMax float64, rMin float64,
	rMax float64, labelIfHighlighted bool, permLabel bool) (ls []*elements.Label) {
	if !point.visible || ((!point.highlighted || !labelIfHighlighted) && !permLabel) {
		return
	}
	if point.valOffset+point.n < phiMin || point.valOffset > phiMax {
		return
	}
	if point.hOffset+point.height < rMin || point.hOffset > rMax {
		return
	}
	point.label.SetText(strconv.FormatFloat(100*(point.n/(2*math.Pi)), 'f', 0, 64) + "%")
	point.label.N = point.valOffset + (point.n / 2)
	point.label.Val = point.hOffset + (point.height / 2)
	ls = append(ls, point.label)
	return
}

type Series struct {
	showText       bool
	data           []*proportionPoint
	tot            float64
	name           string
	visible        bool
	legendEntry    *interact.LegendEntry
	labelStyle     style.ValueLabelStyle
	permanentLabel bool
	chart          *BaseChart
	height         float64
	hOffset        float64
}

func EmptyProportionalSeries(name string) (ser *Series) {
	ser = &Series{
		name:           name,
		visible:        true,
		showText:       true,
		permanentLabel: false,
		labelStyle:     style.DefaultValueLabelStyle(),
	}
	ser.legendEntry = interact.NewLegendEntry(name, "", false, theme.Color(theme.ColorNameForeground), ser.toggleView, nil, nil)
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

func (ser *Series) Bars(xMin float64, xMax float64, yMin float64,
	yMax float64) (fs []*elements.Bar) {
	if ser.chart.IsPolar() {
		return
	}
	for i := range ser.data {
		fs = append(fs, ser.data[i].cartesianBars(xMin, xMax, yMin, yMax)...)
	}
	return
}

func (ser *Series) Labels(nMin float64, nMax float64, valMin float64,
	valMax float64, labelIfHighlighted bool) (ls []*elements.Label) {
	for i := range ser.data {
		if ser.chart.IsPolar() {
			ls = append(ls, ser.data[i].polarLabels(nMin, nMax, valMin, valMax, labelIfHighlighted, ser.permanentLabel)...)
		} else {
			ls = append(ls, ser.data[i].cartesianLabels(nMin, nMax, valMin, valMax, labelIfHighlighted, ser.permanentLabel)...)
		}
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

func (ser *Series) RefreshTheme() {
	for i := range ser.data {
		ser.data[i].refreshTheme()
		ser.data[i].setLabelStyle(ser.labelStyle)
	}
}

func (ser *Series) Hover(n float64, val float64) {
	if ser.chart == nil {
		return
	}
	if !ser.chart.IsPolar() {
		return
	}
	for i := range ser.data {
		ser.data[i].checkIfHoveringPolarBar(n, val)
	}
}

func (ser *Series) SetValueLabelStyle(permLabel bool, labelStyle style.ValueLabelStyle) {
	ser.permanentLabel = permLabel
	ser.labelStyle = labelStyle
	for i := range ser.data {
		ser.data[i].setLabelStyle(labelStyle)
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

func (ser *Series) fadeUnhighlighted() {
	for i := range ser.data {
		ser.data[i].fadeUnhighlighted()
	}
}

func (ser *Series) unFade() {
	for i := range ser.data {
		ser.data[i].unFade()
	}
}

func (ser *Series) highlight() {
	if ser.chart != nil {
		ser.chart.Highlight()
	}
}

func (ser *Series) unhighlight() {
	if ser.chart != nil {
		ser.chart.Unhighlight()
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
		pPoint := emptyProportionPoint(input[i].C, input[i].ColName, ser.labelStyle, ser)
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
