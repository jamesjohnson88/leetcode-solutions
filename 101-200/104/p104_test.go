package p104

import (
	"fmt"
	"reflect"
	"testing"
)

// https://leetcode.com/problems/maximum-depth-of-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		expected int
	}{
		{
			// [3,9,20,null,null,15,7]
			input: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: 3,
		},
		{
			// [1,null,2]
			input: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			expected: 2,
		},
		{
			// empty tree
			input:    nil,
			expected: 0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := maxDepth(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("got %v, expected %v", result, test.expected)
			}
		})
	}
}
