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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) > 2 {
		pre := mergeKLists(lists[:len(lists)/2])
		next := mergeKLists(lists[len(lists)/2:])

		return mergeKLists([]*ListNode{pre, next})
	} else if len(lists) == 2 {
		var r *ListNode
		if lists[0] != nil && lists[1] != nil {
			if lists[0].Val < lists[1].Val {
				r = lists[0]
				lists[0] = lists[0].Next
			} else {
				r = lists[1]
				lists[1] = lists[1].Next
			}
		} else if lists[0] != nil {
			r = lists[0]
			lists[0] = lists[0].Next
		} else if lists[1] != nil {
			r = lists[1]
			lists[1] = lists[1].Next
		}

		head := r
		for lists[0] != nil || lists[1] != nil {
			if lists[0] != nil && lists[1] != nil {
				if lists[0].Val < lists[1].Val {
					r.Next = lists[0]
					r = r.Next
					lists[0] = lists[0].Next
				} else {
					r.Next = lists[1]
					r = r.Next
					lists[1] = lists[1].Next
				}
			} else if lists[0] != nil {
				r.Next = lists[0]
				break
			} else if lists[1] != nil {
				r.Next = lists[1]
				break
			}
		}
		return head
	} else if len(lists) == 1 {
		return lists[0]
	} else {
		return nil
	}

}

func main() {
	v := []*ListNode{{1, &ListNode{4, &ListNode{5, nil}}}, {1, &ListNode{3, &ListNode{4, nil}}}, {2, &ListNode{6, nil}}}

	r := mergeKLists(v)
	fmt.Println(r)
}
