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

func (v *Variables) GetFormattedString() string {
	return fmt.Sprintf("X: %d, Y: %d", v.X, v.Y)
}

func (v *Variables) Add() int {
	return v.X + v.Y
}
