//Exercise 10.1:
//Extend the jpeg program so that it converts any supported input
//format to any output format, using image.Decode to detect the
//input format and a flag to select the output format.

package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"os"
	"strings"
)

func main() {
	outputFormat := flag.String("format", "jpeg", "Output format (jpeg/png)")

	flag.Parse()

	if err := toJPEG(os.Stdin, os.Stdout, *outputFormat); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}

	outputFormat := strings.ToLower(format)
	switch outputFormat {
	case "jpeg":
		fmt.Fprintln(os.Stderr, "Input format =", kind)
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		fmt.Fprintln(os.Stderr, "Input format =", kind)
		return png.Encode(out, img)
	default:
		fmt.Fprintln(os.Stderr, "Input format =", kind)
		return fmt.Errorf("unsupported output format: %s", outputFormat)
	}
}

//go run main.go -format=png < input.jpg > output.png
