package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 8
)

//go:embed cat.jpg
var cat []byte

var (
	CatImg *ebiten.Image
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(CatImg, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(cat))
	if err != nil {
		log.Fatal(err)
	}
	CatImg = ebiten.NewImageFromImage(img)

	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebitengine Demo)")
	if err := ebiten.RunGameWithOptions(&Game{}, &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}); err != nil {
		log.Fatal(err)
	}
}
