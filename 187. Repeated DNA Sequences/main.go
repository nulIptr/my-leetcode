package main

import (
	"fmt"
)

func findRepeatedDnaSequences(s string) []string {
	r := make([]string, 0, 10)
	m := make(map[string]int)
	for i := 0; i < len(s)-9; i++ {
		c := s[i : i+10]
		if v, ok := m[c]; ok {
			if v == 0 {
				m[c] = 1
				r = append(r, c)
			}
		} else {
			m[c] = 0
		}
	}
	return r
}

func main() {

	fmt.Print(findRepeatedDnaSequences("AAAAAAAAAAA"))
}
