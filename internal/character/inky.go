package character

import (
	"pacMan/internal/img"

	"github.com/hajimehoshi/ebiten/v2"
)

type Inky struct {
	X, Y   float64
	Sprite *ebiten.Image
}

var (
	InkyPosUp1        = ebiten.NewImageFromImage(img.CutImage(14, 14, 123, 3))
	InkyPosUp2        = ebiten.NewImageFromImage(img.CutImage(14, 14, 123, 23))
	InkyPosDown1      = ebiten.NewImageFromImage(img.CutImage(14, 14, 123, 43))
	InkyPosDown2      = ebiten.NewImageFromImage(img.CutImage(14, 14, 123, 63))
	InkyPosLeft1      = ebiten.NewImageFromImage(img.CutImage(14, 14, 123, 83))
	InkyPosLeft2      = ebiten.NewImageFromImage(img.CutImage(14, 14, 123, 103))
	InkyPosRight1     = ebiten.NewImageFromImage(img.CutImage(14, 14, 123, 123))
	InkyPosRight2     = ebiten.NewImageFromImage(img.CutImage(14, 14, 123, 143))
	FearInkyPosBlue1  = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 3))
	FearInkyPosBlue2  = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 23))
	FearInkyPosWhite1 = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 43))
	FearInkyPosWhite2 = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 63))
)

func NewInky() (*Inky, error) {
	return &Inky{
		X:      30,
		Y:      10,
		Sprite: InkyPosUp1,
	}, nil
}

func (i *Inky) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(i.X, i.Y)
	op.GeoM.Scale(5, 5)
	screen.DrawImage(i.Sprite, op)
}

func (i *Inky) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		i.Sprite = InkyPosRight2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		i.Sprite = InkyPosUp2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		i.Sprite = InkyPosDown2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		i.Sprite = InkyPosLeft2
	}
	return nil
}
