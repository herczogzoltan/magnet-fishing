package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	playerDrawAnimationSpeed   = 14
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
	StandingImage     *ebiten.Image
	ThrowingImage     *ebiten.Image
	RopeSpinningImage *ebiten.Image
	Options           *ebiten.DrawImageOptions
	Gold              Gold
	Strength          int
	count             int
	Throwing          bool
	ThrowAccuracy     int
}

func NewPlayer() *Player {
	p := &Player{
		StandingImage: &ebiten.Image{},
		Strength:      1,
		Gold:          0,
		Throwing:      false,
	}

	p.StandingImage = LoadImage("assets/player-stand.png")
	p.ThrowingImage = LoadImage("assets/player-throw-release.png")
	p.RopeSpinningImage = LoadImage("assets/player-prepare-throw.png")

	p.Options = &ebiten.DrawImageOptions{}
	p.Options.GeoM.Translate(float64(windowWidth)/2, float64(windowHeight)/2)
	p.Options.GeoM.Translate(float64(windowWidth/5), 0)

	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	if p.isThrowing() {
		p.drawThrowing(screen)
		return
	}

	p.drawStanding(screen)
}

func (p *Player) drawThrowing(screen *ebiten.Image) {
	if p.Throwing {
		sx := (p.count / playerDrawAnimationSpeed) % playerThrownFrameNum * playerThrownFrameWidth

		screen.DrawImage(p.ThrowingImage.SubImage(image.Rect(sx, 0, sx+playerThrownFrameWidth, playerThrownAssetHeight)).(*ebiten.Image), p.Options)
		return
	}

	if isRopeSpinning() {
		sx := p.getAnimationSpeed() * playerPreparingFrameWidth
		screen.DrawImage(p.RopeSpinningImage.SubImage(image.Rect(sx, 0, sx+playerPreparingFrameWidth, playerPreparingAssetHeight)).(*ebiten.Image), p.Options)
		return
	}
}

func (p *Player) drawStanding(screen *ebiten.Image) {
	sx := p.getAnimationSpeed() * playerStandFrameWidth
	screen.DrawImage(p.StandingImage.SubImage(image.Rect(sx, 0, sx+playerStandFrameWidth, playerStandAssetHeight)).(*ebiten.Image), p.Options)
}

func (p *Player) isThrowing() bool {
	return p.Throwing || isRopeSpinning()
}

func (p *Player) IsThrowReleased() bool {
	return p.Throwing && p.count/playerDrawAnimationSpeed%playerThrownFrameNum*playerThrownFrameWidth == 560
}

func (p *Player) getAnimationSpeed() int {
	return (p.count / playerDrawAnimationSpeed) % playerFrameNum
}

func (p *Player) Update() {
	if p.isThrown() {
		p.Throwing = true
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		p.count = 0
	}

	if !p.IsThrowReleased() {
		p.count++
	}

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

func (p *Player) isThrown() bool {
	return !isRopeSpinning() && p.ThrowAccuracy != 0
}

func (p *Player) reset() {
	p.count = 0
	p.Throwing = false
	p.ThrowAccuracy = 0
}
