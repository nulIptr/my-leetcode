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

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	fast, slow := head, head
	// find mid
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow.Next

	//reverse
	slow.Next = nil
	tail := mid
	var preNode *ListNode = nil
	for tail != nil {
		n := tail.Next
		tail.Next = preNode
		preNode = tail
		tail = n
	}
	// merge
	for head != nil && preNode != nil {
		n := head.Next
		head.Next = preNode
		preNode = preNode.Next
		head.Next.Next = n
		head = n
	}
}
func main() {

}
