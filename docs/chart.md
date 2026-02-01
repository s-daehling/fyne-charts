# Chart creation and configuration

The concepts of chart creation and configuration will be explained using `coord.CartesianNumericalChart` as an example.
Most explanations directly apply to other chart types as well.
Important differences are highlighted in the individual sections.

## Creating a new chart

To create a new chart, simply call the `New` function of the required chart type.
For example:

```go
chart := coord.NewCartesianNumericalChart("Title of Example Chart")
```

The `New` function requires the title of the chart as only argument.
It returns a `*coord.CartesianNumericalChart` object.
All chart types in fyne-charts satisfy the `fyne.Widget` interface and can be used like any other widget in fyne.

## Setting title and axis labels

The title of the chart is set during chart creation.
It can also be changed later:

```go
chart.SetTitle("New Chart Title")
```

Axis titles can be set by

```go
chart.SetXAxisLabel("X Axis")
chart.SetYAxisLabel("Y Axis")
```

Note that in other chart types, the axes have different names.
The following table gives an overview.

|chart type|from axis|to axis|
|-|-|-|
|cartesian numerical|XAxis|YAxis|
|cartesian temporal|TAxis|YAxis|
|cartesian categorical|CAxis|YAxis|
|polar numerical|PhiAxis|RAxis|
|polar temporal|TAxis|RAxis|
|polar categorical|CAxis|RAxis|

Proportional charts don't have axes at all.

In case the chart title and/or axes labels are set to an empty string, the space is used to expand the chart itself.

## Legend

By default a legend containing all series of the chart is displayed on the right side.
You can hide the legend by calling

```go
chart.HideLegend()
```

Getting the legend back is done like this

```go
chart.ShowLegend()
```

The legend shows a square next to each series entry colored in the series' color.
Clicking this square hides all elements of the series.
To indicate that a series is hidden, the square turns into a cricle.
By clicking again, the series elements are shown again and the circle turns back into a square.

Hiding one or multiple series in a proportional chart leads to a recalculation of the proportions.
Proportions are always calculated with respect only to visible parts.

## Styling of chart elements

### Title

Text elements in fyne-charts can be styled using the `style.ChartTextStyle` struct.
This provides four parameters:

- `Alignment fyne.TextAlign` (left, center, right)
- `ColorName fyne.ThemeColorName` (color of the text element)
- `SizeName  fyne.ThemeSizeName` (size of the text element)
- `TextStyle fyne.TextStyle` (bold, italic, etc.)

When you want to change the style of a text element, like the title, you should first create an instance of this struct filled with the default values.
This is done by using the `style.Default` functions.
After that you can change the parameters you like to customize and apply the changes to the chart.
Changing the alignment of the title to left-aligned is done like this:

```go
ts := style.DefaultTitleStyle()
ts.Alignment = fyne.TextAlignLeading
chart.SetTitleStyle(ts)
```

Warning: You should avoid creating the style struct like in the example below.

```go
ts := style.ChartTextStyle{
    Alignment: fyne.TextAlignLeading,
}
```

This will set all other parameters to their zero values (not their default values).
As a result, parameters that are supposed to remain at their default values, are changed as well.
Instead, initialize the struct as explained above.

### Axes

The appearance of the axis label and the axis itself can be customized.
For the axis label the `style.ChartTextStyle` struct is used and the same considerations as for the title apply.
The axis is styled using the `style.AxisStyle` struct:

- `LineColorName fyne.ThemeColorName` (color of the axis and tick lines)
- `LineWidth float32` (width of the axis and tick lines)
- `LineShowArrow bool` (show or hide the arrow at the end of the axis line)
- `SupportLineColorName fyne.ThemeColorName` (color of the support line at ticks)
- `SupportLineWidth float32` (width of the support line at ticks)
- `TickColorName fyne.ThemeColorName` (color of the tick text elements)
- `TickSizeName fyne.ThemeSizeName` (size of the tick text elements)
- `TickTextStyle fyne.TextStyle` (bold, italic, etc.)

Similar as explained for styling of the title, there are `style.Default` functions to initialize the struct with default values.
Making the axis label smaller, using a different color for the axis line and changing the tick style to italic is done like this:

```go
ls := style.DefaultAxisLabelStyle()
ls.SizeName = theme.SizeNameText
as := style.DefaultAxisStyle()
as.LineColorName = theme.ColorNamePrimary
as.TickTextStyle = fyne.TextStyle{
    Italic: true,
}
chart.SetXAxisStyle(ls, as)
```

Since proportional charts don't have axes, axis styling is not available for this type of chart.

### Legend

The appearance of the legend can be customized in three aspects:

- location of the legend (using `style.LegendLocation`),
- style of the labels (using again `style.ChartTextStyle`) and
- interactiveness (enable/disable the clickable color indicators next to the series entries)

For example, the legend can be moved to the bottom of the chart with the following code (with all other parameters kept at their default values).

```go
chart.SetLegendStyle(style.LegendLocationBottom, style.DefaultLegendTextStyle(), true)
```

## Automatic or manual data range and axis ticks (only `coord`)

By default the range of the axes (minimum and maximum value) is determined automatically.
This is done such that all data points are visible.
That means that the data range is set to the minimum required value to display all series.
Ticks are determined automatically as well.
The amount of ticks is adjusted dynamically to the available space.
The origin (i.e. the point in where both axes cross each other) is set to (0,0) if it lies within the axis range or to a minimum or maximum value otherwise.

These auto calculations can be deactivated and replaced with fixed values.
Setting a user defined range for the x-axis from 0 to 20 works like this:

```go
err := chart.SetXRange(0, 20)
```

An error is returned if min > max or if a previously user-defined origin is outside the requested range.

Similarly, the origin can be set to a fixed value as well, e.g. to x=5 and y=10.

```go
err = chart.SetOrigin(5,10)
```

An error is returned if the origin lies outside a previously user-defined range.

Setting user-defined ticks is done using `data.NumericalTick`.
In the following example only one tick is set at position x=10 with a support line.

```go
chart.SetXTicks([]data.NumericalTick{
    data.NumericalTick{N: 10, SupportLine: true},
})
```

Analogous methods exist for all chart types.
However, in the case of categorical charts (both cartesian and polar) no method for setting ticks of the categorical axis exist.
Only the range can be user-set by providing a slice with all categories to be shown.

Once the range, origin or ticks have been user-defined, they are no longer determined automatically but will remain at the user defined values.
If you want to return to an automatic calculation at a later time you can do that by calling

```go
chart.SetAutoXRange()
chart.SetAutoOrigin()
chart.SetAutoXTicks(true)
```

This will delete the user defined values and use the automatically calculated values instead.

## Next steps

Learn about how to [create data series and add them to charts](series.md).
