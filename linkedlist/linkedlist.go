package linkedlist

// ListNode - Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func Insert(node *ListNode, val int) *ListNode {
  node.Next = &ListNode{Val: val}
  return node
}
