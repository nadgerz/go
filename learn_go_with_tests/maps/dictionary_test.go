package maps

import (
	"fmt"
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

		_, err := dictionary.Search(testWord)
		want := fmt.Sprintf("could not find the word '%s' you were looking for", testWord)

		if err == nil {
			t.Fatal("Expected to get an error")
		}

		assertStrings(t, err.Error(), want, testWord)
	})

}

func assertStrings(t *testing.T, got, want, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, given)
	}
}
