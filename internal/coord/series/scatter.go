package series

import "image/color"

type ScatterSeries struct {
	dataPointSeries
}

func EmptyScatterSeries(chart chart, name string, color color.Color, polar bool) (ser *ScatterSeries) {
	ser = &ScatterSeries{
		dataPointSeries: dataPointSeries{
			valBase:             0,
			nBarWidth:           0,
			tBarWidth:           0,
			nBarShift:           0,
			tBarShift:           0,
			showDot:             true,
			showFromValBaseLine: false,
			showFromPrevLine:    false,
			showBar:             false,
			sortPoints:          false,
		},
	}
	ser.baseSeries = emptyBaseSeries(chart, name, color, polar, ser.toggleView)
	return
}
