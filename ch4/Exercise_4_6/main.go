//Exercise 4.6:
//Write an in-place function that squashes each run of adjacent Unicode
//spaces (see unicode.IsSpace) in a UTF-8-encoded []byte
//slice into a single ASCII space.

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	b := []byte("abc\r  \n\rdef")
	fmt.Printf("%q\n", string(squashSpaces(b)))
	fmt.Printf("%q\n", b)
}

func squashSpaces(data []byte) []byte {
	space := byte(' ')
	writeIndex := 0
	previousIsSpace := false

	for readIndex := 0; readIndex < len(data); {
		r, size := utf8.DecodeRune(data[readIndex:])

		if unicode.IsSpace(r) {
			if !previousIsSpace {
				data[writeIndex] = space
				writeIndex++
			}
			previousIsSpace = true
		} else {
			copy(data[writeIndex:], data[readIndex:readIndex+size])
			writeIndex += size
			previousIsSpace = false
		}

		readIndex += size
	}

	return data[:writeIndex]
}
