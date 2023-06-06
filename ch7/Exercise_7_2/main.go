//Write a function CountingWriter with the signature below that, given an io.Writer,
//returns a new Writer that wraps the original, and a pointer to an int64 variable that at any
//moment contains the number of bytes written to the new Writer.
//
//func CountingWriter(w io.Writer) (io.Writer, *int64)”

package main

import (
	"bytes"
	"fmt"
	"io"
)

type countingWriter struct {
	writer io.Writer
	count  *int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.writer.Write(p)
	*cw.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	count := int64(0)
	cw := &countingWriter{writer: w, count: &count}
	return cw, cw.count
}

func main() {
	buffer := &bytes.Buffer{}
	writer, count := CountingWriter(buffer)

	// Write some data
	writer.Write([]byte("Hello, world!"))
	fmt.Println(*count) // Output: 13

	// Write more data
	writer.Write([]byte(" This is additional text."))
	fmt.Println(*count) // Output: 38 (Total count of bytes written)

	// The data in the buffer will contain the combined text
	fmt.Println(buffer.String()) // Output: Hello, world! This is additional text.
}
