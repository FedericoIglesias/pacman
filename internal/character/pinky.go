package character

import (
	"pacMan/internal/img"

	"github.com/hajimehoshi/ebiten/v2"
)

type Pinky struct {
	X, Y   float64
	Sprite *ebiten.Image
}

var (
	PinkyPosUp1        = ebiten.NewImageFromImage(img.CutImage(14, 14, 103, 3))
	PinkyPosUp2        = ebiten.NewImageFromImage(img.CutImage(14, 14, 103, 23))
	PinkyPosDown1      = ebiten.NewImageFromImage(img.CutImage(14, 14, 103, 43))
	PinkyPosDown2      = ebiten.NewImageFromImage(img.CutImage(14, 14, 103, 63))
	PinkyPosLeft1      = ebiten.NewImageFromImage(img.CutImage(14, 14, 103, 83))
	PinkyPosLeft2      = ebiten.NewImageFromImage(img.CutImage(14, 14, 103, 103))
	PinkyPosRight1     = ebiten.NewImageFromImage(img.CutImage(14, 14, 103, 123))
	PinkyPosRight2     = ebiten.NewImageFromImage(img.CutImage(14, 14, 103, 143))
	FearPinkyPosBlue1  = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 3))
	FearPinkyPosBlue2  = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 23))
	FearPinkyPosWhite1 = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 43))
	FearPinkyPosWhite2 = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 63))
)

func NewPinky() (*Pinky, error) {
	return &Pinky{
		X:      10,
		Y:      30,
		Sprite: PinkyPosUp1,
	}, nil
}

func (p *Pinky) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	op.GeoM.Scale(5, 5)
	screen.DrawImage(p.Sprite, op)
}

func (p *Pinky) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.Sprite = FearBlinkyPosBlue1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.Sprite = FearBlinkyPosBlue2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.Sprite = FearBlinkyPosWhite1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		p.Sprite = FearBlinkyPosWhite2
	}
	return nil
}
