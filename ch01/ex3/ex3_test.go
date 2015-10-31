// Benchmark implementations of echo
package ex3

import (
	"strings"
	"testing"
)

var (
	input  = make([]string, 0, '~'-'!'+1)
	output string
)

func init() {
	for i := '!'; i <= '~'; i++ {
		input = append(input, string(i))
	}
	output = strings.Join(input, " ")
}

func TestIndex(t *testing.T) {
	result := ConcatIndex(input)
	if result != output {
		t.Fatalf("result %s does not match expected %s", result, output)
	}
}

func TestRange(t *testing.T) {
	result := ConcatRange(input)
	if result != output {
		t.Fatalf("result %s does not match expected %s", result, output)
	}
}

func TestJoin(t *testing.T) {
	result := ConcatJoin(input)
	if result != output {
		t.Fatalf("result %s does not match expected %s", result, output)
	}
}

func BenchmarkIndex(b *testing.B) {
	for i := 1; i < b.N; i++ {
		ConcatIndex(input)
	}
}

func BenchmarkRange(b *testing.B) {
	for i := 1; i < b.N; i++ {
		ConcatRange(input)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 1; i < b.N; i++ {
		ConcatJoin(input)
	}
}
