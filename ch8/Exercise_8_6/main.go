//“Exercise 8.6:
//Add depth-limiting to the concurrent crawler.
//
//That is, if the user sets -depth=3,
//then only URLs reachable by at most three links will be fetched.”

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

var depthLimit int

func main() {
	flag.IntVar(&depthLimit, "depth", 3, "depth limit for fetching URLs")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run main.go [flags] [url]")
		return
	}

	startURL := flag.Args()[0]
	_, err := url.Parse(startURL)
	if err != nil {
		log.Fatal(err)
	}

	seen := make(map[string]bool)
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- []string{startURL} }()

	for i := 0; i < 20; i++ {
		go func() {
			for links := range worklist {
				for _, link := range links {
					if !seen[link] {
						seen[link] = true
						go func(link string) {
							if depthLimit > 0 {
								foundLinks, err := crawl(link)
								if err != nil {
									log.Println(err)
								} else {
									worklist <- foundLinks
								}
							}
						}(link)
					}
				}
			}
		}()
	}

	go func() {
		for link := range unseenLinks {
			worklist <- []string{link}
		}
	}()

	<-worklist
}

func crawl(url string) ([]string, error) {
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
						continue
					}
					if strings.HasPrefix(link.String(), "http") {
						links = append(links, link.String())
					}
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
