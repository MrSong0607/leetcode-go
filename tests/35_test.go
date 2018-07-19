package tests

import "testing"

func TestSearchInsertPosition(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	t.Log(nums, 5, searchInsert(nums, 5))

	nums = []int{1, 3, 5, 6}
	t.Log(nums, 2, searchInsert(nums, 2))

	nums = []int{1, 3, 5, 6}
	t.Log(nums, 7, searchInsert(nums, 7))

	nums = []int{1, 3, 5, 6}
	t.Log(nums, 0, searchInsert(nums, 0))
}

func searchInsert(nums []int, target int) int {
	for index, val := range nums {
		if target <= val {
			return index
		}
	}
	return len(nums)
}
