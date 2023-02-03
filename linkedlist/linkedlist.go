// Package linkedlist - utils for linkedlists
package linkedlist

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Insert - sets Next value of node with val
func Insert(val []int) *ListNode {
  dummy := &ListNode{Val: 0, Next: nil }
  head := dummy
	for v := range val {
		head.Next = &ListNode{Val: v}
		head = head.Next
	}
  return dummy.Next
}
