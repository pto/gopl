// Ex03 prints running times for versions of the echo program.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	timeFunc("Index:", ConcatIndex)
	timeFunc("Range:", ConcatRange)
	timeFunc("Join: ", ConcatJoin)
	timeFunc("NoOp: ", NoOp)
}

// timeFunc runs function f on os.Args and prints the duration.
func timeFunc(heading string, f func([]string) string) {
	start := time.Now()
	f(os.Args[1:])
	duration := time.Now().Sub(start)
	fmt.Println(heading, duration)
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

// ConcatRange concatenates slice elements using a for-range loop.
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

// NoOp does no concatenation to test the overhead of the timing function.
func NoOp(_ []string) string {
	return ""
}
