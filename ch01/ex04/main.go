// Ex04 prints the count, text and filename of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files
// and only uses Go features that have been introduced so far.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string](map[string]int)) // counts[line][file]
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "Stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
	for line, filenames := range counts {
		linecount := 0
		for _, filecount := range filenames {
			linecount += filecount
		}
		if linecount > 1 {
			fmt.Printf("%d\t%s (", linecount, line)
			sep := ""
			for name, _ := range filenames {
				fmt.Printf("%s%s", sep, name)
				sep = " "
			}
			fmt.Println(")")
		}
	}
}

func countLines(f *os.File, filename string, counts map[string](map[string]int)) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		counts[line][filename]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
