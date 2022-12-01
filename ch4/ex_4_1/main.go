// Exercise 4.1:
// Write a function that counts the number of bits that are different in
// two SHA256 hashes.

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(PopCountSHA256(c1, c2))

	c3 := sha256.Sum256([]byte("1"))
	c4 := sha256.Sum256([]byte("1"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c3, c4, c3 == c4, c3)
	fmt.Println(PopCountSHA256(c3, c4))

	c5 := sha256.Sum256([]byte("1"))
	c6 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c5, c6, c5 == c6, c5)
	fmt.Println(PopCountSHA256(c5, c6))

}

func PopCountSHA256(a, b [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		count = count + int(a[i]^b[i])
	}
	return count

}
