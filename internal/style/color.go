package style

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/lucasb-eyer/go-colorful"
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

func LightnessRange(ref color.Color, background color.Color, foreground color.Color, conToBack float64, conToFore float64) (lMin, lMax float64) {
	col, _ := colorful.MakeColor(ref)
	h, c, l := col.Hcl()
	cMax := MaxChroma(h, l)
	cRel := c / cMax
	bgCol, _ := colorful.MakeColor(background)
	_, _, bl := bgCol.Hcl()
	fgCol, _ := colorful.MakeColor(foreground)
	_, _, fl := fgCol.Hcl()
	if fl > bl {
		// background darker than foreground (typical dark mode)
		lMin = bl
		for range 80 {
			cMax = MaxChroma(h, lMin)
			col = colorful.Hcl(h, cRel*cMax, lMin).Clamped()
			if Contrast(col, bgCol) > conToBack {
				lMin -= 0.01
				break
			}
			lMin += 0.01
		}
		lMax = 0.8
		for range 80 {
			cMax = MaxChroma(h, lMax)
			col = colorful.Hcl(h, cRel*cMax, lMax).Clamped()
			if Contrast(col, fgCol) > conToFore {
				lMax += 0.01
				break
			}
			lMax -= 0.01
		}
	} else {
		// background lighter than foreground (typical light mode)
		lMin = 0.25
		for range 80 {
			cMax = MaxChroma(h, lMin)
			col = colorful.Hcl(h, cRel*cMax, lMin).Clamped()
			if Contrast(col, fgCol) > conToFore {
				lMin -= 0.01
				break
			}
			lMin += 0.01
		}
		lMax = bl
		for range 80 {
			cMax = MaxChroma(h, lMax)
			col = colorful.Hcl(h, cRel*cMax, lMax).Clamped()
			if Contrast(col, bgCol) > conToBack {
				lMax += 0.01
				break
			}
			lMax -= 0.01
		}
	}
	return
}

func MaxChroma(h, l float64) (maxC float64) {
	maxC = 0
	for {
		col := colorful.Hcl(h, maxC, l)
		if !col.IsValid() {
			maxC -= 0.01
			return
		}
		maxC += 0.01
	}
}

func LightMediumDark(base fyne.ThemeColorName, step int, settings ThemeSettings) (col colorful.Color) {
	baseCol, _ := colorful.MakeColor(theme.Color(base))
	h, c, l := baseCol.Hcl()
	cRel := c / MaxChroma(h, l)
	switch step {
	case 0:
		if settings.Variant == theme.VariantDark {
			l = settings.MaxL
		} else {
			l = settings.MinL
		}
	case 1:
		l = (settings.MinL + settings.MaxL) / 2
	case 2:
		if settings.Variant == theme.VariantDark {
			l = settings.MinL
		} else {
			l = settings.MaxL
		}
	}
	col = colorful.Hcl(h, cRel*MaxChroma(h, l), l).Clamped()
	return
}

func Neutral(settings ThemeSettings) (col colorful.Color) {
	if settings.Variant == theme.VariantDark {
		col = colorful.Hcl(0, 0, 0.1).Clamped()
	} else {
		col = colorful.Hcl(0, 0, 0.9).Clamped()
	}
	return
}

func EquidistantHue(base fyne.ThemeColorName, step int, totStep int) (col color.Color) {
	baseCol, _ := colorful.MakeColor(theme.Color(base))
	if step == 0 {
		col = baseCol
		return
	}
	hueStep := 360 / float64(totStep)
	h, c, l := baseCol.Hcl()
	cRel := c / MaxChroma(h, l)

	h += hueStep * float64(step)
	if h > 360 {
		h -= 360
	}
	col = colorful.Hcl(h, cRel*MaxChroma(h, l), l).Clamped()
	return
}
