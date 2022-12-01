//Exercise 4.2:
//Write a program that prints the SHA256 hash of its standard input
//by default but supports a command-line flag to print the
//SHA384 or SHA512 hash instead.

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	algPtr := flag.String("alg", "sha256", "enter one of sha256, sha384 or sha512")
	strPtr := flag.String("str", "", "enter a string to calculate the hash of")

	flag.Parse()

	if *strPtr == "" {
		fmt.Println("no string entered")
		return
	}

	switch *algPtr {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(*strPtr)))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(*strPtr)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(*strPtr)))
	}
}
