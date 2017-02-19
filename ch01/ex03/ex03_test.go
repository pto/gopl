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
