package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{name: "no carry", digits: []int{1, 2, 3}, want: []int{1, 2, 4}},
		{name: "single carry", digits: []int{1, 2, 9}, want: []int{1, 3, 0}},
		{name: "all nines", digits: []int{9, 9, 9}, want: []int{1, 0, 0, 0}},
		{name: "single digit", digits: []int{0}, want: []int{1}},
		{name: "single nine", digits: []int{9}, want: []int{1, 0}},
		{name: "mid carry chain", digits: []int{8, 9, 9}, want: []int{9, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.digits))
			copy(input, tt.digits)
			got := plusOne(input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne(%v) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}
