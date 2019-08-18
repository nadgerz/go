package smi

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40

	if got != want {
		t.Errorf("got %0.2f want %0.2f", got, want)
	}
}
