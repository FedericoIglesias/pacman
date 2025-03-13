package rect

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

func (r Rect) Intersects(other Rect) bool {
	return r.X <= other.MaxX() && //r.X = left  and other.MaxX() = right
		other.X <= r.MaxX() && // other.X = left and r.MaxX() = right
		r.Y <= other.MaxY() && // r.Y = top and other.MaxY() = bottom
		other.Y <= r.MaxY() // other.Y = top and r.MaxY() = bottom
}
