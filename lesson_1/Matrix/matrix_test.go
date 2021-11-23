package matrix

import (
	"testing"
)

func TestMatrix(t *testing.T) {
	cases := []struct {
		in   Matrix
		want int
	}{
		{Matrix{{1, 3, 2, 0}, {0, 5, 0, 9}, {1, 2, 3, 4}}, 13},
		{Matrix{{0, 1, 1, 2}, {0, 5, 0, 0}, {2, 0, 3, 3}}, 9},
		{Matrix{{1, 1, 1, 0}, {0, 5, 0, 1}, {2, 1, 3, 10}}, 9},
	}

	for _, c := range cases {
		got := matrixElementSum(c.in)
		if got != c.want {
			t.Errorf(".Expected %d, got %d", c.want, got)
		}
	}
}
