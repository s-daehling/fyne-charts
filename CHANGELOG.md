# Change Log

## [0.2.0] 2026-04-06

### Added

- Add hover functionality for series elements in `coord` and `prop` charts
  - Add method `SetHoverBehavior` (for `*coordChart` in pkg/coord and `*propChart` in pkg/prop)
- Add display of value labels for point series
  - Add method `SetValueLabelStyle` (for `*pointSeries` in pkg/coord)
- Add struct `ValueLabelStyle` for styling of value labels and method `DefaultValueLabelStyle` (in pkg/style)

### Changed

- **Breaking:** Change display of value labels for proportional series
  - Exchange method `SetValueTextStyle` with `SetValueLabelStyle` (for `*Series` in pkg/prop)
- Update docs to new functionality (in docs)

### Removed

- **Breaking:** Remove method `DefaultValueTextStyle` (in pkg/style)
  - use `DefaultValueLLabelStyle` instead

### Fixed

- Fix border padding of chart widgets to Fyne default values

## [0.1.0] 2026-02-09

- first release
