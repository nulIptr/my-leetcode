package main

func twoSum(numbers []int, target int) []int {

	begin := 0
	end := len(numbers) - 1
	for begin < end {
		if numbers[begin]+numbers[end] == target {
			return []int{begin + 1, end + 1}
		} else if numbers[begin]+numbers[end] < target {
			begin++
		} else if numbers[begin]+numbers[end] > target {
			end--
		}
	}
	return nil
}
func main() {

}
