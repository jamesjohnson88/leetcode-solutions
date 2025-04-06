package p136

import (
	"testing"
)

// https://leetcode.com/problems/single-number

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}

	return result
}

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{2, 2, 1},
			expected: 1,
		},
		{
			input:    []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			input:    []int{1},
			expected: 1,
		},
	}

	for _, test := range tests {
		if result := singleNumber(test.input); result != test.expected {
			t.Errorf("expected '%v' and got '%v'", test.expected, result)
		}
	}
}
