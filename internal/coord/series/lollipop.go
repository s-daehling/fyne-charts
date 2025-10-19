package series

import (
	"image/color"
)

type LollipopSeries struct {
	dataPointSeries
}

func EmptyLollipopSeries(chart chart, name string, color color.Color, polar bool) (ser *LollipopSeries) {
	ser = &LollipopSeries{
		dataPointSeries: dataPointSeries{
			valBase:             0,
			nBarWidth:           0,
			tBarWidth:           0,
			nBarShift:           0,
			tBarShift:           0,
			showDot:             true,
			showFromValBaseLine: true,
			showFromPrevLine:    false,
			showBar:             false,
			sortPoints:          false,
		},
	}
	ser.baseSeries = emptyBaseSeries(chart, name, color, polar, ser.toggleView)
	return
}
