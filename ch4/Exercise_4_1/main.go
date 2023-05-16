//“Exercise 4.1:
//Write a function that counts the number of bits that are different in
//two SHA256 hashes.

package main

import (
	"crypto/sha256"
	"log"
)

func main() {
	h := sha256.Sum256([]byte("Hello world\n"))
	m := sha256.Sum256([]byte("Hello Sasha\n"))

	log.Println("H", h)
	log.Println("M", m)

	log.Println("Number of different bits:", bitsDifference(h, m))
}

func bitsDifference(hash1, hash2 [32]byte) int {
	count := 0
	for i := 0; i < len(hash1); i++ {
		log.Println("hash1[i]", hash1[i])
		log.Println("hash2[i]", hash2[i])
		xor := hash1[i] ^ hash2[i]
		log.Println("First XOR", xor)
		for xor != 0 {
			count++
			xor &= xor - 1
			log.Println("XOR in for", xor)
		}
	}
	return count
}
