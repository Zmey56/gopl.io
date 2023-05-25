// “Exercise 5.7:
// Develop startElement and endElement into a general HTML pretty-printer.
// Print comment nodes, text nodes, and the attributes of each element
// (<a href='...'>).  Use short forms like <img/>
// instead of <img></img> when an element has no children.
// Write a test to ensure that the output can be parsed successfully.
// (See Chapter 11.)”

package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
)

func prettyPrintHTML(n *html.Node, w io.Writer) {
	if n.Type == html.ElementNode {
		fmt.Fprintf(w, "<%s", n.Data)
		for _, attr := range n.Attr {
			fmt.Fprintf(w, " %s='%s'", attr.Key, attr.Val)
		}
		if n.FirstChild == nil {
			fmt.Fprintf(w, "/>")
		} else {
			fmt.Fprintf(w, ">")
		}
	} else if n.Type == html.TextNode {
		fmt.Fprint(w, n.Data)
	} else if n.Type == html.CommentNode {
		fmt.Fprintf(w, "<!--%s-->", n.Data)
	}

	if n.FirstChild != nil {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			prettyPrintHTML(c, w)
		}
	}

	if n.Type == html.ElementNode && n.FirstChild != nil {
		fmt.Fprintf(w, "</%s>", n.Data)
	}
}

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

	var buf bytes.Buffer
	prettyPrintHTML(doc, &buf)

	result := buf.String()
	//
	log.Println(result)

	return nil
}
