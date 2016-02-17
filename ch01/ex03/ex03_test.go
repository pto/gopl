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
		t.Fatalf("result %s does not match expected %s", result, expectedOutput)
	}
}

func TestConcatIndex(t *testing.T) {
	result := ConcatIndex(input)
	check(result, t)
}

func TestConcatRange(t *testing.T) {
	result := ConcatRange(input)
	check(result, t)
}

func TestConcatJoin(t *testing.T) {
	result := ConcatJoin(input)
	check(result, t)
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
