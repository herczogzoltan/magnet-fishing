package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	playerStandAnimationSpeed = 14
	playerFrameNum            = 4
	playerStandAssetWidth     = 448
	playerStandAssetHeight    = 118
	playerStandFrameWidth     = (playerStandAssetWidth / playerFrameNum)
	playerThrowAssetWidth     = 517
	playerThrowAssetHeight    = 112
	playerThrowFrameWidth     = playerThrowAssetWidth / playerFrameNum
)

type playerStrength uint8

type Player struct {
	Image    *ebiten.Image
	Options  *ebiten.DrawImageOptions
	Strength playerStrength
	count    int
	Throwing bool
	Thrown   bool
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
		thrownAssetWidth := 700
		throwAssetHeight := 144
		throwFrameWidth := thrownAssetWidth / 5

		sx, sy := 560, 0

		screen.DrawImage(throwingImage.SubImage(image.Rect(sx, sy, sx+throwFrameWidth, sy+throwAssetHeight)).(*ebiten.Image), p.Options)
		return
	}
	if p.Throwing {
		thrownAssetWidth := 700
		throwAssetHeight := 144
		throwFrameWidth := thrownAssetWidth / 5
		if p.count/playerStandAnimationSpeed%5*throwFrameWidth == 560 {
			p.Thrown = true
			return
		}
		throwingImage, _, err := ebitenutil.NewImageFromFile("assets/player-throw-release.png")
		if err != nil {
			log.Fatal(err)
		}

		sx, sy := 0+(p.count/playerStandAnimationSpeed)%5*throwFrameWidth, 0

		screen.DrawImage(throwingImage.SubImage(image.Rect(sx, sy, sx+throwFrameWidth, sy+throwAssetHeight)).(*ebiten.Image), p.Options)
		return
	}
	// Change Animation to throwing
	if isThrowing() {
		throwingImage, _, err := ebitenutil.NewImageFromFile("assets/player-throw.png")
		if err != nil {
			log.Fatal(err)
		}

		sx, sy := 0+p.getAnimationSpeed()*playerThrowFrameWidth, 0
		screen.DrawImage(throwingImage.SubImage(image.Rect(sx, sy, sx+playerThrowFrameWidth, sy+playerThrowAssetHeight)).(*ebiten.Image), p.Options)
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
	p.count++
}
