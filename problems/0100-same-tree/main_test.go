package sametree

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		p    []any
		q    []any
		want bool
	}{
		{
			name: "example1",
			p:    []any{1, 2, 3},
			q:    []any{1, 2, 3},
			want: true,
		},
		{
			name: "different_structure",
			p:    []any{1, 2},
			q:    []any{1, nil, 2},
			want: false,
		},
		{
			name: "different_values",
			p:    []any{1, 2, 1},
			q:    []any{1, 1, 2},
			want: false,
		},
		{
			name: "both_empty",
			p:    []any{},
			q:    []any{},
			want: true,
		},
		{
			name: "one_empty",
			p:    []any{},
			q:    []any{1},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootP := sliceToTree(tt.p)
			rootQ := sliceToTree(tt.q)

			got := isSameTree(rootP, rootQ)

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
