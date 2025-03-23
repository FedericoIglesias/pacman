package character

import (
	"pacMan/internal/global"

	"github.com/hajimehoshi/ebiten/v2"
)

type Square struct {
	Dot  []*Dot
	Wall []*Wall
}

func NewSquare() *Square {

	WALL, DOT := STAGE()

	return &Square{
		Dot:  DOT,
		Wall: WALL,
	}
}

func (s *Square) Draw(screen *ebiten.Image) {
	for _, d := range s.Dot {
		d.Draw(screen)
	}
	for _, w := range s.Wall {
		w.Draw(screen)
	}
}

func STAGE() ([]*Wall, []*Dot) {
	var Wall = []*Wall{}
	var Dot = []*Dot{}
	for r, row := range global.STAGE_BYNARY {
		for c, colum := range row {
			if colum == 1 {
				Wall = append(Wall, NewWall(float64(global.SIDE*c), float64(global.SIDE*r), float64(global.SIDE), float64(global.SIDE)))
			}
			if colum == 0 {
				Dot = append(Dot, NewDot(float64(global.SIDE*c+15), float64(global.SIDE*r+15), 2, 2))
			}
		}
	}
	return Wall, Dot
}
