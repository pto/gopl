// BigRat emits a zoomed-in PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

func main() {
	const (
		cx, cy        = 0.0, 1.0
		zoom          = 1e14
		radius        = 2.0 / zoom
		xmin, ymin    = cx - radius, cy - radius
		xmax, ymax    = cx + radius, cy + radius
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	y := big.NewRat(1, 1)
	h := big.NewRat(1, 1)
	r := big.NewRat(1, 1)
	m := big.NewRat(1, 1)
	x := big.NewRat(1, 1)
	w := big.NewRat(1, 1)
	for py := 0; py < height; py++ {
		y.SetFloat64(float64(py))
		h.SetFloat64(height)
		r.SetFloat64(ymax - ymin)
		m.SetFloat64(ymin)
		y.Quo(y, h)
		y.Mul(y, r)
		y.Add(y, m)
		for px := 0; px < width; px++ {
			x.SetFloat64(float64(px))
			w.SetFloat64(width)
			r.SetFloat64(xmax - xmin)
			m.SetFloat64(xmin)
			x.Quo(x, w)
			x.Mul(x, r)
			x.Add(x, m)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(x, y))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

var four = big.NewRat(4, 1)

func mandelbrot(zr, zi *big.Rat) color.Color {
	const iterations = 200
	const contrast = 15
	vr := big.NewRat(0, 1)
	vi := big.NewRat(0, 1)
	abs := big.NewRat(0, 1)
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
