package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "oneword",
			expected: []string{"oneword"},
		},
		{
			input:    "aLterNATING Text",
			expected: []string{"alternating", "text"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected length was %v, got %v", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("expected word: %s, got %s", word, expectedWord)
			}
		}
	}
}
