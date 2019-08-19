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

func TestAreas(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			name: "Rectangle",
			shape: Rectangle{
				Width:  12,
				Height: 6,
			},
			hasArea: 72.0,
		},
		{
			name: "Circle",
			shape: Circle{
				Radius: 10,
			},
			hasArea: 314.1592653589793,
		},
		{
			name: "Triangle",
			shape: Triangle{
				Base:   12,
				Height: 6,
			},
			hasArea: 36.0,
		},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {

			got, _ := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})
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
