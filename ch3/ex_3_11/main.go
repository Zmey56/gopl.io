//Exercise 3.11:
//Enhance comma so that it deals correctly with floating-point
//numbers and an optional sign.

package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func lenUntilDot(s string) int {
	length := 0
	for ; length < len(s); length++ {
		if s[length] == '.' {
			return length
		}
	}
	return length
}

func comma(s string) string {
	var buf bytes.Buffer

	if lenUntilDot(s) <= 3 {
		return s
	}
	if s[0] == '-' || s[0] == '+' {
		buf.WriteString(s[0:1])
		s = s[1:]
	}
	start := lenUntilDot(s) % 3
	if start == 0 {
		start = 3
	}
	buf.WriteString(s[0:start])
	i := start
	for ; i < lenUntilDot(s); i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}
	buf.WriteString(s[i:])
	return buf.String()
}
