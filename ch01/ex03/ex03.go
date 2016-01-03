// Package ex03 has multiple implementations of string concatenation for
// benchmark testing.
package ex03

import (
	"strings"
)

func ConcatIndex(slice []string) string {
	var s, sep string
	for i := 0; i < len(slice); i++ {
		s += sep + slice[i]
		sep = " "
	}
	return s
}

func ConcatRange(slice []string) string {
	var s, sep string
	for _, d := range slice {
		s += sep + d
		sep = " "
	}
	return s
}

func ConcatJoin(slice []string) string {
	return strings.Join(slice, " ")
}
