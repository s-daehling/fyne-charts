package style

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type LegendLocation string

const (
	LegendLocationTop    LegendLocation = "top"
	LegendLocationBottom LegendLocation = "bottom"
	LegendLocationLeft   LegendLocation = "left"
	LegendLocationRight  LegendLocation = "right"
)

type TextStyle struct {
	Alignment fyne.TextAlign
	ColorName fyne.ThemeColorName
	SizeName  fyne.ThemeSizeName
	TextStyle fyne.TextStyle
}

func DefaultTitleStyle() (titleStyle TextStyle) {
	titleStyle.Alignment = fyne.TextAlignCenter
	titleStyle.ColorName = theme.ColorNameForeground
	titleStyle.SizeName = theme.SizeNameHeadingText
	titleStyle.TextStyle = fyne.TextStyle{}
	return
}

func DefaultAxisLabelStyle() (axisLabelStyle TextStyle) {
	axisLabelStyle.Alignment = fyne.TextAlignCenter
	axisLabelStyle.ColorName = theme.ColorNameForeground
	axisLabelStyle.SizeName = theme.SizeNameSubHeadingText
	axisLabelStyle.TextStyle = fyne.TextStyle{}
	return
}

func DefaultLegendTextStyle() (legendTextStyle TextStyle) {
	legendTextStyle.Alignment = fyne.TextAlignLeading
	legendTextStyle.ColorName = theme.ColorNameForeground
	legendTextStyle.SizeName = theme.SizeNameText
	legendTextStyle.TextStyle = fyne.TextStyle{}
	return
}

func DefaultValueTextStyle() (valueTextStyle TextStyle) {
	valueTextStyle.Alignment = fyne.TextAlignCenter
	valueTextStyle.ColorName = theme.ColorNameForeground
	valueTextStyle.SizeName = theme.SizeNameText
	valueTextStyle.TextStyle = fyne.TextStyle{}
	return
}

type AxisStyle struct {
	LineColorName        fyne.ThemeColorName
	LineWidth            float32
	LineShowArrow        bool
	SupportLineColorName fyne.ThemeColorName
	SupportLineWidth     float32
	TickColorName        fyne.ThemeColorName
	TickSizeName         fyne.ThemeSizeName
	TickTextStyle        fyne.TextStyle
}

func DefaultAxisStyle() (axisStyle AxisStyle) {
	axisStyle.LineColorName = theme.ColorNameForeground
	axisStyle.LineWidth = 1
	axisStyle.LineShowArrow = true
	axisStyle.SupportLineColorName = theme.ColorNameShadow
	axisStyle.SupportLineWidth = 1
	axisStyle.TickColorName = theme.ColorNameForeground
	axisStyle.TickSizeName = theme.SizeNameText
	axisStyle.TickTextStyle = fyne.TextStyle{}
	return
}
