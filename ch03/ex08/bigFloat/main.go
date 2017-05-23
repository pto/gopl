// BigFloat emits a zoomed-in PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

const prec = 64

func main() {
	const (
		x, y                   = 0.0, 1.0
		zoom                   = 2.0
		radius                 = 2.0 / zoom
		xmin, ymin, xmax, ymax = x - radius, y - radius, x + radius, y + radius
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			zr := big.NewFloat(x).SetPrec(prec)
			zi := big.NewFloat(y).SetPrec(prec)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(zr, zi))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(zr, zi *big.Float) color.Color {
	const iterations = 200
	const contrast = 15
	vr, vi := big.NewFloat(0).SetPrec(prec), big.NewFloat(0).SetPrec(prec)
	abs := big.NewFloat(0).SetPrec(prec)
	for n := 0; n < iterations; n++ {
		vr, vi = cmult(vr, vi, vr, vi)
		vr, vi = cadd(vr, vi, zr, zi)
		abs = cabs(vr, vi)
		abs = abs.Sub(abs, big.NewFloat(2).SetPrec(prec))
		if abs.Sign() > 0 {
			return color.Gray{255 - uint8(contrast*n)}
		}
	}
	return color.Black
}

var (
	t1 = big.NewFloat(0).SetPrec(prec)
	t2 = big.NewFloat(0).SetPrec(prec)
	t3 = big.NewFloat(0).SetPrec(prec)
	t4 = big.NewFloat(0).SetPrec(prec)
	zr = big.NewFloat(0).SetPrec(prec)
	zi = big.NewFloat(0).SetPrec(prec)
)

func cadd(xr, xi, yr, yi *big.Float) (*big.Float, *big.Float) {
	t1.Add(xr, yr)
	t2.Add(xi, yi)
	return t1, t2
}

func cmult(xr, xi, yr, yi *big.Float) (*big.Float, *big.Float) {
	t1.Mul(xr, yr)
	t2.Mul(xi, yi)
	zr.Sub(t1, t2)
	t1.Mul(xr, yi)
	t2.Mul(xi, yr)
	zi.Add(t1, t2)
	return zr, zi
}

func cabs(xr, xi *big.Float) *big.Float {
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
		return x.SetInt64(0)
	}
	half := big.NewFloat(0.5).SetPrec(prec)
	guess := big.NewFloat(1).SetPrec(prec)
	for {
		t3.Quo(x, guess)
		t3.Add(guess, t3)
		guess.Mul(half, t3)
		if goodEnough(x, guess) {
			break
		}
	}
	return guess
}

func goodEnough(x, guess *big.Float) bool {
	t4.Mul(guess, guess)
	t4.Sub(t4, x)
	t4.Quo(t4, x)
	delta, _ := t4.Abs(t4).Float64()
	return delta < 1e-15
}
