// Ex10 prints its argument numbers with a comma at each power of 1000.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	pos := (len(s)-1)%3 + 1 // number of digits before first comma
	buf.WriteString(s[:pos])
	for pos < len(s) {
		buf.WriteByte(',')
		buf.WriteString(s[pos : pos+3])
		pos += 3
	}
	return buf.String()
}
