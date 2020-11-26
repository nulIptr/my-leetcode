package main

import "math/rand"

type Solution struct {
	list []int
}

func Constructor(nums []int) Solution {
	return Solution{list: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.list
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	a := make([]int, len(this.list))
	copy(a, this.list)
	l := len(a)

	for i := 0; i < l; i++ {
		j := rand.Intn(l - i)
		a[i], a[j+i] = a[j+i], a[i]
	}

	return a
}

func main() {

}
