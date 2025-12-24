package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Width, Height             = 72, 72
	RenderWidth, RenderHeight = 320, 240
	Speed                     = 4.0
)

type Game struct {
	count     int
	FrameStep float64
	cat       Cat
}

func (g *Game) Update() error {
	g.FrameStep++
	ebiten.SetWindowPosition(
		int(math.Round(g.cat.X)),
		int(math.Round(g.cat.Y)),
	)
	g.cat.Update(24)
	g.cat.Idle = g.cat.Distance < 500
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()

	scaleX := float64(w) / RenderWidth
	scaleY := float64(h) / RenderHeight

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleX, scaleY)

	g.cat.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return RenderWidth, RenderHeight
}
func main() {
	var game = Game{cat: NewCat()}

	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetScreenClearedEveryFrame(true)
	ebiten.SetTPS(60)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowMousePassthrough(true)
	ebiten.SetWindowSize(Width, Height)
	ebiten.SetWindowTitle("Neko")

	if err := ebiten.RunGameWithOptions(&game, &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}); err != nil {
		log.Fatal(err)
	}
}
