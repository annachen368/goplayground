package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})

	for _, num := range nums {
		_, exists := m[num]

		if exists {
			return true
		}

		m[num] = struct{}{}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}
