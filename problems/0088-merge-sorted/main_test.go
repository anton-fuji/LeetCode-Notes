package mergesorted

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		m      int
		nums2  []int
		n      int
		expect []int
	}{
		{
			name:   "example1",
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			expect: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:   "nums1 empty",
			nums1:  []int{0},
			m:      0,
			nums2:  []int{1},
			n:      1,
			expect: []int{1},
		},
		{
			name:   "nums2 empty",
			nums1:  []int{1},
			m:      1,
			nums2:  []int{},
			n:      0,
			expect: []int{1},
		},
		{
			name:   "all nums2 smaller",
			nums1:  []int{4, 5, 6, 0, 0, 0},
			m:      3,
			nums2:  []int{1, 2, 3},
			n:      3,
			expect: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "duplicate values",
			nums1:  []int{2, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{1, 2, 2},
			n:      3,
			expect: []int{1, 2, 2, 2, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)

			if !reflect.DeepEqual(tt.nums1, tt.expect) {
				t.Errorf("got %v, want %v", tt.nums1, tt.expect)
			}
		})
	}
}
