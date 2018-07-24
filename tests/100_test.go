package tests

import (
	"strconv"
	"testing"
)

func TestSameTree(t *testing.T) {
	p, q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	t.Log(isSameTree(p, q))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	a, b := []string{}, []string{}
	a = genFrontScanTree(p, a)
	b = genFrontScanTree(q, b)

	if len(a) != len(b) {
		return false
	}

	for index := range a {
		if a[index] != b[index] {
			return false
		}
	}
	return true
}

func genFrontScanTree(p *TreeNode, seq []string) []string {
	if p == nil {
		seq = append(seq, "")
		return seq
	}

	seq = append(seq, strconv.Itoa(p.Val))
	seq = genFrontScanTree(p.Left, seq)
	seq = genFrontScanTree(p.Right, seq)

	return seq
}
