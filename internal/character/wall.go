package character

import (
	"image/color"
	"pacMan/internal/rect"

	"github.com/hajimehoshi/ebiten/v2"
)

type Wall struct {
	X, Y          float64
	Width, Height float64
	Scale         float64
	Sprite        *ebiten.Image
}

func NewWall(x, y, width, height float64) *Wall {
	return &Wall{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Scale:  1,
		Sprite: ebiten.NewImage(int(width), int(height)),
	}
}

func (w Wall) Collider() rect.Rect {
	return rect.NewRect(w.X, w.Y, w.Width*w.Scale, w.Height*w.Scale)
}

func (w Wall) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(w.Scale, w.Scale)
	op.GeoM.Translate(w.X, w.Y)
	w.Sprite.Fill(color.RGBA{0x00, 0x00, 0xff, 0xff})
	screen.DrawImage(w.Sprite, op)
}

// higurashi
const SIDE = 31

// var STAGE = [][]*Wall{}
func STAGE() []*Wall {
	var STAGE = []*Wall{}
	for r, row := range STAGE_BYNARY {
		for c, colum := range row {
			if colum == 1 {
				STAGE = append(STAGE, NewWall(float64(SIDE*c), float64(SIDE*r), float64(SIDE), float64(SIDE)))
			}
		}
	}
	return STAGE
}

//horizontal
// 	{NewWall(0, 0, side*17, 1), // x y width height
// 		NewWall(0, 157, 93, 1),
// 		NewWall(side, side, side*2, 1),
// 		NewWall(side, side*2, side*2, 1),
// 		NewWall(side, side*3, side*2, 1),
// 		NewWall(side, side*4, side*2, 1),
// 		NewWall(side*4, side, side*3, 1),
// 		NewWall(side*4, side*2, side*3, 1),
// 		NewWall(side*8, side*2, side, 1),
// 		NewWall(side*10, side*2, side*3, 1),
// 		NewWall(side*14, side*2, side*2, 1),
// 		NewWall(side*10, side, side*3, 1),
// 		NewWall(side*14, side, side*2, 1),
// 		NewWall(0, side*19, side*17, 1)}, // 01101110101110110
// 	//vertical
// 	{NewWall(1, 1, 1, side*19),
// 		NewWall(32, 33, 1, 31),
// 		NewWall(32, 95, 1, 31),
// 		NewWall(94, 33, 1, 31),
// 		NewWall(94, 95, 1, 31),
// 		NewWall(side*17, 0, 1, side*19)},
// }

var STAGE_BYNARY = [][]int32{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 0},
	{0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0},
	{1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1},
	{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0},
	{1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
	{1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1},
	{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0},
	{1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0},
	{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	{1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 1},
	{0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0},
	{0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}
