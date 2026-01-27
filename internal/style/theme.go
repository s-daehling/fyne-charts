package style

import (
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
)

type PaletteTheme struct {
	DefTheme fyne.Theme
}

var _ fyne.Theme = (*PaletteTheme)(nil)

func NewColorPaletteTheme(defTheme fyne.Theme) (th fyne.Theme) {
	th = &PaletteTheme{
		DefTheme: defTheme,
	}
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
		case "LightDark":
			if len(params) != 3 {
				return
			}
			step, err := strconv.ParseInt(params[2], 10, 0)
			if err != nil {
				return
			}
			col = LightDark(fyne.ThemeColorName(params[1]), int(step))
			return
		case "LightMediumDark":
			if len(params) != 3 {
				return
			}
			step, err := strconv.ParseInt(params[2], 10, 0)
			if err != nil {
				return
			}
			col = LightMediumDark(fyne.ThemeColorName(params[1]), int(step))
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
		case "Complementary":
			if len(params) != 3 {
				return
			}
			step, err := strconv.ParseInt(params[2], 10, 0)
			if err != nil {
				return
			}
			col = Complementary(fyne.ThemeColorName(params[1]), int(step))
			return
		case "Triadic":
			if len(params) != 3 {
				return
			}
			step, err := strconv.ParseInt(params[2], 10, 0)
			if err != nil {
				return
			}
			col = Triadic(fyne.ThemeColorName(params[1]), int(step))
			return
		case "Tetradic":
			if len(params) != 3 {
				return
			}
			step, err := strconv.ParseInt(params[2], 10, 0)
			if err != nil {
				return
			}
			col = Tetradic(fyne.ThemeColorName(params[1]), int(step))
			return
		}
	}
	col = th.DefTheme.Color(colName, vari)
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
	r = th.DefTheme.Font(ts)
	return
}

func (th PaletteTheme) Icon(iconName fyne.ThemeIconName) (r fyne.Resource) {
	r = th.DefTheme.Icon(iconName)
	return
}

func (th PaletteTheme) Size(sizeName fyne.ThemeSizeName) (s float32) {
	s = th.DefTheme.Size(sizeName)
	return
}

func NewPaletteLightDark(base fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	colNames = []fyne.ThemeColorName{
		fyne.ThemeColorName("_ptStart_LightDark-" + string(base) + "-0_ptEnd_"),
		fyne.ThemeColorName("_ptStart_LightDark-" + string(base) + "-1_ptEnd_"),
	}
	return
}

func NewPaletteLightDarkSet(base []fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	for i := range base {
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_LightDark-"+string(base[i])+"-0_ptEnd_"))
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_LightDark-"+string(base[i])+"-1_ptEnd_"))
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

func NewPaletteDivLightMediumDarkSet(base []fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	for i := range base {
		if i%2 == 0 {
			colNames = append(colNames,
				fyne.ThemeColorName("_ptStart_LightMediumDark-"+string(base[i])+"-0_ptEnd_"))
			colNames = append(colNames,
				fyne.ThemeColorName("_ptStart_LightMediumDark-"+string(base[i])+"-1_ptEnd_"))
			colNames = append(colNames,
				fyne.ThemeColorName("_ptStart_LightMediumDark-"+string(base[i])+"-2_ptEnd_"))
		} else {
			colNames = append(colNames,
				fyne.ThemeColorName("_ptStart_LightMediumDark-"+string(base[i])+"-2_ptEnd_"))
			colNames = append(colNames,
				fyne.ThemeColorName("_ptStart_LightMediumDark-"+string(base[i])+"-1_ptEnd_"))
			colNames = append(colNames,
				fyne.ThemeColorName("_ptStart_LightMediumDark-"+string(base[i])+"-0_ptEnd_"))
		}
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
			fyne.ThemeColorName("_ptStart_Complementary-"+string(base)+"-"+strconv.Itoa(i)+"_ptEnd_"))
	}
	return
}

func NewPaletteTriadic(base fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	for i := range 3 {
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_Triadic-"+string(base)+"-"+strconv.Itoa(i)+"_ptEnd_"))
	}
	return
}

func NewPaletteTetradic(base fyne.ThemeColorName) (colNames []fyne.ThemeColorName) {
	for i := range 4 {
		colNames = append(colNames,
			fyne.ThemeColorName("_ptStart_Tetradic-"+string(base)+"-"+strconv.Itoa(i)+"_ptEnd_"))
	}
	return
}
