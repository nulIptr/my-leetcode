package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {

	curr := 1
	ch := chars[0]
	count := 0
	for i := 0; i < len(chars); i++ {
		if ch == chars[i] {
			count++
		} else {
			if count > 1 {
				str := strconv.Itoa(count)
				for j := 0; j < len(str); j++ {
					chars[curr+j] = str[j]

				}
				curr += len(str)

			}
			chars[curr] = chars[i]
			curr++
			ch = chars[i]
			count = 1
		}
	}

	if count > 1 {
		str := strconv.Itoa(count)
		for j := 0; j < len(str); j++ {
			chars[curr+j] = str[j]

		}
		curr += len(str)

	}
	return curr
}
func main() {
	fmt.Print(compress([]byte{'a', 'b', 'b', 'b', 'b'}))
}
