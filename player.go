package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
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
	playerPreparingAssetWidth  = 539
	playerPreparingAssetHeight = 121
	playerPreparingFrameWidth  = playerPreparingAssetWidth / playerFrameNum
	playerThrownAssetWidth     = 700
	playerThrownAssetHeight    = 144
	playerThrownFrameNum       = 5
	playerThrownFrameWidth     = playerThrownAssetWidth / playerThrownFrameNum
)

type Gold uint64

type Player struct {
	Image         *ebiten.Image
	Options       *ebiten.DrawImageOptions
	Gold          Gold
	Strength      int
	count         int
	Throwing      bool
	Thrown        bool
	ThrownSince   int
	ThrowAccuracy int
}

func NewPlayer() *Player {
	return &Player{
		Image:    &ebiten.Image{},
		Strength: 1,
		Gold:     Gold(0),
		Throwing: false,
		Thrown:   false,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Initialize position
	p.Options = &ebiten.DrawImageOptions{}
	p.Options.GeoM.Translate(float64(windowWidth)/2, float64(windowHeight)/2)
	p.Options.GeoM.Translate(float64(windowWidth/5), 0)

	throwReleaseImage := LoadImage("assets/player-throw-release.png")

	if p.Thrown {
		sx, sy := 560, 0

		screen.DrawImage(throwReleaseImage.SubImage(image.Rect(sx, sy, sx+playerThrownFrameWidth, sy+playerThrownAssetHeight)).(*ebiten.Image), p.Options)
		p.ThrownSince++
		return
	}
	if p.Throwing {
		if p.count/playerStandAnimationSpeed%5*playerThrownFrameWidth == 560 {
			p.Thrown = true
			return
		}

		sx, sy := 0+(p.count/playerStandAnimationSpeed)%5*playerThrownFrameWidth, 0

		screen.DrawImage(throwReleaseImage.SubImage(image.Rect(sx, sy, sx+playerThrownFrameWidth, sy+playerThrownAssetHeight)).(*ebiten.Image), p.Options)
		return
	}
	// Change Animation to throwing
	if isPrepareThrow() {
		preparingImage := LoadImage("assets/player-prepare-throw.png")

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

	// Do not reset click duration when we have a value
	if p.ThrowAccuracy != 0 && inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) == 0 {
		return
	}

	p.setAccuracy(inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft))
}

func (p *Player) setAccuracy(clickDuration int) {
	cd := clickDuration % 100

	if cd <= 50 {
		p.ThrowAccuracy = cd
		return
	}

	p.ThrowAccuracy = 100 - cd
}

func (p *Player) ThrowDistance() int {
	return p.ThrowAccuracy + p.Strength
}

func (p *Player) reset() {
	p.count = 0
	p.Throwing = false
	p.Thrown = false
	p.ThrownSince = 0
	p.ThrowAccuracy = 0
}
