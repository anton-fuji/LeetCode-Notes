// Problem: 108. Convert Sorted Array to Binary Search Tree
// Difficulty: Easy
// Tags: DFS
// URL: https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// Date: 2026-04-28

package sortarr

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) / 2

	root := &TreeNode{Val: nums[mid]}
	root.Left = build(nums, l, mid-1)
	root.Right = build(nums, mid+1, r)

	return root
}
