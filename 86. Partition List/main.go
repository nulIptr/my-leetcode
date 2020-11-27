package main

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

func partition(head *ListNode, x int) *ListNode {
	visualHead := ListNode{Next: head}
	cursor := &visualHead
	bound := &visualHead
	var prev *ListNode

	for cursor.Next != nil {
		prev = cursor
		cursor = cursor.Next
		if cursor.Val < x {
			if cursor != bound.Next {
				prev.Next = cursor.Next
				cursor.Next = bound.Next
				bound.Next = cursor
				cursor = prev
			}
			bound = bound.Next
		}
	}

	return visualHead.Next
}
func main() {

}
