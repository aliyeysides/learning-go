// Package linkedlist - utils for linkedlists
package linkedlist

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// InsertList - inserts values of a []int into LinkedList
func InsertList(arr []int) *ListNode {
  n := len(arr)
  dummy := new(ListNode)
  for i, node := 0, dummy; i < n; node = node.Next {
    node.Next = &ListNode{arr[i], node}
    i++
  }
  return dummy.Next
}
