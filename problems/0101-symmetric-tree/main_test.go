package symmetric

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		in   []any
		want bool
	}{
		{
			name: "example1_true",
			in:   []any{1, 2, 2, 3, 4, 4, 3},
			want: true,
		},
		{
			name: "example2_false",
			in:   []any{1, 2, 2, nil, 3, nil, 3},
			want: false,
		},
		{
			name: "single_node",
			in:   []any{1},
			want: true,
		},
		{
			name: "empty",
			in:   []any{},
			want: true,
		},
		{
			name: "not_symmetric_structure",
			in:   []any{1, 2, 2, nil, 3, 4, nil},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sliceToTree(tt.in)
			got := isSymmetric(root)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func sliceToTree(vals []any) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	if vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// 左
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// 右
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
