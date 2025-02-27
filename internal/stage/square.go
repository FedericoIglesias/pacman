package stage

import "github.com/hajimehoshi/ebiten/v2"

type Square struct {
	IsWall bool
	Rect   *ebiten.Image
	X, Y   int
}

func NewSquare(wall bool, width, height, X, Y int) *Square {
	return &Square{
		IsWall: wall,
		Rect:   ebiten.NewImage(width, height),
		X:      X,
		Y:      Y,
	}
}
