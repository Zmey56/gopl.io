//“Exercise 5.4:
//Extend the visit function so that it extracts other kinds of
//links from the document, such as images, scripts, and style sheets.”

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		var key string
		switch n.Data {
		case "a", "link":
			key = "href"
		case "script", "img":
			key = "src"
		}

		for _, a := range n.Attr {
			if a.Key == key {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
