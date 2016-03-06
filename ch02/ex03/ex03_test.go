package ex03

import "testing"

const testWord uint64 = 0x1234567890abcdef

func TestPopCountLoop1(t *testing.T) {
	result := PopCountLoop1(testWord)
	expected := PopCount(testWord)
	if result != expected {
		t.Fatalf("%d does not match expected result %d", result, expected)
	}
}

func TestPopCountLoop2(t *testing.T) {
	result := PopCountLoop2(testWord)
	expected := PopCount(testWord)
	if result != expected {
		t.Fatalf("%d does not match expected result %d", result, expected)
	}
}

func TestPopCountShift(t *testing.T) {
	result := PopCountShift(testWord)
	expected := PopCount(testWord)
	if result != expected {
		t.Fatalf("%d does not match expected result %d", result, expected)
	}
}

func TestPopCountMask(t *testing.T) {
	result := PopCountMask(testWord)
	expected := PopCount(testWord)
	if result != expected {
		t.Fatalf("%d does not match expected result %d", result, expected)
	}
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
