package main

import "fmt"

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	resultHead := &ListNode{0, head}
	rh := resultHead
	var rt *ListNode = nil
	index := 1
	c := head

	for index <= n+1 {
		if index < m-1 {
			c = c.Next
		} else if index == m-1 {
			if m == 1 {
				rh = resultHead
			} else {
				rh = c
				c = c.Next
			}

		} else if index == m {
			rt = c
		} else if index <= n {
			tm := c.Next
			c.Next = rh.Next
			rh.Next = c
			c = tm
		} else if index == n+1 {
			tm := c.Next
			c.Next = rh.Next
			rh.Next = c
			c = tm
			rt.Next = tm
		}
		index++
	}

	return resultHead.Next

}
func main() {
	v := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	//v := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}

	r := reverseBetween(v, 3, 4)
	fmt.Println(r)
}
