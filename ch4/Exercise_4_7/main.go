//“Exercise 4.7:
//Modify reverse to reverse the characters of a
//[]byte slice that represents a UTF-8-encoded string, in place.
//Can you do it without allocating new memory?”

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, world!"
	data := []byte(str)
	reverse(data)
	fmt.Println(string(data))

	bytes := []byte("Hello, world!")
	reverseBytes(bytes[7:12])
	fmt.Println(string(bytes)) // выведет "Hello, dlrow!"
}

func reverse(data []byte) {
	for i := 0; i < len(data); {
		_, size := utf8.DecodeRune(data[i:])
		reverseBytes(data[i : i+size])
		i += size
	}

	reverseBytes(data)
}

func reverseBytes(bytes []byte) {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}
