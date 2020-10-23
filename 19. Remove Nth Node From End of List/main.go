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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var ln *ListNode
	cur := head
	index := 1
	for cur.Next != nil {
		if index < n {
			cur = cur.Next
			index++
		} else {
			if ln == nil {
				ln = head
			} else {
				ln = ln.Next
			}

			cur = cur.Next
		}
	}

	if ln == nil {
		return head.Next
	} else {
		ln.Next = ln.Next.Next
		return head
	}

}
func main() {

	v := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	//v := [...]int{0, 0, 0, 0}
	r := removeNthFromEnd(v, 2)
	fmt.Println(r)
}
