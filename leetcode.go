// package main
package main

import (
  "github.com/aliyeysides/learning-go/vendor/linkedlist"
  "fmt"
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

func main() {
	// arr := []int{3, 4, 4}
	// target := 7
	// res := twoSum(arr, target)
  l1 := new(ListNode)
  l2 := new(ListNode)

  Insert(l1, 2)
  Insert(l1, 4)
  Insert(l1, 3)

  Insert(l2, 5)
  Insert(l2, 6)
  Insert(l2, 4)

  res = addTwoNumbers(l1, l2)
	fmt.Println(res)
}
