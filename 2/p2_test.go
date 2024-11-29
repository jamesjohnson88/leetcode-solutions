package p2

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/add-two-numbers/description/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  7,
				Next: nil,
			},
		},
	}
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		listNode1 *ListNode
		listNode2 *ListNode
		expected  *ListNode
	}{
		{
			listNode1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			listNode2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			expected: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run("Testing addTwoNumbers", func(t *testing.T) {
			got := addTwoNumbers(test.listNode1, test.listNode2)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf(
					"addTwoNumbers(%v, %v) returned %v; expected %v",
					listNodeToSlice(test.listNode1),
					listNodeToSlice(test.listNode2),
					listNodeToSlice(got),
					listNodeToSlice(test.expected),
				)
			}
		})
	}
}

func listNodeToSlice(node *ListNode) []int {
	var result []int
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}
