// “Exercise 8.5:
// Take an existing CPU-bound sequential program, such as the Mandelbrot
// program of Section 3.3 or the 3-D surface computation of
// Section 3.2, and execute its main loop in parallel using
// channels for communication.
//
// How much faster does it run on a multiprocessor machine?
//
// What is the optimal number of goroutines to use?”
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"runtime"
	"time"
)

type Pixel struct {
	X, Y  int
	Color color.Color
}

func main() {
	start := time.Now()
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	pixelCh := make(chan Pixel, width*height)

	// Determine the number of goroutines to use
	numGoroutines := runtime.NumCPU()

	// Create goroutines to calculate pixels in parallel
	for i := 0; i < numGoroutines; i++ {
		go func(start, end int) {
			for py := start; py < end; py++ {
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					y := float64(py)/height*(ymax-ymin) + ymin
					z := complex(x, y)
					pixelCh <- Pixel{px, py, mandelbrot(z)}
				}
			}
		}(i*height/numGoroutines, (i+1)*height/numGoroutines)
	}

	// Receive calculated pixels and update the image
	for i := 0; i < width*height; i++ {
		pixel := <-pixelCh
		img.Set(pixel.X, pixel.Y, pixel.Color)
	}

	f, err := os.Create("mandelbrot.png")
	if err != nil {
		panic(err)
	}
	png.Encode(f, img) // NOTE: ignoring errors
	f.Close()

	fmt.Println(time.Since(start))
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

//without go - 199.987166ms

//with go - 154.410542ms
