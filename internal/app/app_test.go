package app

import (
	"testing"
)

/*
General Unit Test Workflow
1.) Write a test
2.) Make compiler pass
3.) Run the test, see that it fails and check if error message is meaningful
4.) Write enough code to make test pass
5.) Refactor if needed
*/

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
				t.Errorf("Got: %q, Expected: %q", actual, tc.Expected) // %q used here to wrap values in double quotes for clarity
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
				t.Errorf("Got: %q, Expected: %q", actual, tc.Expected)
			}
		})
	}
}
