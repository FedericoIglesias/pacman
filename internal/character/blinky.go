package character

import (
	"image/color"
	"pacMan/internal/global"
	"pacMan/internal/img"
	"pacMan/internal/rect"

	"github.com/hajimehoshi/ebiten/v2"
)

type Blinky struct {
	X, Y   float64
	Sprite *ebiten.Image
	Scale  float64
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
		X:      30,
		Y:      30,
		Sprite: BlinkyPosUp1,
		Scale:  global.SCALE,
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
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		b.Sprite = BlinkyPosRight2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		b.Sprite = BlinkyPosUp2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		b.Sprite = BlinkyPosDown2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		b.Sprite = BlinkyPosLeft2
	}
	return nil
}

func (b *Blinky) Collider() rect.Rect {
	bound := b.Sprite.Bounds()
	return rect.NewRect(b.X, b.Y, float64(bound.Dx())*b.Scale, float64(bound.Dy())*b.Scale)
}
