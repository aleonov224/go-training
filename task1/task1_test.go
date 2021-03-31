package main

import "testing"

func TestIsAutomorphic(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{1, true},
		{5, true},
		{6, true},
		{25, true},
		{4, false},
	}
	for _, c := range cases {
		got := IsAutomorphic(c.in)
		if got != c.want {
			t.Errorf("isAutomorphic(%d) == %t, want %t", c.in, got, c.want)
		}
	}
}
