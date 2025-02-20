package main

import (
	"log"
	"pacMan/internal"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := internal.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("PACMAN")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
