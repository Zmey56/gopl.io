//“Exercise 8.7: Write a concurrent program that creates a local mirror of a web site,
//fetching each reachable page and writing it to a directory on the
//local disk.
//Only pages within the original domain (for instance, golang.org) should
//be fetched.
//URLs within mirrored pages should be altered as needed so that they
//refer to the mirrored page, not the original.”

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// Mirror represents the website mirror
type Mirror struct {
	domain     string
	outputDir  string
	client     *http.Client
	waitGroup  sync.WaitGroup
	urlsToSave chan string
}

// NewMirror creates a new Mirror instance
func NewMirror(domain, outputDir string) *Mirror {
	return &Mirror{
		domain:     domain,
		outputDir:  outputDir,
		client:     &http.Client{},
		urlsToSave: make(chan string),
	}
}

// Run starts the mirroring process
func (m *Mirror) Run() {
	m.waitGroup.Add(1)
	go m.saveURLs()
	m.mirrorPage(m.domain)
	m.waitGroup.Wait()
	close(m.urlsToSave)
}

// mirrorPage fetches the given URL and recursively mirrors its content
func (m *Mirror) mirrorPage(urlStr string) {
	defer m.waitGroup.Done()

	resp, err := m.client.Get(urlStr)
	if err != nil {
		fmt.Println("Failed to fetch", urlStr, ":", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body for", urlStr, ":", err)
		return
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Failed to parse URL", urlStr, ":", err)
		return
	}

	// Save the HTML page
	savePath := m.getSavePath(baseURL)
	err = os.MkdirAll(filepath.Dir(savePath), 0755)
	if err != nil {
		fmt.Println("Failed to create directory for", urlStr, ":", err)
		return
	}

	err = os.WriteFile(savePath, body, 0644)
	if err != nil {
		fmt.Println("Failed to save", urlStr, ":", err)
		return
	}

	// Parse the HTML to find links
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("Failed to parse HTML for", urlStr, ":", err)
		return
	}

	m.extractLinks(doc, baseURL)
}

// extractLinks finds all anchor tags in the HTML and adds their URLs to be mirrored
func (m *Mirror) extractLinks(n *html.Node, baseURL *url.URL) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for i := range n.Attr {
			if n.Attr[i].Key == "href" {
				linkURL, err := baseURL.Parse(n.Attr[i].Val)
				if err == nil && linkURL.Host == baseURL.Host {
					m.waitGroup.Add(1)
					m.urlsToSave <- linkURL.String()
					break
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		m.extractLinks(c, baseURL)
	}
}

// saveURLs saves the URLs received in parallel
func (m *Mirror) saveURLs() {
	for urlStr := range m.urlsToSave {
		m.mirrorPage(urlStr)
	}
}

// getSavePath returns the path to save the mirrored page
func (m *Mirror) getSavePath(baseURL *url.URL) string {
	filename := baseURL.Host + baseURL.Path
	if filename == "" || filename == "/" {
		filename = "index.html"
	} else if filename[len(filename)-1:] == "/" {
		filename += "index.html"
	}
	filename = strings.ReplaceAll(filename, "/", "_") // Replace slashes with underscores

	return filepath.Join(m.outputDir, filename)
}

func main() {
	mirror := NewMirror("https://golang.org", "mirror")
	mirror.Run()
	fmt.Println("Mirroring complete!")
}
