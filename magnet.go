package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Magnet struct {
	Image       *ebiten.Image
	Options     *ebiten.DrawImageOptions
	Thrown      bool
	ThrownSince int
	Found       bool
	y16, x16    int
	vy16        int
}

func (m *Magnet) Draw(screen *ebiten.Image) {
	if m.Thrown {
		magnetImage, _, err := ebitenutil.NewImageFromFile("assets/magnet.png")
		if err != nil {
			log.Fatal(err)
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(windowWidth)/2, float64(windowHeight)/2)
		op.GeoM.Translate(float64(windowWidth)/5, 0)
		w, h := magnetImage.Size()
		op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)

		cameraX := 10
		cameraY := 10

		op.GeoM.Translate(float64(m.x16/50.0)-float64(cameraX), float64(m.y16/50.0)-float64(cameraY))

		screen.DrawImage(magnetImage, op)
		return
	}

}

func (m *Magnet) Update(g *Game) {
	if g.Player.Thrown {
		m.Thrown = true
	}

	if m.Thrown {
		m.ThrownSince++
		if m.ThrownSince == ThrowReleaseCycle {
			m.Found = true
		}

		if m.ThrownSince <= ThrowReleaseCycle/2 {
			m.y16 -= m.vy16
			m.x16 -= m.vy16
		} else {
			m.y16 += m.vy16
			m.x16 -= m.vy16
		}

		// Gravity
		m.vy16 += 3
		if m.vy16 > 60 {
			m.vy16 = 60
		}
	}

	if m.Found {
		m.Thrown = false
	}

}

func (m *Magnet) reset() {
	m.Found = false
	m.ThrownSince = 0
	m.y16 = 0
	m.x16 = 0
	m.vy16 = 0
}
