package stage

import (
	"pacMan/internal/character"

	"github.com/hajimehoshi/ebiten/v2"
)

type Stage struct {
	Square *character.Square
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
		Square: character.NewSquare(),
		Pacman: Pacman,
		Blinky: Blinky,
	}, nil
}

func (s *Stage) Draw(screen *ebiten.Image) {
	s.Square.Draw(screen)

	s.Pacman.Draw(screen)
	s.Blinky.Draw(screen)
}

func (s *Stage) Update() error {

	if err := s.Pacman.Update(); err != nil {
		return err
	}
	for _, Wall := range s.Square.Wall {
		if s.Pacman.Collider().IntersectsWall(Wall.Collider(), s.Pacman.Dir) {
			s.Pacman.Stop = s.Pacman.Dir
		}
		s.Blinky.Collider().IntersectsSquare(Wall.Collider(), s.Blinky.Dir)
	}
	for i, Dot := range s.Square.Dot {
		if Dot != nil && s.Pacman.Collider().IntersectsGhost(Dot.Collider()) {
			s.Square.Dot[i] = nil
		}
	}
	// if s.Pacman.Collider().IntersectsGhost(s.Blinky.Collider()) {
	// 	s.Pacman.Stop = s.Pacman.Dir
	// }

	if err := s.Blinky.Update(s.Pacman.Collider()); err != nil {
		return err
	}

	s.Pacman.Move()
	return nil
}
