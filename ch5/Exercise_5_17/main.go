//“Exercise 5.17:
//Write a variadic function ElementsByTagName that, given an HTML
//node tree and zero or more names, returns all the elements that
//match one of those names.  Here are two example calls:
//
//
//Click here to view code image
//func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
//
//images := ElementsByTagName(doc, "img")
//headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")”

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	var result []*html.Node
	traverse(doc, names, &result)
	return result
}

func traverse(node *html.Node, names []string, result *[]*html.Node) {
	if node.Type == html.ElementNode {
		for _, name := range names {
			if node.Data == name {
				*result = append(*result, node)
				break
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		traverse(child, names, result)
	}
}

func main() {
	url := "https://golang.org"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Problem", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Println(err)
	}

	// Example 1
	images := ElementsByTagName(doc, "img")
	fmt.Println("Images:")
	for _, img := range images {
		fmt.Println(img.Data)
	}

	// Example 2
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Println("Headings:")
	for _, heading := range headings {
		fmt.Println(heading.Data)
	}
}
