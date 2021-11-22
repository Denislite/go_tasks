package main

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"CivIc", true},
		{"AnNa", true},
		{"anna", true},
		{"abac", false},
		{"aabAa", true},
		{"a", true},
	}

	for _, c := range cases {
		got := palindromeCheck(c.in)
		if got != c.want {
			t.Errorf("%q.Expected %t, got %t", c.in, c.want, got)
		}
	}
}
