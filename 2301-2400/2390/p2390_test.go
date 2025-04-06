package p2390

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/removing-stars-from-a-string

func removeStars(s string) string {
	var stack []rune // bytes also possible but this supports Unicode and reads nicer (to me)
	for _, ch := range s {
		if ch == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return string(stack)
}

func TestRemoveStarts(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{
			s:        "leet**cod*e",
			expected: "lecoe",
		},
		{
			s:        "erase*****",
			expected: "",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := removeStars(test.s)
			if result != test.expected {
				t.Errorf("expected: %s, got: %s", test.expected, result)
			}
		})
	}
}
