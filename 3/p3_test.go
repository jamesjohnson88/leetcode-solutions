package p3

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters

func lengthOfLongestSubstring(s string) int {
	maxLen := 1 // always at least one character
	// from index, look for maxLen+1, update if found, then repeat
	// when repeat chars are found, move pointer along
	// always check that the pointer has enough room to hit a new maxLen
	// if it can't, we should return early

	return maxLen
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "abcabcbb", expected: 3},
		{input: "bbbbb", expected: 1},
		{input: "pwwkew", expected: 3},
	}

	for _, test := range tests {
		t.Run("Testing lengthOfLongestSubstring", func(t *testing.T) {
			got := lengthOfLongestSubstring(test.input)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf(
					"lengthOfLongestSubstring(%v) returned %v; expected %v",
					test.input,
					got,
					test.expected,
				)
			}
		})
	}
}
