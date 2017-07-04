// Ex09 reports the frequency of each word in a text file.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	counts := map[string]int{} // map of words to count
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: wordfreq <filename>")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "wordfreq:", err)
	}

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[strings.ToLower(input.Text())]++
	}
	words := make([]string, 0, len(counts))
	for word := range counts {
		words = append(words, word)
	}
	sort.Strings(words)
	fmt.Printf("%-15s %8s\n", "word", "count")
	for _, word := range words {
		fmt.Printf("%-15s %8d\n", word, counts[word])
	}
}
