package internal

import (
	"pacMan/internal/img"

	"github.com/hajimehoshi/ebiten/v2"
)

type Pacman struct {
	X, Y   float64
	Sprite *ebiten.Image
}

func NewPacman() (*Pacman, error) {
	img.CutImage(32, 32, 0, 0)
	return &Pacman{
		X:      10,
		Y:      10,
		Sprite: ebiten.NewImageFromImage(img.CutImage(32, 32, 0, 0)),
	}, nil
}

func (p *Pacman) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Sprite, op)
}
