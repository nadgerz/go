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

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if want != got {
		t.Errorf("got %q want %q", got, want)
	}
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
		t.Errorf("got error %q want %q given, %q", got, want, given)
	}
}
