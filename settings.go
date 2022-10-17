package gosu

var (
	MusicRoot   = "music"
	WindowSizeX = 1600
	WindowSizeY = 900
)
var (
	// TPS supposed to be multiple of 1000, since only one speed value
	// goes passed per Update, while unit of TransPoint's time is 1ms.
	// TPS affects only on Update(), not on Draw().
	TPS int = 1000 // TPS should be 1000 or greater.

	CursorScale       float64 = 0.1
	ChartItemWidth    float64 = 450
	ChartItemHeight   float64 = 50
	ChartItemShrink   float64 = 0.15
	chartItemshrink   float64 = ChartItemWidth * ChartItemShrink
	chartItemBoxCount int     = int(screenSizeY/ChartItemHeight) + 2 // Gives some margin.

	ScoreScale    float64 = 0.65
	ScoreDigitGap float64 = 0
	MeterWidth    float64 = 4 // The number of pixels per 1ms.
	MeterHeight   float64 = 50
)

// Todo: reset all tick-dependent variables.
// They are mostly at drawer.go or play.go, settings.go
// Keyword: TimeToTick
func SetTPS() {}
