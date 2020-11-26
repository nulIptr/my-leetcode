package main

import (
	"fmt"
	"sort"
	"strings"
)

type MySlice []string

func (receiver MySlice) Len() int {
	return sort.StringSlice(receiver).Len()
}

func (receiver MySlice) Swap(i, j int) {
	sort.StringSlice(receiver).Swap(i, j)
}

func (receiver MySlice) Less(i, j int) bool {
	return !less(receiver[i], receiver[j])
}

func less(s1, s2 string) bool {
	return s1+s2 < s2+s1
}

func largestNumber(nums []int) string {

	allZero := true
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			allZero = false
		}
	}
	if allZero {
		return "0"
	}

	sz := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		sz[i] = fmt.Sprintf("%d", nums[i])
	}
	sort.Sort(MySlice(sz))

	return strings.Join(sz, "")

}
func main() {

	//fmt.Print(largestNumber([]int{3,30,34,5,9}))
	fmt.Print("\n")
	fmt.Print(largestNumber([]int{432, 43243}))
}
