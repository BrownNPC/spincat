package main

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

var configFilePath = filepath.Join(os.Args[0], "..", "spincat-config.json")

var Cfg = LoadConfig(configFilePath)

// Reload config when file is updated
func init() {
	go func() {
		for {
			WatchFile(configFilePath)
			Cfg = LoadConfig(configFilePath)
			ebiten.SetWindowSize(Cfg.Size, Cfg.Size)
			game.cat.SpinSpeed = Cfg.SpinSpeed
		}
	}()
}

const RenderWidth, RenderHeight = 320, 320

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
	g.cat.Update()
	g.cat.SetIdle(g.cat.Distance < 500)
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

var game = Game{cat: NewCat(Cfg.SpinSpeed)}

func main() {
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetScreenClearedEveryFrame(true)
	ebiten.SetTPS(60)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowMousePassthrough(true)
	ebiten.SetWindowSize(Cfg.Size, Cfg.Size)
	ebiten.SetWindowTitle("Spin Cat")

	if err := ebiten.RunGameWithOptions(&game, &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}); err != nil {
		log.Fatal(err)
	}
}
