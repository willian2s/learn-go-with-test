package helloworld

import "testing"

func AssertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Willian", "")
		want := "Hello, Willian"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("say hola to people", func(t *testing.T) {
		got := Hello("Willian", "Spanish")
		want := "Hola, Willian"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("say bonjour to people", func(t *testing.T) {
		got := Hello("Willian", "French")
		want := "Bonjour, Willian"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("say Olá to people", func(t *testing.T) {
		got := Hello("Willian", "Portuguese")
		want := "Olá, Willian"
		AssertCorrectMessage(t, got, want)
	})
}
