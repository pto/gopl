// Ex02 converts its numeric arguments (or standard input) to various units.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pto/gopl/ch02/ex02/conv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := conv.Fahrenheit(t)
		c := conv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, conv.FToC(f), c, conv.CToF(c))
	}
}
