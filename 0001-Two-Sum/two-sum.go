package main

import "fmt"

func twoSum(nums []int, target int) []int {
	dic := map[int]int{}
	for i, num := range nums {
		dif := target - num
		if v, ok := dic[dif]; ok {
			return []int{v, i}
		}
		dic[nums[i]] = i
	}
	return nil
}
func main() {
	nums := []int{1, 2, 3}
	target := 3
	fmt.Println(twoSum(nums, target))
}
