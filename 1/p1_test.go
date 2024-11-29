package p1

import (
	"log"
	"reflect"
	"testing"
)

// https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {
	var compMap = make(map[int]int) // complementary values map
	for i := 0; i < len(nums); i++ {
		compValNeeded := target - nums[i]
		if complement, exists := compMap[compValNeeded]; exists {
			return []int{complement, i}
		} else {
			compMap[nums[i]] = i
		}
	}

	log.Panic("solution not found")
	return nil
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			result: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			result: []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			result: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run("Testing twoSum", func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.result) {
				t.Errorf("twoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.result)
			}
		})
	}
}
