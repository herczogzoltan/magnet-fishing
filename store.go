package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Store struct{}

func (s *Store) Listen(p *Player) {

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		if p.Strength <= 5 {
			p.Gold -= 10
			p.Strength += 1
		}
	}
}
