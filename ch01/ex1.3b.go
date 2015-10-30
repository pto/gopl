// Time a more-efficient implementation of echo.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s := strings.Join(os.Args[1:], " ")
	duration := time.Now().Sub(start)
	fmt.Println(s)
	fmt.Println(duration)
}
