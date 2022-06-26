package main

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	OceanFloor = 800
)

type Game struct {
	Width                  int
	Height                 int
	Player                 *Player
	Throw                  *Throw
	Magnet                 *Magnet
	Found                  bool
	FoundTextTimerFinished bool
}

func NewGame(game *Game) {
	game.Player = NewPlayer()
	game.Throw = &Throw{}
	game.Magnet = NewMagnet()
	game.Found = false
	game.FoundTextTimerFinished = false
}

func (g *Game) Update() error {
	g.Player.Update(g)
	g.Throw.Update(g)
	g.Magnet.Update(g)

	if g.Magnet.Found {
		g.Found = true
		g.reset()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	g.Player.Draw(screen)
	g.Magnet.Draw(screen)

	text.Draw(screen, "Throwing Accuracy:"+strconv.Itoa(int(g.Throw.Accuracy)), mplusNormalFont, g.Width-270, 30, color.Black)

	if g.Found {
		text.Draw(screen, "You found something!", mplusNormalFont, g.Width/2, g.Height/3, color.Black)

		text.Draw(screen, "It's a nice pen! I wonder how did it end up there?", mplusSmallFont, g.Width/2, g.Height/3+30, color.Black)

		go func() { g.timer(5 * time.Second) }()
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || g.FoundTextTimerFinished {
			g.Found = false
		}

	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (windowWidth, windowHeight int) {
	return g.Width, g.Height
}

func (g *Game) isThrown() bool {
	return !isPrepareThrow() && g.Throw.Accuracy != 0
}

func (g *Game) timer(s time.Duration) {
	time.Sleep(s)
	g.FoundTextTimerFinished = true
}

func (g *Game) reset() {
	g.Player.reset()
	g.Throw.reset()
	g.Magnet.reset()
	g.FoundTextTimerFinished = false
}
