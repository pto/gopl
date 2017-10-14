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
	xr := big.NewRat(1, 1)
	xi := big.NewRat(1, 1)
	yr := big.NewRat(1, 1)
	yi := big.NewRat(1, 1)
	for _, x := range cases {
		for _, y := range cases {
			xr.SetFloat64(real(x))
			xi.SetFloat64(imag(x))
			yr.SetFloat64(real(y))
			yi.SetFloat64(imag(y))
			zr, zi := cadd(xr, xi, yr, yi)
			fr, _ := zr.Float64()
			fi, _ := zi.Float64()
			z := x + y
			if !close(fr, real(z)) || !close(fi, imag(z)) {
				t.Errorf("cadd(%v, %v, %v, %v) is %v,%v, should be %v,%v",
					xr, xi, yr, yi, zr, zi, real(z), imag(z))
			}
			zr, zi = cmult(xr, xi, yr, yi)
			fr, _ = zr.Float64()
			fi, _ = zi.Float64()
			z = x * y
			if !close(fr, real(z)) || !close(fi, imag(z)) {
				t.Errorf("cmult(%v, %v, %v, %v) is %g,%g, should be %g,%g",
					xr, xi, yr, yi, fr, fi, real(z), imag(z))
			}
		}
	}
}

func TestAbs(t *testing.T) {
	xr := big.NewRat(1, 1)
	xi := big.NewRat(1, 1)
	for _, x := range cases {
		xr.SetFloat64(real(x))
		xi.SetFloat64(imag(x))
		abs := cabsSquared(xr, xi)
		fabs, _ := abs.Float64()
		if !close(fabs, cmplx.Abs(x)*cmplx.Abs(x)) {
			t.Errorf("cabs(%v, %v) is %g, should be %g",
				xr, xi, fabs, cmplx.Abs(x))
		}
	}
}

func TestMandelbrot(t *testing.T) {
	xr := big.NewRat(1, 1)
	xi := big.NewRat(1, 1)
	for _, x := range cases {
		xr.SetFloat64(real(x))
		xi.SetFloat64(imag(x))
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

func BenchmarkMandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xr := big.NewRat(0, 1)
		xi := big.NewRat(-1, 1)
		mandelbrot(xr, xi)
	}
}
