package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	sz := strings.Fields(s)
	for i := 0; i < len(sz)/2; i++ {
		j := len(sz) - i - 1
		sz[i], sz[j] = sz[j], sz[i]
	}
	r := strings.Join(sz, " ")
	return r
}

func main() {

	fmt.Print(reverseWords("  Bob    Loves  Alice   "))
}
