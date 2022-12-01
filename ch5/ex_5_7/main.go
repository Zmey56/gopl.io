// “Exercise 5.7:
// Develop startElement and endElement into a general HTML pretty-printer.
//
// Print comment nodes, text nodes, and the attributes of each element
// (<a href='...'>).  Use short forms like <img/>
// instead of <img></img> when an element has no children.
//
// Write a test to ensure that the output can be parsed successfully.
// (See Chapter 11.)”
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s>\n", depth*2, "", getEA(n))

		if !isVoidElement(n.Data) {
			depth++
		}

	case html.TextNode:
		for _, line := range strings.Split(n.Data, "\n") {
			if line != "" {
				fmt.Printf("%*s%v\n", depth*2, "", line)
			}
		}
	case html.CommentNode:
		fmt.Printf("%*s<!--%v-->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode && !isVoidElement(n.Data) {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func isVoidElement(elm string) bool {
	// http://w3c.github.io/html/syntax.html#void-elements
	voids := []string{
		"area", "base", "br", "col", "embed", "hr", "img", "input", "link",
		"meta", "param", "source", "track", "wbr",
	}

	for _, v := range voids {
		if elm == v {
			return true
		}
	}
	return false
}

func getEA(n *html.Node) string {
	output := []string{n.Data}
	for _, a := range n.Attr {
		if a.Val != "" {
			output = append(output, fmt.Sprintf("%s=\"%s\"", a.Key, a.Val))
		} else {
			output = append(output, fmt.Sprintf("%s", a.Key))
		}
	}
	if isVoidElement(n.Data) {
		output = append(output, "/")
	}

	return strings.Join(output, " ")
}
