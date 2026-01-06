package coord

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/s-daehling/fyne-charts/internal/coord/series"
)

func (base *BaseChart) Refresh() {
	if base.render != nil {
		base.render.Refresh()
	}
}

func (base *BaseChart) DataChange() {
	base.updateRangeAndOrigin()
	base.updateAxTicks()
	base.updateSeriesVariables()
	base.Refresh()
}

func (base *BaseChart) RasterVisibilityChange() {
	base.rast.Refresh()
}

func (base *BaseChart) ChartSizeChange(fromSpace float32, toSpace float32) {
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
		if base.autoFromRange && base.planeType == CartesianPlane {
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
	}
	base.toAx.AutoNTicks()
}

func (base *BaseChart) updateSeriesVariables() {
	nBarSeries := 0
	maxBoxPoints := 5
	for i := range base.series {
		if ser, ok := base.series[i].(*series.PointSeries); ok {
			if ser.IsBarSeries() {
				nBarSeries++
			}
		} else if _, ok := base.series[i].(*series.StackedSeries); ok {
			nBarSeries++
		} else if bs, ok := base.series[i].(*series.BoxSeries); ok {
			n := bs.NumberOfPoints()
			if n > maxBoxPoints {
				maxBoxPoints = n
			}
		}
	}
	nFromMin, nFromMax := base.fromAx.NRange()
	catSize := (nFromMax - nFromMin) * 0.9
	numCategories := len(base.fromAx.CRange())
	if numCategories > 0 {
		catSize = ((nFromMax - nFromMin) / float64(numCategories)) * 0.9
	}
	barWidth := catSize
	if nBarSeries > 0 {
		barWidth = catSize / float64(nBarSeries)
	}
	barOffset := -barWidth * (0.5 * float64(nBarSeries-1))
	boxWidth := (nFromMax - nFromMin) / float64(maxBoxPoints)
	for i := range base.series {
		if ser, ok := base.series[i].(*series.PointSeries); ok {
			if ser.IsBarSeries() {
				if base.fromType == Categorical {
					ser.SetNumericalBarWidthAndShift(barWidth, barOffset)
					barOffset += barWidth
				}
				if base.planeType == CartesianPlane {
					ser.SetValBaseNumerical(base.toAx.NOrigin())
				}
			} else if ser.IsLollipopSeries() && base.planeType == CartesianPlane {
				ser.SetValBaseNumerical(base.toAx.NOrigin())
			} else if ser.IsAreaSeries() && base.planeType == CartesianPlane {
				ser.SetValBaseNumerical(base.toAx.NOrigin())
			}
		} else if sbs, ok := base.series[i].(*series.StackedSeries); ok {
			if base.fromType == Categorical {
				sbs.SetNumericalBarWidthAndShift(barWidth, barOffset)
				barOffset += barWidth
			}
			sbs.UpdateValOffset()
		} else if bs, ok := base.series[i].(*series.BoxSeries); ok {
			bs.SetWidth(boxWidth)
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
	}
}

func (base *BaseChart) updateHLabelSpacer() {
	base.hLabelRightSpacer.SetMinSize(fyne.NewSize(base.legend.MinSize().Width, 0))
	switch base.planeType {
	case CartesianPlane:
		if base.transposed {
			base.hLabelLeftSpacer.SetMinSize(fyne.NewSize(base.fromAx.Label().Size().Width, 0))
		} else {
			base.hLabelLeftSpacer.SetMinSize(fyne.NewSize(base.toAx.Label().Size().Width, 0))
		}
	case PolarPlane:
		base.hLabelLeftSpacer.SetMinSize(fyne.NewSize(base.fromAx.Label().Size().Width, 0))
	}
}

func (base *BaseChart) RefreshTheme() {
	base.fromAx.RefreshTheme()
	base.toAx.RefreshTheme()
	base.title.TextSize = theme.Size(base.titleStyle.SizeName)
	base.title.Color = theme.Color(base.titleStyle.ColorName)
	base.tooltip.RefreshTheme()
	for i := range base.series {
		base.series[i].RefreshTheme()
	}
}
