package main

func productExceptSelf(nums []int) []int {

	zeroCount := 0

	var t = 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
			if zeroCount > 1 {
				break
			}
		} else {
			t *= nums[i]
		}
	}

	if zeroCount == 0 {
		for i := 0; i < len(nums); i++ {
			tmp := t / nums[i]
			nums[i] = tmp
		}
	} else if zeroCount == 1 {
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				nums[i] = t
			} else {
				nums[i] = 0
			}
		}
	} else {
		for i := 0; i < len(nums); i++ {
			nums[i] = 0
		}
	}

	return nums
}
func main() {

}
