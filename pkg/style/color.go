package style

import "fyne.io/fyne/v2"

type ColorPalette struct {
	i        int
	colNames []fyne.ThemeColorName
}

func (p *ColorPalette) Next() (colName fyne.ThemeColorName) {
	colName = p.colNames[p.i]
	p.i++
	if p.i == len(p.colNames) {
		p.i = 0
	}
	return
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
