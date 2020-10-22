package main

import "fmt"

func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	begin := 0
	end := len(nums)
	cs := 0
	res := 0
	for begin < end {
		cs += nums[begin]
		res += m[cs-k]
		m[cs]++
		begin++
	}
	return res
}

func main() {
	v := [...]int{1, 1, 1}
	//v := [...]int{0, 0, 0, 0}
	r := subarraySum(v[:], 2)
	fmt.Println(r)
}
