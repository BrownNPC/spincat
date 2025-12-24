package main

import (
	"bytes"
	_ "embed"
	"image/gif"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Cat struct {
	X, Y           float64
	idle           bool
	Distance       int
	AnimationFrame int
	// Allow for slower spin speed
	FrameAccumulator float64
	FrameStep        int
	GifFrames        []*ebiten.Image
	SpinSpeed        float64
	GifDelays        []int // ticks to hold each frame for
}

//go:embed spin.gif
var CatSpinGIF []byte

func NewCat(spinSpeed float64) Cat {
	g, err := gif.DecodeAll(bytes.NewReader(CatSpinGIF))
	if err != nil {
		panic("unable to decode spin.gif")
	}
	frames := make([]*ebiten.Image, len(g.Image))
	delays := make([]int, len(g.Delay))
	for i, img := range g.Image {
		frames[i] = ebiten.NewImageFromImage(img)
		// Gif uses 100 ticks per second, our app uses 60 tps
		delays[i] = max(g.Delay[i]*60/100, 1)
	}
	return Cat{
		GifFrames: frames,
		GifDelays: delays,
		// Start on frame 1
		AnimationFrame: 1,
		SpinSpeed:      spinSpeed,
	}
}
func (c *Cat) Draw(screen *ebiten.Image) {
	screen.DrawImage(c.GifFrames[c.AnimationFrame], &ebiten.DrawImageOptions{})
}
func (c *Cat) TickGif() {
	// Wait the encoded amount of frames
	if c.FrameStep >= c.GifDelays[c.AnimationFrame] {
		c.FrameStep = 0
		c.AnimationFrame = max((c.AnimationFrame+1)%len(c.GifFrames), 1)
	}
}
func (c *Cat) SetIdle(v bool) {
	if v == c.idle {
		return
	}
	c.idle = v
	c.FrameStep = 0
	c.FrameAccumulator = 0
	if v {
		c.AnimationFrame = 0
	} else {
		c.AnimationFrame = 1
	}

}
func (c *Cat) Update() {

	if !c.idle {
		// allow for slower or faster spin speeds
		c.FrameAccumulator += c.SpinSpeed

		for c.FrameAccumulator >= 1 {
			c.FrameAccumulator -= 1.0
			c.FrameStep++
			c.TickGif()
		}
		c.TickGif()
	} else if c.idle {
		c.AnimationFrame = 0
	}

	mx, my := ebiten.CursorPosition()
	x := mx - (RenderWidth / 2)
	y := my - (RenderHeight / 2)
	dx, dy := x, y

	if dy < 0 {
		dy = -dy
	}
	if dx < 0 {
		dx = -dx
	}
	c.Distance = dx + dy
	if !c.idle {
		c.FollowCursor(float64(x), float64(y))
	}
}

func (m *Cat) FollowCursor(x, y float64) {
	r := math.Atan2(y, x)
	a := math.Mod((r/math.Pi*180)+360, 360) // Normazing angle to [0, 360)
	switch {
	case a <= 292.5 && a > 247.5:
		//up
		m.Y -= Speed
	case a <= 337.5 && a > 292.5:
		// up right
		m.X += Speed / math.Sqrt2
		m.Y -= Speed / math.Sqrt2
	case a <= 22.5 || a > 337.5:
		// right
		m.X += Speed
	case a <= 67.5 && a > 22.5:
		// down right
		m.X += Speed / math.Sqrt2
		m.Y += Speed / math.Sqrt2
	case a <= 112.5 && a > 67.5:
		// down
		m.Y += Speed
	case a <= 157.5 && a > 112.5:
		// down left
		m.X -= Speed / math.Sqrt2
		m.Y += Speed / math.Sqrt2
	case a <= 202.5 && a > 157.5:
		// left
		m.X -= Speed
	case a <= 247.5 && a > 202.5:
		// up left
		m.X -= Speed
		m.Y -= Speed
	}
}
