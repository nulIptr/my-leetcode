package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	dic := make([]int, 128)

	for i := 0; i < len(t); i++ {
		dic[t[i]]++
	}

	left := 0

	minlen := math.MaxInt32
	minleft := 0
	cnt := 0

	for i := 0; i < len(s); i++ {
		if dic[s[i]]--; dic[s[i]] >= 0 {
			cnt++

		}
		for cnt == len(t) {
			if i-left+1 < minlen {
				minlen = i - left + 1
				minleft = left
			}
			if dic[s[left]]++; dic[s[left]] > 0 {
				cnt--
			}
			left++
		}
	}
	return s[minleft : minleft+minlen]
}
func main() {

	r := minWindow("DEBANC", "ABC")
	fmt.Println(r)
}
