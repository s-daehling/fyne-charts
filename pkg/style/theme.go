package style

import (
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"github.com/s-daehling/fyne-charts/internal/style"
)

type paletteTheme struct {
	defTheme fyne.Theme
}

var _ fyne.Theme = (*paletteTheme)(nil)

func NewPaletteTheme(defTheme fyne.Theme) (th fyne.Theme) {
	th = &paletteTheme{
		defTheme: defTheme,
	}
	return
}

func (th paletteTheme) Color(colName fyne.ThemeColorName, vari fyne.ThemeVariant) (col color.Color) {
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
			col = style.LightDark(fyne.ThemeColorName(params[1]), int(step))
			return
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
			colName = colName[i:]
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

func (th paletteTheme) Font(ts fyne.TextStyle) (r fyne.Resource) {
	r = th.defTheme.Font(ts)
	return
}

func (th paletteTheme) Icon(iconName fyne.ThemeIconName) (r fyne.Resource) {
	r = th.defTheme.Icon(iconName)
	return
}

func (th paletteTheme) Size(sizeName fyne.ThemeSizeName) (s float32) {
	s = th.defTheme.Size(sizeName)
	return
}
