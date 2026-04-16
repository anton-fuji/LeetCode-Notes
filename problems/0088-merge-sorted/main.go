// Problem: 88. Merge Sorted Arrary
// Difficulty: Easy
// Tags: Merge
// URL: https://leetcode.com/problems/merge-sorted-array/
// Date: 2026-04-16

package mergesorted

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := n + m - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}
