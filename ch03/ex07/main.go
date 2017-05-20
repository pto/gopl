// Ex07 emits a PNG image of the solutions to z^4-1=0.
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
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// Generate a color based on the number of iterations needed to come
// close to a root to z^4-1=0 using Newton's method.
//
// Improve guess x0 by calculating x1=x0-f(x0)/f'(x0)
//                                   =x0-(x0^4-1)/(4*x0^3)
//                                   =x0-x0/4+1/(4*x0^3)
func newton(z complex128) color.Color {
	const (
		limit    = 50
		contrast = 20
		epsilon  = 1e-6
	)

	x0 := z
	for n := uint8(0); n < limit; n++ {
		x1 := x0 - x0/4 + 1/(4*x0*x0*x0)
		if cmplx.Abs(x1*x1*x1*x1-1) < epsilon {
			return color.YCbCr{128, 255 - contrast*n, contrast*n - 255}
		}
		x0 = x1
	}
	return color.Black
}
