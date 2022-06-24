package main

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	Width  int
	Height int
	Player *Player
	Throw  *Throw
	Magnet *Magnet
}

func NewGame(game *Game) {
	game.Player = NewPlayer()
	game.Throw = &Throw{}
	game.Magnet = &Magnet{}
}

func (g *Game) Update() error {
	g.Player.Update(g)
	g.Throw.Update(g)
	g.Magnet.Update(g)

	if g.Player.ThrownSince == ThrowReleaseCycle {
		g.reset()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	g.Player.Draw(screen)
	g.Magnet.Draw(screen)

	text.Draw(screen, "Throwing Accuracy:"+strconv.Itoa(int(g.Throw.Accuracy)), mplusNormalFont, g.Width-270, 30, color.Black)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (windowWidth, windowHeight int) {
	return g.Width, g.Height
}

func (g *Game) isThrown() bool {
	return !isPrepareThrow() && g.Throw.Accuracy != 0
}

func (g *Game) reset() {
	g.Player.reset()
	g.Throw.reset()
	g.Magnet.reset()
}
