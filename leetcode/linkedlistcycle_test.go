package leetcode

import "testing"

func TestHasCycle(t *testing.T) {
	cycleList1 := &ListNode{0, nil}
	cycleList1.Next = cycleList1

	cycleNode := &ListNode{-1, nil}
	cycleList2 := &ListNode{0, &ListNode{1, &ListNode{2, cycleNode}}}
	cycleNode.Next = &ListNode{3, &ListNode{4, cycleNode}}
	testCases := []struct {
		head     *ListNode
		hasCycle bool
	}{
		{
			nil,
			false,
		},
		{
			&ListNode{0, nil},
			false,
		},
		{
			&ListNode{0, &ListNode{1, nil}},
			false,
		},
		{
			cycleList1,
			true,
		},
		{
			cycleList2,
			true,
		},
	}
	for i, testCase := range testCases {
		got := hasCycle(testCase.head)
		if got != testCase.hasCycle {
			t.Errorf("#%d got %t, expected %t ", i, got, testCase.hasCycle)
		}
	}

}
