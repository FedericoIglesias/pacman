package character

import (
	"image/color"
	"pacMan/internal/global"
	"pacMan/internal/img"
	"pacMan/internal/rect"
	"pacMan/internal/timer"

	"github.com/hajimehoshi/ebiten/v2"
)

type Pacman struct {
	X, Y, Dx, Dy float64
	Dir          string
	Sprite       *ebiten.Image
	TimerShow    *timer.Timer
	Stop         string
	Scale        float64
}

var (
	child           = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 3, 43))
	leftMouth       = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 3, 4))
	totalLeftMouth  = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 3, 24))
	rightMouth      = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 23, 3))
	totalRightMouth = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 23, 23))
	upMouth         = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 44, 3))
	totalUpMouth    = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 44, 23))
	downMouth       = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 63, 3))
	totalDownMouth  = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 62, 23))
)

func NewPacman() (*Pacman, error) {
	return &Pacman{
		X:         0,
		Y:         0,
		Dx:        0,
		Dy:        0,
		Sprite:    child,
		Dir:       global.LEFT,
		TimerShow: timer.NewTimer(500),
		Stop:      "",
		Scale:     global.SCALE,
	}, nil
}

func (p *Pacman) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(p.Scale, p.Scale)
	op.GeoM.Translate(p.X, p.Y)
	p.Sprite.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	screen.DrawImage(p.Sprite, op)
}

func (p *Pacman) Update() error {
	p.TimerShow.Update()

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.Dx = 0.5
		p.Dy = 0
		p.Dir = global.RIGHT
		p.Sprite = rightMouth
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.Dx = -0.5
		p.Dy = 0
		p.Dir = global.LEFT
		p.Sprite = leftMouth
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Dy = -0.5
		p.Dx = 0
		p.Dir = global.UP
		p.Sprite = upMouth
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Dy = 0.5
		p.Dx = 0
		p.Dir = global.DOWN
		p.Sprite = downMouth
	}

	if p.TimerShow.IsTimerDone() {
		p.ChangeMouth()
		p.TimerShow.Reset()
	}

	return nil
}

func (p *Pacman) Move() {
	if p.Stop == p.Dir {
		p.Dx = 0
	}
	if p.Stop == p.Dir {
		p.Dx = 0
	}
	if p.Stop == p.Dir {
		p.Dy = 0
	}
	if p.Stop == p.Dir {
		p.Dy = 0
	}
	p.X += p.Dx
	p.Y += p.Dy
}

func (p *Pacman) ChangeMouth() {
	if p.Dir == global.RIGHT {
		if p.Sprite == rightMouth {
			p.Sprite = totalRightMouth
		} else {
			p.Sprite = rightMouth
		}
	}

	if p.Dir == global.LEFT {
		if p.Sprite == leftMouth {
			p.Sprite = totalLeftMouth
		} else {
			p.Sprite = leftMouth
		}
	}

	if p.Dir == global.UP {
		if p.Sprite == upMouth {
			p.Sprite = totalUpMouth
		} else {
			p.Sprite = upMouth
		}
	}

	if p.Dir == global.DOWN {
		if p.Sprite == downMouth {
			p.Sprite = totalDownMouth
		} else {
			p.Sprite = downMouth
		}
	}
}

func (p *Pacman) Collider() rect.Rect {
	bound := p.Sprite.Bounds()

	return rect.NewRect(
		p.X,
		p.Y,
		float64(bound.Dx())*p.Scale,
		float64(bound.Dy())*p.Scale,
	)
}
