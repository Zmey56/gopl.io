//Exercise 5.3:
//Write a function to print the contents of all text nodes in an HTML
//document tree.
//
//Do not descend into <script> or <style> elements,
//since their contents are not visible in a web browser.

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

	printTextNode(doc)

}

func printTextNode(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	if n.Data != "script" && n.Data != "style" {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			printTextNode(c)
		}
	}
}
