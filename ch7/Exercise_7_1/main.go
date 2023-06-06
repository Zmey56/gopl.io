//“Exercise 7.1:
//Using the ideas from ByteCounter, implement counters for
//words and for lines. You will find bufio.ScanWords useful.”

package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(p))
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}

	*c += WordCounter(count)

	return count, nil
}

type LineCounter int

func (c *LineCounter) Write(p string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(p))
	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		count++
	}

	*c += LineCounter(count)

	return count, nil
}

func main() {

	var wc WordCounter
	wc.Write("Hello, world! This is a sentence.")
	fmt.Println(wc) // Output: 7

	var lc LineCounter
	lc.Write("Line 1\nLine 2\nLine 3")
	fmt.Println(lc) // Output: 3

}
