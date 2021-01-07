package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]bool)
	for i := 0; i < k && i < len(nums); i++ {
		if _, err := m[nums[i]]; err {
			return true
		}
		m[nums[i]] = true
	}

	for i := k; i < len(nums); i++ {
		if _, err := m[nums[i]]; err {
			return true
		}
		m[nums[i]] = true
		delete(m, nums[i-k])
	}
	return false
}
func main() {
	fmt.Println(containsNearbyDuplicate([]int{99, 99}, 2))
}
