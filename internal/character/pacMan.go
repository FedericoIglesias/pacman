package character

import (
	"fmt"
	"image/color"
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
	Stop         bool
	Scale        float64
}

const (
	LEFT   = "left"
	RIGHT  = "right"
	UP     = "up"
	DOWN   = "down"
	HEIGHT = 13
	WIDTH  = 13
)

var (
	child           = ebiten.NewImageFromImage(img.CutImage(HEIGHT, WIDTH, 3, 43))
	leftMouth       = ebiten.NewImageFromImage(img.CutImage(HEIGHT, WIDTH, 3, 4))
	totalLeftMouth  = ebiten.NewImageFromImage(img.CutImage(HEIGHT, WIDTH, 3, 24))
	rightMouth      = ebiten.NewImageFromImage(img.CutImage(HEIGHT, WIDTH, 23, 3))
	totalRightMouth = ebiten.NewImageFromImage(img.CutImage(HEIGHT, WIDTH, 23, 23))
	upMouth         = ebiten.NewImageFromImage(img.CutImage(HEIGHT, WIDTH, 44, 3))
	totalUpMouth    = ebiten.NewImageFromImage(img.CutImage(HEIGHT, WIDTH, 44, 23))
	downMouth       = ebiten.NewImageFromImage(img.CutImage(HEIGHT, WIDTH, 63, 3))
	totalDownMouth  = ebiten.NewImageFromImage(img.CutImage(HEIGHT, WIDTH, 62, 23))
)

func NewPacman() (*Pacman, error) {
	// scale :=

	return &Pacman{
		X:         0,
		Y:         0,
		Dx:        0,
		Dy:        0,
		Sprite:    child,
		Dir:       "left",
		TimerShow: timer.NewTimer(500),
		Stop:      false,
		Scale:     2,
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

	if p.Stop && p.Dir == LEFT {
		p.Dx = 0
	}
	if p.Stop && p.Dir == RIGHT {
		p.Dx = 0
	}
	if p.Stop && p.Dir == UP {
		p.Dy = 0
	}
	if p.Stop && p.Dir == DOWN {
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
	// PMaxX := p.X + 13*p.Scale
	// PMaxY := p.Y + 13*p.Scale
	// MaxX := X + float64(bound.Bounds().Max.X)
	// MaxY := Y + float64(bound.Bounds().Max.Y)

	// var (
	// 	BetweenY   = p.Y <= MaxY && p.Y >= Y || PMaxY <= MaxY && PMaxY >= Y
	// 	BetweenX   = p.X <= MaxX && p.X >= X || PMaxX <= MaxX && PMaxX >= X
	// 	BoundLeft  = PMaxX >= X && BetweenY && p.X < X
	// 	BoundRight = p.X <= MaxX && BetweenY && PMaxX > MaxX
	// 	BoundUp    = PMaxY >= Y && BetweenX && p.Y < Y
	// 	BoundDown  = p.Y <= MaxY && BetweenX && PMaxY > MaxY
	// 	Collision  = BoundLeft || BoundRight || BoundUp || BoundDown
	// )

	// fmt.Printf("Px:%v, PMaxX:%v,Py:%v, PMaxY:%v,\n", p.X, PMaxX, p.Y, PMaxY)
	fmt.Printf("DX:%v, DY:%v\n", p.Sprite.Bounds().Dy(), p.Sprite.Bounds().Dy())
	//30 - 44
	// if Collision {
	// 	if BoundLeft && p.Dir == RIGHT {
	// 		p.Stop = RIGHT
	// 		return
	// 	}
	// 	if BoundRight && p.Dir == LEFT {
	// 		p.Stop = LEFT
	// 		return
	// 	}
	// 	if BoundUp && p.Dir == DOWN {
	// 		p.Stop = DOWN
	// 		return
	// 	}
	// 	if BoundDown && p.Dir == UP {
	// 		p.Stop = UP
	// 		return
	// 	}
	// } else {
	// 	p.Stop = ""
	// }
}

func (p *Pacman) Collider() rect.Rect {
	bounds := p.Sprite.Bounds()

	return rect.NewRect(
		p.X,
		p.Y,
		float64(bounds.Dx())*p.Scale,
		float64(bounds.Dy())*p.Scale,
	)
}
