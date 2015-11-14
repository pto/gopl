// Order demonstrates referencing identifiers out of order.
package main

import "fmt"

func main() {
	fmt.Println(first)
}

const first = "This is printed first, but declared second."
