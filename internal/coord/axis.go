package coord

import (
	"errors"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"github.com/s-daehling/fyne-charts/internal/renderer"
	"github.com/s-daehling/fyne-charts/pkg/data"
)

func (base *BaseChart) FromAxisElements() (min float64, max float64, origin float64,
	label renderer.Label, ticks []renderer.Tick, arrow renderer.Arrow, show bool) {
	min, max = base.fromAx.NRange()
	origin = base.fromAx.NOrigin()
	label = base.fromAx.Label()
	ticks = base.fromAx.Ticks()
	arrow = base.fromAx.Arrow()
	show = base.fromAx.Visible()
	return
}

func (base *BaseChart) ToAxisElements() (min float64, max float64, origin float64,
	label renderer.Label, ticks []renderer.Tick, arrow renderer.Arrow, show bool) {
	min, max = base.toAx.NRange()
	origin = base.toAx.NOrigin()
	label = base.toAx.Label()
	ticks = base.toAx.Ticks()
	arrow = base.toAx.Arrow()
	show = base.toAx.Visible()
	return
}

func (base *BaseChart) HideFromAxis() {
	base.fromAx.Hide()
}

func (base *BaseChart) ShowFromAxis() {
	base.fromAx.Show()
}

func (base *BaseChart) SetFromAxisLabel(l string) {
	base.fromAx.SetLabel(l)
}

func (base *BaseChart) SetFromAxisLabelStyle(sizeName fyne.ThemeSizeName, colorName fyne.ThemeColorName) {
	base.fromAx.SetAxisLabelStyle(sizeName, colorName)
}

func (base *BaseChart) SetFromAxisStyle(colorName fyne.ThemeColorName) {
	base.fromAx.SetAxisStyle(colorName)
}

func (base *BaseChart) HideToAxis() {
	base.toAx.Hide()
}

func (base *BaseChart) ShowToAxis() {
	base.toAx.Show()
}

func (base *BaseChart) SetToAxisLabel(l string) {
	base.toAx.SetLabel(l)
}

func (base *BaseChart) SetToAxisLabelStyle(sizeName fyne.ThemeSizeName, colorName fyne.ThemeColorName) {
	base.toAx.SetAxisLabelStyle(sizeName, colorName)
}

func (base *BaseChart) SetToAxisStyle(colorName fyne.ThemeColorName) {
	base.toAx.SetAxisStyle(colorName)
}

// -------------------- origin --------------------

func (base *BaseChart) SetAutoOrigin() {
	base.autoOrigin = true
	base.DataChange()
}

func (base *BaseChart) calculateAutoNOrigin() {
	base.fromAx.AutoNOrigin()
	base.toAx.AutoNOrigin()
}

func (base *BaseChart) calculateAutoTOrigin() {
	base.fromAx.AutoTOrigin()
	base.calculateAutoNOrigin()
}

func (base *BaseChart) SetNOrigin(from float64, to float64) (err error) {
	nMinFrom, nMaxFrom := base.fromAx.NRange()
	if !base.autoFromRange && (from > nMaxFrom || from < nMinFrom) {
		err = errors.New("out of user defined range")
		return
	}
	nMinTo, nMaxTo := base.toAx.NRange()
	if !base.autoToRange && (to > nMaxTo || to < nMinTo) {
		err = errors.New("out of user defined range")
		return
	}
	base.autoOrigin = false
	base.toAx.SetNOrigin(to)
	base.fromAx.SetNOrigin(from)
	base.DataChange()
	return
}

func (base *BaseChart) SetTOrigin(from time.Time, to float64) (err error) {
	tMinFrom, tMaxFrom := base.fromAx.TRange()
	if !base.autoFromRange && (from.After(tMaxFrom) || from.Before(tMinFrom)) {
		err = errors.New("t out of user defined range")
		return
	}
	nMinTo, nMaxTo := base.toAx.NRange()
	if !base.autoToRange && (to > nMaxTo || to < nMinTo) {
		err = errors.New("out of user defined range")
		return
	}
	base.autoOrigin = false
	base.toAx.SetNOrigin(to)
	base.fromAx.SetTOrigin(from)
	base.DataChange()
	return
}

// -------------------- from range --------------------

func (base *BaseChart) SetAutoFromRange() {
	base.autoFromRange = true
	base.DataChange()
}

func (base *BaseChart) calculateAutoFromNRange() {
	var min, max float64
	if !base.autoOrigin {
		// if origin was set by user, init range with x origin
		min = base.fromAx.NOrigin()
		max = min
	}
	init := false
	for i := range base.series {
		isEmpty, sMin, sMax := base.series[i].NRange()
		if isEmpty {
			continue
		}
		if !init {
			if base.autoOrigin {
				// range not inited yet and no user set origin -> init range now
				min = sMin
				max = sMax
			}
			init = true
		}
		if min > sMin {
			min = sMin
		}
		if max < sMax {
			max = sMax
		}
	}

	if !init {
		if base.autoOrigin {
			// range around 0
			min = -1
			max = 1
		} else {
			// range around user specified origin
			min -= 1
			max += 1
		}
	}

	// make sure the min and max are not equal
	absMin := math.Abs(min)
	if math.Abs(max) < absMin {
		absMin = math.Abs(max)
	}
	r := math.Abs(max - min)
	if r*1000 < absMin {
		min = max - 1
		max = max + 1
	}
	base.fromAx.SetNRange(min, max)
}

func (base *BaseChart) calculateAutoFromTRange() {
	var min, max time.Time
	if !base.autoOrigin {
		// if origin was set by user, init range with x origin
		min = base.fromAx.TOrigin()
		max = min
	}
	init := false

	for i := range base.series {
		isEmpty, sMin, sMax := base.series[i].TRange()
		// todo check if min and max are really type time.Time
		if isEmpty {
			continue
		}
		if !init {
			if base.autoOrigin {
				// range not inited yet and no user set origin -> init range now
				min = sMin
				max = sMax
			}
			init = true
		}
		if min.After(sMin) {
			min = sMin
		}
		if max.Before(sMax) {
			max = sMax
		}
	}
	if !init {
		if base.autoOrigin {
			// range around 0
			min = time.Now().Add(-time.Hour)
			max = time.Now().Add(time.Hour)
		} else {
			// range around user specified origin
			min = min.Add(-time.Hour)
			max = min.Add(time.Hour)
		}
	}

	// make sure the min and max are not equal
	if min.Equal(max) {
		min = min.Add(-time.Second)
		max = max.Add(time.Second)
	}

	base.fromAx.SetTRange(min, max)
}

func (base *BaseChart) calculateAutoFromCRange() {
	cs := []string{}
	for i := range base.series {
		scs := base.series[i].CRange()
		for j := range scs {
			exist := false
			for k := range cs {
				if scs[j] == cs[k] {
					exist = true
					break
				}
			}
			if !exist {
				cs = append(cs, scs[j])
			}
		}
	}
	base.fromAx.SetCRange(cs)
}

func (base *BaseChart) SetFromNRange(min float64, max float64) (err error) {
	if min > max {
		err = errors.New("invalid range")
		return
	}
	if !base.autoOrigin &&
		(base.fromAx.NOrigin() < min || base.fromAx.NOrigin() > max) {
		err = errors.New("previously defined origin not in range")
		return
	}
	base.autoFromRange = false
	base.fromAx.SetNRange(min, max)
	base.DataChange()
	return
}

func (base *BaseChart) SetFromTRange(min time.Time, max time.Time) (err error) {
	if min.After(max) {
		err = errors.New("invalid range")
		return
	}
	if !base.autoOrigin &&
		(base.fromAx.TOrigin().Before(min) || base.fromAx.TOrigin().After(max)) {
		err = errors.New("previously defined origin not in range")
		return
	}
	base.autoFromRange = false
	base.fromAx.SetTRange(min, max)
	base.DataChange()
	return
}

func (base *BaseChart) SetFromCRange(cs []string) (err error) {
	if len(cs) < 1 {
		err = errors.New("invalid range")
		return
	}
	base.autoFromRange = false
	base.fromAx.SetCRange(cs)
	base.DataChange()
	return
}

// -------------------- from ticks --------------------

func (base *BaseChart) SetAutoFromTicks(autoSupport bool) {
	base.fromAx.SetAutoTicks(autoSupport)
	base.DataChange()
}

func (base *BaseChart) SetFromNTicks(ts []data.NumericalTick) {
	if len(ts) < 1 {
		return
	}
	base.fromAx.SetManualTicks()
	min := ts[0].N
	max := ts[0].N
	for i := range ts {
		if ts[i].N < min {
			min = ts[i].N
		}
		if ts[i].N > max {
			max = ts[1].N
		}
	}
	r := max - min
	orderOfMagn := -100
	// find upper limit for orderOfMagn
	for {
		if math.Pow10(orderOfMagn) < r {
			orderOfMagn++
		} else {
			break
		}
	}
	base.fromAx.SetNTicks(ts, orderOfMagn)
	base.DataChange()
}

func (base *BaseChart) SetFromTTicks(ts []data.TemporalTick, format string) {
	base.fromAx.SetTTicks(ts, format)
	base.DataChange()
}

// -------------------- to range --------------------

func (base *BaseChart) SetAutoToRange() {
	base.autoToRange = true
	base.DataChange()
}

func (base *BaseChart) calculateAutoToRange() {
	var min, max float64
	if !base.autoOrigin {
		min = base.toAx.NOrigin()
		max = min
	}
	init := false
	for i := range base.series {
		isEmpty, sMin, sMax := base.series[i].ValRange()
		if isEmpty {
			continue
		}
		if !init {
			if base.autoOrigin {
				// range not inited yet and no user set origin -> init range now
				min = sMin
				max = sMax
			}
			init = true
		}
		if min > sMin {
			min = sMin
		}
		if max < sMax {
			max = sMax
		}
	}

	if !init {
		if base.autoOrigin {
			// range around 0
			min = -1
			max = 1
		} else {
			// range around user specified origin
			min -= 1
			max += 1
		}
	}

	// make sure the min and max are not equal
	absMin := math.Abs(min)
	if math.Abs(max) < absMin {
		absMin = math.Abs(max)
	}
	r := math.Abs(max - min)
	if r*1000 < absMin {
		min = max - 1
		max = max + 1
	}

	if base.planeType == PolarPlane {
		min = 0.0
	}
	base.toAx.SetNRange(min, max)
}

func (base *BaseChart) SetToRange(min float64, max float64) (err error) {
	if max < 0 {
		err = errors.New("invalid range")
		return
	}
	if !base.autoOrigin &&
		(base.toAx.NOrigin() < min || base.toAx.NOrigin() > max) {
		err = errors.New("previously defined origin not in range")
		return
	}
	base.autoToRange = false
	base.toAx.SetNRange(min, max)
	base.DataChange()
	return
}

// -------------------- to ticks --------------------

func (base *BaseChart) SetAutoToTicks(autoSupport bool) {
	base.toAx.SetAutoTicks(autoSupport)
	base.DataChange()
}

func (base *BaseChart) SetToTicks(ts []data.NumericalTick) {
	if len(ts) < 1 {
		return
	}
	base.toAx.SetManualTicks()
	min := ts[0].N
	max := ts[0].N
	for i := range ts {
		if ts[i].N < min {
			min = ts[i].N
		}
		if ts[i].N > max {
			max = ts[1].N
		}
	}
	r := max - min
	orderOfMagn := -100
	// find upper limit for orderOfMagn
	for {
		if math.Pow10(orderOfMagn) < r {
			orderOfMagn++
		} else {
			break
		}
	}
	base.toAx.SetNTicks(ts, orderOfMagn)
	base.DataChange()
}
