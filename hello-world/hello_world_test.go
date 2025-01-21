package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Willian")
	want := "Hello, Willian!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
