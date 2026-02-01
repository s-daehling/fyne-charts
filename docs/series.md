# Series creation and adding to charts

The concepts of series creation and adding to charts will be explained using `coord.CartesianNumericalChart` and `coord.NumericalPointSeries` as an example.
Most explanations directly apply to other chart and series types as well.
Important differences are highlighted in the individual sections.

## Creating a new series

A new series is created using the `New` functions in `coord` or `prop`.
A new series requires a name, a color and a slice with data points.
Note that wihtin a single chart, series names must be unique.
Creating a new `coord.NumericalPointSeries` can look like this:

```go
numData := []data.NumericalPoint {
    {N: 5, Val: 10},
    {N: 10, Val: 15},
    {N: 15, Val: 20},
}
nps, err := coord.NewNumericalPointSeries("Example Series", theme.ColorNamePrimary, numData)
```

An error is returned if data points are invalid.
Some series types restrict the allowed data ranges.
More information on restrictions can be found below.

## Adding and removing a series to/from a chart

A point series can be visualized in different ways: as scatter, line, area, lollipop or bar plot.
In our example, we want to visualize the series as a scatter plot.

```go
err = chart.AddScatterSeries(nps)
```

An error is returned if another series with the same name already exists in the chart.
Series names must be unique within one chart.

After the series has been added to a chart, it can not be added to another chart.
Trying to do that will result in an error as well.

Removing the series from the chart is done by:

```go
chart.RemoveSeries("Example Series")
```

Now the series could be added to another chart.
This shows, that series live independent from charts.
They can be added and removed from one chart and then reused for another chart.
However, they can only be added to one chart at a time.

### Restrictions

Not every series type can be added to any chart type.
The following table shows which series types are available for the different chart types in `coord`

|(Cartesian / Polar)|Numerical|Temporal|Categorical|
|-|-|-|-|
|Line|y / y|y / y|n / n|
|Area|y / y|y / y|n / n|
|Scatter|y / y|y / y|y / y|
|Lollipop|y / y|y / y|y / y|
|Box|y / n|y / n|y / n|
|Candlestick|y / n|y / n|n / n|
|Bar|y / y|y / y|y / y|
|Stacked Bar|n / n|n / n|y / y|

Moreover, the data range of a series is limited with respect to the data that can be displayed in a certain chart type.
The following table gives an overview of data ranges in all chart types

||Numerical|Temporal|Categorical|
|-|-|-|-|
|Cartesian x-axis|any valid float64|any valid time.Time|any valid string|
|Cartesian y-axis|any valid float64|any valid float64|any valid float64|
|Polar phi-axis|0 <= phi <= 2pi|any valid time.Time|any valid string|
|Polar r-axis|r >= 0|r >= 0|r >= 0|

Note that the y-axis in cartesian charts can contain all valid float64.
In a polar chart the r-axis only allows values equal or greater than zero.
That means, that moving a series from a cartesian to a polar chart can result in an error, if the series contains data points with `Val` < 0.

In `prop` only one series type exists.
This can be added to both chart types.
Data points are restricted to values equal or greater than zero.

## Styling of a series

Depending on the series type different methods for customizing the style are available.
During series creation a color for the series must be provided.
The color choice can be altered afterwards:

```go
nps.SetColor(theme.ColorNameSuccess)
```

In a point series you can change the line width or dot size.

```go
nps.SetLineWidth(2)
nps.SetDotSize(2)
```

Note that these changes will only have an effect if the series is displayed in a way that shows lines or dots.
In our example, changing the line width will not change the appearance of the series, since it is displayed as a scatter plot (dots only).

Other series types can provide different methods.

## Dynamic data manipulation

The data contained in a series can be manipulated during its lifetime.
All series types provide three methods for this.

Adding data to an existing series can be done by

```go
err = nps.AddData([]data.NumericalPoint{
    {N: 20, Val: 25},
    {N: 25, Val: 30},
})
```

The same restrictions as explained above are applied.
If data is invalid, an error is returned.

Data can also be deleted.
To do this the range must be specified (in our example the range refers to the x value).

```go
c := nps.DeleteDataInRange(10, 20)
```

c is the number of data points that were removed by the method.

Finally you can completely reset the series and delete all data points:

```go
nps.Clear()
```

## Next steps

Learn how to use the custom theme of fyne-charts for [series coloring](coloring.md)
