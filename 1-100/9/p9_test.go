package p9

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/palindrome-number

func isPalindrome(x int) bool {
	// negative numbers can never be a palindrome
	// neither can 0-leading numbers, unless 0 itself
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	r := 0
	for x > r {
		// shift r left, insert rightmost value of x
		r = r*10 + x%10
		// remove rightmost value of x
		x /= 10
	}

	return x == r || x == r/10
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{2112, true},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := isPalindrome(test.input)
			if result != test.expected {
				t.Errorf("expected: %v, got: %v", test.expected, result)
			}
		})
	}
}
