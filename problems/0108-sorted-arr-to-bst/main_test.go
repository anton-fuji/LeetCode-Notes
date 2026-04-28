package sortarr

import "testing"

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "example1",
			nums: []int{-10, -3, 0, 5, 9},
		},
		{
			name: "single element",
			nums: []int{1},
		},
		{
			name: "two elements",
			nums: []int{1, 2},
		},
		{
			name: "empty",
			nums: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sortedArrayToBST(tt.nums)

			// ① BSTになっているか
			if !isBST(root, nil, nil) {
				t.Errorf("tree is not a valid BST")
			}

			// ② height-balancedか
			if !isBalanced(root) {
				t.Errorf("tree is not balanced")
			}

			// ③ inorder traversal が元配列と一致するか
			var result []int
			inorder(root, &result)

			if len(result) != len(tt.nums) {
				t.Errorf("length mismatch: got %v, want %v", result, tt.nums)
				return
			}

			for i := range result {
				if result[i] != tt.nums[i] {
					t.Errorf("inorder mismatch: got %v, want %v", result, tt.nums)
					break
				}
			}
		})
	}
}

func isBST(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}
	return isBST(node.Left, min, &node.Val) &&
		isBST(node.Right, &node.Val, max)
}

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	l := height(node.Left)
	if l == -1 {
		return -1
	}

	r := height(node.Right)
	if r == -1 {
		return -1
	}

	if l-r > 1 || r-l > 1 {
		return -1
	}

	if l > r {
		return l + 1
	}
	return r + 1
}

func inorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, res)
	*res = append(*res, node.Val)
	inorder(node.Right, res)
}
