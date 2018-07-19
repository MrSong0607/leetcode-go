package tests

import "testing"

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	a := []int{1, 1, 2}
	t.Log(a, removeDuplicates(a))

	a = []int{1, 1}
	t.Log(a, removeDuplicates(a))

	a = []int{1, 2}
	t.Log(a, removeDuplicates(a))

	a = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(a, removeDuplicates(a))
}

func removeDuplicates(nums []int) int {
	length := len(nums)

	j := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] <= nums[i] {
			if j <= i {
				j = i + 1
			}

			for ; j < len(nums); j++ {
				if nums[j] > nums[i] {
					nums[i+1] = nums[j]
					break
				}
			}
		}

		length = i + 1

		if j == len(nums) {
			break
		}
	}

	if j == 0 {
		length = len(nums)
	}
	return length
}
