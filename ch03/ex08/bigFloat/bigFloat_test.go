package main

import (
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
			xr := big.NewFloat(real(x)).SetPrec(64)
			xi := big.NewFloat(imag(x)).SetPrec(64)
			yr := big.NewFloat(real(y)).SetPrec(64)
			yi := big.NewFloat(imag(y)).SetPrec(64)
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
		xr := big.NewFloat(real(x)).SetPrec(64)
		xi := big.NewFloat(imag(x)).SetPrec(64)
		abs := cabs(xr, xi)
		fabs, _ := abs.Float64()
		if !close(fabs, cmplx.Abs(x)) {
			t.Errorf("cabs(%v, %v) is %g, should be %g",
				xr, xi, fabs, cmplx.Abs(x))
		}
	}
}

func TestSqrt(t *testing.T) {
	for _, x := range cases {
		x := real(x)
		if x < 0 {
			continue
		}
		xr := big.NewFloat(x).SetPrec(64)
		root := bfsqrt(xr)
		froot, _ := root.Float64()
		if !close(froot, math.Sqrt(x)) {
			t.Errorf("bfsqrt(%v) is %g, should be %g", xr, froot, math.Sqrt(x))
		}
	}
}

func close(a, b float64) bool {
	if b == 0 {
		return a == 0
	}
	return math.Abs(a-b)/b < 1e-15
}
