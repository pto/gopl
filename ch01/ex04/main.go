// Ex04 prints the count, text and filename of lines that appear more than once
// in the input.  It reads from Stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type countMap map[string](map[string]int)

func main() {
	counts := make(countMap) // counts[line][filename]
	processFiles(counts)
	printResults(counts)
}

func processFiles(counts countMap) {
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "Stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, "dup2:", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}
}

func printResults(counts countMap) {
	for line, files := range counts {
		linecount := 0
		filenames, sep := "", ""
		for filename, filecount := range files {
			linecount += filecount
			filenames += sep + filename
			sep = ", "
		}
		if linecount > 1 {
			fmt.Printf("%d\t%s (%s)\n", linecount, line, filenames)
		}
	}
}

func countLines(f *os.File, filename string, counts countMap) {
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
