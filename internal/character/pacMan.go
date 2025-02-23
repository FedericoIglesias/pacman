package character

import (
	"pacMan/internal/img"
	"pacMan/internal/timer"

	"github.com/hajimehoshi/ebiten/v2"
)

type Pacman struct {
	X, Y, Dx, Dy float64
	Sprite       *ebiten.Image
	TimerMove    *timer.Timer
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
		X:         10,
		Y:         10,
		Dx:        0,
		Dy:        0,
		Sprite:    child,
		TimerMove: timer.NewTimer(500),
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
		p.Dx = 0.5
		p.Dy = 0
		p.Sprite = rightMouth
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.Dx = -0.5
		p.Dy = 0
		p.Sprite = leftMouth
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Dy = -0.5
		p.Dx = 0
		p.Sprite = upMouth
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Dy = 0.5
		p.Dx = 0
		p.Sprite = downMouth
	}

	p.X += p.Dx
	p.Y += p.Dy

	return nil
}
