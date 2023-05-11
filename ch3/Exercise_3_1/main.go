//“Exercise 3.1:
//If the function f returns a non-finite float64 value,
//the SVG file will contain invalid <polygon>
//elements (although many SVG renderers handle this gracefully).
//
//Modify the program to skip invalid polygons.”

package main

import (
	"fmt"
	"log"
	"math"
	"os"
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
	out, err := os.Create("plot.svg")
	if err != nil {
		log.Fatalf("%s", err)
	}

	//fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
	//	"style='stroke: grey; fill: white; stroke-width: 0.7' "+
	//	"width='%d' height='%d'>", width, height)
	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	out.WriteString(s)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) ||
				math.IsNaN(dx) || math.IsNaN(dy) {
				log.Println(ax, ay, bx, by, cx, cy, dx, dy)
				continue
			} else if math.IsInf(ax, 0) || math.IsInf(ay, 0) || math.IsInf(bx, 0) || math.IsInf(by, 0) ||
				math.IsInf(cx, 0) || math.IsInf(cy, 0) || math.IsInf(dx, 0) || math.IsInf(dy, 0) {
				log.Println(ax, ay, bx, by, cx, cy, dx, dy)
				continue
			} else {
				s = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				out.WriteString(s)
			}
		}
	}
	out.WriteString("</svg>")
	defer out.Close()
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
