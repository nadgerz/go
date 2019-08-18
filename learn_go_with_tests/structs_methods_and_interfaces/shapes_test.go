package smi

import "testing"

type Rectangle struct {
	Width  float64
	Height float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %0.2f want %0.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := Area(rectangle)
	want := 100.0

	if got != want {
		t.Errorf("got %0.2f want %0.2f", got, want)
	}
}
