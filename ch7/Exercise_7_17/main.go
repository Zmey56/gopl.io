//“Exercise 7.17:
//Extend xmlselect so that elements may be selected not just by
//name, but by their attributes too, in the manner of CSS, so that, for
//instance, an element like <div id="page" class="wide">
//could be selected by a matching id or class as well
//as its name.”

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of start elements
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", getElementPath(stack), tok)
			}
		}
	}
}

// containsAll reports whether x contains elements that match the given selectors.
func containsAll(stack []xml.StartElement, selectors []string) bool {
	for _, selector := range selectors {
		found := false
		for _, elem := range stack {
			if matchElement(elem, selector) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// matchElement checks if the given element matches the selector.
func matchElement(elem xml.StartElement, selector string) bool {
	if selector == elem.Name.Local {
		return true
	}
	for _, attr := range elem.Attr {
		if selector == attr.Name.Local && selector == attr.Value {
			return true
		}
	}
	return false
}

// getElementPath returns the full path of the element based on the stack.
func getElementPath(stack []xml.StartElement) string {
	var path string
	for _, elem := range stack {
		path += "/" + elem.Name.Local
	}
	return path
}
