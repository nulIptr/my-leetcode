package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, err := m[nums[i]]; err {
			return true
		}
		m[nums[i]] = true
	}
	return false
}
func main() {

}
