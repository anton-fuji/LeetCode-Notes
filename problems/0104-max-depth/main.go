// Problem: 104. Maximum Depth of Binary Tree
// Difficulty: Easy
// Tags: DFS
// URL: https://leetcode.com/problems/maximum-depth-of-binary-tree/
// Date: 2026-04-27

package maxdepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	if l > r {
		return l + 1
	}
	return r + 1
}
