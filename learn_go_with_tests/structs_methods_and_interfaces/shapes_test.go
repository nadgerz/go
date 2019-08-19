package smi

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got, err := Perimeter(rectangle)
	if err != nil {
	}

	want := 40.0

	if got != want {
		t.Errorf("got %0.2f want %0.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()

		got, err := shape.Area()
		if err != nil {
		}

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}

		checkArea(t, rectangle, 72.0)
	})

	t.Run("bad rectangle - negative widths", func(t *testing.T) {
		rectangle := Rectangle{-1, 6}

		got, err := rectangle.Area()
		if err == nil {
			t.Errorf("Expected Area() to throw an error for a regative width/height [%v]", got)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}

		checkArea(t, circle, 314.1592653589793)
	})
}
