package leetcode

// Given a linked list, determine if it has a cycle in it.
// Follow up: Can you solve it without using extra space?
func hasCycle(head *ListNode) bool {
	for one, two := head, head; two != nil; {
		one = one.Next
		two = two.Next
		if two != nil {
			two = two.Next
		}
		if one != nil && one == two {
			return true
		}
	}
	return false
}
