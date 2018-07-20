package tests

import "testing"

func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	for b := a; b != nil; b = b.Next {
		t.Log(b.Val)
	}

	a = deleteDuplicates(a)
	for b := a; b != nil; b = b.Next {
		t.Log(b.Val)
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p := head
	for p.Next != nil {
		if p.Next.Val == p.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return head
}
