package rect

import "pacMan/internal/global"

type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewRect(x, y, width, height float64) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (r Rect) MaxX() float64 {
	return r.X + r.Width
}

func (r Rect) MaxY() float64 {
	return r.Y + r.Height
}

func (r Rect) IntersectsGhost(other Rect) bool {
	return r.X <= other.MaxX() &&
		other.X <= r.MaxX() &&
		r.Y <= other.MaxY() &&
		other.Y <= r.MaxY()
}

func (r Rect) IntersectsWall(other Rect, dir string) bool {

	var (
		betweenY = (r.Y >= other.Y && r.Y <= other.MaxY()) || (r.MaxY() >= other.Y && r.MaxY() <= other.MaxY()) || (r.Y <= other.Y && r.MaxY() >= other.MaxY())
		betweenX = (r.X >= other.X && r.X <= other.MaxX()) || (r.MaxX() >= other.X && r.MaxX() <= other.MaxX()) || (r.X <= other.X && r.MaxX() >= other.MaxX())
	)

	if dir == global.RIGHT && betweenY {
		return r.MaxX()+1 >= other.X && r.MaxX() <= other.X
	}

	if dir == global.LEFT && betweenY {
		return r.X-1 <= other.MaxX() && r.X >= other.MaxX()
	}

	if dir == global.UP && betweenX {
		return r.Y-1 <= other.MaxY() && r.Y >= other.MaxY()
	}

	if dir == global.DOWN && betweenX {
		return r.MaxY()+1 >= other.Y && r.MaxY() <= other.Y
	}
	return false
}

func (r Rect) IntersectsSquare(other Rect, dir string) []string {
	dirUnavailable := []string{}
	if other.X+15 >= r.X && other.X+15 <= r.MaxX() { // dir = right
		dirUnavailable = append(dirUnavailable, global.RIGHT)
	}
	if other.MaxX()-15 >= r.X && other.MaxX()-15 <= r.MaxX() { // dir = left
		dirUnavailable = append(dirUnavailable, global.LEFT)
	}
	if other.Y-15 >= r.Y && other.Y-15 <= r.MaxY() { // dir = Up
		dirUnavailable = append(dirUnavailable, global.UP)
	}
	if other.MaxY()+15 >= r.Y && other.MaxY()+15 <= r.MaxY() { // dir = Down
		dirUnavailable = append(dirUnavailable, global.DOWN)
	}
	return dirUnavailable
}
