// Ex03 computes an SVG rendering of a 3-D surface function, coloring
// peaks red and valleys blue.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	zmax, zmin := limits()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			r, g, b := color(i, j, zmax, zmin)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
				"style='stroke:rgb(%d,%d,%d)'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
		}
	}
	fmt.Println("</svg>")
}

// Map 2D coordinates to 3D surface
func mapping(i, j int) (x, y, z float64) {
	x = xyrange * (float64(i)/cells - 0.5)
	y = xyrange * (float64(j)/cells - 0.5)
	z = f(x, y)
	return x, y, z
}

func limits() (zmax, zmin float64) {
	zmax = -math.MaxFloat64
	zmin = math.MaxFloat64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, z := mapping(i, j)
			if !math.IsNaN(z) {
				if z > zmax {
					zmax = z
				}
				if z < zmin {
					zmin = z
				}
			}
		}
	}
	return zmax, zmin
}

func corner(i, j int) (sx, sy float64) {
	x, y, z := mapping(i, j)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func color(i, j int, zmax, zmin float64) (r, g, b int) {
	_, _, z := mapping(i, j)

	zrange := zmax - zmin
	r = int(255 * (z - zmin) / zrange)
	g = 0
	b = int(255 * (zmax - z) / zrange)
	return r, g, b
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
