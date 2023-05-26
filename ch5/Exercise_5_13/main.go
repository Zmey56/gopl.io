//“Exercise 5.13:
//Modify crawl to make local copies of the pages it finds,
//creating directories as necessary.
//
//Don’t make copies of pages that come from a different domain.  For
//example, if the original page comes from golang.org, save all files from
//there, but exclude ones from vimeo.com.”

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "https://golang.org"
	domain := "golang.org"
	err := crawl(url, domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "crawl error: %v\n", err)
		os.Exit(1)
	}
}

func crawl(urlStr, domain string) error {
	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to crawl %s: %s", urlStr, resp.Status)
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	host := u.Hostname()

	if !strings.Contains(host, domain) {
		// Skip pages from different domains
		return nil
	}

	dir := host
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link := a.Val
					if isURLRelative(link) {
						absURL := resolveURL(urlStr, link)
						err := crawl(absURL, domain)
						if err != nil {
							fmt.Fprintf(os.Stderr, "crawl error: %v\n", err)
						}
					}
				}
			}
		}
	}

	forEachNode(doc, visitNode)

	filePath := filepath.Join(dir, "index.html")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Saved page: %s\n", filePath)

	return nil
}

func forEachNode(n *html.Node, f func(n *html.Node)) {
	if n == nil {
		return
	}

	f(n)

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, f)
	}
}

func isURLRelative(urlStr string) bool {
	return strings.HasPrefix(urlStr, "/")
}

func resolveURL(baseURL, urlStr string) string {
	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return ""
	}

	absURL := base.ResolveReference(rel).String()
	return absURL
}
