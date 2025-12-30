package coord

import (
	"errors"

	"github.com/s-daehling/fyne-charts/internal/coord/series"
	"github.com/s-daehling/fyne-charts/internal/interact"
)

func (base *BaseChart) addSeriesIfNotExist(ser series.Series) (err error) {
	for i := range base.series {
		if base.series[i].Name() == ser.Name() {
			err = errors.New("series already exists")
			return
		}
	}
	err = ser.BindToChart(base)
	if err != nil {
		return
	}
	base.series = append(base.series, ser)
	base.DataChange()
	return
}

func (base *BaseChart) AddLegendEntry(le *interact.LegendEntry) {
	base.legend.AddEntry(le)
}

func (base *BaseChart) RemoveLegendEntry(name string, super string) {
	base.legend.RemoveEntry(name, super)
}

func (base *BaseChart) AddBarSeries(ls *series.PointSeries) (err error) {
	ls.MakeBar()
	err = base.addSeriesIfNotExist(ls)
	return
}

func (base *BaseChart) AddAreaSeries(ls *series.PointSeries, showDot bool) (err error) {
	ls.MakeArea(showDot)
	err = base.addSeriesIfNotExist(ls)
	return
}

func (base *BaseChart) AddLineSeries(ls *series.PointSeries, showDot bool) (err error) {
	ls.MakeLine(showDot)
	err = base.addSeriesIfNotExist(ls)
	return
}

func (base *BaseChart) AddLollipopSeries(ls *series.PointSeries) (err error) {
	ls.MakeLollipop()
	err = base.addSeriesIfNotExist(ls)
	return
}

func (base *BaseChart) AddScatterSeries(ss *series.PointSeries) (err error) {
	ss.MakeScatter()
	err = base.addSeriesIfNotExist(ss)
	return
}

func (base *BaseChart) AddCandleStickSeries(cs *series.CandleStickSeries) (err error) {
	err = base.addSeriesIfNotExist(cs)
	return
}

func (base *BaseChart) AddBoxSeries(bs *series.BoxSeries) (err error) {
	err = base.addSeriesIfNotExist(bs)
	return
}

func (base *BaseChart) AddStackedBarSeries(sbs *series.StackedSeries) (err error) {
	err = base.addSeriesIfNotExist(sbs)
	return
}
