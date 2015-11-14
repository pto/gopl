// Benchmark implementations of string concatenation as used in the Echo programs.
package ex03

import (
	"strings"
	"testing"
)

var (
	input  = make([]string, 0, '~'-'!'+2)
	output string
)

func init() {
	input = append(input, "Starting with a very long parameter to get things underway in style")
	for i := '!'; i <= '~'; i++ {
		input = append(input, string(i))
	}
	output = strings.Join(input, " ")
}

func TestConcatIndex(t *testing.T) {
	result := ConcatIndex(input)
	if result != output {
		t.Fatalf("result %s does not match expected %s", result, output)
	}
}

func TestConcatRange(t *testing.T) {
	result := ConcatRange(input)
	if result != output {
		t.Fatalf("result %s does not match expected %s", result, output)
	}
}

func TestConcatJoin(t *testing.T) {
	result := ConcatJoin(input)
	if result != output {
		t.Fatalf("result %s does not match expected %s", result, output)
	}
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
