// “Exercise 3.5:
//Implement a full-color Mandelbrot set using the function image.NewRGBA
//and the type color.RGBA or color.YCbCr.”

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
			img.Set(px, py, mandelbrot(z))
		}
	}
	f, err := os.Create("mandelbtot.png")
	if err != nil {
		panic(err)
	}

	png.Encode(f, img) // NOTE: ignoring errors

	f.Close()
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		blue := uint8(255 - contrast*n)
		red := uint8(255 - blue)
		green := uint8(123)

		if cmplx.Abs(v) > 2 {
			return color.RGBA{red, green, blue, 255}
		}
	}
	return color.Black
}
