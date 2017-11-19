package ex03

import (
	"math/bits"
	"testing"
)

func TestPopCounts(t *testing.T) {
	var testWords = [...]uint64{0x1234567890abcdef, 0x0, 0x800000000,
		0x8000000000000000, 0xFFFFFFFFFFFFFFFF}

	check := func(name string, f func(uint64) int) {
		for _, v := range testWords {
			result := f(v)
			expected := PopCount(v)
			if result != expected {
				t.Fatalf("%s: %d does not match expected result %d",
					name, result, expected)
			}
		}
	}

	check("PopCountLoop1", PopCountLoop1)
	check("PopCountLoop2", PopCountLoop2)
	check("PopCountShift", PopCountShift)
	check("PopCountMask", PopCountMask)
	check("OnesCount64", bits.OnesCount64)
}

func BenchmarkPopCount(b *testing.B) {
	for i := 1; i < b.N; i++ {
		PopCount(0x1234567890abcdef)
	}
}

func BenchmarkPopCountLoop1(b *testing.B) {
	for i := 1; i < b.N; i++ {
		PopCountLoop1(0x1234567890abcdef)
	}
}

func BenchmarkPopCountLoop2(b *testing.B) {
	for i := 1; i < b.N; i++ {
		PopCountLoop2(0x1234567890abcdef)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 1; i < b.N; i++ {
		PopCountShift(0x1234567890abcdef)
	}
}

func BenchmarkPopCountMask(b *testing.B) {
	for i := 1; i < b.N; i++ {
		PopCountMask(0x1234567890abcdef)
	}
}

func BenchmarkOnesCount(b *testing.B) {
	for i := 1; i < b.N; i++ {
		bits.OnesCount64(0x1234567890abcdef)
	}
}
