// Package linkedlist - utils for linkedlists
package linkedlist

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Insert - sets Next value of node with val
func Insert(val []int) *ListNode {
  dummy := new(ListNode)
  head := dummy
	for val := range val {
		head.Next = &ListNode{Val: val}
		head = head.Next
	}
  return dummy
}
