package main

import (
	"log"
	"pacMan/internal"
	"pacMan/internal/global"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := internal.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(global.SCREEN_WIDTH, global.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("PACMAN")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
