// package main
package main

import "fmt"

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

func main() {
	arr := []int{3, 4, 4}
	target := 7
	res := twoSum(arr, target)
	fmt.Println(res)
}
