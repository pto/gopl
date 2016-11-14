// Ex04 prints the count, text and filename(s) of lines that appear more than
// once in the input.  It reads from Stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// countMap maps the text of a line to a map that maps a filename to the
// number of occurrences of the text in that file.
type countMap map[string](map[string]int)

func main() {
	counts := make(countMap) // counts[line][filename]
	readInput(counts)
	outputResults(counts)
}

func readInput(counts countMap) {
	filenames := os.Args[1:]
	if len(filenames) == 0 {
		countLines(os.Stdin, "Stdin", counts)
	} else {
		for _, filename := range filenames {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ex04:", err)
				continue
			}
			countLines(file, filename, counts)
			file.Close()
		}
	}
}

func countLines(file *os.File, filename string, counts countMap) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		counts[line][filename]++
	}
	// NOTE: ignoring potential errors from scanner.Err()
}

func outputResults(counts countMap) {
	for line, filemap := range counts {
		linecount := 0
		filenames, sep := "", ""
		for filename, filecount := range filemap {
			linecount += filecount
			filenames += sep + filename
			sep = ", "
		}
		if linecount > 1 {
			fmt.Printf("%d\t%s (%s)\n", linecount, line, filenames)
		}
	}
}
