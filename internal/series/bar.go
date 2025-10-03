package series

import (
	"image/color"
)

type BarSeries struct {
	dataPointSeries
}

func EmptyBarSeries(chart chart, name string, color color.Color, polar bool) (ser *BarSeries) {
	ser = &BarSeries{
		dataPointSeries: dataPointSeries{
			valBase:             0,
			nBarWidth:           0,
			tBarWidth:           0,
			nBarShift:           0,
			tBarShift:           0,
			showDot:             false,
			showFromValBaseLine: false,
			showFromPrevLine:    false,
			showBar:             true,
			sortPoints:          false,
		},
	}
	ser.baseSeries = emptyBaseSeries(chart, name, color, polar, ser.toggleView)
	return
}
