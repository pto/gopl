// Ex08 counts characters in their Unicode categories.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type category struct {
	name     string
	function func(rune) bool
}

var categories = []category{
	{"Control", unicode.IsControl},
	{"Digit", unicode.IsDigit},
	{"Graphic", unicode.IsGraphic},
	{"Letter", unicode.IsLetter},
	{"Lower", unicode.IsLower},
	{"Mark", unicode.IsMark},
	{"Number", unicode.IsNumber},
	{"Print", unicode.IsPrint},
	{"Punct", unicode.IsPunct},
	{"Space", unicode.IsSpace},
	{"Symbol", unicode.IsSymbol},
	{"Title", unicode.IsTitle},
	{"Upper", unicode.IsUpper},
}

func main() {
	counts := make(map[string]int) // count of characters by category
	invalid := 0                   // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex08: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for _, c := range categories {
			if c.function(r) {
				counts[c.name]++
			}
		}
	}
	fmt.Printf("%s%8s\n", "Category", "count")
	for _, c := range categories {
		fmt.Printf("%-8s%8d\n", c.name, counts[c.name])
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
