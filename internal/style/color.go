package style

import (
	"image/color"
	"math"
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
