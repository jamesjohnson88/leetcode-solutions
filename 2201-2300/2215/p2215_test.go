package p2215

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// https://leetcode.com/problems/find-the-difference-of-two-arrays

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})

	for _, n := range nums1 {
		set1[n] = struct{}{}
	}
	for _, n := range nums2 {
		set2[n] = struct{}{}
	}

	diff1 := make([]int, 0)
	diff2 := make([]int, 0)

	for n := range set1 {
		if _, found := set2[n]; !found {
			diff1 = append(diff1, n)
		}
	}
	for n := range set2 {
		if _, found := set1[n]; !found {
			diff2 = append(diff2, n)
		}
	}

	return [][]int{diff1, diff2}
}

func TestFindDifference(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{
			nums1: []int{1, 2, 3},
			nums2: []int{2, 4, 6},
			expected: [][]int{
				{1, 3}, {4, 6},
			},
		},
		{
			nums1: []int{1, 2, 3, 3},
			nums2: []int{1, 1, 2, 2},
			expected: [][]int{
				{3}, {},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := findDifference(test.nums1, test.nums2)

			// normalize nil to empty slice
			for i := range result {
				if result[i] == nil {
					result[i] = []int{}
				}

				// normalize the slices as range usage is non-deterministic
				sort.Ints(result[i])
				sort.Ints(test.expected[i])
			}

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("got %v, expected %v", result, test.expected)
			}
		})
	}
}
