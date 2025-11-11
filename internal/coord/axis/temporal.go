package axis

import (
	"time"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

func (ax *Axis) SetTOrigin(o time.Time) {
	ax.tOrigin = o
}

func (ax *Axis) AutoTOrigin() {
	ax.tOrigin = ax.tMin
}

func (ax *Axis) TOrigin() (o time.Time) {
	o = ax.tOrigin
	return
}

func (ax *Axis) SetTRange(min time.Time, max time.Time) {
	ax.tMin = min
	ax.tMax = max
}

func (ax *Axis) TRange() (min time.Time, max time.Time) {
	min = ax.tMin
	max = ax.tMax
	return
}

func (ax *Axis) SetTTicks(ts []data.TemporalTick, format string) {
	ax.adjustNumberOfTicks(len(ts))
	for i := range ts {
		ax.ticks[i].t = ts[i].T
		ax.ticks[i].label.Text = ts[i].T.Format(format)
		ax.ticks[i].hasSupportLine = ts[i].SupportLine
	}
}

func (ax *Axis) AutoTTicks() {
	if !ax.autoTicks {
		return
	}
	min := ax.tMin
	max := ax.tMax
	ts, format, _ := calculateTTicks(ax.space, min, max, ax.autoSupportLine)
	ax.SetTTicks(ts, format)
}

func calculateTTicks(space float32, min time.Time, max time.Time, supLine bool) (ts []data.TemporalTick, tickFormat string, tipFormat string) {
	minSpacePerLabel := 100
	maxTickNum := int(space / float32(minSpacePerLabel))
	if maxTickNum == 0 {
		maxTickNum = 1
	}
	r := max.Sub(min)
	numYears := int(r.Hours() / (365 * 24))
	numDays := int(r.Hours() / 24)
	var coord time.Time
	var step time.Duration
	addMin := false
	if numYears > (maxTickNum / 2) {
		// #years > #ticks/2 -> use years as ticks
		inc := numYears/maxTickNum + 1
		coord = time.Date(min.Year(), time.January, 1, 0, 0, 0, 0, time.Local)
		step = time.Duration(int(time.Hour) * 24 * 365 * inc)
		tickFormat = "2006"
		tipFormat = "01.2006"
		if min.Month() == time.January && min.Day() <= 7 {
			addMin = true
		}
	} else if numDays/30 > maxTickNum/2 {
		// #months > #ticks/2 -> use months as ticks
		inc := (numDays/30)/maxTickNum + 1
		coord = time.Date(min.Year(), min.Month(), 1, 0, 0, 0, 0, time.Local)
		step = time.Duration(int(time.Hour) * 24 * 31 * inc)
		tickFormat = "01.2006"
		tipFormat = "02.01."
		if min.Day() == 1 {
			addMin = true
		}
	} else if numDays > maxTickNum/2 {
		// #days > #ticks/2 -> days as ticks
		inc := numDays/maxTickNum + 1
		coord = time.Date(min.Year(), min.Month(), min.Day(), 0, 0, 0, 0, time.Local)
		step = time.Duration(int(time.Hour) * 24 * inc)
		tickFormat = "02.01."
		tipFormat = "15h"
		if min.Hour() < 1 {
			addMin = true
		}
	} else if int(r.Hours()) > maxTickNum/2 {
		// #hours > #ticks/2 -> hours as ticks
		inc := int(r.Hours())/maxTickNum + 1
		coord = time.Date(min.Year(), min.Month(), min.Day(), min.Hour(), 0, 0, 0, time.Local)
		step = time.Duration(int(time.Hour) * inc)
		tickFormat = "15h"
		tipFormat = "15:04"
		if min.Minute() < 5 {
			addMin = true
		}
	} else if int(r.Minutes()) > maxTickNum/2 {
		// #mins > #ticks/2 -> mins as ticks
		inc := int(r.Minutes())/maxTickNum + 1
		coord = time.Date(min.Year(), min.Month(), min.Day(), min.Hour(), min.Minute(), 0, 0, time.Local)
		step = time.Duration(int(time.Minute) * inc)
		tickFormat = "15:04"
		tipFormat = "15:04:05"
		if min.Second() < 5 {
			addMin = true
		}
	} else if int(r.Seconds()) > maxTickNum/2 {
		// #secs > #ticks/2 -> secs as ticks
		inc := int(r.Seconds())/maxTickNum + 1
		coord = time.Date(min.Year(), min.Month(), min.Day(), min.Hour(), min.Minute(), min.Second(), 0, time.Local)
		step = time.Duration(int(time.Second) * inc)
		tickFormat = "15:04:05"
		tipFormat = "05.000"
		if min.Sub(coord).Milliseconds() < 25 {
			addMin = true
		}
	} else {
		// #msecs as ticks
		inc := int(r.Milliseconds())/maxTickNum + 1
		coord = time.Date(min.Year(), min.Month(), min.Day(), min.Hour(), min.Minute(), min.Second(), min.Nanosecond(), time.Local)
		step = time.Duration(int(time.Millisecond) * inc)
		tickFormat = "05.000"
		tipFormat = "05.000"
		if min.Sub(coord).Nanoseconds() < 100 {
			addMin = true
		}
	}
	if addMin {
		ts = append(ts, data.TemporalTick{T: min, SupportLine: supLine})
	}
	for {
		coord = coord.Add(step)
		if coord.After(max) {
			break
		}
		ts = append(ts, data.TemporalTick{T: coord, SupportLine: supLine})
	}
	return
}

func (ax *Axis) ConvertTTickstoN() {
	for i := range ax.ticks {
		ax.ticks[i].nLabel = ax.TtoN(ax.ticks[i].t)
		ax.ticks[i].nLine = ax.TtoN(ax.ticks[i].t)
	}
	ax.nOrigin = ax.TtoN(ax.tOrigin)
}

func (ax *Axis) TtoN(t time.Time) (n float64) {
	r := ax.tMax.Sub(ax.tMin).Nanoseconds()
	p := t.Sub(ax.tMin).Nanoseconds()
	n = ax.nMin + (float64(p) / float64(r) * (ax.nMax - ax.nMin))
	return
}

func (ax *Axis) NtoT(n float64) (t time.Time) {
	r := ax.tMax.Sub(ax.tMin).Nanoseconds()
	d := time.Duration(float64(r) * ((n - ax.nMin) / (ax.nMax - ax.nMin)))
	t = ax.tMin.Add(time.Nanosecond * d)
	return
}

func (ax *Axis) TTipFormat() (f string) {
	_, _, f = calculateTTicks(ax.space, ax.tMin, ax.tMax, true)
	return
}
