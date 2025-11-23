package coord

import (
	"errors"

	"github.com/s-daehling/fyne-charts/internal/coord"

	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
)

// CartesianNumericalChart implements a cartesian plane with two numerical axes x and y
type CartesianNumericalChart struct {
	coordChart
}

// NewCartesianNumericalChart returns an initialized CartesianNumericalChart
func NewCartesianNumericalChart(title string) (numChart *CartesianNumericalChart) {
	numChart = &CartesianNumericalChart{
		coordChart: emptyCoordChart(coord.CartesianPlane, coord.Numerical),
	}
	numChart.ExtendBaseWidget(numChart)
	numChart.SetTitle(title)
	return
}

// AddLineSeries adds a series of data which is visualized as line chart.
// If showDots is true, dots are displayed at the osition of the series points.
// The series must have a unique name throughout the chart.
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (numChart *CartesianNumericalChart) AddLineSeries(nps *NumericalPointSeries, showDots bool) (err error) {
	if numChart.base == nil || nps == nil {
		return
	}
	if nps.ser == nil {
		err = errors.New("series not initialized")
		return
	}
	err = numChart.base.AddLineSeries(nps.ser, showDots)
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series must have a unique name throughout the chart.
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (numChart *CartesianNumericalChart) AddScatterSeries(nps *NumericalPointSeries) (err error) {
	if numChart.base == nil || nps == nil {
		return
	}
	if nps.ser == nil {
		err = errors.New("series not initialized")
		return
	}
	err = numChart.base.AddScatterSeries(nps.ser)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series must have a unique name throughout the chart.
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (numChart *CartesianNumericalChart) AddLollipopSeries(nps *NumericalPointSeries) (err error) {
	if numChart.base == nil || nps == nil {
		return
	}
	if nps.ser == nil {
		err = errors.New("series not initialized")
		return
	}
	err = numChart.base.AddLollipopSeries(nps.ser)
	return
}

// AddCandleStickSeries adds a series of data which is visualized as canlde stick chart.
// The series must have a unique name throughout the chart.
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (numChart *CartesianNumericalChart) AddCandleStickSeries(ncs *NumericalCandleStickSeries) (err error) {
	if numChart.base == nil || ncs == nil {
		return
	}
	if ncs.ser == nil {
		err = errors.New("series not initialized")
		return
	}
	err = numChart.base.AddCandleStickSeries(ncs.ser)
	return
}

// AddBoxSeries adds a series of data which is visualized as box chart.
// The series must have a unique name throughout the chart.
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (numChart *CartesianNumericalChart) AddBoxSeries(nbs *NumericalBoxSeries) (err error) {
	if numChart.base == nil || nbs == nil {
		return
	}
	if nbs.ser == nil {
		err = errors.New("series not initialized")
		return
	}
	err = numChart.base.AddBoxSeries(nbs.ser)
	return
}

// AddAreaSeries adds a series of data which is visualized as area chart.
// If showDots is true, dots are displayed at the osition of the series points.
// The series must have a unique name throughout the chart.
// An error is returned,if another series with the same name exists or if the series is already added to another chart
func (numChart *CartesianNumericalChart) AddAreaSeries(nps *NumericalPointSeries, showDots bool) (err error) {
	if numChart.base == nil || nps == nil {
		return
	}
	if nps.ser == nil {
		err = errors.New("series not initialized")
		return
	}
	err = numChart.base.AddAreaSeries(nps.ser, showDots)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series must have a unique name throughout the chart.
// The bars are centered around their N value of the data points. barWidth is the width of the bars.
// An error is returned,if another series with the same name exists, if the series is already added to another chart or if barWidth < 0
func (numChart *CartesianNumericalChart) AddBarSeries(nps *NumericalPointSeries, barWidth float64) (err error) {
	if numChart.base == nil || nps == nil {
		return
	}
	if nps.ser == nil {
		err = errors.New("series not initialized")
		return
	}
	err = nps.SetBarWidth(barWidth)
	if err != nil {
		return
	}
	err = numChart.base.AddBarSeries(nps.ser)
	return
}

// SetYAxisLabel sets the label of the y-axis, which will be displayed at the left side
func (numChart *CartesianNumericalChart) SetYAxisLabel(l string) {
	if numChart.base == nil {
		return
	}
	numChart.base.SetToAxisLabel(l)
}

// SetYRange sets a user defined range for the y-axis;
// an error is returned if min>max or if the origin has been defined by the user before and is outside the given range
func (numChart *CartesianNumericalChart) SetYRange(min float64, max float64) (err error) {
	if numChart.base == nil {
		return
	}
	err = numChart.base.SetToRange(min, max)
	return
}

// SetAutoYRange overrides a previously user defined range and lets the range be calculated automatically
func (numChart *CartesianNumericalChart) SetAutoYRange() {
	if numChart.base == nil {
		return
	}
	numChart.base.SetAutoToRange()
}

// SetYTicks sets the list of user defined ticks to be shown on the y-axis
func (numChart *CartesianNumericalChart) SetYTicks(ts []data.NumericalTick) {
	if numChart.base == nil {
		return
	}
	numChart.base.SetToTicks(ts)
}

// SetAutoYTicks overrides a previously user defined set of y-axis ticks and lets the ticks be calculated automatically
func (numChart *CartesianNumericalChart) SetAutoYTicks(autoSupportLine bool) {
	if numChart.base == nil {
		return
	}
	numChart.base.SetAutoToTicks(autoSupportLine)
}

// SetYAxisStyle changes the style of the Y-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (numChart *CartesianNumericalChart) SetYAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	if numChart.base == nil {
		return
	}
	numChart.base.SetToAxisLabelStyle(labelSize, labelColor)
	numChart.base.SetToAxisStyle(axisColor)
}

// SetOrigin sets a user defined origin (crossing of x and y axis).
// An error is returned, if a range has been defined before and at least one coordinate is outside the range.
func (numChart *CartesianNumericalChart) SetOrigin(x float64, y float64) (err error) {
	if numChart.base == nil {
		return
	}
	err = numChart.base.SetNOrigin(x, y)
	return
}

// SetAutoOrigin resets a previously user defined origin and allows the chart to calculate the ideal origin automatically
func (numChart *CartesianNumericalChart) SetAutoOrigin() {
	if numChart.base == nil {
		return
	}
	numChart.base.SetAutoOrigin()
}

// SetXAxisLabel sets the label of the x-axis, which will be displayed at the bottom
func (numChart *CartesianNumericalChart) SetXAxisLabel(l string) {
	if numChart.base == nil {
		return
	}
	numChart.base.SetFromAxisLabel(l)
}

// SetXRange sets a user defined range for the x-axis.
// An error is returned, if min>max or if the origin has been defined by the user before and is outside the given range
func (numChart *CartesianNumericalChart) SetXRange(min float64, max float64) (err error) {
	if numChart.base == nil {
		return
	}
	err = numChart.base.SetFromNRange(min, max)
	return
}

// SetAutoXRange overrides a previously user defined range and lets the range be calculated automatically
func (numChart *CartesianNumericalChart) SetAutoXRange() {
	if numChart.base == nil {
		return
	}
	numChart.base.SetAutoFromRange()
}

// SetXTicks sets the list of user defined ticks to be shown on the x-axis
func (numChart *CartesianNumericalChart) SetXTicks(ts []data.NumericalTick) {
	if numChart.base == nil {
		return
	}
	numChart.base.SetFromNTicks(ts)
}

// SetAutoXTicks overrides a previously user defined set of x-axis ticks and lets the ticks be calculated automatically
func (numChart *CartesianNumericalChart) SetAutoXTicks(autoSupportLine bool) {
	if numChart.base == nil {
		return
	}
	numChart.base.SetAutoFromTicks(autoSupportLine)
}

// SetXAxisStyle changes the style of the X-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (numChart *CartesianNumericalChart) SetXAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	if numChart.base == nil {
		return
	}
	numChart.base.SetFromAxisLabelStyle(labelSize, labelColor)
	numChart.base.SetFromAxisStyle(axisColor)
}
