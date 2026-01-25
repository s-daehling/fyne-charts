package style

import (
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

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

func NewPaletteLightDark(base fyne.ThemeColorName) (p *ColorPalette) {
	p = &ColorPalette{
		i: 0,
		colNames: []fyne.ThemeColorName{
			fyne.ThemeColorName("_ptStart_LightDark-" + string(base) + "-0_ptEnd_"),
			fyne.ThemeColorName("_ptStart_LightDark-" + string(base) + "-1_ptEnd_"),
		},
	}
	return
}

func NewPaletteLightMediumDark(base fyne.ThemeColorName) (p *ColorPalette) {
	p = &ColorPalette{
		i: 0,
		colNames: []fyne.ThemeColorName{
			fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base) + "-0_ptEnd_"),
			fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base) + "-1_ptEnd_"),
			fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base) + "-2_ptEnd_"),
		},
	}
	return
}
