// Quick echo of command line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:], " ")
}
