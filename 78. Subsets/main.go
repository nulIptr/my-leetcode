package main

import "fmt"

func subsets(nums []int) [][]int {
	var res [][]int
	var curr []int
	res = dfs(nums, curr, res, 0)
	return res
}

func dfs(nums []int, curr []int, res [][]int, start int) [][]int {
	currDup := make([]int, len(curr))
	copy(currDup, curr)
	res = append(res, currDup)

	for i := start; i < len(nums); i++ {
		curr = append(curr, nums[i])
		res = dfs(nums, curr, res, i+1)
		curr = curr[:len(curr)-1]
	}

	return res
}

func main() {

	fmt.Print(subsets([]int{1, 2, 3}))
}
