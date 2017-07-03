// Ex05 reads standard input and eliminates duplicate adjacent lines.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	s := strings.SplitAfter(string(buf), "\n")
	s = dedup(s)
	fmt.Print(strings.Join(s, ""))
}

// dedup eliminates duplicate adjacent strings in slice src, in place.
func dedup(src []string) []string {
	result := src[:0]
	for _, s := range src {
		if len(result) != 0 && s == result[len(result)-1] {
			continue
		}
		result = append(result, s)
	}
	return result
}
