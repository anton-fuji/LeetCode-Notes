// Problem: 94. Binary Tree Inorder Traversal
// Difficulty: Easy
// Tags: DFS
// URL: https://leetcode.com/problems/binary-tree-inorder-traversal/
// Date: 2026-04-17

package inorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return res
}
