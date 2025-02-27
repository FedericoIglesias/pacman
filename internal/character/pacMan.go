package character

import (
	"image/color"
	"pacMan/internal/img"
	"pacMan/internal/timer"

	"github.com/hajimehoshi/ebiten/v2"
)

type Pacman struct {
	X, Y, Dx, Dy float64
	Dir          string
	Sprite       *ebiten.Image
	TimerShow    *timer.Timer
	Stop         bool
}

var (
	child           = ebiten.NewImageFromImage(img.CutImage(13, 13, 3, 43))
	leftMouth       = ebiten.NewImageFromImage(img.CutImage(13, 13, 3, 4))
	totalLeftMouth  = ebiten.NewImageFromImage(img.CutImage(13, 13, 3, 24)) //sizeX 9
	rightMouth      = ebiten.NewImageFromImage(img.CutImage(13, 13, 23, 3))
	totalRightMouth = ebiten.NewImageFromImage(img.CutImage(13, 13, 23, 23))
	upMouth         = ebiten.NewImageFromImage(img.CutImage(13, 13, 44, 3))
	totalUpMouth    = ebiten.NewImageFromImage(img.CutImage(13, 13, 44, 23))
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
		Dir:       "left",
		TimerShow: timer.NewTimer(500),
		Stop:      false,
	}, nil
}

func (p *Pacman) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	op.GeoM.Scale(1, 1)
	p.Sprite.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	screen.DrawImage(p.Sprite, op)
}

func (p *Pacman) Update() error {
	p.TimerShow.Update()

	if p.Stop {
		p.Stop = false
		p.Dx = 0
		p.Dy = 0
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.Dx = 0.5
		p.Dy = 0
		p.Dir = "right"
		p.Sprite = rightMouth
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.Dx = -0.5
		p.Dy = 0
		p.Dir = "left"
		p.Sprite = leftMouth
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Dy = -0.5
		p.Dx = 0
		p.Dir = "up"
		p.Sprite = upMouth
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Dy = 0.5
		p.Dx = 0
		p.Dir = "down"
		p.Sprite = downMouth
	}

	p.X += p.Dx
	p.Y += p.Dy

	if p.TimerShow.IsTimerDone() {
		p.ChangeMouth()
		p.TimerShow.Reset()
	}

	return nil
}

func (p *Pacman) ChangeMouth() {
	if p.Dir == "right" {
		if p.Sprite == rightMouth {
			p.Sprite = totalRightMouth
		} else {
			p.Sprite = rightMouth
		}
	}

	if p.Dir == "left" {
		if p.Sprite == leftMouth {
			p.Sprite = totalLeftMouth
		} else {
			p.Sprite = leftMouth
		}
	}

	if p.Dir == "up" {
		if p.Sprite == upMouth {
			p.Sprite = totalUpMouth
		} else {
			p.Sprite = upMouth
		}
	}

	if p.Dir == "down" {
		if p.Sprite == downMouth {
			p.Sprite = totalDownMouth
		} else {
			p.Sprite = downMouth
		}
	}
}

func (p *Pacman) CheckCollision(bound *ebiten.Image, X, Y float64) bool {
	PMaxX := p.X + 13
	PMaxY := p.Y + 13
	MaxX := X + float64(bound.Bounds().Max.X)
	MaxY := Y + float64(bound.Bounds().Max.Y)

	if p.X <= MaxX && PMaxX >= X && p.Y <= MaxY && PMaxY >= Y {
		p.Stop = true
		return p.Stop
	}
	p.Stop = false
	return p.Stop
}
