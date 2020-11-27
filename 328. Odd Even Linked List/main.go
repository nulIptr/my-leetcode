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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oHead := head
	eHead := head.Next

	r := eHead.Next
	oHead.Next = nil
	eHead.Next = nil

	o := oHead
	e := eHead

	for i := 0; r != nil; i++ {
		if i%2 == 0 {
			o.Next = r
			r = r.Next
			o = o.Next
			o.Next = nil
		} else {
			e.Next = r
			r = r.Next
			e = e.Next
			e.Next = nil
		}
	}

	o.Next = eHead

	return oHead
}
func main() {

}
