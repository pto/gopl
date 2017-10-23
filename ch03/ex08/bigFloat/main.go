// BigFloat emits a zoomed-in PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

const prec = 128

func main() {
	const (
		x, y                   = 0.0, 1.0
		zoom                   = 1e37 // point at which artifacts appear
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
		y.Quo(y, h)
		y.Mul(y, r)
		y.Add(y, m)
		for px := 0; px < width; px++ {
			x := big.NewFloat(float64(px)).SetPrec(prec)
			w := big.NewFloat(width).SetPrec(prec)
			r := big.NewFloat(xmax - xmin).SetPrec(prec)
			m := big.NewFloat(xmin).SetPrec(prec)
			x.Quo(x, w)
			x.Mul(x, r)
			x.Add(x, m)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(x, y))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

var four = big.NewFloat(4).SetPrec(prec)

func mandelbrot(zr, zi *big.Float) color.Color {
	const iterations = 200
	const contrast = 15
	vr := big.NewFloat(0).SetPrec(prec)
	vi := big.NewFloat(0).SetPrec(prec)
	abs := big.NewFloat(0).SetPrec(prec)
	for n := 0; n < iterations; n++ {
		vr, vi = cmult(vr, vi, vr, vi)
		vr, vi = cadd(vr, vi, zr, zi)
		abs = cabsSquared(vr, vi)
		abs.Sub(abs, four)
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

func cabsSquared(xr, xi *big.Float) *big.Float {
	t1 := big.NewFloat(0).SetPrec(prec)
	t2 := big.NewFloat(0).SetPrec(prec)
	t1.Mul(xr, xr)
	t2.Mul(xi, xi)
	return t1.Add(t1, t2)
}
