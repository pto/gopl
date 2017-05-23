// BigFloat emits a zoomed-in PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

const prec = 256

func main() {
	const (
		x, y                   = 0.0, 1.0
		zoom                   = 1e14
		radius                 = 2.0 / zoom
		xmin, ymin, xmax, ymax = x - radius, y - radius, x + radius, y + radius
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := big.NewFloat(float64(py)).SetPrec(prec)
		h := big.NewFloat(height).SetPrec(prec)
		r := big.NewFloat(ymax - ymin).SetPrec(prec)
		m := big.NewFloat(ymin).SetPrec(prec)
		y.Mul(y, h)
		y.Quo(y, r)
		y.Add(y, m)
		for px := 0; px < width; px++ {
			x := big.NewFloat(float64(px)).SetPrec(prec)
			w := big.NewFloat(width).SetPrec(prec)
			r := big.NewFloat(xmax - xmin)
			m := big.NewFloat(xmin)
			x.Mul(x, w)
			x.Quo(x, r)
			x.Add(x, m)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(x, y))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

var two = big.NewFloat(2).SetPrec(prec)

func mandelbrot(zr, zi *big.Float) color.Color {
	const iterations = 200
	const contrast = 15
	vr := big.NewFloat(0).SetPrec(prec)
	vi := big.NewFloat(0).SetPrec(prec)
	abs := big.NewFloat(0).SetPrec(prec)
	for n := 0; n < iterations; n++ {
		vr, vi = cmult(vr, vi, vr, vi)
		vr, vi = cadd(vr, vi, zr, zi)
		abs = cabs(vr, vi)
		abs.Sub(abs, two)
		if abs.Sign() > 0 {
			return color.Gray{255 - uint8(contrast*n)}
		}
	}
	return color.Black
}

func cadd(xr, xi, yr, yi *big.Float) (*big.Float, *big.Float) {
	t1 := big.NewFloat(0).SetPrec(prec)
	t2 := big.NewFloat(0).SetPrec(prec)
	t1.Add(xr, yr)
	t2.Add(xi, yi)
	return t1, t2
}

func cmult(xr, xi, yr, yi *big.Float) (*big.Float, *big.Float) {
	t1 := big.NewFloat(0).SetPrec(prec)
	t2 := big.NewFloat(0).SetPrec(prec)
	zr := big.NewFloat(0).SetPrec(prec)
	zi := big.NewFloat(0).SetPrec(prec)
	t1.Mul(xr, yr)
	t2.Mul(xi, yi)
	zr.Sub(t1, t2)
	t1.Mul(xr, yi)
	t2.Mul(xi, yr)
	zi.Add(t1, t2)
	return zr, zi
}

func cabs(xr, xi *big.Float) *big.Float {
	t1 := big.NewFloat(0).SetPrec(prec)
	t2 := big.NewFloat(0).SetPrec(prec)
	t1.Mul(xr, xr)
	t2.Mul(xi, xi)
	t1.Add(t1, t2)
	return bfsqrt(t1)
}

func bfsqrt(x *big.Float) *big.Float {
	if x.Sign() < 0 {
		panic("bfsqrt cannot handle negative numbers")
	}
	if x.Sign() == 0 {
		return big.NewFloat(0).SetPrec(prec)
	}
	t1 := big.NewFloat(0).SetPrec(prec)
	half := big.NewFloat(0.5).SetPrec(prec)
	guess := big.NewFloat(1).SetPrec(prec)
	for {
		t1.Quo(x, guess)
		t1.Add(guess, t1)
		guess.Mul(half, t1)
		if goodEnough(x, guess) {
			break
		}
	}
	return guess
}

func goodEnough(x, guess *big.Float) bool {
	t1 := big.NewFloat(0).SetPrec(prec)
	t1.Mul(guess, guess)
	t1.Sub(t1, x)
	t1.Quo(t1, x)
	delta, _ := t1.Abs(t1).Float64()
	return delta < 1e-15
}
