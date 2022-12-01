//“Exercise 5.9:
//Write a function expand(s string, f func(string) string) string that
//replaces each substring “$foo” within s by the text returned
//by f("foo").”

package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}
	fmt.Printf("%s\n", expand(os.Args[1], f))
}

func f(s string) string {
	return strings.ToUpper(s)
}

func expand(s string, f func(string) string) string {
	var (
		sb, r    string
		inside   bool = false
		lastChar rune
	)
	if len(s) == 0 {
		return s
	}
	lastChar, size := utf8.DecodeRuneInString(s)
	if lastChar != '$' {
		r += string(lastChar)
	}

	for _, c := range s[size:] {
		if lastChar == '$' {
			if unicode.IsLetter(c) {
				inside = true
			} else {
				// flush if '$' doesn't start an varname
				r += string(lastChar)
			}
		}
		if inside && !(unicode.IsLetter(c)) {
			inside = false
			r += f(sb)
			sb = ""
		}
		lastChar = c
		if c == '$' {
			continue
		}
		if inside {
			sb += string(c)
		} else {
			r += string(c)
		}
	}
	if sb != "" {
		r += f(sb)
	}
	// flush $ at end of input string
	if lastChar == '$' {
		r += string(lastChar)
	}
	return r
}

func usage() {
	fmt.Fprintf(os.Stderr, "\nUsage: %s <string_to_expand>\n",
		os.Args[0])
}
