package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	OceanFloor = 800
)

var (
	catchList []Catch
)

type Catch struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Gold        Gold   `json:"Gold"`
}

type Game struct {
	Width       int
	Height      int
	Store       *Store
	Player      *Player
	Magnet      *Magnet
	Found       bool
	Catch       Catch
	GameStarted bool
	Assets      *embed.FS
}

func NewGame(windowWidth int, windowHeight int, assets *embed.FS) *Game {
	player := NewPlayer()

	loadCatchAsset(assets)

	game := &Game{
		Width:  windowWidth,
		Height: windowHeight,
		Store:  &Store{},
		Player: player,
		Magnet: NewMagnet(),
		Found:  false,
		Catch:  Catch{},
		Assets: assets,
	}

	return game
}

func loadCatchAsset(assets *embed.FS) {
	catchFile, err := assets.Open("assets/catch/catch.json")
	if err != nil {
		panic(err)
	}

	defer catchFile.Close()
	catches, _ := io.ReadAll(catchFile)

	if err := json.Unmarshal(catches, &catchList); err != nil {
		panic(err)
	}
}

func (g *Game) Update() error {
	g.Player.Update()
	g.Magnet.Update(g.Player)
	g.Store.Listen(g.Player)

	if isRopeSpinning() {
		g.GameStarted = true
	}

	if g.Magnet.Found() {
		g.Found = true
		g.Catch = catchList[rand.Intn(len(catchList))]

		g.reset()
	}

	return nil
}

var textFloat = 36

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	g.Player.Draw(screen)
	if g.Player.IsThrowReleased() {
		g.Magnet.Draw(screen)
	}

	text.Draw(screen, "Gold:"+strconv.Itoa(int(g.Player.Gold) + int(g.Catch.Gold)), mplusNormalFont, g.Width-270, 30, color.Black)
	text.Draw(screen, "Throwing Accuracy:"+strconv.Itoa(int(g.Player.ThrowAccuracy)), mplusNormalFont, g.Width-270, 60, color.Black)

	if !g.GameStarted {
		g.displayTutorial(screen)
	}

	if g.Found {
		g.displayCatchMessage(screen)
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

func (g *Game) displayTutorial(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintln("1. Click and hold your epic left mouse button to catch"), mplusNormalFont, g.Width/4, g.Height/3-90, color.Black)
	text.Draw(screen, fmt.Sprintln("2. Check the epic stuff you found"), mplusNormalFont, g.Width/4, g.Height/3-60, color.Black)
	text.Draw(screen, fmt.Sprintln("3. Spend your epic gold on epic stuff"), mplusNormalFont, g.Width/4, g.Height/3-30, color.Black)
	text.Draw(screen, fmt.Sprintln("4. Click and hold epicly to catch again"), mplusNormalFont, g.Width/4, g.Height/3, color.Black)
}

func (g *Game) displayCatchMessage(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("+%d Gold!", g.Catch.Gold), mplusBigFont, g.Width/2, g.Height/3-(66-textFloat), color.Black)
	text.Draw(screen, fmt.Sprintf("I found %s !", g.Catch.Name), mplusNormalFont, g.Width/2, g.Height/3-(36-textFloat), color.Black)

	text.Draw(screen, g.Catch.Description, mplusSmallFont, g.Width/2, g.Height/3-(6-textFloat), color.Black)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (windowWidth, windowHeight int) {
	return g.Width, g.Height
}

func isRopeSpinning() bool {
	return inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) != 0
}

func (g *Game) reset() {
	g.Player.reset()
	g.Magnet.reset()
	textFloat = 36
}
