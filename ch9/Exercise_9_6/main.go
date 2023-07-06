//“Exercise 9.6:
//Measure how the performance of a compute-bound parallel program
//(see Exercise 8.5) varies with GOMAXPROCS.
//What is the optimal value on your computer?
//How many CPUs does your computer have?”

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

	// Determine the number of CPUs
	numCPUs := runtime.NumCPU()
	fmt.Println("Number of CPUs:", numCPUs)

	// Create goroutines to calculate pixels in parallel
	for i := 0; i < numCPUs; i++ {
		go func(start, end int) {
			for py := start; py < end; py++ {
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					y := float64(py)/height*(ymax-ymin) + ymin
					z := complex(x, y)
					pixelCh <- Pixel{px, py, mandelbrot(z)}
				}
			}
		}(i*height/numCPUs, (i+1)*height/numCPUs)
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

	fmt.Println("Execution time:", time.Since(start))
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

//(base) zmey56@Aleksandrs-MacBook-Air Exercise_9_6 % GOMAXPROCS=1 go run main.go
//Number of CPUs: 8
//Execution time: 259.049125ms
//(base) zmey56@Aleksandrs-MacBook-Air Exercise_9_6 % GOMAXPROCS=2 go run main.go
//Number of CPUs: 8
//Execution time: 185.778458ms
//(base) zmey56@Aleksandrs-MacBook-Air Exercise_9_6 % GOMAXPROCS=3 go run main.go
//Number of CPUs: 8
//Execution time: 164.747833ms
//(base) zmey56@Aleksandrs-MacBook-Air Exercise_9_6 % GOMAXPROCS=4 go run main.go
//Number of CPUs: 8
//Execution time: 154.450792ms
//(base) zmey56@Aleksandrs-MacBook-Air Exercise_9_6 % GOMAXPROCS=5 go run main.go
//Number of CPUs: 8
//Execution time: 155.763375ms
//(base) zmey56@Aleksandrs-MacBook-Air Exercise_9_6 % GOMAXPROCS=6 go run main.go
//Number of CPUs: 8
//Execution time: 158.121917ms
//(base) zmey56@Aleksandrs-MacBook-Air Exercise_9_6 % GOMAXPROCS=7 go run main.go
//Number of CPUs: 8
//Execution time: 155.441917ms
//(base) zmey56@Aleksandrs-MacBook-Air Exercise_9_6 % GOMAXPROCS=8 go run main.go
//Number of CPUs: 8
//Execution time: 157.179083ms
