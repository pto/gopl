// Ex06 emits a PNG image of the Mandelbrot fractal in color, using
// supersampling.
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
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	xdelta := float64(1) / width / 4 * (xmax - xmin)
	ydelta := float64(1) / height / 4 * (ymax - ymin)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			avg := (float64(mandelbrot(complex(x-xdelta, y-ydelta))) +
				float64(mandelbrot(complex(x+xdelta, y-ydelta))) +
				float64(mandelbrot(complex(x-xdelta, y+ydelta))) +
				float64(mandelbrot(complex(x+xdelta, y+ydelta)))) / 4
			// Image point (px, py) represents complex value z.
			img.Set(px, py, iterationColor(avg))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) int {
	const iterations = 200

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return n
		}
	}
	return -1
}

func iterationColor(iter float64) color.Color {
	const contrast = 15
	if iter < 0 {
		return color.Black
	}
	return color.YCbCr{128, 255 - uint8(contrast*iter),
		uint8(contrast*iter) - 255}
}
