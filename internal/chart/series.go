package chart

import (
	"errors"
	"image/color"
	"time"

	"github.com/s-daehling/fyne-charts/internal/series"

	"github.com/s-daehling/fyne-charts/pkg/data"
)

func (base *BaseChart) addSeriesIfNotExist(ser series.Series) (err error) {
	base.mutex.Lock()
	for i := range base.series {
		if base.series[i].Name() == ser.Name() {
			err = errors.New("series already exists")
			base.mutex.Unlock()
			return
		}
	}
	base.series = append(base.series, ser)
	base.mutex.Unlock()
	base.DataChange()
	return
}

func (base *BaseChart) AddNumericalLineSeries(name string, points []data.NumericalDataPoint, showDots bool, providerFct func() []data.NumericalDataPoint,
	color color.Color) (ser *series.LineSeries, err error) {
	lSeries := series.EmptyLineSeries(base, name, showDots, color, base.planeType == PolarPlane)
	if providerFct != nil {
		lSeries.AddNumericalUpdateFct(providerFct)
		err = lSeries.UpdateData()
	} else {
		err = lSeries.AddNumericalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(lSeries)
	if err != nil {
		return
	}
	ser = lSeries
	return
}

func (base *BaseChart) AddTemporalLineSeries(name string, points []data.TemporalDataPoint, showDots bool, providerFct func() []data.TemporalDataPoint,
	color color.Color) (ser *series.LineSeries, err error) {
	lSeries := series.EmptyLineSeries(base, name, showDots, color, base.planeType == PolarPlane)
	if providerFct != nil {
		lSeries.AddTemporalUpdateFct(providerFct)
		err = lSeries.UpdateData()
	} else {
		err = lSeries.AddTemporalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(lSeries)
	if err != nil {
		return
	}
	ser = lSeries
	return
}

func (base *BaseChart) AddNumericalScatterSeries(name string, points []data.NumericalDataPoint, providerFct func() []data.NumericalDataPoint,
	color color.Color) (ser *series.ScatterSeries, err error) {
	sSeries := series.EmptyScatterSeries(base, name, color, base.planeType == PolarPlane)
	if providerFct != nil {
		sSeries.AddNumericalUpdateFct(providerFct)
		err = sSeries.UpdateData()
	} else {
		err = sSeries.AddNumericalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(sSeries)
	if err != nil {
		return
	}
	ser = sSeries
	return
}

func (base *BaseChart) AddTemporalScatterSeries(name string, points []data.TemporalDataPoint, providerFct func() []data.TemporalDataPoint,
	color color.Color) (ser *series.ScatterSeries, err error) {
	sSeries := series.EmptyScatterSeries(base, name, color, base.planeType == PolarPlane)
	if providerFct != nil {
		sSeries.AddTemporalUpdateFct(providerFct)
		err = sSeries.UpdateData()
	} else {
		err = sSeries.AddTemporalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(sSeries)
	if err != nil {
		return
	}
	ser = sSeries
	return
}

func (base *BaseChart) AddCategoricalScatterSeries(name string, points []data.CategoricalDataPoint, providerFct func() []data.CategoricalDataPoint,
	color color.Color) (ser *series.ScatterSeries, err error) {
	sSeries := series.EmptyScatterSeries(base, name, color, base.planeType == PolarPlane)
	if providerFct != nil {
		sSeries.AddCategoricalUpdateFct(providerFct)
		err = sSeries.UpdateData()
	} else {
		err = sSeries.AddCategoricalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(sSeries)
	if err != nil {
		return
	}
	ser = sSeries
	return
}

func (base *BaseChart) AddNumericalLollipopSeries(name string, points []data.NumericalDataPoint, providerFct func() []data.NumericalDataPoint,
	color color.Color) (ser *series.LollipopSeries, err error) {
	lSeries := series.EmptyLollipopSeries(base, name, color, base.planeType == PolarPlane)
	if providerFct != nil {
		lSeries.AddNumericalUpdateFct(providerFct)
		err = lSeries.UpdateData()
	} else {
		err = lSeries.AddNumericalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(lSeries)
	if err != nil {
		return
	}
	ser = lSeries
	return
}

func (base *BaseChart) AddTemporalLollipopSeries(name string, points []data.TemporalDataPoint, providerFct func() []data.TemporalDataPoint,
	color color.Color) (ser *series.LollipopSeries, err error) {
	lSeries := series.EmptyLollipopSeries(base, name, color, base.planeType == PolarPlane)
	if providerFct != nil {
		lSeries.AddTemporalUpdateFct(providerFct)
		err = lSeries.UpdateData()
	} else {
		err = lSeries.AddTemporalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(lSeries)
	if err != nil {
		return
	}
	ser = lSeries
	return
}

func (base *BaseChart) AddCategoricalLollipopSeries(name string, points []data.CategoricalDataPoint, providerFct func() []data.CategoricalDataPoint,
	color color.Color) (ser *series.LollipopSeries, err error) {
	lSeries := series.EmptyLollipopSeries(base, name, color, base.planeType == PolarPlane)
	if providerFct != nil {
		lSeries.AddCategoricalUpdateFct(providerFct)
		err = lSeries.UpdateData()
	} else {
		err = lSeries.AddCategoricalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(lSeries)
	if err != nil {
		return
	}
	ser = lSeries
	return
}

func (base *BaseChart) AddNumericalAreaSeries(name string, points []data.NumericalDataPoint, showDots bool, providerFct func() []data.NumericalDataPoint,
	color color.Color) (ser *series.AreaSeries, err error) {
	aSeries := series.EmptyAreaSeries(base, name, showDots, color, base.planeType == PolarPlane)
	if providerFct != nil {
		aSeries.AddNumericalUpdateFct(providerFct)
		err = aSeries.UpdateData()
	} else {
		err = aSeries.AddNumericalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(aSeries)
	if err != nil {
		return
	}
	ser = aSeries
	return
}

func (base *BaseChart) AddTemporalAreaSeries(name string, points []data.TemporalDataPoint, showDots bool, providerFct func() []data.TemporalDataPoint,
	color color.Color) (ser *series.AreaSeries, err error) {
	aSeries := series.EmptyAreaSeries(base, name, showDots, color, base.planeType == PolarPlane)
	if providerFct != nil {
		aSeries.AddTemporalUpdateFct(providerFct)
		err = aSeries.UpdateData()
	} else {
		err = aSeries.AddTemporalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(aSeries)
	if err != nil {
		return
	}
	ser = aSeries
	return
}

func (base *BaseChart) AddNumericalCandleStickSeries(name string,
	points []data.NumericalCandleStick, providerFct func() []data.NumericalCandleStick) (ser *series.CandleStickSeries, err error) {
	csSeries := series.EmptyCandleStickSeries(base, name, base.planeType == PolarPlane)
	if providerFct != nil {
		csSeries.AddNumericalUpdateFct(providerFct)
		err = csSeries.UpdateData()
	} else {
		err = csSeries.AddNumericalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(csSeries)
	if err != nil {
		return
	}
	ser = csSeries
	return
}

func (base *BaseChart) AddTemporalCandleStickSeries(name string,
	points []data.TemporalCandleStick, providerFct func() []data.TemporalCandleStick) (ser *series.CandleStickSeries, err error) {
	csSeries := series.EmptyCandleStickSeries(base, name, base.planeType == PolarPlane)
	if providerFct != nil {
		csSeries.AddTemporalUpdateFct(providerFct)
		err = csSeries.UpdateData()
	} else {
		err = csSeries.AddTemporalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(csSeries)
	if err != nil {
		return
	}
	ser = csSeries
	return
}

func (base *BaseChart) AddNumericalBoxSeries(name string, points []data.NumericalBox, providerFct func() []data.NumericalBox,
	color color.Color) (ser *series.BoxSeries, err error) {
	bSeries := series.EmptyBoxSeries(base, name, color, base.planeType == PolarPlane)
	if providerFct != nil {
		bSeries.AddNumericalUpdateFct(providerFct)
		err = bSeries.UpdateData()
	} else {
		err = bSeries.AddNumericalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(bSeries)
	if err != nil {
		return
	}
	ser = bSeries
	return
}

func (base *BaseChart) AddTemporalBoxSeries(name string, points []data.TemporalBox, providerFct func() []data.TemporalBox,
	color color.Color) (ser *series.BoxSeries, err error) {
	bSeries := series.EmptyBoxSeries(base, name, color, base.planeType == PolarPlane)
	if providerFct != nil {
		bSeries.AddTemporalUpdateFct(providerFct)
		err = bSeries.UpdateData()
	} else {
		err = bSeries.AddTemporalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(bSeries)
	if err != nil {
		return
	}
	ser = bSeries
	return
}

func (base *BaseChart) AddCategoricalBoxSeries(name string, points []data.CategoricalBox, providerFct func() []data.CategoricalBox,
	color color.Color) (ser *series.BoxSeries, err error) {
	bSeries := series.EmptyBoxSeries(base, name, color, base.planeType == PolarPlane)
	if providerFct != nil {
		bSeries.AddCategoricalUpdateFct(providerFct)
		err = bSeries.UpdateData()
	} else {
		err = bSeries.AddCategoricalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(bSeries)
	if err != nil {
		return
	}
	ser = bSeries
	return
}

func (base *BaseChart) AddNumericalBarSeries(name string, points []data.NumericalDataPoint,
	barWidth float64, providerFct func() []data.NumericalDataPoint, color color.Color) (ser *series.BarSeries, err error) {
	bSeries := series.EmptyBarSeries(base, name, color, base.planeType == PolarPlane)
	if providerFct != nil {
		bSeries.AddNumericalUpdateFct(providerFct)
		err = bSeries.UpdateData()
	} else {
		err = bSeries.AddNumericalData(points)
	}
	if err != nil {
		return
	}
	err = bSeries.SetNumericalWidthAndOffset(barWidth, 0)
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(bSeries)
	if err != nil {
		return
	}
	ser = bSeries
	return
}

func (base *BaseChart) AddTemporalBarSeries(name string, points []data.TemporalDataPoint,
	barWidth time.Duration, providerFct func() []data.TemporalDataPoint, color color.Color) (ser *series.BarSeries, err error) {
	bSeries := series.EmptyBarSeries(base, name, color, base.planeType == PolarPlane)
	if providerFct != nil {
		bSeries.AddTemporalUpdateFct(providerFct)
		err = bSeries.UpdateData()
	} else {
		err = bSeries.AddTemporalData(points)
	}
	if err != nil {
		return
	}
	err = bSeries.SetTemporalWidthAndOffset(barWidth, 0)
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(bSeries)
	if err != nil {
		return
	}
	ser = bSeries
	return
}

func (base *BaseChart) AddCategoricalBarSeries(name string, points []data.CategoricalDataPoint, providerFct func() []data.CategoricalDataPoint,
	color color.Color) (ser *series.BarSeries, err error) {
	bSeries := series.EmptyBarSeries(base, name, color, base.planeType == PolarPlane)
	if providerFct != nil {
		bSeries.AddCategoricalUpdateFct(providerFct)
		err = bSeries.UpdateData()
	} else {
		err = bSeries.AddCategoricalData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(bSeries)
	if err != nil {
		return
	}
	ser = bSeries
	return
}

func (base *BaseChart) AddCategoricalStackedBarSeries(name string,
	dataSeries []data.CategoricalDataSeries, providerFct func() []data.CategoricalDataSeries) (ser *series.StackedBarSeries, err error) {
	sbSeries := series.EmptyStackedBarSeries(base, name, base.planeType == PolarPlane)
	if providerFct != nil {
		sbSeries.AddCategoricalUpdateFct(providerFct)
		err = sbSeries.UpdateData()
		if err != nil {
			return
		}
	}
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

func (base *BaseChart) AddProportionalSeries(name string, points []data.ProportionalDataPoint,
	providerFct func() []data.ProportionalDataPoint) (ser *series.ProportionalSeries, err error) {
	pSeries := series.EmptyProportionalSeries(base, name, base.planeType == PolarPlane)
	if providerFct != nil {
		pSeries.AddUpdateFct(providerFct)
		err = pSeries.UpdateData()
	} else {
		err = pSeries.AddData(points)
	}
	if err != nil {
		return
	}
	err = base.addSeriesIfNotExist(pSeries)
	if err != nil {
		return
	}
	ser = pSeries
	return
}
