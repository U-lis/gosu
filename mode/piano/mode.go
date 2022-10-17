package piano

import (
	"github.com/hndada/gosu"
)

var ModePiano4 = gosu.ModeProp{
	Name: "Piano4",
	Mode: gosu.ModePiano4,
	// ChartInfos:     make([]gosu.ChartInfo, 0),      // Zero value.
	// Results:        make(map[[16]byte]gosu.Result), // Zero value.
	// LastUpdateTime: time.Time{},                    // Zero value.
	LoadSkin:     LoadSkin,
	SpeedScale:   &SpeedScale,
	NewChartInfo: NewChartInfo,
	NewScenePlay: NewScenePlay,
	ExposureTime: ExposureTime,
}

var ModePiano7 = gosu.ModeProp{
	Name: "Piano7",
	Mode: gosu.ModePiano7,
	// ChartInfos:     make([]gosu.ChartInfo, 0),      // Zero value.
	// Results:        make(map[[16]byte]gosu.Result), // Zero value.
	// LastUpdateTime: time.Time{},                    // Zero value.
	LoadSkin:     LoadSkin,
	SpeedScale:   &SpeedScale,
	NewChartInfo: NewChartInfo,
	NewScenePlay: NewScenePlay,
	ExposureTime: ExposureTime,
}
