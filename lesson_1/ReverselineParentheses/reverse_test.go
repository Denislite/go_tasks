package reverse

import (
	"testing"
)

func TestReverseParantheses(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"foo(bar(baz))blim", "foobazrabblim"},
		{"foo(bar)baz", "foorabbaz"},
		{"foo(bar)baz(blim)", "foorabbazmilb"},
		{"(bar)", "rab"},
	}

	for _, c := range cases {
		got := reverseParantheses(c.in)
		if got != c.want {
			t.Errorf("%q. Expected %q, got %q", c.in, c.want, got)
		}
	}
}
