package tests

import "testing"

func TestRemoveElement(t *testing.T) {
	a := []int{0, 1, 2, 2, 3, 0, 4, 2}
	t.Log(a, 2)
	t.Log(a, removeElement(a, 2))

	a = []int{3, 2, 2, 3}
	t.Log(a, 3)
	t.Log(a, removeElement(a, 3))

	a = []int{1, 2, 3, 3}
	t.Log(a, 4)
	t.Log(a, removeElement(a, 4))

	a = []int{3, 3, 3, 3}
	t.Log(a, 3)
	t.Log(a, removeElement(a, 3))

	a = []int{4, 5}
	t.Log(a, 4)
	t.Log(a, removeElement(a, 4))

	a = []int{2}
	t.Log(a, 3)
	t.Log(a, removeElement(a, 3))

	a = []int{}
	t.Log(a, 3)
	t.Log(a, removeElement(a, 3))
}

func removeElement(nums []int, val int) int {
	j := len(nums) - 1
	for index := range nums {
		if nums[index] == val {
			for ; j > index; j-- {
				if nums[j] != val {
					nums[index], nums[j] = nums[j], nums[index]
					break
				}
			}
		}

		if j == index {
			if nums[index] == val {
				return index
			}
			return index + 1
		}
	}

	return len(nums)
}
