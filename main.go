// Main - solutions to some leetcode problems with Go.
package main

import (
	"fmt"
	"learning-go/linkedlist"
)

// https://leetcode.com/problems/two-sum/
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

// ListNode - type alias for linkedlist.ListNode
type ListNode = linkedlist.ListNode

// https://leetcode.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, dummy := 0, new(ListNode)

	for node := dummy; l1 != nil || l2 != nil || carry > 0; node = node.Next {
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
	return dummy.Next
}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// func lengthOfLongestSubstring(s string) int {
//   // todo 
// }

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)

	linkedlist.Insert(l1, 2)
	linkedlist.Insert(l1, 4)
	linkedlist.Insert(l1, 3)

	linkedlist.Insert(l2, 5)
	linkedlist.Insert(l2, 6)
	linkedlist.Insert(l2, 4)

	res := addTwoNumbers(l1, l2)

	for node := res; node.Next != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}
