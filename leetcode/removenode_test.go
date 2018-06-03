package leetcode

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []struct {
		head    *ListNode
		nth     int
		newHead *ListNode
	}{
		{
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			2,
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}},
		},
		{
			&ListNode{1, nil},
			1,
			nil,
		},
		{
			&ListNode{1, &ListNode{2, nil}},
			2,
			&ListNode{2, nil},
		},
		{
			&ListNode{1, &ListNode{2, nil}},
			1,
			&ListNode{1, nil},
		},
	}
	for i, tc := range testCases {
		hs := tc.head.String()
		got := removeNthFromEnd(tc.head, tc.nth)
		if got.String() != tc.newHead.String() {
			t.Errorf("#%d list=%s, nth=%d, got %s, expected %s ", i, hs, tc.nth, got, tc.newHead)
		}
	}
}
