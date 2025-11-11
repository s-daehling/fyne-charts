package axis

import (
	"math"
	"strconv"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

func (ax *Axis) SetNOrigin(o float64) {
	ax.nOrigin = o
}

func (ax *Axis) AutoNOrigin() {
	switch ax.typ {
	case CartesianAxis:
		if ax.nMin < 0 && ax.nMax > 0 {
			ax.nOrigin = 0
		} else {
			ax.nOrigin = ax.nMin
		}
	case PolarPhiAxis:
		ax.nOrigin = ax.nMin
	case PolarRAxis:
		ax.nOrigin = ax.nMax
	}
}

func (ax *Axis) NOrigin() (o float64) {
	o = ax.nOrigin
	return
}

func (ax *Axis) SetNRange(min float64, max float64) {
	ax.nMin = min
	ax.nMax = max
}

func (ax *Axis) NRange() (min float64, max float64) {
	min = ax.nMin
	max = ax.nMax
	return
}

func (ax *Axis) SetNTicks(ns []data.NumericalTick, orderOfMagn int) {
	ax.adjustNumberOfTicks(len(ns))
	prec := -orderOfMagn + 1
	if prec < 0 {
		prec = 0
	}
	for i := range ns {
		ax.ticks[i].n = ns[i].N
		ax.ticks[i].nLabel = ns[i].N
		ax.ticks[i].nLine = ns[i].N
		ax.ticks[i].hasSupportLine = ns[i].SupportLine
		if ax.typ == PolarPhiAxis {
			ax.ticks[i].label.Text = strconv.FormatFloat(ns[i].N/math.Pi, 'f', 2, 64) + " pi"
		} else {
			ax.ticks[i].label.Text = strconv.FormatFloat(ns[i].N, 'f', prec, 64)
		}
	}
}

func (ax *Axis) AutoNTicks() {
	if !ax.autoTicks {
		return
	}
	min := ax.nMin
	max := ax.nMax
	if ax.typ == PolarPhiAxis {
		ns := calculatePhiTicks(ax.autoSupportLine)
		ax.SetNTicks(ns, 1)
	} else {
		ns, orderOfMagn := calculateNTicks(ax.space, min, max, ax.autoSupportLine)
		ax.SetNTicks(ns, orderOfMagn)
	}
}

func calculateNTicks(space float32, min float64, max float64,
	supLine bool) (ns []data.NumericalTick, orderOfMagn int) {
	minSpacePerLabel := 50
	maxTickNum := int(space / float32(minSpacePerLabel))
	if maxTickNum == 0 {
		maxTickNum = 1
	}
	r := max - min
	orderOfMagn = 0
	dist := 1.0
	// find upper limit for orderOfMagn
	for {
		if math.Pow10(orderOfMagn) < r {
			orderOfMagn++
		} else {
			break
		}
	}
	for {
		if r/(5.0*math.Pow10(orderOfMagn-1)) < float64(maxTickNum) {
			// more than 5*10^(magnOfOrder-1) ticks fit -> further decrease magnOfOrder
			orderOfMagn -= 1
			if orderOfMagn < -12 {
				break
			}
		} else {
			break
		}
	}
	if r/(1.0*math.Pow10(orderOfMagn)) > float64(maxTickNum) {
		// too many ticks if dist=1*10^magnOfOrder
		if r/(2.0*math.Pow10(orderOfMagn)) > float64(maxTickNum) {
			// too many ticks if dist=2*10^magnOfOrder
			dist = 5.0 * math.Pow10(orderOfMagn)
		} else {
			dist = 2.0 * math.Pow10(orderOfMagn)
		}
	} else {
		dist = 1.0 * math.Pow10(orderOfMagn)
	}

	// if min is not a multiplier of dist, calculate offset
	m := int(min / dist)
	offset := min - float64(m)*dist
	if offset > dist/1.25 {
		offset -= dist
	}

	i := 0
	for {
		coord := min + float64(i)*dist - offset
		// if i > 0 {
		// 	coord -= offset
		// }
		if coord > max+0.001*math.Pow10(orderOfMagn) {
			break
		}
		ns = append(ns, data.NumericalTick{N: coord, SupportLine: supLine})
		i += 1
	}
	return
}

func calculatePhiTicks(supLine bool) (as []data.NumericalTick) {
	// as, _ = calculateNTicks(space, min, max)
	as = []data.NumericalTick{
		{
			N:           0,
			SupportLine: supLine,
		},
		{
			N:           math.Pi / 4,
			SupportLine: supLine,
		},
		{
			N:           math.Pi / 2,
			SupportLine: supLine,
		},
		{
			N:           3 * math.Pi / 4,
			SupportLine: supLine,
		},
		{
			N:           math.Pi,
			SupportLine: supLine,
		},
		{
			N:           5 * math.Pi / 4,
			SupportLine: supLine,
		},
		{
			N:           3 * math.Pi / 2,
			SupportLine: supLine,
		},
		{
			N:           7 * math.Pi / 4,
			SupportLine: supLine,
		},
	}
	return
}

func (ax *Axis) NTipPrecision() (prec int) {
	_, orderOfMagn := calculateNTicks(ax.space, ax.nMin, ax.nMax, true)
	prec = -orderOfMagn + 2
	if prec < 0 {
		prec = 0
	}
	return
}
