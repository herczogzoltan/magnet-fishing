package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	windowWidth  = 1280
	windowHeight = 960
)

func setupEnvironment() {
	ebiten.SetWindowTitle("Magnet Fishing")
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
}

func main() {
	setupEnvironment()

	game := &Game{
		Width:    windowWidth,
		Height:   windowHeight,
		Player:   nil,
		Duration: Duration(0),
	}

	go NewGame(game)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Duration int

type Throw struct {
	Accuracy uint8
	Power    uint8
}

type Game struct {
	Width    int
	Height   int
	Player   *Player
	Duration Duration
	Throw    *Throw
}

func NewGame(game *Game) {
	game.Player = NewPlayer()
}

func (g *Game) Update() error {
	g.Player.Update(g)
	g.Duration = Duration(inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	g.Player.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("\n\nMouseClick duration: %d", inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft)), 0, 0)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (windowWidth, windowHeight int) {
	return g.Width, g.Height
}
