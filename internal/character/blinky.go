package character

import (
	"pacMan/internal/img"

	"github.com/hajimehoshi/ebiten/v2"
)

type Blinky struct {
	X, Y   float64
	Sprite *ebiten.Image
}

// Blinky (red), Pinky (pink), Inky (cyan), and Clyde (orange)

var (
	BlinkyPosUp1    = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 3))
	BlinkyPosUp2    = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 23))
	BlinkyPosDown1  = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 43))
	BlinkyPosDown2  = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 63))
	BlinkyPosLeft1  = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 83))
	BlinkyPosLeft2  = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 103))
	BlinkyPosRight1 = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 123))
	BlinkyPosRight2 = ebiten.NewImageFromImage(img.CutImage(14, 14, 83, 143))
)

func NewBlinky() (*Blinky, error) {
	return &Blinky{
		X:      30,
		Y:      30,
		Sprite: BlinkyPosUp1,
	}, nil
}

func (b *Blinky) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X, b.Y)
	op.GeoM.Scale(5, 5)
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
