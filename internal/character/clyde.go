package character

import (
	"pacMan/internal/img"

	"github.com/hajimehoshi/ebiten/v2"
)

type Clyde struct {
	X, Y   float64
	Sprite *ebiten.Image
}

var (
	ClydePosUp1        = ebiten.NewImageFromImage(img.CutImage(14, 14, 143, 3))
	ClydePosUp2        = ebiten.NewImageFromImage(img.CutImage(14, 14, 143, 23))
	ClydePosDown1      = ebiten.NewImageFromImage(img.CutImage(14, 14, 143, 43))
	ClydePosDown2      = ebiten.NewImageFromImage(img.CutImage(14, 14, 143, 63))
	ClydePosLeft1      = ebiten.NewImageFromImage(img.CutImage(14, 14, 143, 83))
	ClydePosLeft2      = ebiten.NewImageFromImage(img.CutImage(14, 14, 143, 103))
	ClydePosRight1     = ebiten.NewImageFromImage(img.CutImage(14, 14, 143, 123))
	ClydePosRight2     = ebiten.NewImageFromImage(img.CutImage(14, 14, 143, 143))
	FearClydePosBlue1  = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 3))
	FearClydePosBlue2  = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 23))
	FearClydePosWhite1 = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 43))
	FearClydePosWhite2 = ebiten.NewImageFromImage(img.CutImage(14, 14, 163, 63))
)

func NewClyde() (*Clyde, error) {
	return &Clyde{
		X:      50,
		Y:      30,
		Sprite: ClydePosUp1,
	}, nil
}

func (c *Clyde) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(c.X, c.Y)
	op.GeoM.Scale(5, 5)
	screen.DrawImage(c.Sprite, op)
}

func (c *Clyde) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		c.Sprite = ClydePosRight2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		c.Sprite = ClydePosUp2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		c.Sprite = ClydePosDown2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		c.Sprite = ClydePosLeft2
	}
	return nil
}
