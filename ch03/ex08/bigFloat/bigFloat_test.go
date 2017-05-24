package main

import (
	"image/color"
	"math"
	"math/big"
	"math/cmplx"
	"testing"
)

var cases = []complex128{
	complex(0, 0),
	complex(0, 1),
	complex(1, 0),
	complex(1, 1),
	complex(0, -1),
	complex(-1, 0),
	complex(-1, -1),
	complex(0, 0.5),
	complex(0.5, 0),
	complex(0.12345, -0.12345),
	complex(0.444444, 0.555555),
	complex(1.23456, 2.34566),
	complex(-123123, 5343424),
	complex(313241.23, 534213.24),
	complex(1231231, 0),
	complex(1e15, 2e15),
	complex(1e-15, 2e-15),
}

func TestAddMult(t *testing.T) {
	for _, x := range cases {
		for _, y := range cases {
			xr := big.NewFloat(real(x)).SetPrec(prec)
			xi := big.NewFloat(imag(x)).SetPrec(prec)
			yr := big.NewFloat(real(y)).SetPrec(prec)
			yi := big.NewFloat(imag(y)).SetPrec(prec)
			zr, zi := cadd(xr, xi, yr, yi)
			fr, _ := zr.Float64()
			fi, _ := zi.Float64()
			z := x + y
			if !close(fr, real(z)) || !close(fi, imag(z)) {
				t.Error("cadd(%v, %v, %v, %v) is %v,%v, should be %v,%v",
					xr, xi, yr, yi, zr, zi, real(z), imag(z))
			}
			zr, zi = cmult(xr, xi, yr, yi)
			fr, _ = zr.Float64()
			fi, _ = zi.Float64()
			z = x * y
			if !close(fr, real(z)) || !close(fi, imag(z)) {
				t.Errorf("cmult(%f, %f, %f, %f) is %g,%g, should be %g,%g",
					xr, xi, yr, yi, fr, fi, real(z), imag(z))
			}
		}
	}
}

func TestAbs(t *testing.T) {
	for _, x := range cases {
		xr := big.NewFloat(real(x)).SetPrec(prec)
		xi := big.NewFloat(imag(x)).SetPrec(prec)
		abs := cabsSquared(xr, xi)
		fabs, _ := abs.Float64()
		if !close(fabs, cmplx.Abs(x)*cmplx.Abs(x)) {
			t.Errorf("cabs(%v, %v) is %g, should be %g",
				xr, xi, fabs, cmplx.Abs(x))
		}
	}
}

func TestMandelbrot(t *testing.T) {
	for _, x := range cases {
		xr := big.NewFloat(real(x)).SetPrec(prec)
		xi := big.NewFloat(imag(x)).SetPrec(prec)
		m := mandelbrot(xr, xi)
		c128_m := c128_mandelbrot(x)
		if m != c128_m {
			t.Errorf("mandelbrot(%v, %v) is %v, should be %v",
				xr, xi, m, c128_m)
		}
	}
}

func close(a, b float64) bool {
	if b == 0 {
		return a == 0
	}
	return math.Abs(a-b)/b < 1e-15
}

func c128_mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - uint8(contrast*n)}
		}
	}
	return color.Black
}
