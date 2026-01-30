package style

import (
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type ThemeSettings struct {
	Variant fyne.ThemeVariant
	MinL    float64
	MaxL    float64
}

type PaletteTheme struct {
	defTheme      fyne.Theme
	settingsLight ThemeSettings
	settingsDark  ThemeSettings
}

var _ fyne.Theme = (*PaletteTheme)(nil)

func NewColorPaletteTheme(defTheme fyne.Theme) (th fyne.Theme) {
	pth := &PaletteTheme{
		defTheme: defTheme,
		settingsLight: ThemeSettings{
			Variant: theme.VariantLight,
		},
		settingsDark: ThemeSettings{
			Variant: theme.VariantDark,
		},
	}
	pth.settingsLight.MinL, pth.settingsLight.MaxL = LightnessRange(
		defTheme.Color(theme.ColorNamePrimary, theme.VariantLight),
		defTheme.Color(theme.ColorNameBackground, theme.VariantLight),
		defTheme.Color(theme.ColorNameForeground, theme.VariantLight),
		2.5, 1.5)
	pth.settingsDark.MinL, pth.settingsDark.MaxL = LightnessRange(
		defTheme.Color(theme.ColorNamePrimary, theme.VariantDark),
		defTheme.Color(theme.ColorNameBackground, theme.VariantDark),
		defTheme.Color(theme.ColorNameForeground, theme.VariantDark),
		2.5, 1.5)
	th = pth
	return
}

func (th PaletteTheme) Color(colName fyne.ThemeColorName, vari fyne.ThemeVariant) (col color.Color) {
	col = color.Alpha16{}
	if strings.HasPrefix(string(colName), "_ptStart_") && strings.HasSuffix(string(colName), "_ptEnd_") {
		params := parseColName(string(colName)[len("_ptStart_") : len(colName)-len("_ptEnd_")])
		if len(params) == 0 {
			return
		}
		switch params[0] {
		case "LightMediumDark":
			if len(params) != 3 {
				return
			}
			step, err := strconv.ParseInt(params[2], 10, 0)
			if err != nil {
				return
			}
			settings := th.settingsDark
			if vari == theme.VariantLight {
				settings = th.settingsLight
			}
			col = LightMediumDark(fyne.ThemeColorName(params[1]), int(step), settings)
			return
		case "Neutral":
			if len(params) != 1 {
				return
			}
			settings := th.settingsDark
			if vari == theme.VariantLight {
				settings = th.settingsLight
			}
			col = Neutral(settings)
			return
		case "EquidistantHue":
			if len(params) != 4 {
				return
			}
			step, err := strconv.ParseInt(params[2], 10, 0)
			if err != nil {
				return
			}
			totStep, err := strconv.ParseInt(params[3], 10, 0)
			if err != nil {
				return
			}
			col = EquidistantHue(fyne.ThemeColorName(params[1]), int(step), int(totStep))
			return
			// case "AnalogousFive":
			// 	if len(params) != 3 {
			// 		return
			// 	}
			// 	step, err := strconv.ParseInt(params[2], 10, 0)
			// 	if err != nil {
			// 		return
			// 	}
			// 	col = AnalgogousFive(fyne.ThemeColorName(params[1]), int(step))
			// 	return
		}
	}
	col = th.defTheme.Color(colName, vari)
	return
}

func parseColName(colName string) (params []string) {
	for {
		if colName == "" {
			return
		}
		if strings.HasPrefix(colName, "_ptStart_") {
			i := strings.Index(colName, "_ptEnd_") + len("_ptEnd_")
			if i == -1 {
				params = nil
				return
			}
			params = append(params, colName[:i])
			colName = colName[i+1:]
		} else {
			i := strings.Index(colName, "-")
			if i == -1 {
				params = append(params, colName)
				colName = ""
			} else {
				params = append(params, colName[:i])
				colName = colName[i+1:]
			}
		}
	}
}

func (th PaletteTheme) Font(ts fyne.TextStyle) (r fyne.Resource) {
	r = th.defTheme.Font(ts)
	return
}

func (th PaletteTheme) Icon(iconName fyne.ThemeIconName) (r fyne.Resource) {
	r = th.defTheme.Icon(iconName)
	return
}

func (th PaletteTheme) Size(sizeName fyne.ThemeSizeName) (s float32) {
	s = th.defTheme.Size(sizeName)
	return
}

func NewPaletteLightDark(base fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	colNames = []fyne.ThemeColorName{
		fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base) + "-0_ptEnd_"),
		fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base) + "-2_ptEnd_"),
	}
	return
}

func NewPaletteLightDarkSet(base []fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	for i := range base {
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_LightMediumDark-"+string(base[i])+"-0_ptEnd_"))
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_LightMediumDark-"+string(base[i])+"-2_ptEnd_"))
	}
	return
}

func NewPaletteLightMediumDark(base fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	colNames = []fyne.ThemeColorName{
		fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base) + "-0_ptEnd_"),
		fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base) + "-1_ptEnd_"),
		fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base) + "-2_ptEnd_"),
	}
	return
}

func NewPaletteLightMediumDarkSet(base []fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	for i := range base {
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_LightMediumDark-"+string(base[i])+"-0_ptEnd_"))
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_LightMediumDark-"+string(base[i])+"-1_ptEnd_"))
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_LightMediumDark-"+string(base[i])+"-2_ptEnd_"))
	}
	return
}

func NewPaletteDivergentLightMediumDark(base1 fyne.ThemeColorName, base2 fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	colNames = []fyne.ThemeColorName{
		fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base1) + "-0_ptEnd_"),
		fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base1) + "-1_ptEnd_"),
		fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base1) + "-2_ptEnd_"),
		fyne.ThemeColorName("_ptStart_Neutral_ptEnd_"),
		fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base2) + "-2_ptEnd_"),
		fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base2) + "-1_ptEnd_"),
		fyne.ThemeColorName("_ptStart_LightMediumDark-" + string(base2) + "-0_ptEnd_"),
	}
	return
}

func NewPaletteEquidistantHue(base fyne.ThemeColorName, numCols int) (colNames []fyne.ThemeColorName) {
	for i := range numCols {
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_EquidistantHue-"+string(base)+"-"+strconv.Itoa(i)+"-"+strconv.Itoa(numCols)+"_ptEnd_"))
	}
	return
}

func NewPaletteComplementary(base fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	for i := range 2 {
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_EquidistantHue-"+string(base)+"-"+strconv.Itoa(i)+"-2_ptEnd_"))
	}
	return
}

func NewPaletteTriadic(base fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	for i := range 3 {
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_EquidistantHue-"+string(base)+"-"+strconv.Itoa(i)+"-3_ptEnd_"))
	}
	return
}

func NewPaletteQuadratic(base fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	for i := range 4 {
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_EquidistantHue-"+string(base)+"-"+strconv.Itoa(i)+"-4_ptEnd_"))
	}
	return
}

func NewPaletteHexadic(base fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	for i := range 6 {
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_EquidistantHue-"+string(base)+"-"+strconv.Itoa(i)+"-6_ptEnd_"))
	}
	return
}

func NewPaletteAnalogousThree(base fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	for i := range 3 {
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_AnalogousFive-"+string(base)+"-"+strconv.Itoa(i+1)+"_ptEnd_"))
	}
	return
}

func NewPaletteAnalogousFive(base fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	for i := range 5 {
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_AnalogousFive-"+string(base)+"-"+strconv.Itoa(i)+"_ptEnd_"))
	}
	return
}
