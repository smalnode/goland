package leetcode

// Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes
// of the first two lists.
// Example: Input: 1->2->4, 1->3->4
//         Output: 1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}
