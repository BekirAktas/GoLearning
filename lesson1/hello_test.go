package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Deniz", "")
		want := "Hello, Deniz"

		assertCorrectMessage(t, got, want)
	})

	t.Run("return 'Hello, world' when an empty string is given to func", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello name with given spanish", func(t *testing.T) {
		got := Hello("Deniz", "Spanish")
		want := "Hola, Deniz"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello name with given turkish", func(t *testing.T) {
		got := Hello("Deniz", "Turkish")
		want := "Merhaba, Deniz"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}