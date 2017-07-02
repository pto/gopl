// Ex10 prints its argument numbers with a comma at each power of 1000.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	pos := (len(s)-1)%3 + 1 // initially, # digits without a comma
	buf.WriteString(s[:pos])
	for pos < len(s) {
		buf.WriteByte(',')
		buf.WriteString(s[pos : pos+3])
		pos += 3
	}
	return buf.String()
}
