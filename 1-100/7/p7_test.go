package p7

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

// https://leetcode.com/problems/reverse-integer

func reverse(x int) int {
	s := fmt.Sprintf("%d", x)

	start := 0
	if s[0] == '-' {
		start = 1
	}

	runes := []rune(s)
	for i, j := start, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	result, _ := strconv.Atoi(string(runes))
	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}

	return result
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{
			input:    123,
			expected: 321,
		},
		{
			input:    -123,
			expected: -321,
		},
		{
			input:    120,
			expected: 21,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := reverse(test.input)
			if result != test.expected {
				t.Errorf("expected: %d, got: %d", test.expected, result)
			}
		})
	}
}
