package _1

import (
	"testing"
)

// https://leetcode.com/problems/container-with-most-water

func maxArea(height []int) int {
	largestArea := 0
	li, ri := 0, len(height)-1

	area := func(v1, v2 int) {
		a := v1 * v2
		if a > largestArea {
			largestArea = a
		}
	}

	for li < ri {
		width := ri - li
		if height[li] < height[ri] {
			area(height[li], width)
			li++
		} else {
			area(height[ri], width)
			ri--
		}
	}

	return largestArea
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			input:    []int{1, 1},
			expected: 1,
		},
		{
			input:    []int{1, 2, 3, 1000, 9},
			expected: 9,
		},
		{
			input:    []int{1, 2, 1},
			expected: 2,
		},
	}

	for _, test := range tests {
		actual := maxArea(test.input)
		if actual != test.expected {
			t.Errorf("got '%d', expected '%d'", actual, test.expected)
		}
	}
}
