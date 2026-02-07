# Coloring of series elements using the color palette theme

The concepts of series creation and adding to charts will be explained using `coord.CartesianNumericalChart` and `coord.NumericalPointSeries` as an example.
Most explanations directly apply to other chart and series types as well.
Important differences are highlighted in the individual sections.

## Purpose of the color palette theme

Series elements are colored by the theme according to the series' color name.
The default theme offers only a limited amount of colors that might not be sufficient, e.g., if more colors are needed.

The color palette theme provided in the `style` package extends an existing theme (the fyne default theme or another custom theme).
It can interpret color names that are generated dynamically by a `style.ColorPalette`.
A color palettes requires one or more color names as a basis and generates additional colors from that.
Different algorithm for color generation exist and are described below.

## Loading the theme and creating color palettes

Loading the color palette theme requires another theme to be extended.
In the following example the fyne default theme is used:

```go
myApp := app.New()
myApp.Settings().SetTheme(style.NewColorPaletteTheme(theme.DefaultTheme()))
```

Now we can create a color palette.
As an example, a color palette containing one light and one dark shade of the primary color is created:

```go
pal := style.NewPaletteLightDark(theme.ColorNamePrimary)
```

You can rotate through the color palette and receive the next color name by

```go
colName := pal.Next()
```

Rotating in reversed order is done by

```go
colName := pal.Previous()
```

Colors are generated using the [HCL color space](https://en.wikipedia.org/wiki/HCL_color_space).
Different color palette generators for either varying the lightness (L) or hue (H) exist within `style`

Color palettes with varying lightness:

|Generator Function|Description|Number of Colors|
|-|-|-|
|`NewPaletteLightDark(base fyne.ThemeColorName)`|light and dark shade of base|2|
|`NewPaletteLightMediumDark(base fyne.ThemeColorName)`|light, medium and dark shade of base|3|
|`NewPaletteLightDarkSet(base []fyne.ThemeColorName)`|light and dark shade of each base|2*len(base)|
|`NewPaletteLightMediumDarkSet(base []fyne.ThemeColorName)`|light, medium and dark shade of each base|3*len(base)|
|`NewPaletteDivergentLightMediumDarkSet(base1, base2 fyne.ThemeColorName)`|light, medium and dark shade of each base with neutral in the middle|7|

Color palettes with varying hue:

|Generator Function|Description|Number of Colors|
|-|-|-|
|`NewPaletteComplementary(base fyne.ThemeColorName)`|base and complementary in HCL|2|
|`NewPaletteTriadic(base fyne.ThemeColorName)`|base and two colors with distance of 120° in HCL|3|
|`NewPaletteQuadratic(base fyne.ThemeColorName)`|base and three colors with distance of 90° in HCL|4|
|`NewPaletteHexadic(base fyne.ThemeColorName)`|base and fice colors with distance of 60° in HCL|6|

It is also possible to create combined palettes.
The following example first creates a palette with three different color names using `NewPaletteTriadic`.
Subsequenty a new palette with a light and a dark shade from each of the three colors is created:

```go
pal := style.NewPaletteTriadic(theme.ColorNamePrimary)
pal = style.NewPaletteLightDarkSet(pal.Names())
```

The resulting palette has six color names.

## Applying color palettes to charts

Putting it all together we can create charts using dynamically created color palettes.
Here is a full example of a pie chart with six categories.

```go
package main

import (
    "strconv"

    "github.com/s-daehling/fyne-charts/pkg/data"
    "github.com/s-daehling/fyne-charts/pkg/prop"
    "github.com/s-daehling/fyne-charts/pkg/style"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/theme"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Color Palette Theme Example")
    myApp.Settings().SetTheme(style.NewColorPaletteTheme(theme.DefaultTheme()))

    pal := style.NewPaletteComplementary(theme.ColorNamePrimary)
    pal = style.NewPaletteLightMediumDarkSet(pal.Names())

    chart := prop.NewPieChart("Example Chart")
    serData := make([]data.ProportionalPoint, 0)
    for i := range 6 {
        serData = append(serData, data.ProportionalPoint{
            Val:     1,
            C:       "Cat" + strconv.Itoa(i),
            ColName: pal.Next(),
        })
    }
    ps, err := prop.NewSeries("Proportion", serData)
    if err != nil {
        panic(err)
    }
    err = chart.AddSeries(ps)
    if err != nil {
        panic(err)
    }

    myWindow.SetContent(chart)
    myWindow.Resize(fyne.NewSize(500, 300))
    myWindow.ShowAndRun()
}
```
