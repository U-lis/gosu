package gosu

import (
	"embed"
	"fmt"
	"path/filepath"
	"runtime/debug"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hndada/gosu/format/osr"
)

type Game struct {
	Scene
}
type Scene interface {
	Update() any
	Draw(screen *ebiten.Image)
}

// func init() {
// 	if runtime.GOOS == "windows" {
// 		os.Setenv("EBITEN_GRAPHICS_LIBRARY", "opengl")
// 		fmt.Println("OpenGL mode has enabled.")
// 	}
// }

var (
	chartInfos  [][]ChartInfo
	modeProps   []ModeProp
	sceneSelect *SceneSelect
)

// Todo: load settings
func NewGame(props []ModeProp) *Game {
	modeProps = props
	g := &Game{}

	// 1. Load chart info and score data
	// 2. Check removed chart
	// 3. Check added chart
	// Each mode scans Music root independently.
	// LoadChartInfosSet(props)
	// TidyChartInfosSet(props)
	// for i, prop := range modeProps {
	// 	modeProps[i].ChartInfos = prop.LoadNewChartInfos(MusicRoot)
	// }
	// SaveChartInfosSet(props) // 4. Save chart infos to local file
	LoadGeneralSkin()
	for _, mode := range modeProps {
		mode.LoadSkin()
	}
	LoadHandlers(props)
	ebiten.SetWindowTitle("gosu")
	ebiten.SetWindowSize(WindowSizeX, WindowSizeY)
	ebiten.SetTPS(TPS)
	modeHandler.Max = len(props)
	sceneSelect = NewSceneSelect()
	// ebiten.SetCursorMode(ebiten.CursorModeHidden)
	return g
}

func NewGameWithEmbed(props []ModeProp, skin, music embed.FS) *Game {
	modeProps = props
	g := &Game{}
	dirs, err := music.ReadDir("")
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		fs, err := music.ReadDir(dir.Name())
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, f := range fs {
			if f.IsDir() {
				continue
			}
			ChartFileModeByFile(f)
			ChartFileMode(f)
			cpath := filepath.Join(dir.Name(), f.Name())
			info, err := prop.NewChartInfo(cpath) // First load should be done with no mods
			if err != nil {
				fmt.Printf("error at %s: %s\n", filepath.Base(cpath), err)
				continue
			}
			chartInfos = PutChartInfo(chartInfos, info)
		}
	}
	for i, prop := range modeProps {
		modeProps[i].ChartInfos = prop.LoadNewChartInfos(music)
	}
	LoadGeneralSkin()
	for _, mode := range modeProps {
		mode.LoadSkin()
	}
	LoadHandlers(props)
	ebiten.SetWindowTitle("gosu")
	ebiten.SetWindowSize(WindowSizeX, WindowSizeY)
	ebiten.SetTPS(TPS)
	modeHandler.Max = len(props)
	sceneSelect = NewSceneSelect()
	// ebiten.SetCursorMode(ebiten.CursorModeHidden)
	return g
}

func (g *Game) Update() (err error) {
	MusicVolumeKeyHandler.Update()
	EffectVolumeKeyHandler.Update()
	SpeedScaleKeyHandler.Update()
	if g.Scene == nil {
		g.Scene = sceneSelect
	}
	args := g.Scene.Update()
	switch args := args.(type) {
	case error:
		return args
	case PlayToResultArgs: // Todo: SceneResult
		EffectVolume = 0.25 // Todo: resolve delayed effect sound playing
		ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
		debug.SetGCPercent(100)
		g.Scene = sceneSelect
		ebiten.SetWindowTitle("gosu")
	case SelectToPlayArgs:
		EffectVolume = 0 // Todo: resolve delayed effect sound playing
		ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
		debug.SetGCPercent(0)
		prop := modeProps[currentMode]
		g.Scene, err = prop.NewScenePlay(args.Path, args.Replay)
		if err != nil {
			return
		}
	}
	return
}
func (g *Game) Draw(screen *ebiten.Image) {
	g.Scene.Draw(screen)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenSizeX, screenSizeY
}

type SelectToPlayArgs struct {
	// Mode int
	// Mods   Mods
	Path   string
	Replay *osr.Format
}

type PlayToResultArgs struct {
	Result
}
