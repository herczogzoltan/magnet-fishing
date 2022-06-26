package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	magnetImage, _, err := ebitenutil.NewImageFromFile("assets/magnet.png")

	if err != nil {
		log.Fatal(err)
	}

	op := &ebiten.DrawImageOptions{}
	// Setup starting coordinates
	op.GeoM.Translate(float64(windowWidth)/2, float64(windowHeight)/2)
	op.GeoM.Translate(float64(windowWidth)/5, 0)

	return &Magnet{
		Image:         magnetImage,
		Options:       op,
		Thrown:        false,
		ThrownSince:   0,
		ThrowDistance: 0,
		Found:         false,
	}
}

func (m *Magnet) Draw(screen *ebiten.Image) {

	ty := m.calculateY() + 5
	if m.Thrown {
		if m.ThrownSince >= int(m.ThrowDistance*130) {
			m.Options.GeoM.Translate(-float64(2), +ty*float64((m.ThrownSince/15)))
		} else if m.ThrownSince >= int(m.ThrowDistance*100) && m.ThrownSince <= int(m.ThrowDistance*130) {
			m.Options.GeoM.Translate(-float64(2), 0)
		} else {
			m.Options.GeoM.Translate(-float64(2), -ty+float64(m.ThrownSince/10))
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
		if m.ThrownSince != 0 && m.Options.GeoM.Element(1, 2) >= OceanFloor {
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
