package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("word exists", func(t *testing.T) {

		testWord := "test"

		got, _ := dictionary.Search(testWord)
		want := "this is just a test"

		assertStrings(t, got, want, testWord)
	})

	t.Run("word does not exist", func(t *testing.T) {

		testWord := "glibble"

		_, got := dictionary.Search(testWord)

		assertError(t, got, ErrNotFound, testWord)
	})

}

func assertStrings(t *testing.T, got, want, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, given)
	}
}

func assertError(t *testing.T, got, want error, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, given)
	}
}
