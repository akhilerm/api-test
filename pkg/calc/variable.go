package calc

import "fmt"

type Variables struct {
	X int
	Y int
}

func NewVariables(x, y int) *Variables {
	return &Variables{
		X: x,
		Y: y,
	}
}

func (s *Variables) GetFormattedString() string {
	return fmt.Sprintf("X: %d, Y: %d", s.X, s.Y)
}
