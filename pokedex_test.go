package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "pikachu bulbasaur charmander",
			expected: []string{"pikachu", "bulbasaur", "charmander"},
		},
		{
			input:    "All tests go inside TestXXX functions that take a *testing.T argument:",
			expected: []string{"all", "tests", "go", "inside", "testxxx", "functions", "that", "take", "a", "*testing.t", "argument:"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Len do not match. \nlen(actual) = %d\nlen(c.expected) = %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("words do not match. \nword = %s\nexpectedWord = %s\n actual slice = %v\n expected slice = %v", word, expectedWord, actual, c.expected)
			}
		}
	}
}
