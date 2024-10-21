package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertString(t, got, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})
}

func TestDictionaryErr(t *testing.T) {
	tests := []struct {
		name     string
		err      DictionaryErr
		expected string
	}{
		{
			name:     "ErrNotFound",
			err:      ErrNotFound,
			expected: "could not find the word you were looking for",
		},
		{
			name:     "ErrWordExists",
			err:      ErrWordExists,
			expected: "cannot add word because it already exists",
		},
		{
			name:     "ErrWordDoesNotExist",
			err:      ErrWordDoesNotExist,
			expected: "cannot update word because it does not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.expected {
				t.Errorf("got %q, expected %q", got, tt.expected)
			}
		})
	}
}

func assertString(t testing.TB, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got: %q, expected: %q", got, expected)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should fing added word: ", err)
	}

	assertString(t, got, definition)
}

func assertError(t testing.TB, got, expected error) {
	t.Helper()

	if got != expected {
		t.Errorf("got: %q, expected: %q", got, expected)
	}
}
