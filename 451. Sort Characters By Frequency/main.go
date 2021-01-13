package main

import (
	"fmt"
	"sort"
)

type Obj struct {
	char  byte
	count int
}

type Slice []Obj

func (o Slice) Len() int {
	return len(o)
}

func (o Slice) Less(i, j int) bool {
	return o[i].count > o[j].count
}

func (o Slice) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func frequencySort(s string) string {

	sz := make([]Obj, 128, 128)
	for i := byte(0); i < 128; i++ {
		sz[i].char = i
	}

	for i := 0; i < len(s); i++ {
		sz[s[i]].count++
	}

	sort.Sort(Slice(sz))

	r := make([]byte, 0, len(s))

	for i := 0; i < 128; i++ {
		if sz[i].count == 0 {
			break
		}
		for j := 0; j < sz[i].count; j++ {
			r = append(r, sz[i].char)
		}
	}

	return string(r)
}
func main() {
	fmt.Print(frequencySort("acabca"))
}
