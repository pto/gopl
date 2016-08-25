package ex03

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
	expected := strings.Join(input, " ")
	check := func(result string) {
		if result != expected {
			t.Fatalf("result %q does not match expected %q", result, expected)
		}
	}
	check(ConcatIndex(input))
	check(ConcatRange(input))
	check(ConcatJoin(input))
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
