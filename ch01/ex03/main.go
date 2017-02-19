// Ex03 prints running times for versions of the echo program.
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var input []string
	for i := '!'; i <= '~'; i++ {
		input = append(input, strings.Repeat(string(i), 20))
	}
	timeFunc("Index:", ConcatIndex, input)
	timeFunc("Range:", ConcatRange, input)
	timeFunc("Join:", ConcatJoin, input)
}

// timeFunc runs function f on slice s and prints the duration.
func timeFunc(heading string, f func([]string) string, s []string) {
	start := time.Now()
	f(s)
	fmt.Println(heading, time.Now().Sub(start))
}

// ConcatIndex concatenates slice elements using a traditional for loop.
func ConcatIndex(slice []string) string {
	var s, sep string
	for i := 0; i < len(slice); i++ {
		s += sep + slice[i]
		sep = " "
	}
	return s
}

// ConcatRange concatenates slice elements using a for range loop.
func ConcatRange(slice []string) string {
	var s, sep string
	for _, elem := range slice {
		s += sep + elem
		sep = " "
	}
	return s
}

// ConcatJoin concatenates slice elements using strings.Join.
func ConcatJoin(slice []string) string {
	return strings.Join(slice, " ")
}
