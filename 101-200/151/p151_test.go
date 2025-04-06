package p151

import (
	"slices"
	"strings"
	"testing"
)

// https://leetcode.com/problems/reverse-words-in-a-string

func reverseWords(s string) string {
	fields := strings.Fields(s)
	slices.Reverse(fields)
	return strings.Join(fields, " ")
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			input:    "  hello world  ",
			expected: "world hello",
		},
		{
			input:    "a good   example",
			expected: "example good a",
		},
	}

	for _, test := range tests {
		if result := reverseWords(test.input); result != test.expected {
			t.Errorf("result %q, expected %q", result, test.expected)
		}
	}
}
