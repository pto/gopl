// Time a less-efficient implementation of echo.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	duration := time.Now().Sub(start)
	fmt.Println(s)
	fmt.Println(duration)
}
