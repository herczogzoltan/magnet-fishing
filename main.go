package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	windowWidth  = 1280
	windowHeight = 960
)

var (
	mplusNormalFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func setupEnvironment() {
	ebiten.SetWindowTitle("Magnet Fishing")
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
}

func main() {
	setupEnvironment()

	game := &Game{
		Width:         windowWidth,
		Height:        windowHeight,
		Player:        nil,
		Throw:         nil,
		ClickDuration: ClickDuration(0),
	}

	go NewGame(game)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type ClickDuration int

type ThrowPower uint8

type ThrowAccuracy uint8

type Throw struct {
	Accuracy ThrowAccuracy
	Power    ThrowPower
}

func (t *Throw) setAccuracy(cd ClickDuration) {
	durationString := strconv.Itoa(int(cd))
	accuracy, err := t.calculateAccuracy(durationString)

	if err != nil {
		log.Fatal(err)
	}

	t.Accuracy = ThrowAccuracy(accuracy)
}

func (t *Throw) calculateAccuracy(duration string) (int, error) {
	dsLen := len(duration)

	if dsLen >= 2 {
		return strconv.Atoi(duration[dsLen-2:])
	}

	return strconv.Atoi(duration)
}

func (t *Throw) setPower(s playerStrength) {
	t.Power = ThrowPower(s)
}

func (t *Throw) Update(g *Game) {
	t.setPower(g.Player.Strength)
	t.setAccuracy(g.ClickDuration)
}

type Game struct {
	Width         int
	Height        int
	Player        *Player
	ClickDuration ClickDuration
	Throw         *Throw
}

func NewGame(game *Game) {
	game.Player = NewPlayer()
	game.Throw = &Throw{}
}

func (g *Game) Update() error {
	g.Player.Update(g)
	g.Throw.Update(g)

	// Do not reset click duration when we have a value
	if g.ClickDuration != 0 && inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) == 0 {
		return nil
	}

	g.ClickDuration = ClickDuration(inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	g.Player.Draw(screen)

	text.Draw(screen, "Throwing Accuracy:"+strconv.Itoa(int(g.Throw.Accuracy)), mplusNormalFont, g.Width-270, 30, color.Black)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (windowWidth, windowHeight int) {
	return g.Width, g.Height
}
