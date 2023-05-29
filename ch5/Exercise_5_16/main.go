//“Exercise 5.16:
//Write a variadic version of strings.Join.”

package main

import (
	"fmt"
)

func joinStrings(sep string, strings ...string) string {
	result := ""
	for i, s := range strings {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}

func main() {
	result := joinStrings("-", "Hello", "world", "from", "Go")
	fmt.Println(result) // Output: Hello-world-from-Go
}
