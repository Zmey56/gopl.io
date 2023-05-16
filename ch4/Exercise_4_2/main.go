//Exercise 4.2:
//Write a program that prints the SHA256 hash of its standard input
//by default but supports a command-line flag to print the
//SHA384 or SHA512 hash instead.”

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Parse command-line flags
	algorithm := flag.String("algo", "sha256", "Hash algorithm: sha256, sha384, or sha512")
	flag.Parse()

	// Read standard input
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Calculate hash based on the chosen algorithm
	switch *algorithm {
	case "sha256":
		hash := sha256.Sum256(data)
		fmt.Printf("SHA256 hash: %x\n", hash)
	case "sha384":
		hash := sha512.Sum384(data)
		fmt.Printf("SHA384 hash: %x\n", hash)
	case "sha512":
		hash := sha512.Sum512(data)
		fmt.Printf("SHA512 hash: %x\n", hash)
	default:
		fmt.Println("Invalid algorithm specified.")
		os.Exit(1)
	}
}
