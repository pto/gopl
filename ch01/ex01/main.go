// Ex01 prints os.Args[0], the name of the command that invoked it, and all
// command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
