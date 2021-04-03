package main

import "testing"

func TestCompress(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"кооооооординатааааа", "к#7#ординат#5#а"},
		{"AAABBB", "AAABBB"},
		{"#AAABBB", "#AAABBB"},
	}
	for _, c := range cases {
		got := Compress(c.in)
		if got != c.want {
			t.Errorf("Compress(%s) == %s, want %s", c.in, got, c.want)
		}
	}
}

//Decompress

func TestDecompress(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"к#7#ординат#5#а3", "кооооооординатааааа3"},
		{"AAABBB", "AAABBB"},
		{"#AAABBB", "#AAABBB"},
	}
	for _, c := range cases {
		got := Decompress(c.in)
		if got != c.want {
			t.Errorf("Decompress(%s) == %s, want %s", c.in, got, c.want)
		}
	}
}
