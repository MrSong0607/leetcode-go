package tests

import "testing"

func TestMergeSortedArray(t *testing.T) {
	nums1, m, nums2, n := []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3
	t.Log(nums1, m, nums2, n)
	merge(nums1, m, nums2, n)
	t.Log(nums1)

	nums1, m, nums2, n = []int{0}, 0, []int{1}, 1
	t.Log(nums1, m, nums2, n)
	merge(nums1, m, nums2, n)
	t.Log(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, l := m-1, n-1, m+n-1
	for ; i >= 0 && j >= 0; l-- {
		if nums1[i] >= nums2[j] {
			nums1[l] = nums1[i]
			i--
			continue
		}

		nums1[l] = nums2[j]
		j--
	}

	for i >= 0 {
		nums1[l] = nums1[i]
		i--
		l--
	}

	for j >= 0 {
		nums1[l] = nums2[j]
		j--
		l--
	}
}
