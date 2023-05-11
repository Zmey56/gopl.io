//“Exercise 3.4:
//Following the approach of the Lissajous example in Section 1.7,
//construct a web server that computes surfaces and writes SVG data to
//the client.
//
//The server must set the Content-Type header like this:
//
//w.Header().Set("Content-Type", "image/svg+xml")”
//

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
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
	http.HandleFunc("/", createCorner)
	log.Fatal(http.ListenAndServe("localhost:8070", nil))
}

func createCorner(w http.ResponseWriter, r *http.Request) {
	// Set the content type to "image/svg+xml"
	w.Header().Set("Content-Type", "image/svg+xml")

	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	io.WriteString(w, s)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) ||
				math.IsNaN(dx) || math.IsNaN(dy) {
				log.Println(ax, ay, bx, by, cx, cy, dx, dy)
				continue
			} else if math.IsInf(ax, 0) || math.IsInf(ay, 0) || math.IsInf(bx, 0) || math.IsInf(by, 0) ||
				math.IsInf(cx, 0) || math.IsInf(cy, 0) || math.IsInf(dx, 0) || math.IsInf(dy, 0) {
				log.Println(ax, ay, bx, by, cx, cy, dx, dy)
				continue
			} else {
				if z > 0 {
					s = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: #ff0000'/>\n",
						ax, ay, bx, by, cx, cy, dx, dy)
				} else {
					s = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: #0000ff'/>\n",
						ax, ay, bx, by, cx, cy, dx, dy)
				}
				io.WriteString(w, s)
			}
		}
	}
	io.WriteString(w, "</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
