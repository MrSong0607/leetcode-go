package tests

import "testing"

func TestMergeTwoSortedLists(t *testing.T) {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	b := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	c := mergeTwoLists(a, b)
	for c != nil {
		t.Log(c.Val)
		c = c.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var l3, q *ListNode

	for l1 != nil && l2 != nil {
		var p *ListNode

		if l1.Val > l2.Val {
			p = l2
			l2 = l2.Next
		} else {
			p = l1
			l1 = l1.Next
		}

		if l3 == nil {
			l3 = p
			q = l3
		} else {
			l3.Next = p
			l3 = l3.Next
		}
	}

	if l1 != nil {
		l3.Next = l1
	} else {
		l3.Next = l2
	}

	return q
}
