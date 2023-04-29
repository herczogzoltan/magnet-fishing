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
	ThrowDistance Distance
	Found         bool
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
		Found:         false,
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

		if m.ThrownSince >= int(m.ThrowDistance*130) {
			m.Options.GeoM.Translate(-float64(tx), +ty*float64((m.ThrownSince/10)))
		} else if m.ThrownSince >= int(m.ThrowDistance*100) && m.ThrownSince <= int(m.ThrowDistance*130) {
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
	if m.ThrowDistance == Distance(0) && g.Throw.Distance != Distance(0) {
		m.ThrowDistance = g.Throw.Distance
	}

	if m.Thrown {
		m.ThrownSince++

		// Found when hit ocean floor
		if m.Options.GeoM.Element(1, 2) >= OceanFloor {
			m.Found = true
		}
	}

	if m.Found {

		m.Thrown = false
	}
}

func (m *Magnet) calculateY() float64 {
	x := m.Options.GeoM.Element(0, 2)
	diameter := int(m.ThrowDistance)
	radius := float64(diameter / 2)
	h := x - radius

	return radius + math.Sqrt(math.Pow(radius, 2)-math.Pow(x-h, 2))
}

func (m *Magnet) reset() {
	m.Found = false
	m.ThrownSince = 0
	m.ThrowDistance = 0

	// Setup starting coordinates
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(windowWidth)/2, float64(windowHeight)/2)
	op.GeoM.Translate(float64(windowWidth)/5, 0)
	m.Options = op
}
