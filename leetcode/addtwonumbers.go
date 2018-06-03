package leetcode

// You are given two non-empty linked lists representing
// tow non-negative integers.
// The digits are store in reverse order and each of their
// nodes contain a single digit. Add the tow numbers and
// return it as a linked list.
// You may assume the two numbers do not contain any leading
// zero, except the number 0 itself.

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	e1 := l1
	e2 := l2
	carryover := 0
	zero := &ListNode{
		0,
		nil,
	}
	var res, last *ListNode
	for e1 != zero || e2 != zero {
		v := e1.Val + e2.Val + carryover
		carryover = v / 10
		v = v % 10
		n := &ListNode{
			v,
			nil,
		}
		if last == nil {
			res, last = n, n
		} else {
			last.Next, last = n, n
		}
		if e1.Next == nil {
			e1 = zero
		} else {
			e1 = e1.Next
		}
		if e2.Next == nil {
			e2 = zero
		} else {
			e2 = e2.Next
		}
	}
	if carryover > 0 {
		last.Next = &ListNode{
			carryover,
			nil,
		}
		last = last.Next
	}
	return res
}
