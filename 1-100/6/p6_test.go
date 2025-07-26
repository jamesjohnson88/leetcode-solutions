package p6

import (
	"fmt"
	"slices"
	"testing"
)

// https://leetcode.com/problems/zigzag-conversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	arr := make([][]byte, numRows)
	row := 1
	dir := 1

	for _, char := range s {
		arr[row-1] = append(arr[row-1], byte(char))

		if (row == numRows && dir == 1) || (row == 1 && dir == -1) {
			dir = -dir
		}
		row += dir
	}

	return string(slices.Concat(arr...))
}

func TestConvert(t *testing.T) {
	tests := []struct {
		input    string
		numRows  int
		expected string
	}{
		{
			input:    "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			input:    "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		{
			input:    "A",
			numRows:  1,
			expected: "A",
		},
		{
			input:    "AB",
			numRows:  1,
			expected: "AB",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := convert(test.input, test.numRows)
			if result != test.expected {
				t.Errorf("expected: %s, got: %s", test.expected, result)
			}
		})
	}
}
