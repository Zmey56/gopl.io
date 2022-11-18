//Exercise 3.7: Another simple fractal uses Newton’s method to find complex
//solutions to a function such as z4−1 = 0.
//
//Shade each starting point by the number of iterations required
//to get close to one of the four roots.
//
//Color each point by the root it approaches.

package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
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
	f, err := os.Create("mandelbtot.png")
	if err != nil {
		panic(err)
	}

	png.Encode(f, img) // NOTE: ignoring errors

	f.Close()
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			if math.Abs(1-real(z)) < 1e-1 && math.Abs(imag(z)) < 1e-1 {
				return color.RGBA{255 - contrast*i, 0, 0, 0xff}
			} else if math.Abs(-1-real(z)) < 1e-1 && math.Abs(imag(z)) < 1e-1 {
				return color.RGBA{0, 255 - contrast*i, 0, 0xff}
			} else if math.Abs(real(z)) < 1e-1 && math.Abs(1-imag(z)) < 1e-1 {
				return color.RGBA{0, 0, 255 - contrast*i, 0xff}
			} else {
				return color.RGBA{0, 255 - contrast*i, 255 - contrast*i, 0xff}
			}
		}
	}
	return color.Black
}
