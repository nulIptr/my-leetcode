package main

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	arr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		arr[s[i]-97]++
	}

	for i := 0; i < len(s); i++ {
		if arr[s[i]-97] == 1 {
			return i
		}
	}
	return -1
}
func main() {

}
