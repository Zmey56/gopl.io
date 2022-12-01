//“Exercise 5.3:
//Write a function to print the contents of all text nodes in an HTML
//document tree.
//
//Do not descend into <script> or <style> elements,
//since their contents are not visible in a web browser.”

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse err: %v\n", err)
	}

	for _, texts := range visit(nil, doc) {
		fmt.Println(texts)
	}
}

func visit(texts []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data != "script" && n.Data != "style" {
		for _, text := range n.Attr {
			texts = append(texts, text.Val)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = visit(texts, c)
	}

	return texts
}
