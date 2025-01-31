package app

import "testing"

func TestGetUserInput(t *testing.T) {
	cases := []struct{
		Name string,
		Input, Expected int
	}{
		{"Test Case #1", "1", 1},
		{"Test Case #2", "-1", -1},
	}
	for i, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {

		})
	}
	got := getUserInput()
	want := 3

	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
