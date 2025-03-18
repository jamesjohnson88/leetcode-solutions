package p1207

import (
	"testing"
)

// https://leetcode.com/problems/unique-number-of-occurrences

// first version, more readable, not fancy
func uniqueOccurrences1(arr []int) bool {
	occurrences := map[int]int{}
	for _, i := range arr {
		occurrences[i]++
	}

	unique := map[int]bool{}
	for _, j := range occurrences {
		unique[j] = true
	}

	return len(occurrences) == len(unique)
}

// optimized version, less readable but more efficient
func uniqueOccurrences2(arr []int) bool {
	occurrences := make(map[int]int)
	unique := make(map[int]struct{})

	for _, num := range arr {
		occurrences[num]++
	}

	for _, freq := range occurrences {
		if _, exists := unique[freq]; exists {
			return false
		}
		unique[freq] = struct{}{}
	}

	return true
}

func TestUniqueOccurrences1(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 2, 2, 1, 1, 3},
			expected: true,
		},
		{
			input:    []int{1, 2},
			expected: false,
		},
		{
			input:    []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			expected: true,
		},
	}

	for _, test := range tests {
		actual := uniqueOccurrences1(test.input)
		if actual != test.expected {
			t.Errorf("input: %v, actual: %v, expected: %v", test.input, actual, test.expected)
		}
	}
}

func TestUniqueOccurrences2(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 2, 2, 1, 1, 3},
			expected: true,
		},
		{
			input:    []int{1, 2},
			expected: false,
		},
		{
			input:    []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			expected: true,
		},
	}

	for _, test := range tests {
		actual := uniqueOccurrences2(test.input)
		if actual != test.expected {
			t.Errorf("input: %v, actual: %v, expected: %v", test.input, actual, test.expected)
		}
	}
}
