//“Exercise 1.12:
//Modify the Lissajous server to read parameter values from the URL.
//For example, you might arrange it so that a URL like
//http://localhost:8000/?cycles=20
//sets the number of cycles to 20 instead of the default 5.
//
//Use the strconv.Atoi function to convert the string
//parameter into an integer.

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"strconv"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
)

//!+main

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	cycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))

	if err != nil {
		fmt.Fprint(w, "Error, query parameter \"cycles\" is required an must be an integer")
		return
	}

	log.Println(cycles)
	lissajous(w, cycles)
	//lissajous(w)

}

func lissajous(out io.Writer, c int) {
	const (
		//cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	cycles := 5

	if c != 5 {
		cycles = c
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
