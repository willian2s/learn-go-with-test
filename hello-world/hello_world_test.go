package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Willian", "")
		want := "Hello, Willian!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Elodie", "Portuguese")
		want := "Olá, Elodie!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Elodie", "Japanese")
		want := "こんにちは, Elodie!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// testing.TB is an interface that is available in the testing package and
	// and means that we can pass in anything that satisfies the testing.TB interface.
	// In this case, we pass in *testing.T.
	// t.Helper() tells the test suite that this method is a helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
