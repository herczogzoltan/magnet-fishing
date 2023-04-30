package main

import (
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ClickDuration int

type Throw struct {
	Accuracy int
	Power    int
}

func (t *Throw) setAccuracy(cd ClickDuration) {
	durationString := strconv.Itoa(int(cd))
	accuracy, err := t.calculateAccuracy(durationString)

	if err != nil {
		log.Fatal(err)
	}

	if accuracy <= 50 {
		t.Accuracy = accuracy
		return
	}

	t.Accuracy = 100 - accuracy
}

func (t *Throw) calculateAccuracy(duration string) (int, error) {
	dsLen := len(duration)

	if dsLen >= 2 {
		return strconv.Atoi(duration[dsLen-2:])
	}

	return strconv.Atoi(duration)
}

func (t *Throw) Update(g *Game) {
	t.Power = g.Player.Strength

	// Do not reset click duration when we have a value
	if t.Accuracy != 0 && inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) == 0 {
		return
	}

	t.setAccuracy(ClickDuration(inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft)))
}

func (t *Throw) Distance() int {
	return t.Accuracy + t.Power
}

func (t *Throw) reset() {
	t.Accuracy = 0
	t.Power = 0
}
