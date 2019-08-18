package smi

import (
	"errors"
	"fmt"
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
	if rect.Width < 0 || rect.Height < 0 {
		return -1, errors.New(fmt.Sprintf("Neither Width nor Height can be < 0 [%v, %v]", rect.Width, rect.Height))
	}

	return rect.Width * rect.Height, nil
}
