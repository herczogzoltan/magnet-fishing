package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Magnet struct {
	Image         *ebiten.Image
	Options       *ebiten.DrawImageOptions
	Thrown        bool
	ThrownSince   int
	ThrowDistance int
}

func NewMagnet() *Magnet {
	op := &ebiten.DrawImageOptions{}
	// Setup starting coordinates
	op.GeoM.Translate(float64(windowWidth)/2, float64(windowHeight)/2)
	op.GeoM.Translate(float64(windowWidth)/5, 0)

	return &Magnet{
		Image:         LoadImage("assets/magnet.png"),
		Options:       op,
		Thrown:        false,
		ThrownSince:   0,
		ThrowDistance: 0,
	}
}

func (m *Magnet) Draw(screen *ebiten.Image) {

	if m.Thrown {
		diff := float64(m.ThrowDistance / 5)
		if diff >= 4 {
			diff = 4
		}
		ty := m.calculateY() + diff/5
		tx := 2 + diff

		if m.ThrownSince >= m.ThrowDistance*130 {
			m.Options.GeoM.Translate(-float64(tx), +ty*float64((m.ThrownSince/10)))
		} else if m.ThrownSince >= m.ThrowDistance*100 && m.ThrownSince <= m.ThrowDistance*130 {
			m.Options.GeoM.Translate(-float64(tx), 0)
		} else {
			m.Options.GeoM.Translate(-float64(tx), -ty+float64(m.ThrownSince/4))
		}

		screen.DrawImage(m.Image, m.Options)
		return
	}

}

func (m *Magnet) Update(g *Game) {
	if g.Player.Thrown {
		m.Thrown = true
	}

	// do not override it while throwing
	if m.ThrowDistance != g.Player.ThrowDistance() {
		m.ThrowDistance = g.Player.ThrowDistance()
	}

	if m.Thrown {
		m.ThrownSince++
	}
}

func (m *Magnet) Found() bool {
	return m.Thrown && m.Options.GeoM.Element(1, 2) >= OceanFloor
}

func (m *Magnet) calculateY() float64 {
	x := m.Options.GeoM.Element(0, 2)
	diameter := m.ThrowDistance
	radius := float64(diameter / 2)
	h := x - radius

	return radius + math.Sqrt(math.Pow(radius, 2)-math.Pow(x-h, 2))
}

func (m *Magnet) reset() {
	m.ThrownSince = 0
	m.ThrowDistance = 0
	m.Thrown = false

	// Setup starting coordinates
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(windowWidth)/2, float64(windowHeight)/2)
	op.GeoM.Translate(float64(windowWidth)/5, 0)
	m.Options = op
}
