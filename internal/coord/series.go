package coord

import (
	"errors"

	"github.com/s-daehling/fyne-charts/internal/coord/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

func (base *BaseChart) addSeriesIfNotExist(ser series.Series) (err error) {
	for i := range base.series {
		if base.series[i].Name() == ser.Name() {
			err = errors.New("series already exists")
			return
		}
	}
	err = ser.Bind(base)
	if err != nil {
		return
	}
	base.series = append(base.series, ser)
	base.DataChange()
	return
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

func (base *BaseChart) AddCategoricalStackedBarSeries(name string,
	dataSeries []data.CategoricalDataSeries) (ser *series.StackedBarSeries, err error) {
	sbSeries := series.EmptyStackedBarSeries(name)
	for i := range dataSeries {
		err = sbSeries.AddCategoricalSeries(dataSeries[i])
		if err != nil {
			return
		}
	}
	err = base.addSeriesIfNotExist(sbSeries)
	if err != nil {
		return
	}
	ser = sbSeries
	return
}
