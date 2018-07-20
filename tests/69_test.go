package tests

import "testing"

func TestSqrt(t *testing.T) {
	t.Log(4, mySqrt(4))
	t.Log(8, mySqrt(8))
	t.Log(10000, mySqrt(10000))
	t.Log(1000000, mySqrt(1000000))
	t.Log(10000000000, mySqrt(10000000000))
	t.Log(1000000000000000000, mySqrt(1000000000000000000))
}

func mySqrt(x int) int {
	if x == 0 {
		return x
	}

	if x < 4 {
		return 1
	}

	i := 0
	for ; i*i <= x; i++ {
	}

	return i - 1
}
