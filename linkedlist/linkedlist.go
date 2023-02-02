// Package linkedlist - utils for linkedlists
package linkedlist

// ListNode - Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

// Insert - sets Next value of node with val
func Insert(node *ListNode, val int) *ListNode {
  node.Next = &ListNode{Val: val}
  return node.Next
}

