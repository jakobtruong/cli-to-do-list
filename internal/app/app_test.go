package app

import (
	"testing"
)

// var options = []string{
// 	"Show Tasks",
// 	"Add Task",
// 	"Mark Task as Completed",
// 	"Save Tasks to File",
// }

func TestGetMenu(t *testing.T) {
	cases := []struct {
		Name     string
		Input    []string
		Expected string
	}{
		{"Test Case #1", options,
			"\nMenu:\n" +
				"1. Show Tasks\n" +
				"2. Add Task\n" +
				"3. Mark Task as Completed\n" +
				"4. Save Tasks to File\n" +
				"5. Exit\n\n", // Requires "," here since without it, a semicolon is automatically inserted into token stream of non-blank line if line's final token is an integer, floating point, imaginary, rune, or string literal
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := getMenu(options)
			if actual != tc.Expected {
				t.Errorf("Got: %v, Expected: %v", actual, tc.Expected)
			}
		})
	}
}
func TestGetOption(t *testing.T) {
	cases := []struct {
		Name, Input string
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
			actual := getOption(tc.Input, options)
			if actual != tc.Expected {
				t.Errorf("Got: %v, Expected: %v", actual, tc.Expected)
			}
		})
	}
}
