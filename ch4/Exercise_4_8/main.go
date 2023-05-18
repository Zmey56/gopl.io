//“Exercise 4.8:
//Modify charcount to count letters, digits, and so on
//in their Unicode categories, using functions like unicode.IsLetter.”

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if unicode.IsLetter(r) {
			counts["Letter"]++
		} else if unicode.IsDigit(r) {
			counts["Digit"]++
		} else if unicode.IsSpace(r) {
			counts["Space"]++
		} else if unicode.IsPunct(r) {
			counts["Punctuation"]++
		} else {
			counts["Other"]++
		}
	}

	fmt.Println("Character counts by category:")
	for category, count := range counts {
		fmt.Printf("%s: %d\n", category, count)
	}
}
