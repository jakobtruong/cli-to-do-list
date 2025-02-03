package app

import (
	"testing"
)

func TestGetOption(t *testing.T) {
	cases := []struct {
		Name, input string
		Expected    int
	}{
		{"Test Case #1", "1", 1},
		{"Test Case #2", "-1", -1},
		{"Test Case #3", "abc", -1},
		{"Test Case #4", "", -1},
		{"Test Case #5", "3", 3},
		{"Test Case #6", "1000", -1},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := getOption(tc.input, Options)
			if actual != tc.Expected {
				t.Fatal("Failure")
			}
		})
	}
}
