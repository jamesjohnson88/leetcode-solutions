package p4

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ml := len(nums1) + len(nums2)
	mergedArray := make([]int, ml)
	n1i, n2i, mi := 0, 0, 0

	for n1i < len(nums1) && n2i < len(nums2) {
		if nums1[n1i] < nums2[n2i] {
			mergedArray[mi] = nums1[n1i]
			n1i++
		} else {
			mergedArray[mi] = nums2[n2i]
			n2i++
		}
		mi++
	}

	// one of these will be obsolete but branch-less is likely still more performant
	copy(mergedArray[mi:], nums1[n1i:])
	copy(mergedArray[mi:], nums2[n2i:])

	if ml%2 == 0 {
		v1 := mergedArray[ml/2-1]
		v2 := mergedArray[ml/2]
		return float64(v1+v2) / 2
	}

	return float64(mergedArray[ml/2])
}

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2,
		},
		{
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.5,
		},
		{
			nums1:    []int{1, 2, 3, 4, 5},
			nums2:    []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
			expected: 9,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := findMedianSortedArrays(test.nums1, test.nums2)
			if result != test.expected {
				t.Errorf("got %f, expected %f", result, test.expected)
			}
		})
	}
}
