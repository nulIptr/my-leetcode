package main

import "container/list"

type innerInt struct {
	Value int
	Ptr   *list.Element
}

type LRUCache struct {
	inner map[int]innerInt
	l     *list.List

	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		inner:    make(map[int]innerInt, capacity),
		l:        list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.inner[key]
	if ok {
		this.l.MoveToFront(v.Ptr)
		return v.Value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.inner[key]
	if ok {
		this.l.MoveToFront(v.Ptr)
		v.Ptr.Value = key
		delete(this.inner, key)
		v.Value = value
		this.inner[key] = v
	} else {
		ptr := this.l.PushFront(key)
		this.inner[key] = innerInt{
			Value: value,
			Ptr:   ptr,
		}
	}

	if this.l.Len() > this.capacity {
		p := this.l.Back()
		this.l.Remove(p)
		delete(this.inner, p.Value.(int))
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {

}
