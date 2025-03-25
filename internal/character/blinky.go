package character

import (
	"image/color"
	"pacMan/internal/global"
	"pacMan/internal/img"
	"pacMan/internal/rect"

	"github.com/hajimehoshi/ebiten/v2"
)

type Blinky struct {
	X, Y     float64
	Sprite   *ebiten.Image
	Scale    float64
	Dir      string // direction
	MovTicks int
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
		Y:        global.SIDE*7 + 1,
		Sprite:   BlinkyPosUp1,
		Scale:    global.SCALE,
		Dir:      global.RIGHT,
		MovTicks: 0,
	}, nil
}

func (b *Blinky) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.Scale, b.Scale)
	op.GeoM.Translate(b.X, b.Y)
	b.Sprite.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	screen.DrawImage(b.Sprite, op)
}

func (b *Blinky) Update() error {
	b.MovTicks++
	b.Destination()
	if b.MovTicks == 60 {
		b.MovTicks = 0
	}
	return nil
}

func (b *Blinky) Collider() rect.Rect {
	bound := b.Sprite.Bounds()
	return rect.NewRect(b.X, b.Y, float64(bound.Dx())*b.Scale, float64(bound.Dy())*b.Scale)
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

func (b *Blinky) Destination() {
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
