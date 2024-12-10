package p3

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		maxCount := 0
		dict := make(map[rune]bool)
		subIndex := i

		for subIndex < len(s) {
			if exists := dict[rune(s[subIndex])]; !exists {
				dict[rune(s[subIndex])] = true
				maxCount++
				subIndex++
				if maxCount > maxLen {
					maxLen = maxCount
				}
			} else {
				break
			}
		}
	}

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
