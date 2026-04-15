package removedup

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "example1",
			in:   []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			name: "example2",
			in:   []int{1, 1, 2, 3, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "no_duplicates",
			in:   []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "all_duplicates",
			in:   []int{1, 1, 1, 1},
			want: []int{1},
		},
		{
			name: "empty",
			in:   []int{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.in)
			got := deleteDuplicates(head)
			gotSlice := listToSlice(got)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("got %v, want %v", gotSlice, tt.want)
			}
		})
	}
}

func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
