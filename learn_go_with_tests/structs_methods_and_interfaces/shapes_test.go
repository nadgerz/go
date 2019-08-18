package smi

import (
	"fmt"
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

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}

		got, err := Area(rectangle)
		if err != nil {
		}

		fmt.Printf("%#v\n", err)
		fmt.Printf("%T\n", err)

		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("bad rectangle - negative widths", func(t *testing.T) {
		rectangle := Rectangle{-1, 6}

		got, err := Area(rectangle)
		if err != nil {
		}

		fmt.Printf("%#v\n", err)
		fmt.Printf("%T\n", err)

		// TBD: Would rather pass back an error 'object'
		want := -1.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}

		got, err := circle.Area(circle.Radius)
		if err != nil {
		}

		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
