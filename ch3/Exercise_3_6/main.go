//“Exercise 3.6:
//Supersampling is a technique to reduce the effect of pixelation by
//computing the color value at several points within each pixel and
//taking the average.  The simplest method is to divide each pixel into
//four “subpixels.”  Implement it”

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
		for px := 0; px < width; px++ {
			subpixelColors := make([]color.Color, 0)
			for sub_py := 0; sub_py <= 1; sub_py++ {
				for sub_px := 0; sub_px <= 1; sub_px++ {
					y := float64(py*2+sub_py)/(height*2)*(ymax-ymin) + ymin
					x := float64(px*2+sub_px)/(width*2)*(xmax-xmin) + xmin
					z := complex(x, y)
					color := mandelbrot(z)
					subpixelColors = append(subpixelColors, color)
				}

			}
			img.Set(px, py, avgColor(subpixelColors))
		}
	}
	f, err := os.Create("mandelbtot.png")
	if err != nil {
		panic(err)
	}

	png.Encode(f, img) // NOTE: ignoring errors

	f.Close()
}

func avgColor(colors []color.Color) color.Color {
	var average_red uint32 = 0
	var average_green uint32 = 0
	var average_blue uint32 = 0
	var average_alpha uint32 = 0

	for _, color := range colors {
		red, green, blue, alpha := color.RGBA()
		// Because for some reasons, the go color library returns
		// (color.COMPONENT | (color.COMPONENT << 8)), we truncate the value to uint8.
		average_red += uint32(uint8(red))
		average_green += uint32(uint8(green))
		average_blue += uint32(uint8(blue))
		average_alpha += uint32(uint8(alpha))
	}
	return color.RGBA{
		uint8(average_red / uint32(len(colors))),
		uint8(average_green / uint32(len(colors))),
		uint8(average_blue / uint32(len(colors))),
		uint8(average_alpha / uint32(len(colors))),
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0, contrast * n, 0, 0xff}
		}
	}
	return color.Black
}
