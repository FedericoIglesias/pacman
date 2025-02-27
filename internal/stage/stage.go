package stage

import (
	"fmt"
	"pacMan/internal/character"
	"pacMan/internal/global"

	"github.com/hajimehoshi/ebiten/v2"
)

type Stage struct {
	Stage  [][]Square
	Pacman *character.Pacman
	Blinky *character.Blinky
}

func NewMap() (*Stage, error) {
	Pacman, err := character.NewPacman()
	if err != nil {
		return nil, err
	}

	Blinky, err := character.NewBlinky()
	if err != nil {
		return nil, err
	}

	return &Stage{
		Stage:  MakeStage(),
		Pacman: Pacman,
		Blinky: Blinky,
	}, nil
}

func (s *Stage) Draw(screen *ebiten.Image) {
	s.Blinky.Draw(screen)
	// stop := s.Pacman.X > s.Blinky.X && s.Pacman.Y > s.Blinky.Y

	// if stop {
	// 	fmt.Println("Stop")
	// } else {
	// 	fmt.Println("Move")
	// }
	s.Pacman.Draw(screen)
	// stop := s.Pacman.CheckCollision(s.Blinky.Sprite, s.Blinky.X, s.Blinky.Y)
	// squareX := global.SCREEN_WIDTH / len(s.Stage[0])
	// squareY := global.SCREEN_HEIGHT / len(s.Stage)
	// rect := ebiten.NewImage(squareX, squareY)
	// for r := 0; r < len(s.Stage); r++ {
	// 	for col := 0; col < len(s.Stage[r]); col++ {
	// 		op := &ebiten.DrawImageOptions{}
	// 		if s.Stage[r][col].IsWall {
	// 			rect.Fill(color.RGBA{0, 0, 255, 0xff})
	// 		} else {
	// 			rect.Fill(color.RGBA{0, 0, 0, 0xff})
	// 		}
	// 		op.GeoM.Translate(float64(col*squareX), float64(r*squareY))
	// 		op.GeoM.Scale(1, 1)
	// 		screen.DrawImage(rect, op)
	// 	}
	// }
}

func MakeStage() [][]Square {
	st := [][]Square{}

	SquareWidth := global.SCREEN_WIDTH / len(global.STAGE_1[0]) //31
	SquareHeight := global.SCREEN_HEIGHT / len(global.STAGE_1)  // 27

	fmt.Printf("SquareWidth: %v\n", SquareWidth)
	fmt.Printf("SquareHeight: %v\n", SquareHeight)

	for row := 0; row < len(global.STAGE_1); row++ {
		newRow := []Square{}
		for col := 0; col < len(global.STAGE_1[row]); col++ {
			isWall := global.STAGE_1[row][col] == "b"
			square := NewSquare(isWall, SquareWidth, SquareHeight, col*SquareWidth, row*SquareHeight)
			newRow = append(newRow, *square)
		}
		st = append(st, newRow)
	}
	return st
}

func (s *Stage) Update() error {
	s.Pacman.CheckCollision(s.Blinky.Sprite, s.Blinky.X, s.Blinky.Y)

	if err := s.Pacman.Update(); err != nil {
		return err
	}
	if err := s.Blinky.Update(); err != nil {
		return err
	}

	return nil
}
