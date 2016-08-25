// Package ex03 has multiple implementations of string concatenation for
// benchmark testing.
package ex03

import "strings"

// ConcatIndex concatenates slice elements using a traditional for loop.
func ConcatIndex(slice []string) string {
	var s, sep string
	for i := 0; i < len(slice); i++ {
		s += sep + slice[i]
		sep = " "
	}
	return s
}

// ConcatRange concatenates slice elements using a for range loop.
func ConcatRange(slice []string) string {
	var s, sep string
	for _, elem := range slice {
		s += sep + elem
		sep = " "
	}
	return s
}

// ConcatJoin concatenates slice elements using strings.Join.
func ConcatJoin(slice []string) string {
	return strings.Join(slice, " ")
}
