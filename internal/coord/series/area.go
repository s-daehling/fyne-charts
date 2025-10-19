package series

import (
	"image/color"
)

type AreaSeries struct {
	dataPointSeries
}

func EmptyAreaSeries(chart chart, name string, showDots bool, color color.Color, polar bool) (ser *AreaSeries) {
	ser = &AreaSeries{
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

func (ser *AreaSeries) RasterColorCartesian(x float64, y float64) (col color.Color) {
	col = ser.baseSeries.RasterColorCartesian(x, y)
	if !ser.visible {
		return
	}
	// find first data point with x higher
	for i := range ser.data {
		if ser.data[i].n > x {
			if i == 0 {
				break
			}
			x1 := ser.data[i-1].n
			x2 := ser.data[i].n
			y1 := ser.data[i-1].val
			y2 := ser.data[i].val
			// interpolate
			yS := y1 + (((x - x1) / (x2 - x1)) * (y2 - y1))
			if yS > ser.valBase && y > ser.valBase && y < yS {
				r, g, b, _ := ser.color.RGBA()
				col = color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: 0x8888}

			} else if yS < ser.valBase && y < ser.valBase && y > yS {
				r, g, b, _ := ser.color.RGBA()
				col = color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: 0x8888}
			}
			break
		}
	}
	return
}

func (ser *AreaSeries) RasterColorPolar(phi float64, r float64, x float64, y float64) (col color.Color) {
	col = ser.baseSeries.RasterColorPolar(phi, r, x, y)
	if !ser.visible {
		return
	}
	red, green, blue, _ := ser.color.RGBA()
	colArea := color.RGBA64{R: uint16(red), G: uint16(green), B: uint16(blue), A: 0x8888}
	// find first data point with x higher
	for i := range ser.data {
		if ser.data[i].n > phi {
			if i == 0 {
				break
			}
			phi1 := ser.data[i-1].n
			phi2 := ser.data[i].n
			r1 := ser.data[i-1].val
			r2 := ser.data[i].val
			R := r1 + (((phi - phi1) / (phi2 - phi1)) * (r2 - r1))
			if r < R {
				col = colArea
			}
			break
		}
	}
	return
}

func (ser *AreaSeries) toggleView() {
	ser.dataPointSeries.toggleView()
	ser.chart.RasterVisibilityChange()
}
