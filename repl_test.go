package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world   ",
			expected: []string{"hello", "world"},
		},
	}

	for i, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Fatalf("number of outputs mismatch, expected: %d, got: %d", len(c.expected), len(actual))
		}

		for ii := range actual {
			if actual[ii] != c.expected[ii] {
				t.Fatalf("fail: case %d: expected %v: got: %v", i+1, c.expected[i], actual[i])
			}
		}
	}
}
