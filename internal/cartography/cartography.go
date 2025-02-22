package cartography

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Cartography struct {
	cart [][]string
	// Walls []Wall
	// Dots  []Dot
}

func NewMap() (*Cartography, error) {
	return &Cartography{

		cart: [][]string{
			{"1", "2", "3", "1", "2", "3"},
			{"2", "3", "1", "2", "3", "1"},
			{"3", "1", "2", "3", "1", "2"},
		},
	}, nil
}

func (c *Cartography) Draw(screen *ebiten.Image) {
	rect := ebiten.NewImage(20, 20)
	square := 20
	for r := 0; r < len(c.cart); r++ {
		for col := 0; col < len(c.cart[r]); col++ {
			op := &ebiten.DrawImageOptions{}
			if c.cart[r][col] == "1" {
				rect.Fill(color.RGBA{0, 0, 0, 0xff})
			}
			if c.cart[r][col] == "2" {
				rect.Fill(color.RGBA{128, 128, 128, 0xff})
			}
			if c.cart[r][col] == "3" {
				rect.Fill(color.RGBA{255, 255, 255, 0xff})
			}
			op.GeoM.Translate(float64(col*square), float64(r*square))
			op.GeoM.Scale(3, 3)
			screen.DrawImage(rect, op)
		}
	}
}
