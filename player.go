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
	playerFrameWidth          = (playerStandAssetWidth / playerFrameNum)
)

type Player struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions
	count   int
}

func NewPlayer() *Player {
	playerImage, _, err := ebitenutil.NewImageFromFile("assets/player-stand.png")
	if err != nil {
		log.Fatal(err)
	}

	return &Player{
		Image: playerImage,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Initialize position
	p.Options = &ebiten.DrawImageOptions{}
	p.Options.GeoM.Translate(float64(windowWidth)/2, float64(windowHeight)/2)
	p.Options.GeoM.Translate(float64(windowWidth/5), 0)

	// Animate standing
	sx, sy := 0+p.getAnimationSpeed()*playerFrameWidth, 0
	screen.DrawImage(p.Image.SubImage(image.Rect(sx, sy, sx+playerFrameWidth, sy+playerStandAssetHeight)).(*ebiten.Image), p.Options)
}

func (p *Player) getAnimationSpeed() int {
	return (p.count / playerStandAnimationSpeed) % playerFrameNum
}

func (p *Player) Update(g *Game) {
	p.count++
}
