package internal

import (
	"image/color"
	"pacMan/internal/stage"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	// Pacman *character.Pacman
	// Blinky *character.Blinky
	// Pinky  *character.Pinky
	// Inky   *character.Inky
	// Clyde  *character.Clyde
	Stage *stage.Stage
	// Menu   *Menu
	// Point  *Point
	// Dot    *Dot
	// Wall   *Wall
}

func NewGame() (*Game, error) {
	// Blinky, err := character.NewBlinky()
	// if err != nil {
	// 	panic(err)
	// }

	// Pinky, err := character.NewPinky()
	// if err != nil {
	// 	panic(err)
	// }

	// Inky, err := character.NewInky()
	// if err != nil {
	// 	panic(err)
	// }

	// Clyde, err := character.NewClyde()
	// if err != nil {
	// 	panic(err)
	// }

	Stage, err := stage.NewMap()
	if err != nil {
		panic(err)
	}

	g := &Game{
		// Blinky: Blinky,
		// Pinky:  Pinky,
		// Inky:   Inky,
		// Clyde:  Clyde,
		Stage: Stage,
	}
	return g, nil
}

func (g *Game) Update() error {
	if err := g.Stage.Update(); err != nil {
		return err
	}
	// if err := g.Blinky.Update(); err != nil {
	// 	return err
	// }

	// if err := g.Pinky.Update(); err != nil {
	// 	return err
	// }

	// if err := g.Inky.Update(); err != nil {
	// 	return err
	// }

	// if err := g.Clyde.Update(); err != nil {
	// 	return err
	// }

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	g.Stage.Draw(screen)
	// g.Pacman.Draw(screen)
	// g.Blinky.Draw(screen)
	// g.Pinky.Draw(screen)
	// g.Inky.Draw(screen)
	// g.Clyde.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 600
}
