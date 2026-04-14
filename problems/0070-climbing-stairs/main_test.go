package climbing_stairs

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "example1",
			n:    2,
			want: 2,
		},
		{
			name: "example2",
			n:    3,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			if got != tt.want {
				t.Errorf("climbingStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
