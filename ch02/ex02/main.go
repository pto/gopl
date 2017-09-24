// Ex02 converts its numeric arguments to various units.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pto/gopl/ch02/ex02/conv"
)

func main() {
	for _, arg := range os.Args[1:] {
		display(arg)
	}
}

var sep string

func display(val string) {
	fmt.Print(sep)
	sep = "\n"
	x, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex02: %v\n", err)
		return
	}

	f := conv.Fahrenheit(x)
	c := conv.Celsius(x)
	fmt.Printf("%s = %s, %s = %s\n", f, conv.FToC(f), c, conv.CToF(c))

	ft := conv.Feet(x)
	m := conv.Meters(x)
	fmt.Printf("%s = %s, %s = %s\n", ft, conv.FToM(ft), m, conv.MToF(m))

	p := conv.Pounds(x)
	k := conv.Kilograms(x)
	fmt.Printf("%s = %s, %s = %s\n", p, conv.PToK(p), k, conv.KToP(k))

	a := conv.Acres(x)
	h := conv.Hectares(x)
	fmt.Printf("%s = %s, %s = %s\n", a, conv.AToH(a), h, conv.HToA(h))

	g := conv.Gallons(x)
	l := conv.Liters(x)
	fmt.Printf("%s = %s, %s = %s\n", g, conv.GToL(g), l, conv.LToG(l))
}
