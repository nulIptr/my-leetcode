package main

func plusOne(digits []int) []int {
	for index := len(digits) - 1; index >= 0; index-- {
		if digits[index] == 9 {
			digits[index] = 0
		} else {
			digits[index] += 1
			return digits
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}

func main() {

}
