package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is a test"}
	//got := dict.Search("test")
	//want := "this is a test"
	//
	//assertStrings(t, got, want)
	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})
	
	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("unknown")
		assertErrors(t, got, ErrNotFound)
	})
}

func assertErrors(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}