package tests

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	a, b := "1010", "1011"
	t.Log(a, b, addBinary(a, b))

	a, b = "11", "1"
	t.Log(a, b, addBinary(a, b))

	a, b = "1", "111"
	t.Log(a, b, addBinary(a, b))

	a, b = "100", "110010"
	t.Log(a, b, addBinary(a, b))

	a, b = "110010", "10111"
	t.Log(a, b, addBinary(a, b))
}

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}
	shortLength := len(b)

	up, sum := false, ""

	for i := 1; shortLength-i >= 0; i++ {
		aIndex, bIndex := len(a)-i, shortLength-i
		if up {
			if a[aIndex] == '1' && b[bIndex] == '1' {
				sum = "1" + sum
			} else {
				if a[aIndex] == '1' || b[bIndex] == '1' {
					sum = "0" + sum
				} else {
					sum = "1" + sum
					up = false
				}
			}
		} else {
			if a[aIndex] == '1' && b[bIndex] == '1' {
				sum = "0" + sum
				up = true
			} else {
				if a[aIndex] == '1' || b[bIndex] == '1' {
					sum = "1" + sum
				} else {
					sum = "0" + sum
				}
			}
		}
	}

	for i := len(a) - shortLength - 1; i >= 0; i-- {
		if up {
			if a[i] == '1' {
				sum = "0" + sum
			} else {
				sum = "1" + sum
				up = false
			}
		} else {
			sum = a[:i+1] + sum
			break
		}
	}

	if up {
		sum = "1" + sum
	}

	return sum
}
