package stage

type Square struct {
	IsWall bool
	X, Y   int
}

func NewSquare(w bool, x, y int) *Square {
	return &Square{
		IsWall: w,
		X:      x,
		Y:      y,
	}
}
