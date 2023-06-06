//“Exercise 7.5:
//The LimitReader function in the io package accepts an
//io.Reader r and a number of bytes n, and
//returns another Reader that reads from r
//but reports an end-of-file condition after n bytes. Implement it.
//	func LimitReader(r io.Reader, n int64) io.Reader”

package main

import (
	"bytes"
	"fmt"
	"io"
)

type limitReader struct {
	r     io.Reader
	limit int64
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	if lr.limit <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > lr.limit {
		p = p[:lr.limit]
	}

	n, err = lr.r.Read(p)
	lr.limit -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{
		r:     r,
		limit: n,
	}
}

func main() {
	// Example usage
	input := []byte("This is a test input string.")
	reader := LimitReader(bytes.NewReader(input), 10)

	buf := make([]byte, 5)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error:", err)
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
