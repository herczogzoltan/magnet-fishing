package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadImage(name string) *ebiten.Image {
	file, err := assets.Open(name)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	image, _, err := ebitenutil.NewImageFromReader(file)

	if err != nil {
		log.Fatal(err)
	}

	return image
}
