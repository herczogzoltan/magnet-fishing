package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	OceanFloor = 800
)

type Game struct {
	Width       int
	Height      int
	Player      *Player
	Throw       *Throw
	Magnet      *Magnet
	Found       bool
	Catch       Catch
	GameStarted bool
}

func NewGame(game *Game) {
	textFloat = 36
	game.Player = NewPlayer()
	game.Throw = &Throw{}
	game.Magnet = NewMagnet()
	game.Found = false
	game.Catch = Catch{}
	game.GameStarted = false
}

func (g *Game) Update() error {
	g.Player.Update(g)
	g.Throw.Update(g)
	g.Magnet.Update(g)

	if isPrepareThrow() {
		g.GameStarted = true
	}

	if g.Magnet.Found {
		g.Found = true
		g.Catch = catchList.Catches[rand.Intn(len(catchList.Catches))]

		g.reset()
	}

	return nil
}

var textFloat int

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	g.Player.Draw(screen)
	g.Magnet.Draw(screen)

	text.Draw(screen, "Gold:"+strconv.Itoa(int(g.Player.Gold)), mplusNormalFont, g.Width-270, 30, color.Black)
	text.Draw(screen, "Throwing Accuracy:"+strconv.Itoa(int(g.Throw.Accuracy)), mplusNormalFont, g.Width-270, 60, color.Black)

	if !g.GameStarted {
		text.Draw(screen, fmt.Sprintln("1. Click and hold your epic left mouse button to catch"), mplusNormalFont, g.Width/4, g.Height/3-90, color.Black)
		text.Draw(screen, fmt.Sprintln("2. Check the epic stuff you found"), mplusNormalFont, g.Width/4, g.Height/3-60, color.Black)
		text.Draw(screen, fmt.Sprintln("3. Spend your epic gold on epic stuff"), mplusNormalFont, g.Width/4, g.Height/3-30, color.Black)
		text.Draw(screen, fmt.Sprintln("4. Click and hold epicly to catch again"), mplusNormalFont, g.Width/4, g.Height/3, color.Black)
	}

	if g.Found {
		text.Draw(screen, fmt.Sprintf("+%d Gold!", g.Catch.Gold), mplusBigFont, g.Width/2, g.Height/3-(66-textFloat), color.Black)
		text.Draw(screen, fmt.Sprintf("I found %s !", g.Catch.Name), mplusNormalFont, g.Width/2, g.Height/3-(36-textFloat), color.Black)

		text.Draw(screen, g.Catch.Description, mplusSmallFont, g.Width/2, g.Height/3-(6-textFloat), color.Black)

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.Player.Gold += g.Catch.Gold
			g.Found = false
			g.Catch = Catch{}
		}
		if textFloat != 0 {
			textFloat -= 1
		}
	}
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
	textFloat = 36
}
