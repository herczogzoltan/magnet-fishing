package main

import (
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ThrowPower uint8
type ThrowAccuracy uint8
type ClickDuration int
type Distance int

type Throw struct {
	Accuracy ThrowAccuracy
	Power    ThrowPower
	Distance Distance
}

func (t *Throw) setAccuracy(cd ClickDuration) {
	durationString := strconv.Itoa(int(cd))
	accuracy, err := t.calculateAccuracy(durationString)

	if err != nil {
		log.Fatal(err)
	}

	t.Accuracy = ThrowAccuracy(accuracy)
}

func (t *Throw) getAccuracy() ThrowAccuracy {
	return t.Accuracy
}

func (t *Throw) calculateAccuracy(duration string) (int, error) {
	dsLen := len(duration)

	if dsLen >= 2 {
		return strconv.Atoi(duration[dsLen-2:])
	}

	return strconv.Atoi(duration)
}

func (t *Throw) setPower(s playerStrength) {
	t.Power = ThrowPower(s)
}

func (t *Throw) Update(g *Game) {
	t.setPower(g.Player.Strength)

	// Do not reset click duration when we have a value
	if t.Accuracy != 0 && inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) == 0 {
		return
	}

	t.setAccuracy(ClickDuration(inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft)))
	t.calculateDistance()
}

func (t *Throw) calculateDistance() {
	t.Distance = Distance(t.Accuracy+ThrowAccuracy(t.Power)) * 2
}

func (t *Throw) reset() {
	t.Accuracy = 0
	t.Power = 0
}
