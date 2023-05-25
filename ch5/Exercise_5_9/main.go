// “Exercise 5.9:
// Write a function expand(s string, f func(string) string) string that
// replaces each substring “$foo” within s by the text returned
// by f("foo").”
package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "$", "$$") // Replace "$" with "$$" to prevent unintended substitution

	var result strings.Builder
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '$' && i+1 < len(s) && isAlphaNumeric(s[i+1]) {
			// Found a "$" followed by a valid identifier character
			result.WriteString(s[start:i])
			start = i
			end := i + 1
			for end < len(s) && isAlphaNumeric(s[end]) {
				end++
			}
			identifier := s[i+1 : end]
			result.WriteString(f(identifier))
			start = end
			i = end - 1
		}
	}

	result.WriteString(s[start:])
	return result.String()
}

func isAlphaNumeric(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')
}

func main() {
	s := "Hello, $name! Today is $day."
	result := expand(s, func(identifier string) string {
		if identifier == "name" {
			return "John Doe"
		} else if identifier == "day" {
			return "Monday"
		}
		return ""
	})
	fmt.Println(result)
}
