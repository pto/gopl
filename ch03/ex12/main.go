// Ex12 checks its two arguments to see if they are anagrams of each other.
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ex12 <string1> <string2>")
		os.Exit(1)
	}
	fmt.Println(anagrams(os.Args[1], os.Args[2]))
}

func anagrams(str1, str2 string) bool {
	s1 := []byte(strings.Map(letter, str1))
	s2 := []byte(strings.Map(letter, str2))
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	return string(s1) == string(s2)
}

// letter returns lower-cased letters, or -1 for non-letters
func letter(r rune) rune {
	if unicode.IsLetter(r) {
		return unicode.ToLower(r)
	} else {
		return -1
	}
}
