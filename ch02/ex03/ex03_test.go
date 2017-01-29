package ex03

import "testing"

const testWord uint64 = 0x1234567890abcdef

func TestPopCounts(t *testing.T) {
	expected := PopCount(testWord)
	check := func(f string, result int) {
		if result != expected {
			t.Fatalf("%s: %d does not match expected result %d", f, result, expected)
		}
	}

	check("PopCountLoop1", PopCountLoop1(testWord))
	check("PopCountLoop2", PopCountLoop2(testWord))
	check("PopCountShift", PopCountShift(testWord))
	check("PopCountMask", PopCountMask(testWord))
}

func BenchmarkPopCount(b *testing.B) {
	for i := 1; i < b.N; i++ {
		PopCount(testWord)
	}
}

func BenchmarkPopCountLoop1(b *testing.B) {
	for i := 1; i < b.N; i++ {
		PopCountLoop1(testWord)
	}
}

func BenchmarkPopCountLoop2(b *testing.B) {
	for i := 1; i < b.N; i++ {
		PopCountLoop2(testWord)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 1; i < b.N; i++ {
		PopCountShift(testWord)
	}
}

func BenchmarkPopCountMask(b *testing.B) {
	for i := 1; i < b.N; i++ {
		PopCountMask(testWord)
	}
}
