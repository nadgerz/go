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

	t.Run("new word", func(t *testing.T) {

		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
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
