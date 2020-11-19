package main

//差集		a & (~b)
func singleNumber(nums []int) int {
	t := 0
	r := 0
	for i := 0; i < len(nums); i++ {
		r = (r ^ nums[i]) & (^t)
		t = (t ^ nums[i]) & (^r)
	}
	return r
}
func main() {

}
