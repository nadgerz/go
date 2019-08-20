package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("word exists", func(t *testing.T) {

		got := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("word does not exist", func(t *testing.T) {

		testWord := "glibble"

		_, err := dictionary.Search(testWord)
		want := "could not find the word 'testWord' you were looking for"

		if err == nil {
			t.Fatal("Expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})

}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
