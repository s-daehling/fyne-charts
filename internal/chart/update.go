package chart

import "github.com/s-daehling/fyne-charts/internal/series"

func (base *BaseChart) DataChange() {
	base.updateRangeAndOrigin()
	base.updateAxTicks()
	base.updateSeriesVariables()
	base.render.Refresh()
}

func (base *BaseChart) RasterVisibilityChange() {
	base.rast.Refresh()
}

func (base *BaseChart) resize(fromSpace float32, toSpace float32) {
	base.fromAx.SetSpace(fromSpace)
	base.toAx.SetSpace(toSpace)
	base.updateAxTicks()
}

func (base *BaseChart) updateRangeAndOrigin() {
	switch base.fromType {
	case Numerical:
		if base.autoToRange {
			base.calculateAutoToRange()
		}
		if base.autoFromRange || base.planeType == CartesianPlane {
			base.calculateAutoFromNRange()
		}
		if base.autoOrigin {
			base.calculateAutoNOrigin()
		}
	case Temporal:
		if base.autoToRange {
			base.calculateAutoToRange()
		}
		if base.autoFromRange {
			base.calculateAutoFromTRange()
		}
		if base.autoOrigin {
			base.calculateAutoTOrigin()
		}
	case Categorical:
		if base.autoToRange {
			base.calculateAutoToRange()
		}
		if base.autoFromRange {
			base.calculateAutoFromCRange()
		}
		if base.autoOrigin {
			base.calculateAutoNOrigin()
		}
	case Proportional:
		if base.autoOrigin {
			base.calculateAutoNOrigin()
		}
	}
}

func (base *BaseChart) updateAxTicks() {
	switch base.fromType {
	case Numerical:
		base.fromAx.AutoNTicks()
	case Temporal:
		base.fromAx.AutoTTicks()
		base.fromAx.ConvertTTickstoN()
	case Categorical:
		base.fromAx.AutoCTicks()
		base.fromAx.ConvertCTickstoN()
	case Proportional:
		base.fromAx.AutoNTicks()
	}
	base.toAx.AutoNTicks()
}

func (base *BaseChart) updateSeriesVariables() {
	nBarSeries := 0
	nPropSeries := 0
	maxBoxPoints := 5
	for i := range base.series {
		if _, ok := base.series[i].(*series.BarSeries); ok {
			nBarSeries++
		} else if _, ok := base.series[i].(*series.StackedBarSeries); ok {
			nBarSeries++
		} else if bs, ok := base.series[i].(*series.BoxSeries); ok {
			n := bs.NumberOfPoints()
			if n > maxBoxPoints {
				maxBoxPoints = n
			}
		} else if _, ok := base.series[i].(*series.ProportionalSeries); ok {
			nPropSeries++
		}
	}
	nFromMin, nFromMax := base.fromAx.NRange()
	nToMin, nToMax := base.toAx.NRange()
	catSize := (nFromMax - nFromMin) * 0.9
	numCategories := len(base.fromAxis().CRange())
	if numCategories > 0 {
		catSize = ((nFromMax - nFromMin) / float64(numCategories)) * 0.9
	}
	barWidth := catSize
	if nBarSeries > 0 {
		barWidth = catSize / float64(nBarSeries)
	}
	barOffset := -barWidth * (0.5 * float64(nBarSeries-1))
	propHeight := (nToMax - nToMin) / float64(nPropSeries)
	propOffset := 0.0
	boxWidth := (nFromMax - nFromMin) / float64(maxBoxPoints)
	for i := range base.series {
		if ls, ok := base.series[i].(*series.LollipopSeries); ok {
			if base.planeType == CartesianPlane {
				ls.SetValBaseNumerical(base.toAx.NOrigin())
			}
		} else if bs, ok := base.series[i].(*series.BarSeries); ok {
			if base.fromType == Categorical {
				bs.SetNumericalBarWidthAndShift(barWidth, barOffset)
				barOffset += barWidth
			}
		} else if sbs, ok := base.series[i].(*series.StackedBarSeries); ok {
			if base.fromType == Categorical {
				sbs.SetNumericalBarWidthAndShift(barWidth, barOffset)
				barOffset += barWidth
			}
			sbs.UpdateValOffset()
		} else if bs, ok := base.series[i].(*series.BoxSeries); ok {
			bs.SetWidth(boxWidth)
		} else if as, ok := base.series[i].(*series.AreaSeries); ok {
			as.SetValBaseNumerical(base.toAx.NOrigin())
		} else if ps, ok := base.series[i].(*series.ProportionalSeries); ok {
			ps.SetHeightAndOffset(propHeight*0.9, propOffset)
			propOffset += propHeight
			// ps.UpdateValOffest()
		}
	}

	switch base.fromType {
	case Temporal:
		for i := range base.series {
			base.series[i].ConvertTtoN(base.fromAx.TtoN)
		}
	case Categorical:
		for i := range base.series {
			base.series[i].ConvertCtoN(base.fromAx.CtoN)
		}
	case Proportional:
		for i := range base.series {
			base.series[i].ConvertPtoN(base.fromAx.PtoN)
		}
	}
}
