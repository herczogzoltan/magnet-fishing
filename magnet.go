package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Magnet struct {
	Image       *ebiten.Image
	Options     *ebiten.DrawImageOptions
	flyDuration int
	flyDistance int
}

func NewMagnet() *Magnet {
	op := &ebiten.DrawImageOptions{}
	// Setup starting coordinates
	op.GeoM.Translate(float64(windowWidth)/2, float64(windowHeight)/2)
	op.GeoM.Translate(float64(windowWidth)/5, 0)

	return &Magnet{
		Image:       LoadImage("assets/magnet.png"),
		Options:     op,
		flyDistance: 0,
		flyDuration: 0,
	}
}

func (m *Magnet) Draw(screen *ebiten.Image) {
	diff := float64(m.flyDistance / 5)
	if diff >= 4 {
		diff = 4
	}
	ty := m.calculateY() + diff/5
	tx := 2 + diff

	if m.flyDuration >= m.flyDistance*130 {
		m.Options.GeoM.Translate(-float64(tx), +ty*float64((m.flyDuration/10)))
	} else if m.flyDuration >= m.flyDistance*100 && m.flyDuration <= m.flyDistance*130 {
		m.Options.GeoM.Translate(-float64(tx), 0)
	} else {
		m.Options.GeoM.Translate(-float64(tx), -ty+float64(m.flyDuration/4))
	}

	screen.DrawImage(m.Image, m.Options)
}

func (m *Magnet) Update(g *Game) {
	// do not override it while throwing
	if m.flyDistance != g.Player.ThrowDistance() {
		m.flyDistance = g.Player.ThrowDistance()
	}

	if g.Player.Thrown {
		m.flyDuration++
	}
}

func (m *Magnet) Found() bool {
	return m.Options.GeoM.Element(1, 2) >= OceanFloor
}

func (m *Magnet) calculateY() float64 {
	x := m.Options.GeoM.Element(0, 2)
	diameter := m.flyDistance
	radius := float64(diameter / 2)
	h := x - radius

	return radius + math.Sqrt(math.Pow(radius, 2)-math.Pow(x-h, 2))
}

func (m *Magnet) reset() {
	m.flyDuration = 0
	m.flyDistance = 0

	// Setup starting coordinates
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(windowWidth)/2, float64(windowHeight)/2)
	op.GeoM.Translate(float64(windowWidth)/5, 0)
	m.Options = op
}
