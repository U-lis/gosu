package game

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Sprite struct {
	src  *ebiten.Image
	W, H int // desired w, h
	X, Y int

	fixed bool // a sprite that never moves once appears
	op    *ebiten.DrawImageOptions

	Saturation float64
	Dimness    float64
	// Theta      float64
	Color color.Color

	BornTime time.Time
	// LifeTime time.Time // zero value goes eternal
	ebiten.CompositeMode
}

func (s Sprite) IsOut(screenSize image.Point) bool {
	return (s.X+s.W < 0 || s.X > screenSize.X ||
		s.Y+s.H < 0 || s.Y > screenSize.Y)
}

func (s Sprite) Draw(screen *ebiten.Image) {
	if s.src == nil {
		log.Fatal("s.src is nil")
	}
	if s.IsOut(screen.Bounds().Max) {
		return
	}
	if s.fixed {
		screen.DrawImage(s.src, s.op)
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(s.ScaleW(), s.ScaleH())
		op.GeoM.Translate(float64(s.X), float64(s.Y))
		op.ColorM.ChangeHSV(0, s.Saturation, s.Dimness)
		if s.CompositeMode != 0 {
			op.CompositeMode = s.CompositeMode
		}
		screen.DrawImage(s.src, op)
	}
}

func (s *Sprite) SetImage(i image.Image) {
	switch i.(type) {
	case *ebiten.Image:
		s.src = i.(*ebiten.Image)
	default:
		i2, err := ebiten.NewImageFromImage(i, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		s.src = i2
	}
}

func (s Sprite) ScaleW() float64 {
	w1, _ := s.src.Size()
	return float64(s.W) / float64(w1)
}
func (s Sprite) ScaleH() float64 {
	_, h1 := s.src.Size()
	return float64(s.H) / float64(h1)
}

// Suppose minor parameter has already been set
func (s *Sprite) SetFixedOp(w, h, x, y int) {
	s.W = w
	s.H = h
	s.X = x
	s.Y = y
	op := &ebiten.DrawImageOptions{}
	// op.GeoM.Rotate(s.Theta)
	op.GeoM.Scale(s.ScaleW(), s.ScaleH())
	op.GeoM.Translate(float64(x), float64(y))
	if s.Color != nil {
		r, g, b, _ := s.Color.RGBA()
		op.ColorM.Scale(0, 0, 0, 1) // reset
		op.ColorM.Translate(
			float64(r)/0xff,
			float64(g)/0xff,
			float64(b)/0xff,
			0, // temp
		)
	}
	s.op = op
	s.fixed = true
}

// for debugging
func (s Sprite) PrintWHXY(comment string) {
	fmt.Println(comment, s.W, s.H, s.X, s.Y)
}

func NewSprite(src image.Image) Sprite {
	var sprite Sprite
	sprite.SetImage(src)

	sprite.BornTime = time.Now()
	sprite.Saturation = 1
	sprite.Dimness = 1
	return sprite
}

// each frame has same interval
type Animation struct {
	srcs   []*ebiten.Image
	Sprite // no use src
	// Duration time.Time // temp: using global duration

	Rep int64 // 반복 횟수
	ebiten.CompositeMode
}

const RepInfinite = -1

func NewAnimation(srcs []*ebiten.Image) Animation {
	var a Animation
	a.srcs = make([]*ebiten.Image, len(srcs))
	// temp: only []*ebiten.Image can be passed
	for i, src := range srcs {
		a.srcs[i] = src
	}
	a.BornTime = time.Now()
	a.Saturation = 1
	a.Dimness = 1
	return a
}

// 800ms 마다 1회 재생되는 8프레임짜리 애니메이션
// 0~99 100~199... 2 3 4 5 6 7
// 현재 2000ms, 4번째 frame이 보여지면 됨
// 2010ms에도 4번째 frame이 보여지면 됨
// todo: frameTime을 int로 구해버리고 그걸로 다시 나누면 실제 AnimatinoDuration보다 작아지는 효과
// duration이 100ms이고 6프레임이면
// 프레임당 16.6ms
const AnimationDuration = 450 // temp: global duration in ms

func (a Animation) Draw(screen *ebiten.Image) {
	frameTime := AnimationDuration / float64(len(a.srcs))
	elapsedTime := time.Since(a.BornTime).Milliseconds()
	rep := elapsedTime / AnimationDuration
	if a.Rep != RepInfinite && rep >= a.Rep {
		return
	}
	t := elapsedTime % AnimationDuration
	i := int(float64(t) / frameTime)
	// temp: suppose all animation goes not fixed
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(a.ScaleW(), a.ScaleH())
	op.GeoM.Translate(float64(a.X), float64(a.Y))
	op.ColorM.ChangeHSV(0, a.Saturation, a.Dimness)
	if a.CompositeMode != 0 {
		op.CompositeMode = a.CompositeMode
	}
	screen.DrawImage(a.srcs[i], op)
}

// temp: suppose all frames have same size in an animation
func (a Animation) ScaleW() float64 {
	w1, _ := a.srcs[0].Size()
	return float64(a.W) / float64(w1)
}
func (a Animation) ScaleH() float64 {
	_, h1 := a.srcs[0].Size()
	return float64(a.H) / float64(h1)
}

type LongSprite struct {
	Sprite
	Vertical bool
}

// temp: no need to be method of LongSprite, to make sure only LongSprite uses this
func (s LongSprite) IsOut(w, h, x, y int, screenSize image.Point) bool {
	return x+w < 0 || x > screenSize.X || y+h < 0 || y > screenSize.Y
}

// 사이즈 제한 있어서 *ebiten.Image로 직접 그리면 X
func (s LongSprite) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w1, h1 := s.src.Size()
	switch s.Vertical {
	case true:
		op.GeoM.Scale(s.ScaleW(), 1) // height 쪽은 굳이 scale 하지 않는다
		// important: op is not AB = BA
		x, y := s.X, s.Y
		op.GeoM.Translate(float64(x), float64(y))
		q, r := s.H/h1, s.H%h1+1 // quotient, remainder // temp: +1

		first := s.src.Bounds()
		w, h := s.W, r
		first.Min = image.Pt(0, h1-r)
		if !s.IsOut(w, h, x, y, screen.Bounds().Size()) {
			screen.DrawImage(s.src.SubImage(first).(*ebiten.Image), op)
		}
		op.GeoM.Translate(0, float64(h))
		y += h
		h = h1
		for i := 0; i < q; i++ {
			if !s.IsOut(w, h, x, y, screen.Bounds().Size()) {
				screen.DrawImage(s.src, op)
			}
			op.GeoM.Translate(0, float64(h))
			y += h
		}

	default:
		op.GeoM.Scale(1, s.ScaleH())
		op.GeoM.Translate(float64(s.X), float64(s.Y))
		q, r := s.W/w1, s.W%w1+1 // temp: +1

		for i := 0; i < q; i++ {
			screen.DrawImage(s.src, op)
			op.GeoM.Translate(float64(w1), 0)
		}

		last := s.src.Bounds()
		last.Max = image.Pt(r, h1)
		screen.DrawImage(s.src.SubImage(last).(*ebiten.Image), op)
	}
}
