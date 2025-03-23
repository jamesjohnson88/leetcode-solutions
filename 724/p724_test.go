package p724

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/find-pivot-index

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	leftSum := 0
	for i, num := range nums {
		rightSum := total - leftSum - num
		if leftSum == rightSum {
			return i
		}
		leftSum += num
	}

	return -1
}

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			nums:     []int{1, 2, 3},
			expected: -1,
		},
		{
			nums:     []int{2, 1, -1},
			expected: 0,
		},
		{
			nums:     []int{-1, -1, -1, -1, -1, -1},
			expected: -1,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := pivotIndex(test.nums)
			if test.expected != result {
				t.Errorf("expected: %d, got: %d", test.expected, result)
			}
		})
	}
}
