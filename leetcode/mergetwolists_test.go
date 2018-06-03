package leetcode

import "testing"

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		l1  *ListNode
		l2  *ListNode
		res *ListNode
	}{
		{
			nil,
			nil,
			nil,
		},
		{
			&ListNode{1, &ListNode{2, nil}},
			nil,
			&ListNode{1, &ListNode{2, nil}},
		},
		{
			nil,
			&ListNode{1, &ListNode{2, nil}},
			&ListNode{1, &ListNode{2, nil}},
		},
		{
			&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
	}
	for i, tc := range testCases {
		l1s := tc.l1.String()
		l2s := tc.l2.String()
		got := mergeTwoLists(tc.l1, tc.l2)
		if got.String() != tc.res.String() {
			t.Errorf("#%d l1=%s, l2=%s, got %s, expected %s ", i, l1s, l2s, got, tc.res)
		}
	}
}
