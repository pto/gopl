// Echo4 uses the default format for slices to print command line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
}
