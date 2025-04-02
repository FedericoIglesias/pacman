package character

import (
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
	Speed        float64
	StateMouth   int
}

var (
	CHILD             = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 3, 43))
	LEFT_MOUTH        = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 3, 3))
	TOTAL_LEFT_MOUTH  = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 3, 23))
	RIGHT_MOUTH       = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 23, 3))
	TOTAL_RIGHT_MOUTH = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 23, 23))
	UP_MOUTH          = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 43, 3))
	TOTAL_UP_MOUTH    = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 43, 23))
	DOWN_MOUTH        = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 63, 3))
	TOTAL_DOWN_MOUTH  = ebiten.NewImageFromImage(img.CutImage(global.PACMAN_HEIGHT, global.PACMAN_WIDTH, 62, 23))
	STATE_UP          = []*ebiten.Image{CHILD, UP_MOUTH, TOTAL_UP_MOUTH, UP_MOUTH}
	STATE_DOWN        = []*ebiten.Image{CHILD, DOWN_MOUTH, TOTAL_DOWN_MOUTH, DOWN_MOUTH}
	STATE_LEFT        = []*ebiten.Image{CHILD, LEFT_MOUTH, TOTAL_LEFT_MOUTH, LEFT_MOUTH}
	STATE_RIGHT       = []*ebiten.Image{CHILD, RIGHT_MOUTH, TOTAL_RIGHT_MOUTH, RIGHT_MOUTH}
)

func NewPacman() (*Pacman, error) {
	return &Pacman{
		X:          (global.SIDE * 9) + 2.5,
		Y:          (global.SIDE * 15) + 2,
		Dx:         0,
		Dy:         0,
		Sprite:     CHILD,
		Dir:        global.LEFT,
		TimerShow:  timer.NewTimer(250),
		Stop:       "",
		Scale:      global.SCALE,
		Speed:      1,
		StateMouth: 0,
	}, nil
}

func (p *Pacman) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(p.Scale, p.Scale)
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Sprite, op)
}

func (p *Pacman) Update() error {
	p.Stop = ""
	p.TimerShow.Update()

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.Dx = p.Speed
		p.Dy = 0
		p.Dir = global.RIGHT
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.Dx = -p.Speed
		p.Dy = 0
		p.Dir = global.LEFT
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Dy = -p.Speed
		p.Dx = 0
		p.Dir = global.UP
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Dy = p.Speed
		p.Dx = 0
		p.Dir = global.DOWN
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
	if p.StateMouth == 3 {
		p.StateMouth = 0
	} else {
		p.StateMouth++
	}
	if p.Dir == global.RIGHT {
		p.Sprite = STATE_RIGHT[p.StateMouth]
	}

	if p.Dir == global.LEFT {
		p.Sprite = STATE_LEFT[p.StateMouth]
	}

	if p.Dir == global.UP {
		p.Sprite = STATE_UP[p.StateMouth]
	}

	if p.Dir == global.DOWN {
		p.Sprite = STATE_DOWN[p.StateMouth]
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
