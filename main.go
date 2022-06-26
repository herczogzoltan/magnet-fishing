package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	windowWidth  = 1280
	windowHeight = 960
)

var (
	mplusBigFont    font.Face
	mplusNormalFont font.Face
	mplusSmallFont  font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    36,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	mplusSmallFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    20,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}

	loadCatchAsset()
}

func setupEnvironment() {
	ebiten.SetWindowTitle("Magnet Fishing")
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
}

func main() {
	setupEnvironment()

	game := &Game{
		Width:  windowWidth,
		Height: windowHeight,
		Player: nil,
		Throw:  nil,
	}

	go NewGame(game)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func isPrepareThrow() bool {
	return inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) != 0
}
