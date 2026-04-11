package addbinary

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{
			name: "example1",
			a:    "11",
			b:    "1",
			want: "100",
		},
		{
			name: "example2",
			a:    "1010",
			b:    "1011",
			want: "10101",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addBinary(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("addBinary(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
