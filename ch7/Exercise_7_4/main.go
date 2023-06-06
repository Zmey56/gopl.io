//“Exercise 7.4:
//The strings.NewReader function returns a value that satisfies
//the io.Reader interface (and others) by reading from its
//argument, a string. Implement a simple version of NewReader
//yourself, and use it to
//make the HTML parser (§5.2) take input from a
//string.”

package main

import (
	"fmt"
	"io"
)

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

type StringReader struct {
	data  string
	index int
}

func NewReader(s string) *StringReader {
	return &StringReader{data: s, index: 0}
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	if sr.index >= len(sr.data) {
		return 0, io.EOF
	}

	n = copy(p, sr.data[sr.index:])
	sr.index += n
	return n, nil
}

func ParseHTML(html string) *Node {
	reader := NewReader(html)
	node := &Node{
		Type: DocumentNode,
	}

	stack := []*Node{node}
	current := node

	for {
		token, err := getNextToken(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error:", err)
			break
		}

		switch token.Type {
		case StartTagToken:
			element := &Node{
				Type: ElementNode,
				Data: token.Data,
			}
			current.FirstChild = element
			current = element
			stack = append(stack, current)
		case EndTagToken:
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				current = stack[len(stack)-1]
			}
		case TextToken:
			textNode := &Node{
				Type: TextNode,
				Data: token.Data,
			}
			current.FirstChild = textNode
		}
	}

	return node
}

type TokenType int

const (
	StartTagToken TokenType = iota
	EndTagToken
	TextToken
)

type Token struct {
	Type TokenType
	Data string
}

func getNextToken(reader *StringReader) (Token, error) {
	var token Token

	for {
		b := make([]byte, 1)
		_, err := reader.Read(b)
		if err != nil {
			return token, err
		}

		if b[0] == '<' {
			token.Type = StartTagToken
			token.Data, err = readTag(reader)
			return token, err
		} else if b[0] == '>' {
			token.Type = EndTagToken
			return token, nil
		} else {
			token.Type = TextToken
			token.Data = string(b)
			return token, nil
		}
	}
}

func readTag(reader *StringReader) (string, error) {
	var tag string

	for {
		b := make([]byte, 1)
		_, err := reader.Read(b)
		if err != nil {
			return tag, err
		}

		if b[0] == '>' {
			break
		}

		tag += string(b)
	}

	return tag, nil
}

func main() {
	html := "<html><body><h1>Hello, World!</h1></body></html>"
	node := ParseHTML(html)

	fmt.Println(node)
}
