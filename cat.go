package main

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewCat() Cat {
	var CatGif []*ebiten.Image
	entry, err := CatSpinGIF.ReadDir("cat-spin")
	if err != nil {
		panic(err)
	}
	entries := len(entry)
	for i := range entries {
		frame, err := CatSpinGIF.ReadFile(fmt.Sprintf("cat-spin/frame_%04d.png", i))
		if err != nil {
			panic(err)
		}
		img, _, err := image.Decode(bytes.NewReader(frame))
		CatGif = append(CatGif, ebiten.NewImageFromImage(img))
	}
	return Cat{Animation: CatGif}
}
func (c *Cat) Draw(screen *ebiten.Image) {
	screen.DrawImage(c.Animation[c.AnimationFrame], &ebiten.DrawImageOptions{})

}
func (c *Cat) Update(FPS int) {
	c.FrameStep++
	if !c.Idle {
		if c.FrameStep > 60/FPS {
			c.FrameStep = 0
			c.AnimationFrame = (c.AnimationFrame + 1) % len(c.Animation)
		}
	} else if c.Idle {
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
	if !c.Idle {
		c.FollowCursor(float64(x), float64(y))
	}
}

type Cat struct {
	X, Y           float64
	Idle           bool
	Distance       int
	AnimationFrame int
	FrameStep      int
	Animation      []*ebiten.Image
}

//go:embed cat-spin
var CatSpinGIF embed.FS

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
