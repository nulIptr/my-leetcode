package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	l := len(nums) - 1
	res := make([][]int, 0, 1024)
	sort.Ints(nums)
	for i := 0; i < l; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 {
			break
		}

		begin := i + 1
		end := l
		for begin < end {
			if nums[begin]+nums[end]+nums[i] == 0 {

				res = append(res, []int{nums[i], nums[begin], nums[end]})
				begin++
				for nums[begin] == nums[begin-1] && begin < end {
					begin++
				}
				end--
				for nums[end] == nums[end+1] && begin < end {
					end--
				}
			} else if nums[begin]+nums[end]+nums[i] < 0 {
				begin++
			} else if nums[begin]+nums[end]+nums[i] > 0 {
				end--
			}
		}

	}
	return res
}

func main() {
	v := [...]int{-1, 0, 1, 2, -1, -4}
	//v := [...]int{0, 0, 0, 0}
	r := threeSum(v[:])
	fmt.Println(r)
}
