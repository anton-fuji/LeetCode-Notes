package maxdepth

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		root []any
		want int
	}{
		{
			name: "example1",
			root: []any{3, 9, 20, nil, nil, 15, 7},
			want: 3,
		},
		{
			name: "single node",
			root: []any{1},
			want: 1,
		},
		{
			name: "empty tree",
			root: []any{},
			want: 0,
		},
		{
			name: "skewed tree",
			root: []any{1, 2, nil, 3, nil, 4},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.root)
			got := maxDepth(root)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func buildTree(vals []any) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// left
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// right
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
