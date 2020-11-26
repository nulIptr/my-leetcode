package main

func maxArea(height []int) int {
	result := 0
	l := 0
	r := len(height) - 1
	for l < r {
		result = max(result, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a

}
func main() {

}
