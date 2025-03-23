package character

import (
	"image/color"
	"pacMan/internal/rect"

	"github.com/hajimehoshi/ebiten/v2"
)

type Dot struct {
	X, Y          float64
	Width, Height float64
	Sprite        *ebiten.Image
}

func NewDot(x, y, width, height float64) *Dot {
	return &Dot{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Sprite: ebiten.NewImage(int(width), int(height)),
	}
}

func (d Dot) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(d.X, d.Y)
	d.Sprite.Fill(color.RGBA{0xF0, 0xE5, 0x91, 0xff})
	screen.DrawImage(d.Sprite, op)
}

func (d Dot) Collider() rect.Rect {
	return rect.NewRect(d.X, d.Y, d.Width, d.Height)
}
