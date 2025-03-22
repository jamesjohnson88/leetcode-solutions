package p735

import (
	"fmt"
	"reflect"
	"testing"
)

// https://leetcode.com/problems/asteroid-collision

func asteroidCollision(asteroids []int) []int {
	// optional, but saves float conversion
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var stack []int
	for _, a := range asteroids {
		for len(stack) > 0 && (stack[len(stack)-1] >= 0 && a < 0) {
			top := abs(stack[len(stack)-1])
			curr := abs(a)

			if curr > top {
				stack = stack[:len(stack)-1]
			} else if curr < top {
				a = 0
			} else {
				stack = stack[:len(stack)-1]
				a = 0
			}
		}

		if a != 0 {
			stack = append(stack, a)
		}
	}

	return stack
}

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{5, 10, -5},
			expected: []int{5, 10},
		},
		{
			input:    []int{8, -8},
			expected: []int{},
		},
		{
			input:    []int{10, 2, -5},
			expected: []int{10},
		},
		{
			input:    []int{-2, -1, 1, 2},
			expected: []int{-2, -1, 1, 2},
		},
		{
			input:    []int{-2, -2, 1, -2},
			expected: []int{-2, -2, -2},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			if got := asteroidCollision(test.input); !reflect.DeepEqual(got, test.expected) {
				t.Errorf("got: %v, expected: %v", got, test.expected)
			}
		})
	}
}
