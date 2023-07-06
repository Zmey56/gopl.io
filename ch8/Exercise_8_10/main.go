// “Exercise 8.10:
// HTTP requests may be cancelled by closing the optional Cancel
// channel in the http.Request struct.
//
// Modify the web crawler of Section 8.6 to
// support cancellation.
//
// Hint: the http.Get convenience function does not give you an
// opportunity to customize a Request.
//
// Instead, create the request using http.NewRequest, set
// its Cancel field, then perform the request by calling
// http.DefaultClient.Do(req).”
package main

import (
	"context"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
)

func crawl(ctx context.Context, url string) []string {
	fmt.Println(url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
		return []string{}
	}
	defer resp.Body.Close()

	links, err := extractLinks(resp.Body)
	if err != nil {
		log.Print(err)
	}
	return links
}

func extractLinks(r io.Reader) ([]string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
					break
				}
			}
		}
	}

	forEachNode(doc, visitNode, nil)
	return links, nil
}

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

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(ctx, link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

	cancel()
}
