package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	r := 0
	for i := 0; i < len(citations); i++ {
		if citations[i] >= i+1 {
			r++
		} else {
			return i
		}
	}

	return r

}
func main() {
	fmt.Println(hIndex([]int{0, 0, 0}))
	fmt.Println(hIndex([]int{1}))
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}
