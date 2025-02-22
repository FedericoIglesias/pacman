package internal

import (
	"pacMan/internal/img"

	"github.com/hajimehoshi/ebiten/v2"
)

type Pacman struct {
	X, Y   float64
	Sprite *ebiten.Image
}

var (
	child           = ebiten.NewImageFromImage(img.CutImage(13, 13, 3, 43))
	leftMouth       = ebiten.NewImageFromImage(img.CutImage(13, 13, 3, 4))
	totalLeftMouth  = ebiten.NewImageFromImage(img.CutImage(13, 13, 3, 27)) //sizeX 9
	rightMouth      = ebiten.NewImageFromImage(img.CutImage(13, 13, 23, 3))
	totalRightMouth = ebiten.NewImageFromImage(img.CutImage(13, 13, 23, 23))
	upMouth         = ebiten.NewImageFromImage(img.CutImage(13, 13, 44, 3))
	totalUpMouth    = ebiten.NewImageFromImage(img.CutImage(13, 13, 47, 23))
	downMouth       = ebiten.NewImageFromImage(img.CutImage(13, 13, 63, 3))
	totalDownMouth  = ebiten.NewImageFromImage(img.CutImage(13, 13, 62, 23))
)

func NewPacman() (*Pacman, error) {
	return &Pacman{
		X:      10,
		Y:      10,
		Sprite: child,
	}, nil
}

func (p *Pacman) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	op.GeoM.Scale(5, 5)
	screen.DrawImage(p.Sprite, op)
}

func (p *Pacman) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.Sprite = rightMouth
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && p.Sprite == rightMouth {
		p.Sprite = totalRightMouth
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.Sprite = leftMouth
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && p.Sprite == leftMouth {
		p.Sprite = totalLeftMouth
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Sprite = downMouth
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && p.Sprite == downMouth {
		p.Sprite = totalDownMouth
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Sprite = upMouth
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.Sprite == upMouth {
		p.Sprite = totalUpMouth
	}

	return nil
}
