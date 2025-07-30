package p14

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/longest-common-prefix

func longestCommonPrefix(strs []string) string {
	strsLen := len(strs)
	if strsLen == 0 {
		return ""
	}

	if strsLen == 1 {
		return strs[0]
	}

	str1 := strs[0]
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < strsLen; j++ {
			if i >= len(strs[j]) || strs[j][i] != str1[i] {
				return str1[:i]
			}
		}
	}

	return str1
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			input:    []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			input:    []string{"ab", "a"},
			expected: "a",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := longestCommonPrefix(test.input)
			if result != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, result)
			}
		})
	}
}
