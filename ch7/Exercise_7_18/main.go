//Exercise 7.18 Using the token-based decoder API, write a program that will read an arbitrary XML document
//and  construct a tree of generic nodes that represents it. Nodes are of two kinds: CharData nodes represent
//text strings, and Element nodes represent named elements and their attributes. Each element node has a slice
//of child nodes.
//
//You may find the following declarations helpful.
//
//import "encoding/xml"
//type Node interface{} // CharData or *Element
//type CharData string
//type Element struct {
//    Type     xml.Name
//    Attr     []xml.Attr
//    Children []Node
//}

package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func buildTree(data []byte) (Node, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))

	var stack []*Element
	var root *Element

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch token := token.(type) {
		case xml.StartElement:
			element := &Element{
				Type: token.Name,
				Attr: token.Attr,
			}
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, element)
			} else {
				root = element
			}
			stack = append(stack, element)
		case xml.EndElement:
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			} else {
				stack = nil
			}
		case xml.CharData:
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, CharData(token))
			}
		}
	}

	return root, nil
}

func main() {
	url := "http://www.w3.org/TR/2006/REC-xml11-20060816"

	response, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	tree, err := buildTree(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Print the tree for demonstration purposes
	printTree(tree, 0)
}

func printTree(node Node, indent int) {
	switch node := node.(type) {
	case *Element:
		printIndent(indent)
		fmt.Printf("Element: %s\n", node.Type.Local)
		for _, attr := range node.Attr {
			printIndent(indent + 1)
			fmt.Printf("Attr: %s=%s\n", attr.Name.Local, attr.Value)
		}
		for _, child := range node.Children {
			printTree(child, indent+1)
		}
	case CharData:
		printIndent(indent)
		fmt.Printf("CharData: %s\n", string(node))
	}
}

func printIndent(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
}
