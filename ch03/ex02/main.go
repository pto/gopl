// Ex02 outputs an SVG rendering of a 3-D surface function, ignoring NaN,
// coloring peaks red and valleys blue, and automatically adjusting the z-scale.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyRange       = 30.0                // axis range (max-min)
	xyScale       = width / 2 / xyRange // pixels per x or y unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var zMin, zMax, zScale float64 = zLimits() // z-axis limits

var f = eggBox // Function to graph

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			r, g, b := colorForCell(i, j)

			if !(math.IsNaN(ax) || math.IsNaN(ay) ||
				math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) ||
				math.IsNaN(dx) || math.IsNaN(dy)) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
					"style='stroke:rgb(%d,%d,%d)' />\n",
					ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
			}
		}
	}

	drawAxes()

	fmt.Println("</svg>\n")
}

// Find z minimum, maximum and scale
func zLimits() (zMin, zMax, zScale float64) {
	zMin = math.MaxFloat64
	zMax = -math.MaxFloat64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			z := f(cellToXY(i, j))
			if z < zMin {
				zMin = z
			}
			if z > zMax {
				zMax = z
			}
		}
	}

	zScale = height / 2 / math.Max(math.Abs(zMin), math.Abs(zMax))
	zScale /= 2 // Give some breathing room

	return zMin, zMax, zScale
}

// Find SVG display position (sx, sy) for corner (i, j)
func corner(i, j int) (sx, sy float64) {
	x, y := cellToXY(i, j)
	z := f(x, y)
	sx, sy = project(x, y, z)
	return sx, sy
}

// Color for cell (i, j)
func colorForCell(i, j int) (r, g, b int) {
	z := f(cellToXY(i, j))
	zRange := zMax - zMin
	r = int(255 * (z - zMin) / zRange)
	b = int(255 * (zMax - z) / zRange)
	if r < 0 {
		r = 0
	}
	if b < 0 {
		b = 0
	}
	return r, 0, b
}

// Find point (x, y) at corner (i, j).
func cellToXY(i, j int) (x, y float64) {
	x = xyRange * (float64(i)/cells - 0.5)
	y = xyRange * (float64(j)/cells - 0.5)
	return x, y
}

// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy).
func project(x, y, z float64) (sx, sy float64) {
	sx = width/2 + (x-y)*math.Cos(angle)*xyScale
	sy = height/2 + (x+y)*math.Sin(angle)*xyScale - z*zScale
	return sx, sy
}

// Draw x, y and z axes
func drawAxes() {
	minVal := -xyRange / 2
	maxVal := xyRange / 2
	drawLine(minVal, 0.0, 0.0, maxVal, 0.0, 0.0)
	drawLine(0.0, minVal, 0.0, 0.0, maxVal, 0.0)
	zLimit := height / 2 * zScale
	drawLine(0.0, 0.0, -zLimit, 0.0, 0.0, zLimit)
}

// Draw line (x1, y1, z1) to (x2, y2, z2)
func drawLine(x1, y1, z1, x2, y2, z2 float64) {
	ax, ay := project(x1, y1, z1)
	bx, by := project(x2, y2, z2)
	fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
		"style='stroke:gray;stroke-width:2' />\n",
		ax, ay, bx, by, bx, by, ax, ay)
}

// Functions to graph

func original(x, y float64) (z float64) {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func eggBox(x, y float64) (z float64) {
	const top, bottom = 1.5, 0
	z = math.Sin(x) + math.Cos(y)
	// Outer border
	if x > 5*math.Pi/2+0.5 || x < -3*math.Pi/2-0.5 ||
		y > 4*math.Pi+0.5 || y < -4*math.Pi-0.5 {
		return bottom
	}
	// Long edges
	if x > 5*math.Pi/2 || x < -3*math.Pi/2 {
		return top
	}
	// Short edges
	if y > 4*math.Pi || y < -4*math.Pi {
		return top
	}
	// Flattened tops and bottoms
	if z > top {
		return top
	} else if z < bottom {
		return bottom
	} else {
		return z
	}
}

func moguls(x, y float64) (z float64) {
	return math.Sin(y) - y/2
}

func saddle(x, y float64) (z float64) {
	z = x*x - y*y
	if x > xyRange/2-1 || x < -xyRange/2+1 {
		z = 0
	}
	return z
}
