package stage

type Square struct {
	Type *Type
	X, Y int
}

func NewSquare(t *Type, x, y int) *Square {
	return &Square{
		Type: t,
		X:    x,
		Y:    y,
	}
}

func (s *Square) IsWall() bool {
	return s.Type.IsWall
}
