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

// ChartTextStyle for styling of various text elements in a chart
type ChartTextStyle struct {
	Alignment fyne.TextAlign
	ColorName fyne.ThemeColorName
	SizeName  fyne.ThemeSizeName
	TextStyle fyne.TextStyle
}

// DefaultTitleStyle returns the default style of the chart title
func DefaultTitleStyle() (titleStyle ChartTextStyle) {
	titleStyle.Alignment = fyne.TextAlignCenter
	titleStyle.ColorName = theme.ColorNameForeground
	titleStyle.SizeName = theme.SizeNameHeadingText
	titleStyle.TextStyle = fyne.TextStyle{}
	return
}

// DefaultAxisLabelStyle returns the default style of axis labels
func DefaultAxisLabelStyle() (axisLabelStyle ChartTextStyle) {
	axisLabelStyle.Alignment = fyne.TextAlignCenter
	axisLabelStyle.ColorName = theme.ColorNameForeground
	axisLabelStyle.SizeName = theme.SizeNameSubHeadingText
	axisLabelStyle.TextStyle = fyne.TextStyle{}
	return
}

// DefaultLegendTextStyle returns the default style of legend entries
func DefaultLegendTextStyle() (legendTextStyle ChartTextStyle) {
	legendTextStyle.Alignment = fyne.TextAlignLeading
	legendTextStyle.ColorName = theme.ColorNameForeground
	legendTextStyle.SizeName = theme.SizeNameText
	legendTextStyle.TextStyle = fyne.TextStyle{}
	return
}

// AxisStyle for styling axes
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

// DefaultAxisStyle returns the default style of axes
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

// ValueLabelStyle for styling of labels showing the value of data points
type ValueLabelStyle struct {
	ValueTextStyle      ChartTextStyle
	BackgroundColorName fyne.ThemeColorName
	StrokeColorName     fyne.ThemeColorName
	StrokeWidth         float32
}

// DefaultValueLabelStyle returns the default style of value labels
func DefaultValueLabelStyle() (labelStyle ValueLabelStyle) {
	labelStyle.ValueTextStyle.Alignment = fyne.TextAlignCenter
	labelStyle.ValueTextStyle.ColorName = theme.ColorNameForeground
	labelStyle.ValueTextStyle.SizeName = theme.SizeNameText
	labelStyle.ValueTextStyle.TextStyle = fyne.TextStyle{}
	labelStyle.BackgroundColorName = theme.ColorNameBackground
	labelStyle.StrokeColorName = theme.ColorNameForeground
	labelStyle.StrokeWidth = 1
	return
}
