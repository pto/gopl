package main

import (
	"strings"
	"testing"
)

var input []string

func init() {
	for i := '!'; i <= '~'; i++ {
		input = append(input, strings.Repeat(string(i), 20))
	}
}

func TestConcat(t *testing.T) {
	cases := []struct {
		input    []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{" "}, " "},
		{[]string{" ", " "}, "   "},
		{[]string{"a"}, "a"},
		{[]string{" a "}, " a "},
		{[]string{"first", "second"}, "first second"},
		{[]string{"first ", " second"}, "first   second"},
		{input, strings.Join(input, " ")},
	}
	for _, c := range cases {
		result := ConcatIndex(c.input)
		if result != c.expected {
			t.Errorf("ConcatIndex(%q) is %q, want %q",
				c.input, result, c.expected)
		}
		result = ConcatRange(c.input)
		if result != c.expected {
			t.Errorf("ConcatRange(%q) is %q, want %q",
				c.input, result, c.expected)
		}
		result = ConcatJoin(c.input)
		if result != c.expected {
			t.Errorf("ConcatJoin(%q) is %q, want %q",
				c.input, result, c.expected)
		}
	}
	s1 := ConcatIndex(input)
	s2 := ConcatRange(input)
	s3 := ConcatJoin(input)
	if s1 != s2 || s2 != s3 {
		t.Error("concat functions differ")
	}
}

func BenchmarkConcatIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatIndex(input)
	}
}

func BenchmarkConcatRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatRange(input)
	}
}

func BenchmarkConcatJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatJoin(input)
	}
}

func BenchmarkNoOp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoOp(input)
	}
}
