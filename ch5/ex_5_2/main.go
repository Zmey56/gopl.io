//“Exercise 5.2:
//Write a function to populate a mapping from element names—p, div,
//span, and so on—to the number of elements with that name
//in an HTML document tree.”

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

var count = make(map[string]int)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findElements: %v\n", err)
	}

	visit(doc)

	for t, c := range count {
		fmt.Println(t, c)
	}
}

func visit(n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
