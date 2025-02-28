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
	Stop         string
}

const (
	LEFT  = "left"
	RIGHT = "right"
	UP    = "up"
	DOWN  = "down"
)

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
		Stop:      "",
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

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.Dx = 0.5
		p.Dy = 0
		p.Dir = RIGHT
		p.Sprite = rightMouth
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.Dx = -0.5
		p.Dy = 0
		p.Dir = LEFT
		p.Sprite = leftMouth
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Dy = -0.5
		p.Dx = 0
		p.Dir = UP
		p.Sprite = upMouth
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Dy = 0.5
		p.Dx = 0
		p.Dir = DOWN
		p.Sprite = downMouth
	}

	if p.Stop == LEFT && p.Dx < 0 {
		p.Dx = 0
	}
	if p.Stop == RIGHT && p.Dx > 0 {
		p.Dx = 0
	}
	if p.Stop == UP && p.Dy < 0 {
		p.Dy = 0
	}
	if p.Stop == DOWN && p.Dy > 0 {
		p.Dy = 0
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
	if p.Dir == RIGHT {
		if p.Sprite == rightMouth {
			p.Sprite = totalRightMouth
		} else {
			p.Sprite = rightMouth
		}
	}

	if p.Dir == LEFT {
		if p.Sprite == leftMouth {
			p.Sprite = totalLeftMouth
		} else {
			p.Sprite = leftMouth
		}
	}

	if p.Dir == UP {
		if p.Sprite == upMouth {
			p.Sprite = totalUpMouth
		} else {
			p.Sprite = upMouth
		}
	}

	if p.Dir == DOWN {
		if p.Sprite == downMouth {
			p.Sprite = totalDownMouth
		} else {
			p.Sprite = downMouth
		}
	}
}

func (p *Pacman) CheckCollision(bound *ebiten.Image, X, Y float64) {
	PMaxX := p.X + 13
	PMaxY := p.Y + 13
	MaxX := X + float64(bound.Bounds().Max.X)
	MaxY := Y + float64(bound.Bounds().Max.Y)

	if p.X <= MaxX && PMaxX >= X && p.Y <= MaxY && PMaxY >= Y {

		if p.X <= MaxX {
			p.Stop = LEFT
		}
		if PMaxX >= X {
			p.Stop = RIGHT
		}
		if p.Y <= MaxY {
			p.Stop = UP
		}
		if PMaxY >= Y {
			p.Stop = DOWN
		}
	}
}
