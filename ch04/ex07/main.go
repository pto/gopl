// Ex07 reverses a UTF-8 byte slice in place without allocating new memory.
package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ex07 <UTF-8 string>")
		os.Exit(1)
	}
	fmt.Println(string(reverse([]byte(os.Args[1]))))
}

// reverse reverses a UTF-8 byte slice in place.
func reverse(s []byte) []byte {
	i, j := 0, len(s)
	for i < j {
		firstRune, firstSize := utf8.DecodeRune(s[i:])
		lastRune, lastSize := utf8.DecodeLastRune(s[:j])
		shiftToFront := firstSize - lastSize
		if shiftToFront != 0 {
			middle := s[i+firstSize : j-lastSize]
			copy(s[i+firstSize-shiftToFront:], middle)
		}
		utf8.EncodeRune(s[i:], lastRune)
		utf8.EncodeRune(s[j-firstSize:], firstRune)
		i += lastSize
		j -= firstSize
	}
	return s
}
