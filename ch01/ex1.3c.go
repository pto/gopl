// Time another less-efficient implementation of echo
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	duration := time.Now().Sub(start)
	fmt.Println(s)
	fmt.Println(duration)
}

//!-
