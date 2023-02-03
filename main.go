// Main - solutions to some leetcode problems with Go.
package main

import (
	"fmt"
	"learning-go/linkedlist"
	"learning-go/utils"
	"sort"
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
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	seen := make(map[rune]int, n)
	record, start := 0, 0

	for i, c := range s {
		if _, ok := seen[c]; ok == true && seen[c] >= start {
			start = seen[c] + 1
		} else {
			record = utils.Max(record, i-start+1)
		}
		seen[c] = i
	}
	return record
}

// https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {
	n := len(height)
	record, l, r := 0, 0, n-1

	for l < r {
		minH := utils.Min(height[l], height[r])
		width := r - l
		area := minH * width
		record = utils.Max(record, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return record
}

// https://leetcode.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	record, n := "", len(s)
	for i := range s {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			txt := s[l : r+1]
			if len(txt) > len(record) {
				record = txt
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			txt := s[l : r+1]
			if len(txt) > len(record) {
				record = txt
			}
			l--
			r++
		}
	}
	return record
}

// https://leetcode.com/problems/3sum/
func threeSum(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := range nums {
    if i > 0 && nums[i] == nums[i-1] {
      continue
    }
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans := []int{nums[i], nums[l], nums[r]}
				res = append(res, ans)
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			}
		}
	}
	return res
}

func main() {
	arg := []int{-1, 0, 1, 2, -1, -4}
	// arg := []int{0,1,1}
	// arg := []int{0,0,0}
	// res := threeSum(arg)
	res := threeSum(arg)
	fmt.Println(res)
}
