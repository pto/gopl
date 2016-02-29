package ex03

import (
	"strings"
	"testing"
)

var (
	input          = make([]string, 0, '~'-'!'+2) // ! to ~ inclusive, plus 1
	expectedOutput string
)

func init() {
	input = append(input, "Starting with a very long parameter just for fun")
	for i := '!'; i <= '~'; i++ {
		input = append(input, string(i))
	}
	expectedOutput = strings.Join(input, " ")
}

func check(result string, t *testing.T) {
	if result != expectedOutput {
		t.Fatalf("result %q does not match expected %q", result, expectedOutput)
	}
}

func TestConcatIndex(t *testing.T) {
	check(ConcatIndex(input), t)
}

func TestConcatRange(t *testing.T) {
	check(ConcatRange(input), t)
}

func TestConcatJoin(t *testing.T) {
	check(ConcatJoin(input), t)
}

func BenchmarkConcatIndex(b *testing.B) {
	for i := 1; i < b.N; i++ {
		ConcatIndex(input)
	}
}

func BenchmarkConcatRange(b *testing.B) {
	for i := 1; i < b.N; i++ {
		ConcatRange(input)
	}
}

func BenchmarkConcatJoin(b *testing.B) {
	for i := 1; i < b.N; i++ {
		ConcatJoin(input)
	}
}
