package p8

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

// https://leetcode.com/problems/string-to-integer-atoi

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return 0
	}

	negative := false
	start := 0

	if s[0] == '-' {
		negative = true
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	result := 0
	for i := start; i < len(s); i++ {
		char := rune(s[i])
		if char >= '0' && char <= '9' {
			digit := int(char - '0')
			if result > (math.MaxInt32-digit)/10 {
				if negative {
					return math.MinInt32
				}
				return math.MaxInt32
			}
			result = result*10 + digit
		} else {
			break
		}
	}

	if negative {
		result = -result
	}
	return result
}

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"42", 42},
		{" -042", -42},
		{"1337c0d3", 1337},
		{"0-1", 0},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"+1", 1},
		{"9223372036854775808", 2147483647},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := myAtoi(test.input)
			if result != test.expected {
				t.Errorf("expected: %d, got: %d", test.expected, result)
			}
		})
	}
}
