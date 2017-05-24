package main

import (
	"fmt"
	"image/color"
	"math/big"
)

func main() {
	x := big.NewRat(0, 1)
	y := big.NewRat(1, 2)
	fmt.Println(mandelbrot(x, y))
}

var four = big.NewRat(4, 1)

func mandelbrot(zr, zi *big.Rat) color.Color {
	const iterations = 200
	const contrast = 15
	vr := big.NewRat(0, 1)
	vi := big.NewRat(0, 1)
	abs := big.NewRat(0, 1)
	for n := 0; n < iterations; n++ {
		fmt.Println("iteration", n, "v:", vr, vi)
		vr, vi = cmult(vr, vi, vr, vi)
		vr, vi = cadd(vr, vi, zr, zi)
		abs = cabsSquared(vr, vi)
		fmt.Println("abs", abs)
		abs.Sub(abs, four)
		fmt.Println("abs-4:", abs)
		if abs.Sign() > 0 {
			return color.Gray{255 - uint8(contrast*n)}
		}
	}
	return color.Black
}

func cadd(xr, xi, yr, yi *big.Rat) (*big.Rat, *big.Rat) {
	t1 := big.NewRat(0, 1)
	t2 := big.NewRat(0, 1)
	t1.Add(xr, yr)
	t2.Add(xi, yi)
	return t1, t2
}

func cmult(xr, xi, yr, yi *big.Rat) (*big.Rat, *big.Rat) {
	t1 := big.NewRat(0, 1)
	t2 := big.NewRat(0, 1)
	zr := big.NewRat(0, 1)
	zi := big.NewRat(0, 1)
	t1.Mul(xr, yr)
	t2.Mul(xi, yi)
	zr.Sub(t1, t2)
	t1.Mul(xr, yi)
	t2.Mul(xi, yr)
	zi.Add(t1, t2)
	return zr, zi
}

func cabsSquared(xr, xi *big.Rat) *big.Rat {
	t1 := big.NewRat(0, 1)
	t2 := big.NewRat(0, 1)
	t1.Mul(xr, xr)
	t2.Mul(xi, xi)
	return t1.Add(t1, t2)
}
