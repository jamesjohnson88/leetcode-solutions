package p118

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/pascals-triangle

func generate(numRows int) [][]int {
	var result [][]int

	if numRows == 0 {
		return result
	}

	for i := 0; i < numRows; i++ {
		var row []int
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				row = append(row, result[i-1][j-1]+result[i-1][j])
			}
		}
		result = append(result, row)
	}

	return result
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		input    int
		expected [][]int
	}{
		{
			input:    1,
			expected: [][]int{{1}},
		},
		{
			input:    5,
			expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			input:    3,
			expected: [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			input: 7,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
				{1, 6, 15, 20, 15, 6, 1},
			},
		},
		{
			input: 10,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
				{1, 6, 15, 20, 15, 6, 1},
				{1, 7, 21, 35, 35, 21, 7, 1},
				{1, 8, 28, 56, 70, 56, 28, 8, 1},
				{1, 9, 36, 84, 126, 126, 84, 36, 9, 1},
			},
		},
		{
			input:    2,
			expected: [][]int{{1}, {1, 1}},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := generate(test.input)
			if !equalSlices(result, test.expected) {
				t.Errorf("expected: %v, got: %v", test.expected, result)
			}
		})
	}
}

// equalSlices is a helper function to compare 2D slices
func equalSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
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
