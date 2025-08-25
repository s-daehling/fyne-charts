[![Go Reference](https://pkg.go.dev/badge/github.com/s-daehling/fyne-charts.svg)](https://pkg.go.dev/github.com/s-daehling/fyne-charts)
[![Go Report Card](https://goreportcard.com/badge/github.com/s-daehling/fyne-charts)](https://goreportcard.com/report/github.com/s-daehling/fyne-charts)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

# fyne-charts

fyne-charts provides widgets for data visualization for the fyne UI toolkit.
fyne-charts uses native fyne CanvasObjects like canvas.Line, canvas.Circle or canvas.Rectangle.
All objects of the chart widget scale with the widget size.
Properties like the axis range are determined automatically or can be set by the user.

Currently supported are widgets to visualize data in a 2D cartesian or polar plane.
More widgets are planned.

## 2D Cartesian and Polar Plane Widgets

For each type of plane four different widgets are provided to represent 2D data.
Each widget can be used to visualize a random number of data series.
Depending on the widget type different types of data series are possible.

Following list gives an overview of the widgets and supported series types.

||Cartesian Numerical|Cartesian Temporal|Cartesian Categorical|Cartesian Proportional|Polar Angular|Polar Temporal| Polar Categorical|Polar Proportional|
|-|-|-|-|-|-|-|-|-|
|Data Range|x(float64) -> y(float64)|t(time.Time) -> y(float64)|c(string) -> y(float64)|c(string) -> p(float64(>=0))|phi(float64[0,2pi]) -> r(float64(>=0))|t(time.Time) -> r(float64(>=0))|c(string) -> r(float64(>=0))|c(string) -> p(float64(>0))|
|Line|x|x|||x|x|||
|Area|x|x|||x|x|||
|Scatter|x|x|x||x|x|x||
|Lollipop|x|x|x||x|x|x||
|Box|x|x|x||||||
|Candlestick|x|x|||||||
|Bar|x|x|x||x|x|x||
|Stacked Bar|x|x|x||x|x|x||
|Proportion||||x||||x|

![example](docs/example.png "Example")

## Getting started

Include fyne-charts into your project

``
go get github.com/s-daehling/fyne-charts
``

A demo of the provided widgets and series can be found in ``cmd/main.go``

## Documentation

Documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/s-daehling/fyne-charts)

## License

The project is licensed under BSD 3-Clause License.
