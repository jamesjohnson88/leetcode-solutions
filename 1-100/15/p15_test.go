package p15

import (
	"fmt"
	"sort"
	"testing"
)

// https://leetcode.com/problems/3sum

func threeSum(nums []int) [][]int {
	var results [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		fixedVal := nums[i]
		if i > 0 && fixedVal == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			total := fixedVal + nums[left] + nums[right]
			if total == 0 {
				results = append(results, []int{fixedVal, nums[left], nums[right]})

				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if total < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return results
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input: []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			input:    []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			input: []int{0, 0, 0},
			expected: [][]int{
				{0, 0, 0},
			},
		},
		{
			input: []int{-1, 0, 0, 1, 1, 1},
			expected: [][]int{
				{-1, 0, 1},
			},
		},
		{
			input: []int{1, -1, -1, 0},
			expected: [][]int{
				{-1, 0, 1},
			},
		},
		{
			input: []int{-1, 0, 1, 2, -1, -4, -1},
			expected: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := threeSum(test.input)
			if !equalTriplets(result, test.expected) {
				t.Errorf("expected: %v, got: %v", test.expected, result)
			}
		})
	}
}

// equalTriplets compares equality between two slices of triplets
func equalTriplets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	sortTriplets(a)
	sortTriplets(b)

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

// sortTriplets sorts triplets for consistent comparison
func sortTriplets(triplets [][]int) {
	for _, triplet := range triplets {
		sort.Ints(triplet)
	}

	sort.Slice(triplets, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if triplets[i][k] != triplets[j][k] {
				return triplets[i][k] < triplets[j][k]
			}
		}
		return false
	})
}
