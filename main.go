package main

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth  = 1280
	windowHeight = 960
)

//go:embed assets/catch/*.json assets/*.png
var assets embed.FS

func setupEnvironment() {
	ebiten.SetWindowTitle("Magnet Fishing")
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
}

func main() {
	setupEnvironment()

	game := NewGame(windowWidth, windowHeight, &assets)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
