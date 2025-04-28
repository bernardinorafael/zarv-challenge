package manhattan

import "testing"

func TestDistance(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{
			name:   "empty matrix",
			matrix: [][]int{},
			want:   -1,
		},
		{
			name:   "matrix with empty row",
			matrix: [][]int{{}},
			want:   -1,
		},
		{
			name: "success manhattan distance",
			matrix: [][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 1},
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.matrix); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
