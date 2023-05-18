//Exercise 4.9:
//Write a program wordfreq to report the frequency of each word
//in an input text file. Call input.Split(bufio.ScanWords) before the first
//call to Scan to break the input into words instead of lines.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("example.txt")
	if err != nil {
		log.Println("Can't read a file", err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)
	wordFreq := make(map[string]int)

	for scan.Scan() {
		word := scan.Text()
		wordFreq[word]++
	}

	if scan.Err() != nil {
		fmt.Println("Error scanning file:", scan.Err())
		return
	}

	fmt.Println("Word Frequencies:")
	for word, freq := range wordFreq {
		fmt.Printf("%s: %d\n", word, freq)
	}
}
