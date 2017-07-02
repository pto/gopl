// Comma prints its integer or floating point signed argument numbers with a
// comma at each power of 1000.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in an optionally-signed decimal string.
func comma(s string) string {
	if len(s) == 0 {
		return ""
	}
	var sign string
	var fraction string
	if s[0] == '+' || s[0] == '-' {
		sign = s[:1]
		s = s[1:]
	}
	if dot := strings.Index(s, "."); dot >= 0 {
		fraction = s[dot:]
		s = s[:dot]
	}
	return sign + intComma(s) + fraction
}

// intComma inserts commas in an unsigned integer string.
func intComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return intComma(s[:n-3]) + "," + s[n-3:]
}
