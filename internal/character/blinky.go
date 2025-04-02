package character

import (
	"pacMan/internal/global"
	"pacMan/internal/img"
	"pacMan/internal/rect"

	"github.com/hajimehoshi/ebiten/v2"
)

type Blinky struct {
	X, Y            float64
	Sprite          *ebiten.Image
	Scale           float64
	Dir, DirX, DirY string // direction
	MovTicks        int
	ForbidenMov     []string
	PriorityMov     string
}

const (
	WIDTH  = 14
	HEIGHT = 14
)

var (
	BlinkyPosUp1        = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 3))
	BlinkyPosUp2        = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 23))
	BlinkyPosDown1      = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 43))
	BlinkyPosDown2      = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 63))
	BlinkyPosLeft1      = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 83))
	BlinkyPosLeft2      = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 103))
	BlinkyPosRight1     = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 123))
	BlinkyPosRight2     = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 143))
	FearBlinkyPosBlue1  = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 3))
	FearBlinkyPosBlue2  = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 23))
	FearBlinkyPosWhite1 = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 43))
	FearBlinkyPosWhite2 = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 63))
)

func NewBlinky() (*Blinky, error) {
	return &Blinky{
		X:        global.SIDE * 9,
		Y:        global.SIDE * 7,
		Sprite:   BlinkyPosUp1,
		Scale:    global.SCALE,
		Dir:      global.LEFT,
		MovTicks: 0,
	}, nil
}

func (b *Blinky) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.Scale, b.Scale)
	op.GeoM.Translate(b.X, b.Y)
	screen.DrawImage(b.Sprite, op)
}

func (b *Blinky) Update() error {
	b.MovTicks++
	if b.MovTicks == 60 {
		b.MovTicks = 0
		b.CheckPosition()
		b.CalculateDistance(rect.NewRect(b.X, b.Y, 0, 0))
		b.ChoseMov()
	}
	b.Move()
	return nil
}

func (b *Blinky) Collider() rect.Rect {
	bound := b.Sprite.Bounds()
	return rect.NewRect(b.X, b.Y, float64(bound.Dx())*b.Scale, float64(bound.Dy())*b.Scale)
}

func (b *Blinky) Move() {
	if b.Dir == global.RIGHT {
		b.MoveRight()
	}
	if b.Dir == global.LEFT {
		b.MoveLeft()
	}
	if b.Dir == global.UP {
		b.MoveUp()
	}
	if b.Dir == global.DOWN {
		b.MoveDown()
	}
}

func (b *Blinky) MoveRight() {
	b.X += 0.5
}
func (b *Blinky) MoveLeft() {
	b.X -= 0.5
}
func (b *Blinky) MoveUp() {
	b.Y -= 0.5
}
func (b *Blinky) MoveDown() {
	b.Y += 0.5
}

func (b *Blinky) CalculateDistance(objetive rect.Rect) {
	distanceX := b.X - 600
	ditanceY := b.Y - 0

	if distanceX < 0 {
		b.DirX = global.RIGHT
	} else {
		b.DirX = global.LEFT
	}

	if ditanceY < 0 {
		b.DirY = global.DOWN
	} else {
		b.DirY = global.UP
	}

	if distanceX > ditanceY {
		b.PriorityMov = "X"
	} else {
		b.PriorityMov = "Y"
	}
}

func (b *Blinky) CheckPosition() {
	b.ForbidenMov = []string{}
	X := b.X / global.SIDE
	Y := b.Y / global.SIDE

	if global.STAGE_BYNARY[int(Y)][int(X+1)] == 1 {
		b.ForbidenMov = append(b.ForbidenMov, global.RIGHT)
	}
	if global.STAGE_BYNARY[int(Y)][int(X-1)] == 1 {
		b.ForbidenMov = append(b.ForbidenMov, global.LEFT)
	}
	if global.STAGE_BYNARY[int(Y+1)][int(X)] == 1 {
		b.ForbidenMov = append(b.ForbidenMov, global.DOWN)
	}
	if global.STAGE_BYNARY[int(Y-1)][int(X)] == 1 {
		b.ForbidenMov = append(b.ForbidenMov, global.UP)
	}
}

func (b *Blinky) ChoseMov() {

	if b.Dir == global.RIGHT {
		b.ForbidenMov = append(b.ForbidenMov, global.LEFT)
	}
	if b.Dir == global.LEFT {
		b.ForbidenMov = append(b.ForbidenMov, global.RIGHT)
	}
	if b.Dir == global.UP {
		b.ForbidenMov = append(b.ForbidenMov, global.DOWN)
	}
	if b.Dir == global.DOWN {
		b.ForbidenMov = append(b.ForbidenMov, global.UP)
	}

	if b.PriorityMov == "X" && !global.Include(b.ForbidenMov, b.DirX) {
		b.Dir = b.DirX
	}

	if b.PriorityMov == "Y" && !global.Include(b.ForbidenMov, b.DirY) {
		b.Dir = b.DirY
	}

	if global.Include(b.ForbidenMov, b.Dir) {
		if !global.Include(b.ForbidenMov, global.DOWN) {
			b.Dir = global.DOWN
		}
		if !global.Include(b.ForbidenMov, global.UP) {
			b.Dir = global.UP
		}
		if !global.Include(b.ForbidenMov, global.LEFT) {
			b.Dir = global.LEFT
		}
		if !global.Include(b.ForbidenMov, global.RIGHT) {
			b.Dir = global.RIGHT
		}
	}
}
