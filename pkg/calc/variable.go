package calc

import "fmt"

type Variables struct {
	X int
	Y int
	Z int
}

func NewVariables(x, y, z int) *Variables {
	return &Variables{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v *Variables) GetFormattedString() string {
	return fmt.Sprintf("X: %d, Y: %d, Z: %d", v.X, v.Y, v.Z)
}

func (v *Variables) Add() int {
	return v.X + v.Y + v.Z
}

func (v *Variables) Multiply() int {
	return v.X * v.Y * v.Z
}
