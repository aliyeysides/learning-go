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

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func removeHelper(head *ListNode, n int) (*ListNode, int) {
	if head == nil {
		return head, 0
	}
	i := 0
	head.Next, i = removeHelper(head.Next, n)
	if i+1 == n {
		return head.Next, i + 1
	}
	return head, i + 1
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res, _ := removeHelper(head, n)
	return res
}

// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	rules := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := make([]rune, 0)
	for _, char := range s {
		if _, ok := rules[char]; ok {
			stack = append(stack, char)
		} else if len(stack) == 0 || rules[stack[len(stack)-1]] != char {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	node := dummy
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			node.Next = list2
			list2 = list2.Next
		} else {
			node.Next = list1
			list1 = list1.Next
		}
		node = node.Next
	}
	if list1 != nil {
		node.Next = list1
	} else {
		node.Next = list2
	}
	return dummy.Next
}

// https://leetcode.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	mid := n / 2
	l, r := mergeKLists(lists[:mid]), mergeKLists(lists[mid:])
	return mergeTwoLists(l, r)
}

// https://leetcode.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		} else if nums[mid] <= nums[r] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	return -1
}

// https://leetcode.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	result := new([][]int)

	var dfs func(start int, remaining int, path []int)

	dfs = func(i int, remaining int, path []int) {
		if remaining == 0 {
			*result = append(*result, append(make([]int, 0), path...))
			return
		}

		for j := i; j < len(candidates); j++ {
			if remaining-candidates[j] >= 0 {
				dfs(j, remaining-candidates[j], append(path, candidates[j]))
			}
		}
		return
	}

	dfs(0, target, nil)
	return *result
}

func main() {
	data := []int{7, 3, 2}
	target := 18
	fmt.Println(combinationSum(data, target))
}
