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

func hasCycle(head *ListNode) bool {

	m := make(map[*ListNode]bool)

	for true {
		if head != nil {
			if _, err := m[head]; !err {
				m[head] = true
				head = head.Next
				continue
			}
			return true
		}
		return false
	}
	return false
}
func main() {
	hasCycle(&ListNode{
		Val:  1,
		Next: nil,
	})
}
