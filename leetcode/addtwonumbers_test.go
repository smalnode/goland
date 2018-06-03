package leetcode

import "testing"
import "fmt"

func (l *ListNode) String() string {
	var s string
	for e := l; e != nil; e = e.Next {
		if e == l {
			s = fmt.Sprintf("%d", e.Val)
		} else {
			s = fmt.Sprintf("%s->%d", s, e.Val)
		}
	}
	return s
}

func TestAddTwoNumbers(t *testing.T) {
	testcases := []struct {
		l1  *ListNode
		l2  *ListNode
		res *ListNode
	}{
		{
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			&ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{
			&ListNode{0, nil},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
		},
		{
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{0, nil},
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		},
		{
			&ListNode{5, &ListNode{2, nil}},
			&ListNode{5, &ListNode{7, nil}},
			&ListNode{0, &ListNode{0, &ListNode{1, nil}}},
		},
	}

	for i, c := range testcases {
		got := addTwoNumbers(c.l1, c.l2).String()
		expected := c.res.String()
		if expected != got {
			t.Errorf("#%d %s + %s expected %s, got %s ", i, c.l1, c.l2, expected, got)
		} else {
			t.Log(got)
		}
	}
}
