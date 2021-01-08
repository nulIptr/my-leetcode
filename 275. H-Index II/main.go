package main

func hIndex(citations []int) int {

	r := 0
	for i := 0; i < len(citations); i++ {
		if citations[len(citations)-1-i] >= i+1 {
			r++
		} else {
			return r
		}
	}

	return r

}
func main() {

}
