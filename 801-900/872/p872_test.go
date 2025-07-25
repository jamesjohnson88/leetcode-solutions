package p872

import (
	"fmt"
	"slices"
	"testing"
)

// https://leetcode.com/problems/leaf-similar-trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return slices.Equal(findLeaves(root1), findLeaves(root2))
}

func findLeaves(t *TreeNode) []int {
	if t == nil {
		return []int{}
	}

	// leaf node
	if t.Left == nil && t.Right == nil {
		return []int{t.Val}
	}

	return append(findLeaves(t.Left), findLeaves(t.Right)...)
}

func TestLeafSimilar(t *testing.T) {
	tests := []struct {
		root1    *TreeNode
		root2    *TreeNode
		expected bool
	}{
		{
			// Example 1: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
			// Both have leaf sequence [6,7,4,9,8]
			root1: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			root2: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 9,
						},
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
			expected: true,
		},
		{
			// Example 2: root1 = [1,2,3], root2 = [1,3,2]
			// root1 has leaf sequence [2,3], root2 has leaf sequence [3,2]
			root1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			root2: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			actual := leafSimilar(test.root1, test.root2)
			if actual != test.expected {
				t.Errorf("Test case %d: actual: %v, expected: %v", i+1, actual, test.expected)
			}
		})
	}
}
