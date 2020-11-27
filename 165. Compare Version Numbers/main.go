package main

import (
	"strconv"
)

func compareVersion(version1 string, version2 string) int {
	l1 := len(version1)
	l2 := len(version2)
	if l1 == 0 && l2 == 0 {
		return 0
	}
	if l1 == 0 {
		return compareVersion("0", version2)
	}
	if l2 == 0 {
		return compareVersion(version1, "0")
	}
	tk1, version1 := next(version1)
	tk2, version2 := next(version2)
	num1, _ := strconv.Atoi(tk1)
	num2, _ := strconv.Atoi(tk2)
	if num1 > num2 {
		return 1
	} else if num1 < num2 {
		return -1
	} else {
		return compareVersion(version1, version2)
	}
}

func next(s string) (string, string) {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return s[:i], s[i+1:]
		}
	}
	return s, ""
}
func main() {

}
