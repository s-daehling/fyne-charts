package style

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	ftheme "fyne.io/fyne/v2/theme"
)

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
			col = col.DecreaseLightnessUntilContrast(1, fg)
			if baseCol.l > col.l {
				col = baseCol
			}
		} else {
			// col.l = float32(math.Min(0.15, float64(fg.l)))
			col.l = 0.2
			col = col.IncreaseLightnessUntilContrast(1, fg)
			if baseCol.l < col.l {
				col = baseCol
			}
		}
	} else {
		if fromBtoF {
			// col.l = float32(math.Min(0.15, float64(bg.l)))
			col.l = 0.2
			col = col.IncreaseLightnessUntilContrast(2, bg)
			if baseCol.l < col.l {
				col = baseCol
			}
		} else {
			// col.l = float32(math.Max(0.85, float64(bg.l)))
			col.l = 0.8
			col = col.DecreaseLightnessUntilContrast(2, bg)
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
	if step == 1 {
		baseCol.h += 180
	}
	if baseCol.h > 360 {
		baseCol.h -= 360
	}
	col = baseCol
	return
}

func Triadic(base fyne.ThemeColorName, step int) (col color.Color) {
	baseCol := ToHSLA(ftheme.Color(base))
	if step == 1 {
		baseCol.h += 120
	} else if step == 2 {
		baseCol.h += 240
	}
	if baseCol.h > 360 {
		baseCol.h -= 360
	}
	col = baseCol
	return
}

func Tetradic(base fyne.ThemeColorName, step int) (col color.Color) {
	baseCol := ToHSLA(ftheme.Color(base))
	if step == 1 {
		baseCol.h += 180
	} else if step == 2 {
		baseCol.h += 90
	} else if step == 3 {
		baseCol.h += 270
	}
	if baseCol.h > 360 {
		baseCol.h -= 360
	}
	col = baseCol
	return
}
