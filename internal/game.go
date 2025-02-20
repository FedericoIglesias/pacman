package internal

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Pacman *Pacman
	// Menu   *Menu
	// Ghost  *Ghost
	// Point  *Point
	// Dot    *Dot
	// Wall   *Wall
}

func NewGame() (*Game, error) {
	Pacman, err := NewPacman()
	if err != nil {
		panic(err)
	}
	g := &Game{
		Pacman: Pacman,
	}
	return g, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	g.Pacman.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 600
}
