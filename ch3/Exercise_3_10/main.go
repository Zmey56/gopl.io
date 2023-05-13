//“Exercise 3.10:
//Write a non-recursive version of comma, using bytes.Buffer
//instead of string concatenation.”

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

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}

	start := n % 3

	if start == 0 {
		start = 3
	}

	buf.WriteString(s[0:start])
	for i := start; i < len(s); i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
