package coord

import (
	"errors"

	"github.com/s-daehling/fyne-charts/internal/coord"
	"github.com/s-daehling/fyne-charts/pkg/data"
	"github.com/s-daehling/fyne-charts/pkg/style"
)

// PolarCategoricalChart implements a polar plane with one categorical c-axis and one numerical r-axis
type PolarCategoricalChart struct {
	coordChart
}

// NewPolarCategoricalChart returns an initialized PolarCategoricalChart
func NewPolarCategoricalChart(title string) (catChart *PolarCategoricalChart) {
	catChart = &PolarCategoricalChart{
		coordChart: emptyCoordChart(coord.PolarPlane, coord.Categorical),
	}
	catChart.ExtendBaseWidget(catChart)
	catChart.SetTitle(title)
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// An error is returned,if another series with the same name exists, if the series is already added to another chart or if Val < 0 for one or more points
func (catChart *PolarCategoricalChart) AddScatterSeries(cps *CategoricalPointSeries) (err error) {
	if catChart.base == nil || cps == nil {
		return
	}
	if cps.ser == nil {
		err = errors.New("series not initialized")
		return
	}
	err = catChart.base.AddScatterSeries(cps.ser)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// An error is returned,if another series with the same name exists, if the series is already added to another chart or if Val < 0 for one or more points
func (catChart *PolarCategoricalChart) AddLollipopSeries(cps *CategoricalPointSeries) (err error) {
	if catChart.base == nil || cps == nil {
		return
	}
	if cps.ser == nil {
		err = errors.New("series not initialized")
		return
	}
	err = catChart.base.AddLollipopSeries(cps.ser)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// An error is returned,if another series with the same name exists, if the series is already added to another chart or if Val < 0 for one or more points
func (catChart *PolarCategoricalChart) AddBarSeries(cps *CategoricalPointSeries) (err error) {
	if catChart.base == nil || cps == nil {
		return
	}
	if cps.ser == nil {
		err = errors.New("series not initialized")
		return
	}
	err = catChart.base.AddBarSeries(cps.ser)
	return
}

// AddStackedBarSeries adds a series of data which is visualized as stacked bar chart.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (catChart *PolarCategoricalChart) AddStackedBarSeries(css *CategoricalStackedSeries) (err error) {
	if catChart.base == nil || css == nil {
		return
	}
	if css.ser == nil {
		err = errors.New("series not initialized")
		return
	}
	err = catChart.base.AddStackedBarSeries(css.ser)
	return
}

// SetRAxisLabel sets the label of the r-axis, which will be displayed at the bottom
func (catChart *PolarCategoricalChart) SetRAxisLabel(l string) {
	if catChart.base == nil {
		return
	}
	catChart.base.SetToAxisLabel(l)
}

// SetRRange sets a user defined range for the r-axis;
// an error is returned if max<0 or if the origin has been defined by the user before and is outside the given range
func (catChart *PolarCategoricalChart) SetRRange(max float64) (err error) {
	if catChart.base == nil {
		return
	}
	err = catChart.base.SetToRange(0.0, max)
	return
}

// SetAutoRRange overrides a previously user defined range and lets the range be calculated automatically
func (catChart *PolarCategoricalChart) SetAutoRRange() {
	if catChart.base == nil {
		return
	}
	catChart.base.SetAutoToRange()
}

// SetRTicks sets the list of user defined ticks to be shown on the r-axis
func (catChart *PolarCategoricalChart) SetRTicks(ts []data.NumericalTick) {
	if catChart.base == nil {
		return
	}
	catChart.base.SetToTicks(ts)
}

// SetAutoRTicks overrides a previously user defined set of r-axis ticks and lets the ticks be calculated automatically
func (catChart *PolarCategoricalChart) SetAutoRTicks(autoSupportLine bool) {
	if catChart.base == nil {
		return
	}
	catChart.base.SetAutoToTicks(autoSupportLine)
}

// SetRAxisStyle changes the style of the R-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (catChart *PolarCategoricalChart) SetRAxisStyle(labelStyle style.LabelStyle,
	axisStyle style.AxisStyle) {
	if catChart.base == nil {
		return
	}
	catChart.base.SetToAxisLabelStyle(labelStyle)
	catChart.base.SetToAxisStyle(axisStyle)
}

// SetCAxisLabel sets the label of the c-axis, which will be displayed at the left side
func (catChart *PolarCategoricalChart) SetCAxisLabel(l string) {
	if catChart.base == nil {
		return
	}
	catChart.base.SetFromAxisLabel(l)
}

// SetCRange sets a user defined range for the c-axis.
// This will also determine the ticks and their order on the c-axis.
// An error is returned if cs is empty.
func (catChart *PolarCategoricalChart) SetCRange(cs []string) (err error) {
	if catChart.base == nil {
		return
	}
	err = catChart.base.SetFromCRange(cs)
	return
}

// SetAutoCRange overrides a previously user defined range and lets the range be calculated automatically
func (catChart *PolarCategoricalChart) SetAutoCRange() {
	if catChart.base == nil {
		return
	}
	catChart.base.SetAutoFromRange()
}

// SetCAxisStyle changes the style of the C-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (catChart *PolarCategoricalChart) SetCAxisStyle(labelStyle style.LabelStyle,
	axisStyle style.AxisStyle) {
	if catChart.base == nil {
		return
	}
	catChart.base.SetFromAxisLabelStyle(labelStyle)
	catChart.base.SetFromAxisStyle(axisStyle)
}
