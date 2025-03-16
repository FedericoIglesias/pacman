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
	Wall   *character.Wall
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
		Wall:   character.NewWall(50, 50, 1, 20),
	}, nil
}

func (s *Stage) Draw(screen *ebiten.Image) {
	s.Pacman.Draw(screen)
	s.Wall.Draw(screen)
	// s.Blinky.Draw(screen)
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
	if s.Pacman.Collider().IntersectsWall(s.Wall.Collider()) {
		s.Pacman.Stop = s.Pacman.Dir
	}

	// if s.Pacman.Collider().IntersectsGhost(s.Blinky.Collider()) {
	// 	s.Pacman.Stop = s.Pacman.Dir
	// }

	if err := s.Pacman.Update(); err != nil {
		return err
	}
	// if err := s.Blinky.Update(); err != nil {
	// 	return err
	// }

	return nil
}
