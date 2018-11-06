// Ex04 prints the count, text and filename(s) of lines that appear more than
// once in the input.  It reads from Stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int) // counts[line][filename]
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex04: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}

	for line, filemap := range counts {
		linecount := 0
		filenames := []string{}
		for filename, filecount := range filemap {
			linecount += filecount
			filenames = append(filenames, filename)
		}
		if linecount > 1 {
			sort.Strings(filenames) // make the order deterministic for tests
			fmt.Printf("%d\t%s (%s)\n", linecount, line,
				strings.Join(filenames, ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int, name string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		counts[line][name]++
	}
	// NOTE: ignoring potential errors from scanner.Err()
}
