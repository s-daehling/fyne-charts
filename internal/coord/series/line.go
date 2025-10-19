package series

import (
	"image/color"
)

type LineSeries struct {
	dataPointSeries
}

func EmptyLineSeries(chart chart, name string, showDots bool, color color.Color, polar bool) (ser *LineSeries) {
	ser = &LineSeries{
		dataPointSeries: dataPointSeries{
			valBase:             0,
			nBarWidth:           0,
			tBarWidth:           0,
			nBarShift:           0,
			tBarShift:           0,
			showDot:             showDots,
			showFromValBaseLine: false,
			showFromPrevLine:    true,
			showBar:             false,
			sortPoints:          true,
		},
	}
	ser.baseSeries = emptyBaseSeries(chart, name, color, polar, ser.toggleView)
	return
}
