package app

import "testing"

func TestGetUserInput(t *testing.T) {
	got := getUserInput()
	want := 3

	if got != want {
		t.Errorf("got %s want %s given", got, want)
	}
}
