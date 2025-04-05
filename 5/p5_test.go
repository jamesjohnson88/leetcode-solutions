package p5

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/longest-palindromic-substring

func longestPalindrome(s string) string {
	longest := ""
	longestLen := 0

	for i := 0; i < len(s); i++ {
		l, r := i, i
		// handle odd length palindromes
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > longestLen {
				longest = s[l : r+1]
				longestLen = r - l + 1
			}
			l--
			r++
		}
		// handle even length palindromes
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > longestLen {
				longest = s[l : r+1]
				longestLen = r - l + 1
			}
			l--
			r++
		}
	}

	return longest
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "babad",
			expected: "bab",
		},
		{
			input:    "cbbd",
			expected: "bb",
		},
		{
			input:    "bb",
			expected: "bb",
		},
		{
			input:    "caba",
			expected: "aba",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := longestPalindrome(test.input)
			if result != test.expected {
				t.Errorf("expected: %s, got: %s", test.expected, result)
			}
		})
	}
}
