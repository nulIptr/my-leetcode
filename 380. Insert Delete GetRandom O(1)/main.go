package main

import "math/rand"

type RandomizedSet struct {
	l []int
	m map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {

	return RandomizedSet{m: make(map[int]int), l: make([]int, 0, 3)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {

	if _, ok := this.m[val]; ok {
		return false
	} else {

		this.l = append(this.l, val)
		this.m[val] = len(this.l) - 1
		return true
	}

}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.m[val]; ok {
		delete(this.m, val)

		if index == len(this.l)-1 {
			this.l = this.l[0:index]
		} else {
			this.m[this.l[len(this.l)-1]] = index
			this.l[index] = this.l[len(this.l)-1]
			this.l = this.l[0 : len(this.l)-1]
		}

		return true
	} else {
		return false
	}

}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {

	return this.l[rand.Int()%len(this.m)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	obj := Constructor()
	_ = obj.Insert(1)
	_ = obj.Insert(2)
	_ = obj.Remove(1)
	_ = obj.GetRandom()
}
