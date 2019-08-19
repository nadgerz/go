package smi

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() (float64, error)
}

func Perimeter(rect Rectangle) (float64, error) {
	return 2 * (rect.Width + rect.Height), nil
}

func (r Rectangle) Area() (float64, error) {
	if r.Width <= 0 || r.Height <= 0 {
		return -1, fmt.Errorf("Neither Rectangle Width nor Height can be <= 0 [%v, %v]", r.Width, r.Height)
	}

	return r.Width * r.Height, nil
}

func (c Circle) Area() (float64, error) {
	if c.Radius <= 0 {
		return -1, fmt.Errorf("Circle radius can not be <= 0 [%v]", c.Radius)
	}

	return math.Pi * c.Radius * c.Radius, nil
}
