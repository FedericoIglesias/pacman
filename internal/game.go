package internal

import (
	"image/color"
	"pacMan/internal/stage"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Stage *stage.Stage
}

func NewGame() (*Game, error) {

	Stage, err := stage.NewMap()
	if err != nil {
		panic(err)
	}

	g := &Game{
		Stage: Stage,
	}
	return g, nil
}

func (g *Game) Update() error {
	if err := g.Stage.Update(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0x00, 0x00, 0xff})
	g.Stage.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 600
}
