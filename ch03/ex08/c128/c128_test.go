package main

import (
	"testing"
)

func BenchmarkMandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mandelbrot(0 - 1i)
	}
}
