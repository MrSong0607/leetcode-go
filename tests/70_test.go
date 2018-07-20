package tests

import (
	"testing"
)

func TestClimbingStairs(t *testing.T) {
	t.Log(2, climbStairs(2))
	t.Log(3, climbStairs(3))
	t.Log(5, climbStairs(5))
	t.Log(35, climbStairs(35))
	t.Log(44, climbStairs(44))
}

func climbStairs(n int) int {
	count := 0

	for x := n; x >= 0; x -= 2 {
		y, jc := (n-x)/2, 1

		for i := 1; i <= y+x; i++ {
			if i > y && i > x {
				jc *= i

				if x > y {
					jc /= (i - x)
				} else {
					jc /= (i - y)
				}
			}
		}

		count += jc
	}

	return count
}
