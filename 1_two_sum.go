package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // store num -> index

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		j, exists := m[diff]

		if exists {
			return []int{j, i}
		}

		m[nums[i]] = i
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
