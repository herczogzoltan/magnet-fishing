package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	ThrowReleaseCycle = 100
)

const (
	playerStandAnimationSpeed  = 14
	playerFrameNum             = 4
	playerStandAssetWidth      = 448
	playerStandAssetHeight     = 118
	playerStandFrameWidth      = (playerStandAssetWidth / playerFrameNum)
	playerPreparingAssetWidth  = 517
	playerPreparingAssetHeight = 112
	playerPreparingFrameWidth  = playerPreparingAssetWidth / playerFrameNum
	playerThrownAssetWidth     = 700
	playerThrownAssetHeight    = 144
	playerThrownFrameNum       = 5
	playerThrownFrameWidth     = playerThrownAssetWidth / playerThrownFrameNum
)

type playerStrength uint8

type Player struct {
	Image       *ebiten.Image
	Options     *ebiten.DrawImageOptions
	Strength    playerStrength
	count       int
	Throwing    bool
	Thrown      bool
	ThrownSince int
}

func NewPlayer() *Player {
	playerImage, _, err := ebitenutil.NewImageFromFile("assets/player-stand.png")
	if err != nil {
		log.Fatal(err)
	}

	return &Player{
		Image:    playerImage,
		Strength: playerStrength(1),
		Throwing: false,
		Thrown:   false,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Initialize position
	p.Options = &ebiten.DrawImageOptions{}
	p.Options.GeoM.Translate(float64(windowWidth)/2, float64(windowHeight)/2)
	p.Options.GeoM.Translate(float64(windowWidth/5), 0)

	if p.Thrown {

		throwingImage, _, err := ebitenutil.NewImageFromFile("assets/player-throw-release.png")
		if err != nil {
			log.Fatal(err)
		}

		sx, sy := 560, 0

		screen.DrawImage(throwingImage.SubImage(image.Rect(sx, sy, sx+playerThrownFrameWidth, sy+playerThrownAssetHeight)).(*ebiten.Image), p.Options)
		p.ThrownSince++
		return
	}
	if p.Throwing {
		if p.count/playerStandAnimationSpeed%5*playerThrownFrameWidth == 560 {
			p.Thrown = true
			return
		}
		thrownImage, _, err := ebitenutil.NewImageFromFile("assets/player-throw-release.png")
		if err != nil {
			log.Fatal(err)
		}

		sx, sy := 0+(p.count/playerStandAnimationSpeed)%5*playerThrownFrameWidth, 0

		screen.DrawImage(thrownImage.SubImage(image.Rect(sx, sy, sx+playerThrownFrameWidth, sy+playerThrownAssetHeight)).(*ebiten.Image), p.Options)
		return
	}
	// Change Animation to throwing
	if isPrepareThrow() {
		preparingImage, _, err := ebitenutil.NewImageFromFile("assets/player-prepare-throw.png")
		if err != nil {
			log.Fatal(err)
		}

		sx, sy := 0+p.getAnimationSpeed()*playerPreparingFrameWidth, 0
		screen.DrawImage(preparingImage.SubImage(image.Rect(sx, sy, sx+playerPreparingFrameWidth, sy+playerPreparingAssetHeight)).(*ebiten.Image), p.Options)
		return
	}

	sx, sy := 0+p.getAnimationSpeed()*playerStandFrameWidth, 0
	screen.DrawImage(p.Image.SubImage(image.Rect(sx, sy, sx+playerStandFrameWidth, sy+playerStandAssetHeight)).(*ebiten.Image), p.Options)
}

func (p *Player) getAnimationSpeed() int {
	return (p.count / playerStandAnimationSpeed) % playerFrameNum
}

func (p *Player) Update(g *Game) {
	if g.isThrown() {
		p.Throwing = true
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		p.count = 0
	}

	p.count++
}

func (p *Player) reset() {
	p.count = 0
	p.Throwing = false
	p.Thrown = false
	p.ThrownSince = 0
}
