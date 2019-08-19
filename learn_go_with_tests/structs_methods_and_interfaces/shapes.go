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

func Perimeter(rect Rectangle) (float64, error) {
	return 2 * (rect.Width + rect.Height), nil
}

func Area(rect Rectangle) (float64, error) {
	if rect.Width <= 0 || rect.Height <= 0 {
		return -1, fmt.Errorf("Neither Rectangle Width nor Height can be <= 0 [%v, %v]", rect.Width, rect.Height)
	}

	return rect.Width * rect.Height, nil
}

func Area(circle Circle) (float64, error) {
	return 0.0, nil
}

func (c *Circle) Area(radius float64) (float64, error) {
	if radius <= 0 {
		return -1, fmt.Errorf("Circle radius can not be <= 0 [%v]", radius)
	}

	return math.Pi * radius * radius, nil
}
