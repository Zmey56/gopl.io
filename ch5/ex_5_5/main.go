//“Exercise 5.5:
//Implement countWordsAndImages. (See Exercise 4.9 for word-splitting.)”
//

package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "ex 5.5: %v\n", err)
			continue
		}
		fmt.Printf("%s: Words=%d, Images=%d\n", url, words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	// If the current node is an image element, increament the imgage count
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	// If the current node is a TextNode, split the Data string and count the words
	if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			scanner.Text()
			words++
		}
	}

	// Aggregate the count of words and images from the current node with the count from all of it's children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}

	// Return the count of words and images
	return
}
