package figures

import (
	"testing"
)

func TestCircleLength(t *testing.T) {
	cases := []struct {
		in   Circle
		want float64
	}{
		{Circle{Dot{1, 2, 3}, 1, 4.0}, 4.00},
		{Circle{Dot{7, 2, 1}, 1, 8.0}, 5.66},
		{Circle{Dot{6, 4, 4}, 1, 5.1}, 4.52},
		{Circle{Dot{3, 7, 2}, 1, 7.0}, 5.29},
		{Circle{Dot{9, 4, 6}, 1, 9.2}, 6.07},
	}

	for _, c := range cases {
		got := c.in.FindLength()
		if got != c.want {
			t.Errorf("Expected %f, got %f", c.want, got)
		}
	}
}

func TestCircleRadius(t *testing.T) {
	cases := []struct {
		in   Circle
		want float64
	}{
		{Circle{Dot{1, 2, 3}, 1, 4.0}, 1.13},
		{Circle{Dot{7, 2, 1}, 1, 8.0}, 1.60},
		{Circle{Dot{6, 4, 4}, 1, 5.1}, 1.27},
		{Circle{Dot{3, 7, 2}, 1, 7.0}, 1.49},
		{Circle{Dot{9, 4, 6}, 1, 19.2}, 2.47},
	}

	for _, c := range cases {
		got := c.in.FindRadius()
		if got != c.want {
			t.Errorf("Expected %f, got %f", c.want, got)
		}
	}
}

func TestCircleColor(t *testing.T) {
	cases := []struct {
		in       Circle
		field_in Field
		want     Color
	}{
		{Circle{Dot{1, 2, 3}, 1, 5.8}, Field{2, Parallelogram{}, Circle{}, [3]Dot{}}, 2},
		{Circle{Dot{7, 2, 1}, 1, 9.8}, Field{3, Parallelogram{}, Circle{}, [3]Dot{}}, 3},
		{Circle{Dot{6, 4, 4}, 1, 5.1}, Field{1, Parallelogram{}, Circle{}, [3]Dot{}}, 5},
		{Circle{Dot{3, 7, 2}, 1, 12.0}, Field{4, Parallelogram{}, Circle{}, [3]Dot{}}, 3},
		{Circle{Dot{9, 4, 6}, 1, 10.0}, Field{5, Parallelogram{}, Circle{}, [3]Dot{}}, 2},
	}

	for _, c := range cases {
		c.in.SetColor(c.field_in)
		got := c.in.color
		if got != c.want {
			t.Errorf("Expected %s, got %s", c.want, got)
		}
	}
}

func TestDotColor(t *testing.T) {
	cases := []struct {
		in   Color
		want Color
	}{
		{5, 5},
		{4, 4},
		{3, 3},
		{2, 2},
		{1, 1},
	}

	test_dot := Dot{}

	for _, c := range cases {
		test_dot.SetColor(c.in)
		got := test_dot.color
		if got != c.want {
			t.Errorf("Expected %s, got %s", c.want, got)
		}
	}
}

func TestParalCenter(t *testing.T) {
	cases := []struct {
		in   Parallelogram
		want Dot
	}{
		{Parallelogram{[4]Dot{{1, 1, 3}, {2, 4, 5}, {6, 4, 5}, {5, 1, 3}}}, Dot{3.5, 2.5, 3}},
		{Parallelogram{[4]Dot{{3, 1, 3}, {2, 4, 5}, {6, 4, 5}, {7, 1, 3}}}, Dot{4.5, 2.5, 3}},
		{Parallelogram{[4]Dot{{0, 1, 3}, {2, 4, 5}, {6, 4, 5}, {4, 1, 3}}}, Dot{3, 2.5, 2}},
	}
	for _, c := range cases {
		got := c.in.FindCenter()
		if got != c.want {
			t.Errorf("Expected %v, got %v", c.want, got)
		}
	}
}

func TestParalPeak(t *testing.T) {
	cases := []struct {
		in   Parallelogram
		want Dot
	}{
		{Parallelogram{[4]Dot{{1, 1, 3}, {2, 4, 5}, {6, 4, 5}, {5, 1, 3}}}, Dot{5, 1, 3}},
		{Parallelogram{[4]Dot{{3, 1, 3}, {2, 4, 5}, {6, 4, 5}, {7, 1, 3}}}, Dot{7, 1, 4}},
		{Parallelogram{[4]Dot{{0, 1, 3}, {2, 4, 5}, {6, 4, 5}, {4, 1, 3}}}, Dot{4, 1, 2}},
	}
	for _, c := range cases {
		got := c.in.CalculatePeak()
		if got != c.want {
			t.Errorf("Expected %v, got %v", c.want, got)
		}
	}
}

func TestParalArea(t *testing.T) {
	cases := []struct {
		in   Parallelogram
		want float64
	}{
		{Parallelogram{[4]Dot{{1, 1, 3}, {2, 4, 5}, {8, 4, 5}, {7, 1, 3}}}, 18},
		{Parallelogram{[4]Dot{{3, 1, 3}, {2, 4, 5}, {6, 4, 5}, {7, 1, 3}}}, 12},
		{Parallelogram{[4]Dot{{0, 1, 3}, {2, 4, 5}, {4, 4, 5}, {2, 1, 3}}}, 6},
	}
	for _, c := range cases {
		got := c.in.CalculateArea()
		if got != c.want {
			t.Errorf("Expected %v, got %v", c.want, got)
		}
	}
}
