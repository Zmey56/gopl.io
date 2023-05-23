//“Exercise 5.5:
//Implement countWordsAndImages.
//(See Exercise 4.9 for word-splitting.)”

package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {
		words, images, err := countWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}

		fmt.Println("Words:" + strconv.Itoa(words) + " Images:" + strconv.Itoa(images))
	}
}

func countWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	elemCount := make(map[string]int)

	getValues(elemCount, doc)

	return elemCount["Words"], elemCount["Images"], nil
}

func getValues(elemCount map[string]int, n *html.Node) {

	switch n.Type {
	case html.TextNode:
		if n.Parent.Data != "script" && n.Parent.Data != "style" {
			elemCount["Words"] = elemCount["Words"] + wordCount(n.Data)
		}
	case html.ElementNode:
		if n.Data == "img" {
			elemCount["Images"] = elemCount["Images"] + 1
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getValues(elemCount, c)
	}
}

func wordCount(s string) int {
	n := 0
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n++
	}
	return n
}
