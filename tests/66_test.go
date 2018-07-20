package tests

import "testing"

func TestPlusOne(t *testing.T) {
	a := []int{1, 2, 3}
	t.Log(a)
	t.Log(plusOne(a))

	a = []int{4, 3, 2, 1}
	t.Log(a)
	t.Log(plusOne(a))

	a = []int{}
	t.Log(a)
	t.Log(plusOne(a))

	a = []int{9}
	t.Log(a)
	t.Log(plusOne(a))

	a = []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	t.Log(a)
	t.Log(plusOne(a))
}

func plusOne(digits []int) []int {
	for j := len(digits) - 1; j >= 0; j-- {
		if j == len(digits)-1 {
			digits[j]++
		}

		if digits[j] == 10 {
			digits[j] = 0
			if j == 0 {
				digits = append([]int{1}, digits...)
			} else {
				digits[j-1]++
			}
		} else {
			break
		}
	}

	return digits
}
