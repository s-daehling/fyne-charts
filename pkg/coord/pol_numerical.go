package coord

import (
	"github.com/s-daehling/fyne-charts/internal/coord"
	"github.com/s-daehling/fyne-charts/pkg/data"

	"fyne.io/fyne/v2"
)

// PolarNumericalChart implements a polar plane with one numerical phi-axis and one numerical r-axis
type PolarNumericalChart struct {
	coordChart
}

// NewPolarNumericalChart returns an initialized PolarNumericalChart
func NewPolarNumericalChart(title string) (numChart *PolarNumericalChart) {
	numChart = &PolarNumericalChart{
		coordChart: emptyCoordChart(coord.PolarPlane, coord.Numerical),
	}
	numChart.ExtendBaseWidget(numChart)
	numChart.SetTitle(title)
	return
}

// AddLineSeries adds a series of data which is visualized as line chart.
// If showDots is true, dots are displayed at the osition of the series points.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// Only points in the range of 0 <= N <= 2pi are displayed
// An error is returned,if another series with the same name exists, if the series is already added to another chart or if Val < 0 for one or more points
func (numChart *PolarNumericalChart) AddLineSeries(nps *NumericalPointSeries, showDots bool) (err error) {
	err = numChart.base.AddLineSeries(nps.ser, showDots)
	return
}

// AddScatterSeries adds a series of data which is visualized as scatter chart.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// Only points in the range of 0 <= N <= 2pi are displayed
// An error is returned,if another series with the same name exists, if the series is already added to another chart or if Val < 0 for one or more points
func (numChart *PolarNumericalChart) AddScatterSeries(nps *NumericalPointSeries) (err error) {
	err = numChart.base.AddScatterSeries(nps.ser)
	return
}

// AddLollipopSeries adds a series of data which is visualized as lollipop chart.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// Only points in the range of 0 <= N <= 2pi are displayed
// An error is returned,if another series with the same name exists, if the series is already added to another chart or if Val < 0 for one or more points
func (numChart *PolarNumericalChart) AddLollipopSeries(nps *NumericalPointSeries) (err error) {
	err = numChart.base.AddLollipopSeries(nps.ser)
	return
}

// AddAreaSeries adds a series of data which is visualized as area chart.
// If showDots is true, dots are displayed at the osition of the series points.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// Only points in the range of 0 <= N <= 2pi are displayed
// An error is returned,if another series with the same name exists, if the series is already added to another chart or if Val < 0 for one or more points
func (numChart *PolarNumericalChart) AddAreaSeries(nps *NumericalPointSeries, showDots bool) (err error) {
	err = numChart.base.AddAreaSeries(nps.ser, showDots)
	return
}

// AddBarSeries adds a series of data which is visualized as bar chart.
// The series must have a unique name throughout the chart.
// Only points with a Val equal or greater than zero can be added
// Only points in the range of 0 <= N <= 2pi are displayed
// The bars are centered around their N value of the data points. barWidth is the width of the bars.
// An error is returned,if another series with the same name exists, if the series is already added to another chart, if Val < 0 for one or more points or if barWidth < 0
func (numChart *PolarNumericalChart) AddBarSeries(nps *NumericalPointSeries, barWidth float64) (err error) {
	err = nps.SetBarWidth(barWidth)
	if err != nil {
		return
	}
	err = numChart.base.AddBarSeries(nps.ser)
	return
}

// SetRAxisLabel sets the label of the r-axis, which will be displayed at the bottom-right
func (numChart *PolarNumericalChart) SetRAxisLabel(l string) {
	numChart.base.SetToAxisLabel(l)
}

// SetRRange sets a user defined range for the r-axis;
// an error is returned if max<0 or if the origin has been defined by the user before and is outside the given range
func (numChart *PolarNumericalChart) SetRRange(max float64) (err error) {
	err = numChart.base.SetToRange(0.0, max)
	return
}

// SetAutoRRange overrides a previously user defined range and lets the range be calculated automatically
func (numChart *PolarNumericalChart) SetAutoRRange() {
	numChart.base.SetAutoToRange()
}

// SetRTicks sets the list of user defined ticks to be shown on the r-axis
func (numChart *PolarNumericalChart) SetRTicks(ts []data.NumericalTick) {
	numChart.base.SetToTicks(ts)
}

// SetAutoRTicks overrides a previously user defined set of r-axis ticks and lets the ticks be calculated automatically
func (numChart *PolarNumericalChart) SetAutoRTicks(autoSupportLine bool) {
	numChart.base.SetAutoToTicks(autoSupportLine)
}

// SetRAxisStyle changes the style of the R-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (numChart *PolarNumericalChart) SetRAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	numChart.base.SetToAxisLabelStyle(labelSize, labelColor)
	numChart.base.SetToAxisStyle(axisColor)
}

// SetOrigin sets a user defined origin (crossing of phi and r axis).
// An error is returned, if a range has been defined before and at least one coordinate is outside the range.
func (numChart *PolarNumericalChart) SetOrigin(phi float64, r float64) (err error) {
	err = numChart.base.SetNOrigin(phi, r)
	return
}

// SetAutoOrigin resets a previously user defined origin and allows the chart to calculate the ideal origin automatically
func (numChart *PolarNumericalChart) SetAutoOrigin() {
	numChart.base.SetAutoOrigin()
}

// SetPhiAxisLabel sets the label of the phi-axis, which will be displayed at the left side
func (numChart *PolarNumericalChart) SetPhiAxisLabel(l string) {
	numChart.base.SetFromAxisLabel(l)
}

// SetPhiTicks sets the list of user defined ticks to be shown on the phi-axis
func (numChart *PolarNumericalChart) SetPhiTicks(ts []data.NumericalTick) {
	numChart.base.SetFromNTicks(ts)
}

// SetAutoPhiTicks overrides a previously user defined set of phi-axis ticks and lets the ticks be calculated automatically
func (numChart *PolarNumericalChart) SetAutoPhiTicks(autoSupportLine bool) {
	numChart.base.SetAutoFromTicks(autoSupportLine)
}

// SetPhiAxisStyle changes the style of the Phi-axis
// default value label size: theme.SizeNameSubHeadingText
// default value label color: theme.ColorNameForeground
// default value axis color: theme.ColorNameForeground
func (numChart *PolarNumericalChart) SetPhiAxisStyle(labelSize fyne.ThemeSizeName,
	labelColor fyne.ThemeColorName, axisColor fyne.ThemeColorName) {
	numChart.base.SetFromAxisLabelStyle(labelSize, labelColor)
	numChart.base.SetFromAxisStyle(axisColor)
}
