// Problem: 101. Symmetric Tree
// Difficulty: Easy
// Tags: DFS
// URL: https://leetcode.com/problems/symmetric-tree/
// Date: 2026-04-22

package symmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}

	return isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
}
