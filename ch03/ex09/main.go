// Ex09 is a web server for Mandelbrot images.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var x, y, zoom float64
	var err error
	if err := r.ParseForm(); err != nil {
		// ignore errors
	}
	x, err = strconv.ParseFloat(r.Form.Get("x"), 64)
	if err != nil {
		x = 0.0
	}
	y, err = strconv.ParseFloat(r.Form.Get("y"), 64)
	if err != nil {
		y = -1.0
	}
	zoom, err = strconv.ParseFloat(r.Form.Get("zoom"), 64)
	if err != nil {
		zoom = 1
	}
	writeImage(w, x, y, zoom)
}

func writeImage(w io.Writer, x, y, zoom float64) {
	const width, height = 1024, 1024
	var radius = 2.0 / zoom
	var xmin, ymin, xmax, ymax = x - radius, y - radius, x + radius, y + radius

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
