package character

import (
	"image/color"
	"pacMan/internal/global"
	"pacMan/internal/rect"

	"github.com/hajimehoshi/ebiten/v2"
)

type Wall struct {
	X, Y          float64
	Width, Height float64
	Scale         float64
	Sprite        *ebiten.Image
}

func NewWall(x, y, width, height float64) Wall {
	return Wall{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Scale:  global.SCALE,
	}
}

func (w Wall) Collider() rect.Rect {
	return rect.NewRect(w.X, w.Y, w.Width, w.Height)
}

func (w Wall) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(w.Scale, w.Scale)
	op.GeoM.Translate(w.X, w.Y)
	w.Sprite.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff})
	screen.DrawImage(w.Sprite, op)
}
