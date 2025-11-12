package coord

import (
	"errors"
	"image/color"
	"time"

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
	base.series = append(base.series, ser)
	base.DataChange()
	return
}

func (base *BaseChart) AddNumericalLineSeries(name string, points []data.NumericalPoint, showDots bool,
	color color.Color) (ser *series.PointSeries, err error) {
	lSeries := series.EmptyLineSeries(base, name, showDots, color, base.planeType == PolarPlane)
	err = lSeries.AddNumericalData(points)
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

func (base *BaseChart) AddTemporalLineSeries(name string, points []data.TemporalPoint, showDots bool,
	color color.Color) (ser *series.PointSeries, err error) {
	lSeries := series.EmptyLineSeries(base, name, showDots, color, base.planeType == PolarPlane)
	err = lSeries.AddTemporalData(points)
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

func (base *BaseChart) AddNumericalScatterSeries(name string, points []data.NumericalPoint,
	color color.Color) (ser *series.PointSeries, err error) {
	sSeries := series.EmptyScatterSeries(base, name, color, base.planeType == PolarPlane)
	err = sSeries.AddNumericalData(points)
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

func (base *BaseChart) AddTemporalScatterSeries(name string, points []data.TemporalPoint,
	color color.Color) (ser *series.PointSeries, err error) {
	sSeries := series.EmptyScatterSeries(base, name, color, base.planeType == PolarPlane)
	err = sSeries.AddTemporalData(points)
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

func (base *BaseChart) AddCategoricalScatterSeries(name string, points []data.CategoricalPoint,
	color color.Color) (ser *series.PointSeries, err error) {
	sSeries := series.EmptyScatterSeries(base, name, color, base.planeType == PolarPlane)
	err = sSeries.AddCategoricalData(points)
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

func (base *BaseChart) AddNumericalLollipopSeries(name string, points []data.NumericalPoint,
	color color.Color) (ser *series.PointSeries, err error) {
	lSeries := series.EmptyLollipopSeries(base, name, color, base.planeType == PolarPlane)
	err = lSeries.AddNumericalData(points)
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

func (base *BaseChart) AddTemporalLollipopSeries(name string, points []data.TemporalPoint,
	color color.Color) (ser *series.PointSeries, err error) {
	lSeries := series.EmptyLollipopSeries(base, name, color, base.planeType == PolarPlane)
	err = lSeries.AddTemporalData(points)
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

func (base *BaseChart) AddCategoricalLollipopSeries(name string, points []data.CategoricalPoint,
	color color.Color) (ser *series.PointSeries, err error) {
	lSeries := series.EmptyLollipopSeries(base, name, color, base.planeType == PolarPlane)
	err = lSeries.AddCategoricalData(points)
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

func (base *BaseChart) AddNumericalAreaSeries(name string, points []data.NumericalPoint, showDots bool,
	color color.Color) (ser *series.PointSeries, err error) {
	aSeries := series.EmptyAreaSeries(base, name, showDots, color, base.planeType == PolarPlane)
	err = aSeries.AddNumericalData(points)
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

func (base *BaseChart) AddTemporalAreaSeries(name string, points []data.TemporalPoint, showDots bool,
	color color.Color) (ser *series.PointSeries, err error) {
	aSeries := series.EmptyAreaSeries(base, name, showDots, color, base.planeType == PolarPlane)
	err = aSeries.AddTemporalData(points)
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
	points []data.NumericalCandleStick) (ser *series.CandleStickSeries, err error) {
	csSeries := series.EmptyCandleStickSeries(base, name, base.planeType == PolarPlane)
	err = csSeries.AddNumericalData(points)
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
	points []data.TemporalCandleStick) (ser *series.CandleStickSeries, err error) {
	csSeries := series.EmptyCandleStickSeries(base, name, base.planeType == PolarPlane)
	err = csSeries.AddTemporalData(points)
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

func (base *BaseChart) AddNumericalBoxSeries(name string, points []data.NumericalBox,
	color color.Color) (ser *series.BoxSeries, err error) {
	bSeries := series.EmptyBoxSeries(base, name, color, base.planeType == PolarPlane)
	err = bSeries.AddNumericalData(points)
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

func (base *BaseChart) AddTemporalBoxSeries(name string, points []data.TemporalBox,
	color color.Color) (ser *series.BoxSeries, err error) {
	bSeries := series.EmptyBoxSeries(base, name, color, base.planeType == PolarPlane)
	err = bSeries.AddTemporalData(points)
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

func (base *BaseChart) AddCategoricalBoxSeries(name string, points []data.CategoricalBox,
	color color.Color) (ser *series.BoxSeries, err error) {
	bSeries := series.EmptyBoxSeries(base, name, color, base.planeType == PolarPlane)
	err = bSeries.AddCategoricalData(points)
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

func (base *BaseChart) AddNumericalBarSeries(name string, points []data.NumericalPoint,
	barWidth float64, color color.Color) (ser *series.PointSeries, err error) {
	bSeries := series.EmptyBarSeries(base, name, color, base.planeType == PolarPlane)
	err = bSeries.AddNumericalData(points)
	if err != nil {
		return
	}
	err = bSeries.SetNumericalBarWidthAndShift(barWidth, 0)
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

func (base *BaseChart) AddTemporalBarSeries(name string, points []data.TemporalPoint,
	barWidth time.Duration, color color.Color) (ser *series.PointSeries, err error) {
	bSeries := series.EmptyBarSeries(base, name, color, base.planeType == PolarPlane)
	err = bSeries.AddTemporalData(points)
	if err != nil {
		return
	}
	err = bSeries.SetTemporalBarWidthAndShift(barWidth, 0)
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

func (base *BaseChart) AddCategoricalBarSeries(name string, points []data.CategoricalPoint,
	color color.Color) (ser *series.PointSeries, err error) {
	bSeries := series.EmptyBarSeries(base, name, color, base.planeType == PolarPlane)
	err = bSeries.AddCategoricalData(points)
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
	dataSeries []data.CategoricalDataSeries) (ser *series.StackedBarSeries, err error) {
	sbSeries := series.EmptyStackedBarSeries(base, name, base.planeType == PolarPlane)
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
