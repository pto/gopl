// Ex01 prints all command line arguments, including os.Args[0], the
// name of the command that invoked it.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
