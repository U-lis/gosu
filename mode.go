package gosu

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/hndada/gosu/format/osr"
	"github.com/hndada/gosu/format/osu"
)

const (
	ModeNone   = iota - 1
	ModePiano4 // 1, 2, 3, 4 Key
	ModePiano7 // 5, 6, 7+ Key
	ModeDrum
	ModeKaraoke
)

// ModeProp stands for Mode properties.
type ModeProp struct {
	Name string
	Mode int
	// ChartInfos     []ChartInfo
	// Cursor         int                 // Todo: custom chart infos - custom cursor
	// Results        map[[16]byte]Result // md5.Size = 16
	// LastUpdateTime time.Time
	LoadSkin func()
	// Skin interface{ Load() } // Todo: use this later
	// SpeedKeyHandler ctrl.KeyHandler
	SpeedScale   *float64
	NewChartInfo func(string) (ChartInfo, error)
	NewScenePlay func(cpath string, rf *osr.Format) (Scene, error)
	ExposureTime func(float64) float64
}

// Mode determines a mode of chart file by its path.
func ChartFileMode(fpath string) int {
	switch strings.ToLower(filepath.Ext(fpath)) {
	case ".osu":
		mode, keyCount := osu.Mode(fpath)
		switch mode {
		case osu.ModeMania:
			if keyCount <= 4 || keyCount == 6 {
				return ModePiano4
			}
			return ModePiano7
		case osu.ModeTaiko:
			return ModeDrum
		default:
			return ModeNone
		}
	case ".ojn", ".bms":
		return ModePiano7
	}
	return ModeNone
}

func ChartFileModeByFile(f fs.File) int {
	info, err := f.Stat()
	if err != nil {
		return -1
	}
	switch strings.ToLower(filepath.Ext(info.Name())) {
	case ".osu":
		mode, keyCount := osu.ModeByFile(f)
		switch mode {
		case osu.ModeMania:
			if keyCount <= 4 || keyCount == 6 {
				return ModePiano4
			}
			return ModePiano7
		case osu.ModeTaiko:
			return ModeDrum
		default:
			return ModeNone
		}
	case ".ojn", ".bms":
		return ModePiano7
	}
	return ModeNone
}
