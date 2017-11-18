// Ex01 prints all command line arguments, including os.Args[0], the
// executable name.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
