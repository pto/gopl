// Ex06 reads standard input and eliminates duplicate Unicode spaces.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	buf = squash(buf)
	fmt.Print(string(buf))
}

// squash consolidates Unicode spaces into a singe space.
func squash(src []byte) []byte {
	result := src[:0]
	gotSpace := false
	i := 0
	for i < len(src) {
		r, size := utf8.DecodeRune(src[i:])
		isSpace := unicode.IsSpace(r)
		if !isSpace || !gotSpace {
			result = append(result, src[i:i+size]...)
		}
		gotSpace = isSpace
		i += size
	}
	return result
}
