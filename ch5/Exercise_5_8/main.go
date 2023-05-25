//“Exercise 5.8:
//Modify forEachNode so that the pre and post
//functions return a boolean result indicating whether to continue the
//traversal.
//
//Use it to write a function ElementByID with the following signature
//that finds the first HTML element with the specified id attribute.
//The function should stop the traversal as soon as a match is found.
//
//func ElementByID(doc *html.Node, id string) *html.Node”

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if !pre(n) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := forEachNode(c, pre, post)
		if result != nil {
			return result
		}
	}

	if post != nil {
		if !post(n) {
			return n
		}
	}

	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var result *html.Node

	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					result = n
					return false // Stop traversal
				}
			}
		}
		return true // Continue traversal
	}

	post := func(n *html.Node) bool {
		return true // Continue traversal
	}

	forEachNode(doc, pre, post)

	return result
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <url> <id>")
		return
	}

	url := os.Args[1]
	id := os.Args[2]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to retrieve HTML: %v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	node := ElementByID(doc, id)
	if node != nil {
		fmt.Printf("Found element with ID '%s': %v\n", id, node)
	} else {
		fmt.Printf("Element with ID '%s' not found\n", id)
	}
}
