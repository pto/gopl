// Channel emits a zoomed-in PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		x, y                   = 0.0, 1.0
		zoom                   = 1e14 // point at which artifacts appear
		radius                 = 2.0 / zoom
		xmin, ymin, xmax, ymax = x - radius, y - radius, x + radius, y + radius
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	type result struct {
		py     int
		colors [width]color.Color
	}
	ch := make(chan result, height)

	for py := 0; py < height; py++ {
		go func(py int) {
			var colors [width]color.Color
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				colors[px] = mandelbrot(z)
			}
			ch <- result{py, colors}
		}(py)
	}

	for py := 0; py < height; py++ {
		r := <-ch
		for px := 0; px < width; px++ {
			img.Set(px, r.py, r.colors[px])
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
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
