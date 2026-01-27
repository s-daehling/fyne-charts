package style

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	ftheme "fyne.io/fyne/v2/theme"
)

// Contrast caluclates the contrast between a foreground and background color
// formula from W3C: https://www.w3.org/TR/2008/REC-WCAG20-20081211/#contrast-ratiodef
func Contrast(fg color.Color, bg color.Color) (c float64) {
	bgR, bgG, bgB, _ := bg.RGBA()
	bgsR := float64(bgR) / float64(0xffff)
	bgsG := float64(bgG) / float64(0xffff)
	bgsB := float64(bgB) / float64(0xffff)
	// bgAN := float64(bgA) / float64(0xffff)
	var bgRV, bgGV, bgBV float64
	if bgsR <= 0.03928 {
		bgRV = bgsR / 12.92
	} else {
		bgRV = math.Pow((bgsR+0.055)/1.055, 2.4)
	}
	if bgsG <= 0.03928 {
		bgGV = bgsG / 12.92
	} else {
		bgGV = math.Pow((bgsG+0.055)/1.055, 2.4)
	}
	if bgsB <= 0.03928 {
		bgBV = bgsB / 12.92
	} else {
		bgBV = math.Pow((bgsB+0.055)/1.055, 2.4)
	}
	bL := 0.2126*bgRV + 0.7152*bgGV + 0.0722*bgBV

	fgR, fgG, fgB, fgA := fg.RGBA()
	fgsA := float64(fgA) / float64(0xffff)
	fgsR := (float64(fgR) / float64(0xffff)) + ((1 - fgsA) * bgsR)
	fgsG := float64(fgG)/float64(0xffff) + ((1 - fgsA) * bgsG)
	fgsB := float64(fgB)/float64(0xffff) + ((1 - fgsA) * bgsB)
	var fgRV, fgGV, fgBV float64
	if fgsR <= 0.03928 {
		fgRV = fgsR / 12.92
	} else {
		fgRV = math.Pow((fgsR+0.055)/1.055, 2.4)
	}
	if fgsG <= 0.03928 {
		fgGV = fgsG / 12.92
	} else {
		fgGV = math.Pow((fgsG+0.055)/1.055, 2.4)
	}
	if fgsB <= 0.03928 {
		fgBV = fgsB / 12.92
	} else {
		fgBV = math.Pow((fgsB+0.055)/1.055, 2.4)
	}
	fL := 0.2126*fgRV + 0.7152*fgGV + 0.0722*fgBV

	if bL > fL {
		c = (bL + 0.05) / (fL + 0.05)
	} else {
		c = (fL + 0.05) / (bL + 0.05)
	}

	return
}

type HSLA struct {
	h, s, l, a float32
}

func NewHSLA(h, s, l, a float32) (col HSLA) {
	col.h = h - float32(int(h/360)*360)
	col.s = s
	if s < 0 {
		col.s = 0
	} else if s > 1 {
		col.s = 1
	}
	col.l = l
	if l < 0 {
		col.l = 0
	} else if l > 1 {
		col.l = 1
	}
	col.a = a
	if a < 0 {
		col.a = 0
	} else if a > 1 {
		col.a = 1
	}
	return
}

func ToHSLA(col color.Color) (hsla HSLA) {
	r, g, b, a := col.RGBA()
	hsla.a = float32(a) / float32(0xffff)
	sr := (float64(r) / float64(0xffff)) / float64(hsla.a)
	sg := (float64(g) / float64(0xffff)) / float64(hsla.a)
	sb := (float64(b) / float64(0xffff)) / float64(hsla.a)
	max := math.Max(math.Max(sr, sg), sb)
	min := math.Min(math.Min(sr, sg), sb)

	if math.Abs(max-min) < 0.000001 {
		hsla.h = 0
	} else if math.Abs(max-sr) < 0.000001 {
		hsla.h = float32(60 * ((sg - sb) / (max - min)))
	} else if math.Abs(max-sg) < 0.000001 {
		hsla.h = float32(60 * (2 + (sb-sr)/(max-min)))
	} else if math.Abs(max-sb) < 0.000001 {
		hsla.h = float32(60 * (4 + (sr-sg)/(max-min)))
	}
	if hsla.h < 0 {
		hsla.h += 360
	}
	if math.Abs(max-min) < 0.000001 {
		hsla.s = 0
	} else {
		hsla.s = float32((max - min) / (1 - math.Abs(max+min-1)))
	}
	hsla.l = float32((max + min) / 2)
	return
}

func (hsla HSLA) RGBA() (r, g, b, a uint32) {
	c := (1 - math.Abs(2*float64(hsla.l)-1)) * float64(hsla.s)
	h_ := float64(hsla.h) / 60
	x := c * (1 - math.Abs(math.Mod(h_, 2)-1))
	m := float64(hsla.l) - (c / 2)

	cUInt := uint32((c + m) * float64(0xffff) * float64(hsla.a))
	xUInt := uint32((x + m) * float64(0xffff) * float64(hsla.a))
	mUInt := uint32(m * float64(0xffff) * float64(hsla.a))
	if h_ < 1 {
		r = cUInt
		g = xUInt
		b = mUInt
	} else if h_ < 2 {
		r = xUInt
		g = cUInt
		b = mUInt
	} else if h_ < 3 {
		r = mUInt
		g = cUInt
		b = xUInt
	} else if h_ < 4 {
		r = mUInt
		g = xUInt
		b = cUInt
	} else if h_ < 5 {
		r = xUInt
		g = mUInt
		b = cUInt
	} else {
		r = cUInt
		g = mUInt
		b = xUInt
	}
	a = uint32(hsla.a * float32(0xffff))
	return
}

func (hsla HSLA) WithLightness(l float32) (ret HSLA) {
	ret = hsla
	ret.l = l
	return
}

func (hsla HSLA) WithSaturation(s float32) (ret HSLA) {
	ret = hsla
	ret.s = s
	return
}

func (hsla HSLA) WithHue(h float32) (ret HSLA) {
	ret = hsla
	ret.h = h
	return
}

func (hsla HSLA) ShiftHue(h float32) (ret HSLA) {
	ret = hsla
	ret.h += h
	if ret.h > 360 {
		ret.h -= 360
	} else if ret.h < 0 {
		ret.h += 360
	}
	return
}

func (hsla HSLA) IncreaseLightnessUntilContrast(con float64, refCol color.Color) (ret HSLA) {
	ret = hsla
	for i := range 100 {
		c := Contrast(ret, refCol)
		if c > con {
			break
		}
		ret.l += float32(i) / 100
		if ret.l > 1 {
			ret.l = 1
			break
		}
	}
	return
}

func (hsla HSLA) DecreaseLightnessUntilContrast(con float64, refCol color.Color) (ret HSLA) {
	ret = hsla
	for i := range 100 {
		c := Contrast(ret, refCol)
		if c > con {
			break
		}
		ret.l -= float32(i) / 100
		if ret.l < 0 {
			ret.l = 0
			break
		}
	}
	return
}

func LightDark(base fyne.ThemeColorName, step int) (col HSLA) {
	baseCol := ToHSLA(ftheme.Color(base))
	bg := ToHSLA(ftheme.Color(ftheme.ColorNameBackground))
	fg := ToHSLA(ftheme.Color(ftheme.ColorNameForeground))
	fromBtoF := true
	if fg.l < bg.l {
		fromBtoF = false
	}
	col = baseCol
	if step == 0 {
		if fromBtoF {
			// col.l = float32(math.Max(0.85, float64(fg.l)))
			col.l = 0.8
			col = col.DecreaseLightnessUntilContrast(1.2, fg)
			if baseCol.l > col.l {
				col = baseCol
			}
		} else {
			// col.l = float32(math.Min(0.15, float64(fg.l)))
			col.l = 0.2
			col = col.IncreaseLightnessUntilContrast(1.2, fg)
			if baseCol.l < col.l {
				col = baseCol
			}
		}
	} else {
		if fromBtoF {
			// col.l = float32(math.Min(0.15, float64(bg.l)))
			col.l = 0.2
			col = col.IncreaseLightnessUntilContrast(2.2, bg)
			if baseCol.l < col.l {
				col = baseCol
			}
		} else {
			// col.l = float32(math.Max(0.85, float64(bg.l)))
			col.l = 0.8
			col = col.DecreaseLightnessUntilContrast(2.2, bg)
			if baseCol.l > col.l {
				col = baseCol
			}
		}
	}
	return
}

func LightMediumDark(base fyne.ThemeColorName, step int) (col color.Color) {
	if step == 0 {
		col = LightDark(base, 0)
	} else if step == 1 {
		l := LightDark(base, 0)
		// fmt.Println(l)
		d := LightDark(base, 1)
		// fmt.Println(d)
		l.l = float32(math.Abs(float64(l.l+d.l)) / 2)
		// fmt.Println(l)
		col = l
	} else {
		col = LightDark(base, 1)
	}

	return
}

func DivergentLightDarkWithNeutral(base1 fyne.ThemeColorName, base2 fyne.ThemeColorName, step int) (col color.Color) {
	if step == 0 || step == 1 {
		col = LightDark(base1, step)
	} else if step == 3 {
		col = LightDark(base2, 1)
	} else if step == 4 {
		col = LightDark(base2, 0)
	} else {
		bg := ToHSLA(ftheme.Color(ftheme.ColorNameBackground))
		fg := ToHSLA(ftheme.Color(ftheme.ColorNameForeground))
		if fg.l < bg.l {
			bg.l -= 0.075
		} else {
			bg.l += 0.075
		}
		col = bg

		// col1 := LightDark(base1, 1)
		// col2 := LightDark(base2, 0)
		// col1.h = (col1.h + col2.h) / 2
		// col1.s = 0
		// col = col1
	}
	return
}

func DivergentLightMediumDarkWithNeutral(base1 fyne.ThemeColorName, base2 fyne.ThemeColorName, step int) (col color.Color) {
	if step < 3 {
		col = LightMediumDark(base1, step)
	} else if step == 4 {
		col = LightMediumDark(base2, 2)
	} else if step == 5 {
		col = LightMediumDark(base2, 1)
	} else if step == 6 {
		col = LightMediumDark(base2, 0)
	} else {
		bg := ToHSLA(ftheme.Color(ftheme.ColorNameBackground))
		fg := ToHSLA(ftheme.Color(ftheme.ColorNameForeground))
		if fg.l < bg.l {
			bg.l -= 0.075
		} else {
			bg.l += 0.075
		}
		col = bg

		// col1 := LightDark(base1, 1)
		// col2 := LightDark(base2, 0)
		// col1.h = (col1.h + col2.h) / 2
		// col1.s = 0
		// col = col1
	}
	return
}

func Monochromatic(base fyne.ThemeColorName, step int, totStep int) (col color.Color) {
	inc := 0.7 / float32(totStep-1)
	baseCol := ToHSLA(ftheme.Color(base))
	bg := ToHSLA(ftheme.Color(ftheme.ColorNameBackground))
	fg := ToHSLA(ftheme.Color(ftheme.ColorNameForeground))
	if fg.l < bg.l {
		baseCol.l = 0.85 - (float32(step) * inc)
	} else {
		baseCol.l = 0.15 + (float32(step) * inc)
	}
	col = baseCol
	return
}

func Blend(base1 fyne.ThemeColorName, base2 fyne.ThemeColorName, step int, totStep int) (col color.Color) {
	col1 := ftheme.Color(base1)
	col2 := ftheme.Color(base2)

	scale := float32(step) / float32(totStep)

	r1, g1, b1, a1 := col1.RGBA()
	sa1 := float32(a1) / float32(0xffff)
	sr1 := (float32(r1) / float32(0xffff)) / sa1
	sg1 := (float32(g1) / float32(0xffff)) / sa1
	sb1 := (float32(b1) / float32(0xffff)) / sa1

	r2, g2, b2, a2 := col2.RGBA()
	sa2 := float32(a2) / float32(0xffff)
	sr2 := (float32(r2) / float32(0xffff)) / sa2
	sg2 := (float32(g2) / float32(0xffff)) / sa2
	sb2 := (float32(b2) / float32(0xffff)) / sa2

	a := (sa1 * (1 - scale)) + (sa2 * scale)
	r := ((sr1 * (1 - scale)) + (sr2 * scale)) * a
	g := ((sg1 * (1 - scale)) + (sg2 * scale)) * a
	b := ((sb1 * (1 - scale)) + (sb2 * scale)) * a

	col = color.RGBA{R: uint8(r * 0xff), G: uint8(g * 0xff), B: uint8(b * 0xff), A: uint8(a * 0xff)}

	return
}

func Complementary(base fyne.ThemeColorName, step int) (col color.Color) {
	baseCol := ToHSLA(ftheme.Color(base))
	allCols := make([]HSLA, 0)
	allCols = append(allCols, baseCol)
	if baseCol.s < 0.5 {
		baseCol.s = 0.5
	}
	baseCol.h += 180
	if baseCol.h > 360 {
		baseCol.h -= 360
	}
	allCols = append(allCols, baseCol)
	allCols = optimizeContrastAlt(allCols)
	col = allCols[step]
	return
}

func Triadic(base fyne.ThemeColorName, step int) (col color.Color) {
	baseCol := ToHSLA(ftheme.Color(base))
	allCols := make([]HSLA, 0)
	allCols = append(allCols, baseCol)
	if baseCol.s < 0.5 {
		baseCol.s = 0.5
	}
	for range 2 {
		baseCol.h += 120
		if baseCol.h > 360 {
			baseCol.h -= 360
		}
		allCols = append(allCols, baseCol)
	}
	allCols = optimizeContrastAlt(allCols)
	col = allCols[step]
	return
}

func Tetradic(base fyne.ThemeColorName, step int) (col color.Color) {
	baseCol := ToHSLA(ftheme.Color(base))
	allCols := make([]HSLA, 0)
	allCols = append(allCols, baseCol)
	if baseCol.s < 0.5 {
		baseCol.s = 0.5
	}
	for range 3 {
		baseCol.h += 90
		if baseCol.h > 360 {
			baseCol.h -= 360
		}
		allCols = append(allCols, baseCol)
	}
	allCols = optimizeContrastAlt(allCols)
	col = allCols[step]
	return
}

func EquidistantHue(base fyne.ThemeColorName, step int, totStep int) (col color.Color) {
	baseCol := ToHSLA(ftheme.Color(base))
	hueStep := 360 / float32(totStep)
	nextCol := baseCol
	if nextCol.s < 0.5 {
		nextCol.s = 0.5
	}
	allCols := make([]HSLA, 0)
	allCols = append(allCols, baseCol)
	for range totStep - 1 {
		nextCol.h += hueStep
		if nextCol.h > 360 {
			nextCol.h -= 360
		}
		allCols = append(allCols, nextCol)
	}
	allCols = optimizeContrastAlt(allCols)
	col = allCols[step]
	return
}

type colVar struct {
	col       HSLA
	bgCon     float64
	fgCon     float64
	minSetCon float64
}

func optimizeContrast(in []HSLA) (out []HSLA) {
	if len(in) == 0 {
		return
	}
	out = append(out, in[0])

	for i := range len(in) - 1 {
		nextCol := in[i+1]
		nextCol.l = 0.3
		colVars := make([]colVar, 0)
		colVarsWithContrast := make([]colVar, 0)
		for j := range 3 {
			cv := colVar{col: nextCol}
			cv.col.l += float32(j) * 0.2
			cv.bgCon = Contrast(cv.col, theme.Color(theme.ColorNameBackground))
			cv.fgCon = Contrast(cv.col, theme.Color(theme.ColorNameForeground))
			minSetCon := 100.0
			for k := range out {
				con := Contrast(cv.col, out[k])
				if con < minSetCon {
					minSetCon = con
				}
			}
			cv.minSetCon = minSetCon
			colVars = append(colVars, cv)
			if cv.bgCon > 2.2 && cv.fgCon > 1.2 {
				colVarsWithContrast = append(colVarsWithContrast, cv)
			}
		}
		if len(colVarsWithContrast) > 0 {
			// find variant with highest min contrast to set
			cvInd := 0
			for j := range colVarsWithContrast {
				if colVarsWithContrast[j].minSetCon > colVarsWithContrast[cvInd].minSetCon {
					cvInd = j
				}
			}
			nextCol = colVarsWithContrast[cvInd].col
		} else {
			// find variant with highest contrast to background
			cvInd := 0
			for j := range colVars {
				if colVars[j].bgCon > colVars[cvInd].bgCon {
					cvInd = j
				}
			}
			nextCol = colVars[cvInd].col
		}
		out = append(out, nextCol)
	}
	return
}

type setVariant struct {
	cols      []HSLA
	minBgCon  float64
	minFgCon  float64
	minSetCon float64
}

func optimizeContrastAlt(in []HSLA) (out []HSLA) {
	if len(in) == 0 {
		return
	}
	set := make([]HSLA, 0)
	setVars := make([]setVariant, 0)
	setVarsWithContrast := make([]setVariant, 0)
	for i := range len(in) {
		if i == 0 {
			continue
		}
		set = append(set, in[i])
	}
	numVars := int(math.Pow(3, float64(len(set))))
	for i := range numVars {
		setVar := setVariant{
			cols:      []HSLA{in[0]},
			minBgCon:  100.0,
			minFgCon:  100.0,
			minSetCon: 100.0,
		}
		for j := range set {
			col := set[j]
			col.l = 0.3 + 0.2*float32((i/int(math.Pow(3, float64(j))))%3)
			setVar.cols = append(setVar.cols, col)
		}
		for j := range setVar.cols {
			setVar.minBgCon = math.Min(setVar.minBgCon, Contrast(setVar.cols[j], theme.Color(theme.ColorNameBackground)))
			setVar.minFgCon = math.Min(setVar.minFgCon, Contrast(setVar.cols[j], theme.Color(theme.ColorNameForeground)))
			minCon := 100.0
			for k := range setVar.cols {
				if k == j {
					continue
				}
				con := Contrast(setVar.cols[j], setVar.cols[k])
				minCon = math.Min(minCon, con)
			}
			setVar.minSetCon = math.Min(setVar.minSetCon, minCon)
		}
		setVars = append(setVars, setVar)
		if setVar.minBgCon > 2.2 && setVar.minFgCon > 1.2 {
			setVarsWithContrast = append(setVarsWithContrast, setVar)
		}
	}
	// for i := range setVars {
	// 	fmt.Println(setVars[i])
	// }
	if len(setVarsWithContrast) > 0 {
		varInd := 0
		for i := range setVarsWithContrast {
			if setVarsWithContrast[i].minSetCon > setVarsWithContrast[varInd].minSetCon {
				varInd = i
			}
		}
		out = setVarsWithContrast[varInd].cols
		// fmt.Println("with", setVarsWithContrast[varInd])
	} else {
		varInd := 0
		for i := range setVars {
			if setVars[i].minBgCon > setVars[varInd].minBgCon {
				varInd = i
			}
		}
		out = setVars[varInd].cols
		// fmt.Println("without", setVars[varInd])
	}
	return
}
