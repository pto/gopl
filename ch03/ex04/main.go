// Ex04 is a web server that displays an SVG rendering of a 3-D surface
// function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	height, err := strconv.Atoi(r.Form.Get("height"))
	if err != nil {
		height = 320
	}
	width, err := strconv.Atoi(r.Form.Get("width"))
	if err != nil {
		width = 600
	}
	cells, err := strconv.Atoi(r.Form.Get("cells"))
	if err != nil {
		cells = 100
	}
	color, err := strconv.ParseInt(r.Form.Get("color"), 0, 0)
	if err != nil {
		color = 0
	}
	red := (color >> 16) & 0xFF
	green := (color >> 8) & 0xFF
	blue := color & 0xFF

	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:rgb(%d,%d,%d); fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", red, green, blue, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, height, width, cells)
			bx, by := corner(i, j, height, width, cells)
			cx, cy := corner(i, j+1, height, width, cells)
			dx, dy := corner(i+1, j+1, height, width, cells)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j, height, width, cells int) (float64, float64) {
	xyscale := float64(width) / 2 / xyrange
	zscale := float64(height) * 0.4

	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
