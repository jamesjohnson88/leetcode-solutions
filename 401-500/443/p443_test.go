package p443

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

// https://leetcode.com/problems/string-compression

func compress(chars []byte) int {
	nChars := len(chars)
	writeIndex := 0
	i := 0

	for i < nChars {
		currentChar := chars[i]
		count := 0

		for i < nChars && chars[i] == currentChar {
			count++
			i++
		}

		if count == 1 {
			chars[writeIndex] = currentChar
			writeIndex++
		} else {
			chars[writeIndex] = currentChar
			writeIndex++

			for _, digit := range strconv.Itoa(count) {
				chars[writeIndex] = byte(digit)
				writeIndex++
			}
		}
	}
	return writeIndex
}

func TestCompress(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
	}{
		{
			input:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: 6,
		},
		{
			input:    []byte{'a'},
			expected: 1,
		},
		{
			input:    []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expected: 4,
		},
		{
			input:    []byte{'a', 'b', 'c'},
			expected: 3,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := compress(test.input)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("got %v, expected %v", result, test.expected)
			}
		})
	}
}
