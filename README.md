[![Go Reference](https://pkg.go.dev/badge/github.com/s-daehling/fyne-charts.svg)](https://pkg.go.dev/github.com/s-daehling/fyne-charts)
[![Go Report Card](https://goreportcard.com/badge/github.com/s-daehling/fyne-charts)](https://goreportcard.com/report/github.com/s-daehling/fyne-charts)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

# fyne-charts

fyne-charts provides widgets for data visualization for the [Fyne UI toolkit](https://fyne.io/).
It does not rely on other chart libraries.
Instead widgets are composed of native Fyne-CanvasObjects like canvas.Line, canvas.Circle or canvas.Rectangle.
They adapt to the available size, can be updated dynamically and provide some user interaction.

Currently supported are widgets to visualize data in a 2D coordinate systems (cartesian or polar) and proportional data.
More widgets are planned.

**Note: fyne-charts is in an early development stage. The API might still change in the future.**

## Prerequesites and Getting Started

fyne-charts is an extension of Fyne.
Before starting with fyne-charts, make sure Fyne is correctly installed.
You can find more information on how to get started with Fyne on their [website](https://docs.fyne.io/started/).

After setting up Fyne, inlcude fyne-charts in your project:

```sh
go get github.com/s-daehling/fyne-charts
```

Now you can create a simple chart with the following code example:

```go
package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/theme"

    "github.com/s-daehling/fyne-charts/pkg/coord"
    "github.com/s-daehling/fyne-charts/pkg/data"
)

func main() {
    a := app.New()
    w := a.NewWindow("fyne-charts demo")

    chart := coord.NewCartesianCategoricalChart("Example Chart")
    chart.SetCAxisLabel("Categories")
    chart.SetYAxisLabel("Value")

    cps, err := coord.NewCategoricalPointSeries("Bar Series", theme.ColorNamePrimary, []data.CategoricalPoint{{
        C:   "One",
        Val: 15,
    }, {
        C:   "Two",
        Val: 30,
    }, {
        C:   "Three",
        Val: 45,
    }})
    if err != nil {
        panic(err)
    }
    err = chart.AddBarSeries(cps)
    if err != nil {
        panic(err)
    }

    w.SetContent(chart)
    w.Resize(fyne.NewSize(500, 300))
    w.ShowAndRun()
}
```

Demos of fyne-charts can be found in ``cmd/``

## Widgets

### Two-Dimensional Coordinate System Charts

![example](docs/coord-example.png "Example of Coordinate System Charts")

fyne-charts provides six widgets for visualization of two-dimensional data in a coordinate system located in `github.com/s-daehling/fyne-charts/pkg/coord`.
They differ in the coordinate system that is used and in the kind of data that is mapped.

Two possible coordinate systems:

* Cartesian (mapping from a x-axis to an orthogonal y-axis)
* Polar (mapping from a phi-axis to a radial r-axis)

In both coordinate systems different data types can be used for the "from-axis" (x or phi):

* Numerical: data that is represented by a number (implemented as float64),
* Temporal: data that is represented by a timestamp (implemented as time.Time) and
* Categorical: data that is represented by a name (implemented as string).

### Proportional Data Charts

![example](docs/prop-example.png "Example of Proportional Data Charts")

Analogous to the coordinate system charts, also proportional data charts can be drawn in two ways:

* Horizontal Bars
* Pie/Doughnut

The two corresponding chart widgets are provided by the package `github.com/s-daehling/fyne-charts/pkg/prop`:

## Documentation

The following tutorials help you to get started:

- [Chart creation and configuration](docs/chart.md)
- [Series creation and adding to charts](docs/series.md)
- [Coloring of series elements using the color palette theme](docs/coloring.md)

Code documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/s-daehling/fyne-charts)

## License

The project is licensed under BSD 3-Clause License.
