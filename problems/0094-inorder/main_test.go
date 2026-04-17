package inorder

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		in   []any
		want []int
	}{
		{
			name: "example1",
			in:   []any{1, nil, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "single_node",
			in:   []any{1},
			want: []int{1},
		},
		{
			name: "balanced",
			in:   []any{1, 2, 3},
			want: []int{2, 1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sliceToTree(tt.in)
			got := inorderTraversal(root)

			if !reflect.DeepEqual(got, tt.want) {
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
