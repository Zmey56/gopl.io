//Exercise 5.2:
//Write a function to populate a mapping from element names—p, div,
//span, and so on—to the number of elements with that name
//in an HTML document tree.”

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	elemCount := make(map[string]int)
	countElementDoc(elemCount, doc)

	for element, count := range elemCount {
		fmt.Printf("%s: %d\n", element, count)
	}
}

func countElementDoc(elemCount map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		elemCount[n.Data]++

	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countElementDoc(elemCount, c)
	}

}
