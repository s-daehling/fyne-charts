package style

import (
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/s-daehling/fyne-charts/internal/style"
)

func NewColorPaletteTheme(defTheme fyne.Theme) (th fyne.Theme) {
	th = &style.PaletteTheme{
		DefTheme: defTheme,
	}
	return
}

type ColorPalette struct {
	i        int
	colNames []fyne.ThemeColorName
}

func (p *ColorPalette) Next() (colName fyne.ThemeColorName) {
	if len(p.colNames) == 0 {
		colName = theme.ColorNameForeground
		return
	}
	colName = p.colNames[p.i]
	p.i++
	if p.i == len(p.colNames) {
		p.i = 0
	}
	return
}

func (p *ColorPalette) Add(colName fyne.ThemeColorName) {
	p.colNames = append(p.colNames, colName)
}

func (p *ColorPalette) Remove(colName fyne.ThemeColorName) {
	p.colNames = slices.DeleteFunc(p.colNames, func(e fyne.ThemeColorName) bool {
		return e == colName
	})
}

func (p *ColorPalette) Names() (n []fyne.ThemeColorName) {
	n = append(n, p.colNames...)

	return
}

func NewPaletteLightDark(base fyne.ThemeColorName) (p *ColorPalette) {
	p = &ColorPalette{
		i:        0,
		colNames: style.NewPaletteLightDark(base),
	}
	return
}

func NewPaletteLightDarkSet(base []fyne.ThemeColorName) (p *ColorPalette) {
	p = &ColorPalette{
		i:        0,
		colNames: style.NewPaletteLightDarkSet(base),
	}
	return
}

func NewPaletteLightMediumDark(base fyne.ThemeColorName) (p *ColorPalette) {
	p = &ColorPalette{
		i:        0,
		colNames: style.NewPaletteLightMediumDark(base),
	}
	return
}

func NewPaletteLightMediumDarkSet(base []fyne.ThemeColorName) (p *ColorPalette) {
	p = &ColorPalette{
		i:        0,
		colNames: style.NewPaletteLightMediumDarkSet(base),
	}
	return
}

func NewPaletteDivLightMediumDarkSet(base []fyne.ThemeColorName) (p *ColorPalette) {
	p = &ColorPalette{
		i:        0,
		colNames: style.NewPaletteDivLightMediumDarkSet(base),
	}
	return
}

func NewPaletteEquidistantHue(base fyne.ThemeColorName, numCols int) (p *ColorPalette) {
	p = &ColorPalette{
		i:        0,
		colNames: style.NewPaletteEquidistantHue(base, numCols),
	}
	return
}

func NewPaletteComplementary(base fyne.ThemeColorName) (p *ColorPalette) {
	p = &ColorPalette{
		i:        0,
		colNames: style.NewPaletteComplementary(base),
	}
	return
}

func NewPaletteTriadic(base fyne.ThemeColorName) (p *ColorPalette) {
	p = &ColorPalette{
		i:        0,
		colNames: style.NewPaletteTriadic(base),
	}
	return
}

func NewPaletteTetradic(base fyne.ThemeColorName) (p *ColorPalette) {
	p = &ColorPalette{
		i:        0,
		colNames: style.NewPaletteTetradic(base),
	}
	return
}
