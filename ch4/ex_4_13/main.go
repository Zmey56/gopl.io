//“Exercise 4.13:
//The JSON-based web service of the Open Movie Database
//lets you search https://omdbapi.com/ for a movie by name and
//download its poster image.
//
//
//Write a tool poster that downloads the poster image for the
//movie named on the command line.”

package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const apikey = "8879001f"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("usage: go run main.go tt1285016") // The tt12.. arg is a valid IMDB ID
		os.Exit(1)
	}

	url := "http://img.omdbapi.com/?i=" + args[0] + "&apikey=" + apikey
	poster(url)
}

func poster(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	contentTypeRaw := resp.Header["Content-Type"][0]
	var f *os.File
	var contentType string

	if !strings.Contains(contentTypeRaw, "image") {
		fmt.Println("wrong url")
		os.Exit(1)
	} else if strings.Contains(contentTypeRaw, "jpeg") {
		contentType = "jpeg"
	} else if strings.Contains(contentTypeRaw, "jpg") {
		contentType = "jpg"
	} else if strings.Contains(contentTypeRaw, "png") {
		contentType = "png"
	}

	f, err = os.Create("poster." + contentType)

	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	img, err := jpeg.Decode(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	jpeg.Encode(f, img, nil)
}
