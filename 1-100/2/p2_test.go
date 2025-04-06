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
	returnNode := new(ListNode)
	carry := 0 // used for carrying between loops

	for node := returnNode; l1 != nil || l2 != nil || carry != 0; node = node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		node.Next = &ListNode{Val: carry % 10}
		carry /= 10
	}

	return returnNode.Next
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		listNode1 *ListNode
		listNode2 *ListNode
		expected  *ListNode
	}{
		{
			listNode1: sliceToListNode([]int{2, 4, 3}),
			listNode2: sliceToListNode([]int{5, 6, 4}),
			expected:  sliceToListNode([]int{7, 0, 8}),
		},
		{
			listNode1: sliceToListNode([]int{0}),
			listNode2: sliceToListNode([]int{0}),
			expected:  sliceToListNode([]int{0}),
		},
		{
			listNode1: sliceToListNode([]int{9, 9, 9, 9, 9, 9, 9}),
			listNode2: sliceToListNode([]int{9, 9, 9, 9}),
			expected:  sliceToListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
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

// Helpers to reduce clutter from setups

func sliceToListNode(nums []int) *ListNode {
	var dummy = &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

func listNodeToSlice(node *ListNode) []int {
	var result []int
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}
