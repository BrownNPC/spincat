package main

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
	Speed        = 2.0
)

//go:embed cat-spin
var CatSpinGIF embed.FS

type Cat struct {
	X, Y float64
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

var (
	CatGif []*ebiten.Image
)

type Game struct {
	count     int
	FrameStep float64
}

func (g *Game) Update() error {
	g.FrameStep++

	mx, my := ebiten.CursorPosition()

	// 24 fps
	if g.FrameStep > 60/24 {
		g.FrameStep = 0
		g.count = (g.count + 1) % len(CatGif)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()

	scaleX := float64(w) / screenWidth
	scaleY := float64(h) / screenHeight

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleX, scaleY)

	screen.DrawImage(CatGif[g.count], op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	var game = Game{}
	{
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
	}

	// ebiten.SetWindowDecorated(false)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetScreenClearedEveryFrame(true)
	ebiten.SetTPS(60)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowMousePassthrough(true)
	ebiten.SetWindowSize(72, 72)
	ebiten.SetWindowTitle("Neko")

	if err := ebiten.RunGameWithOptions(&game, &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}); err != nil {
		log.Fatal(err)
	}
}

var (
	prevX, prevY int
	firstFrame   = true
)

func CursorDelta() (dx, dy int) {
	x, y := ebiten.CursorPosition()

	if firstFrame {
		prevX, prevY = x, y
		firstFrame = false
		return 0, 0
	}

	dx = x - prevX
	dy = y - prevY
	prevX, prevY = x, y
	return
}
