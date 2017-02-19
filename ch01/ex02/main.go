// Ex02 prints the index and value of each command line argument.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i+1, arg)
	}
}
