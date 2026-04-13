package sqrtx

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "example1",
			x:    4,
			want: 2,
		},
		{
			name: "example2",
			x:    8,
			want: 2,
		},
		{
			name: "example3",
			x:    0,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mySqrt(tt.x)
			if got != tt.want {
				t.Errorf("mySqrt(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}
