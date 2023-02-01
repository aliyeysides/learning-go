// package main
package main

import "fmt"

func twoSum(arr []int, target int) []int {
	seen := make(map[int]int)

	for i, val := range arr {
		remainder := target - val

		if _, ok := seen[remainder]; ok {
			return []int{seen[remainder], i}
		}

		seen[val] = i
	}
	return nil
}

func main() {
	arr := []int{3, 4, 4}
	target := 7
	res := twoSum(arr, target)
	fmt.Println(res)
}
