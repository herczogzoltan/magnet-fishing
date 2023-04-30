package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ClickDuration int

type Throw struct {
	Accuracy int
	Power    int
}

func (t *Throw) setAccuracy(clickDuration int) {
	cd := clickDuration % 100

	if cd <= 50 {
		t.Accuracy = cd
		return
	}

	t.Accuracy = 100 - cd
}

func (t *Throw) Update(g *Game) {
	t.Power = g.Player.Strength

	// Do not reset click duration when we have a value
	if t.Accuracy != 0 && inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) == 0 {
		return
	}

	t.setAccuracy(inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft))
}

func (t *Throw) Distance() int {
	return t.Accuracy + t.Power
}

func (t *Throw) reset() {
	t.Accuracy = 0
	t.Power = 0
}
