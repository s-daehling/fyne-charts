package coord

import (
	"errors"

	"github.com/s-daehling/fyne-charts/internal/coord"
	"github.com/s-daehling/fyne-charts/pkg/data"
	"github.com/s-daehling/fyne-charts/pkg/style"
)

// CartesianCategoricalChart implements a cartesian plane with a categorical c-axis and a numerical y-axis
type CartesianCategoricalChart struct {
	coordChart
}

// NewCartesianCategoricalChart returns an initialized CategoricalChart
func NewCartesianCategoricalChart(title string) (catChart *CartesianCategoricalChart) {
	catChart = &CartesianCategoricalChart{
		coordChart: emptyCoordChart(coord.CartesianPlane, coord.Categorical),
	}
	catChart.ExtendBaseWidget(catChart)
	catChart.SetTitle(title)
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series must have a unique name throughout the chart.
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (catChart *CartesianCategoricalChart) AddScatterSeries(cps *CategoricalPointSeries) (err error) {
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

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series must have a unique name throughout the chart.
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (catChart *CartesianCategoricalChart) AddBarSeries(cps *CategoricalPointSeries) (err error) {
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
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (catChart *CartesianCategoricalChart) AddStackedBarSeries(css *CategoricalStackedSeries) (err error) {
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

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series must have a unique name throughout the chart.
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (catChart *CartesianCategoricalChart) AddLollipopSeries(cps *CategoricalPointSeries) (err error) {
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

// AddBoxSeries adds a series of data which is visualized as box chart.
// The series must have a unique name throughout the chart.
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (catChart *CartesianCategoricalChart) AddBoxSeries(cbs *CategoricalBoxSeries) (err error) {
	if catChart.base == nil || cbs == nil {
		return
	}
	if cbs.ser == nil {
		err = errors.New("series not initialized")
		return
	}
	err = catChart.base.AddBoxSeries(cbs.ser)
	return
}

// SetOrientation defines the orientation of axes
// if transposed is false, the C-axis is the horizontal axis and the Y-axis is the vertical axis
// if transposed is true, the C-axis is the vertical axis and the Y-axis is the horizontal axis
func (catChart *CartesianCategoricalChart) SetOrientation(transposed bool) {
	catChart.base.SetCartesianOrientantion(transposed)
}

// SetYAxisLabel sets the label of the y-axis, which will be displayed at the left side
func (catChart *CartesianCategoricalChart) SetYAxisLabel(l string) {
	if catChart.base == nil {
		return
	}
	catChart.base.SetToAxisLabel(l)
}

// SetYRange sets a user defined range for the y-axis;
// an error is returned if min>max or if the origin has been defined by the user before and is outside the given range
func (catChart *CartesianCategoricalChart) SetYRange(min float64, max float64) (err error) {
	if catChart.base == nil {
		return
	}
	err = catChart.base.SetToRange(min, max)
	return
}

// SetAutoYRange overrides a previously user defined range and lets the range be calculated automatically
func (catChart *CartesianCategoricalChart) SetAutoYRange() {
	if catChart.base == nil {
		return
	}
	catChart.base.SetAutoToRange()
}

// SetYTicks sets the list of user defined ticks to be shown on the y-axis
func (catChart *CartesianCategoricalChart) SetYTicks(ts []data.NumericalTick) {
	if catChart.base == nil {
		return
	}
	catChart.base.SetToTicks(ts)
}

// SetAutoYTicks overrides a previously user defined set of y-axis ticks and lets the ticks be calculated automatically
func (catChart *CartesianCategoricalChart) SetAutoYTicks(autoSupportLine bool) {
	if catChart.base == nil {
		return
	}
	catChart.base.SetAutoToTicks(autoSupportLine)
}

// SetYAxisStyle changes the style of the Y-axis
func (catChart *CartesianCategoricalChart) SetYAxisStyle(labelStyle style.ChartTextStyle,
	axisStyle style.AxisStyle) {
	if catChart.base == nil {
		return
	}
	catChart.base.SetToAxisLabelStyle(labelStyle)
	catChart.base.SetToAxisStyle(axisStyle)
}

// SetOrigin sets a user defined origin (crossing of c and y axis).
// An error is returned, if a range has been defined before and at least one coordinate is outside the range.
func (catChart *CartesianCategoricalChart) SetOrigin(y float64) (err error) {
	if catChart.base == nil {
		return
	}
	err = catChart.base.SetNOrigin(0.0, y)
	return
}

// SetAutoOrigin resets a previously user defined origin and allows the chart to calculate the ideal origin automatically
func (catChart *CartesianCategoricalChart) SetAutoOrigin() {
	if catChart.base == nil {
		return
	}
	catChart.base.SetAutoOrigin()
}

// SetCAxisLabel sets the label of the c-axis, which will be displayed at the bottom
func (catChart *CartesianCategoricalChart) SetCAxisLabel(l string) {
	if catChart.base == nil {
		return
	}
	catChart.base.SetFromAxisLabel(l)
}

// SetCRange sets a user defined range for the c-axis.
// This will also determine the ticks and their order on the c-axis.
// An error is returned if cs is empty.
func (catChart *CartesianCategoricalChart) SetCRange(cs []string) (err error) {
	if catChart.base == nil {
		return
	}
	err = catChart.base.SetFromCRange(cs)
	return
}

// SetAutoCRange overrides a previously user defined set of categories to be shown and lets the set be calculated automatically
func (catChart *CartesianCategoricalChart) SetAutoCRange() {
	if catChart.base == nil {
		return
	}
	catChart.base.SetAutoFromRange()
}

// SetCAxisStyle changes the style of the C-axis
func (catChart *CartesianCategoricalChart) SetCAxisStyle(labelStyle style.ChartTextStyle,
	axisStyle style.AxisStyle) {
	if catChart.base == nil {
		return
	}
	catChart.base.SetFromAxisLabelStyle(labelStyle)
	catChart.base.SetFromAxisStyle(axisStyle)
}
