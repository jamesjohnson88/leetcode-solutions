package p88

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/merge-sorted-array

func merge(nums1 []int, m int, nums2 []int, n int) {
	num1Idx := m - 1
	num2Idx := n - 1
	writeIdx := m + n - 1

	for num1Idx >= 0 && num2Idx >= 0 {
		if nums1[num1Idx] > nums2[num2Idx] {
			nums1[writeIdx] = nums1[num1Idx]
			num1Idx--
		} else {
			nums1[writeIdx] = nums2[num2Idx]
			num2Idx--
		}
		writeIdx--
	}

	for num2Idx >= 0 {
		nums1[writeIdx] = nums2[num2Idx]
		num2Idx--
		writeIdx--
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
		{
			nums1:    []int{4, 5, 6, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			nums1:    []int{1, 3, 5, 0, 0},
			m:        3,
			nums2:    []int{2, 4},
			n:        2,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			nums1:    []int{2, 0},
			m:        1,
			nums2:    []int{1},
			n:        1,
			expected: []int{1, 2},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			merge(test.nums1, test.m, test.nums2, test.n)

			if !equalSlices(test.nums1, test.expected) {
				t.Errorf("expected: %v, got: %v", test.expected, test.nums1)
			}
		})
	}
}

// equalSlices is a helper function for slice comparison
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
