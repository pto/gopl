// Ex04 rotates slices in a single pass.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		var count int
		for i, s := range strings.Fields(input.Text()) {
			if i == 0 {
				var err error
				count, err = strconv.Atoi(s)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue outer
				}
				continue
			}
			x, err := strconv.Atoi(s)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		if count > len(ints) || count < 0 {
			fmt.Fprintln(os.Stderr, "ex04: invalid rotate count")
			continue outer
		}
		rotate(ints, count)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

// rotate rotates a slice of ints by count places, in place, in a single pass.
func rotate(s []int, count int) {
	initial := make([]int, count)
	copy(initial, s[:count])
	copy(s[:], s[count:])
	copy(s[len(s)-count:], initial)
}
