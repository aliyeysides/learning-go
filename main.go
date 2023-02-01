package main

import (
	"fmt"
	"learning-go/linkedlist"
)

func twoSum(arr []int, target int) []int {
	seen := make(map[int]int)

	for i, val := range arr {
		remainder := target - val

		if k, ok := seen[remainder]; ok {
			return []int{k, i}
		}

		seen[val] = i
	}
	return nil
}

func addTwoNumbers(l1 *linkedlist.ListNode, l2 *linkedlist.ListNode) *linkedlist.ListNode {
	carry, dummy := 0, new(linkedlist.ListNode)

	for node := dummy; l1 != nil || l2 != nil || carry > 0; node = node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		node.Next = &linkedlist.ListNode{Val: carry % 10}
		carry /= 10
	}
	return dummy.Next
}

func main() {
	l1 := new(linkedlist.ListNode)
	l2 := new(linkedlist.ListNode)

	linkedlist.Insert(l1, 2)
	linkedlist.Insert(l1, 4)
	linkedlist.Insert(l1, 3)

	linkedlist.Insert(l2, 5)
	linkedlist.Insert(l2, 6)
	linkedlist.Insert(l2, 4)

	res := addTwoNumbers(l1, l2)
  
  for node := res; res != nil; node = node.Next {
    fmt.Println(node.Val)
  }
}
