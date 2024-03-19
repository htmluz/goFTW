package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Luiz", "")
		want := "Hello, Luiz"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say hello when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Luiz", "Spanish")
		want := "Hola, Luiz"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Luiz", "Portuguese")
		want := "Ola, Luiz"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
