package internal

import (
	"image/color"
	"pacMan/internal/character"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Pacman *character.Pacman
	Blinky *character.Blinky
	// Menu   *Menu
	// Point  *Point
	// Dot    *Dot
	// Wall   *Wall
}

func NewGame() (*Game, error) {
	Pacman, err := character.NewPacman()
	if err != nil {
		panic(err)
	}
	Blinky, err := character.NewBlinky()
	if err != nil {
		panic(err)
	}

	g := &Game{
		Pacman: Pacman,
		Blinky: Blinky,
	}
	return g, nil
}

func (g *Game) Update() error {
	if err := g.Pacman.Update(); err != nil {
		return err
	}
	if err := g.Blinky.Update(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	g.Pacman.Draw(screen)
	g.Blinky.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 600
}
