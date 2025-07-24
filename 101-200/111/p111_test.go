package p111

import (
	"fmt"
	"reflect"
	"testing"
)

// https://leetcode.com/problems/minimum-depth-of-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if l == 0 || r == 0 {
		// both invalid (zero) -> leaf node (count as 1)
		// only one is invalid -> ignore it by taking max
		return max(l, r) + 1
	}
	// both valid, take smallest
	return min(l, r) + 1
}

func TestMinDepth(t *testing.T) {
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
			expected: 2,
		},
		{
			// [2,null,3,null,4,null,5,null,6]
			input: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val:   5,
							Right: &TreeNode{Val: 6},
						},
					},
				},
			},
			expected: 5,
		},
		{
			// empty tree
			input:    nil,
			expected: 0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := minDepth(test.input)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("got %v, expected %v", result, test.expected)
			}
		})
	}
}
