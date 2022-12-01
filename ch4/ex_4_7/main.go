//“Exercise 4.7:
//Modify reverse to reverse the characters of a
//[]byte slice that represents a UTF-8-encoded string, in place.
//Can you do it without allocating new memory?”

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout // modified during testing

func main() {
	s := "TestAndTest"
	b := []byte(s)
	fmt.Fprintf(stdout, "%p, %s\n", b, b)
	b = reverse(b)
	fmt.Fprintf(stdout, "%p, %s\n", b, b)
}

// reverse reverses a slice of ints in place.
func reverse(b []byte) []byte {
	runes := bytes.Runes(b)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return []byte(string(runes))
}
