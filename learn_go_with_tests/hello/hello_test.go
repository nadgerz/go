package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Steve", "")
		want := "Hello, Steve"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {

		got := Hello("Estaban", "Spanish")
		want := "Hola, Estaban"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {

		got := Hello("Etienne", "French")
		want := "Bonjour, Etienne"

		assertCorrectMessage(t, got, want)
	})

}
