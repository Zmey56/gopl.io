package main

import "fmt"

func main() {
	var tests = []struct {
		s1, s2 string
		want   bool
	}{
		{"a", "a", true},
		{"a", "b", false},
		{"ab", "ba", true},
		{"debit card", "bad credit", true},
		{"punishments", "nine thumps", true},
		{"ab", "baba", false},
	}
	for _, test := range tests {
		got := isAnagram(test.s1, test.s2)
		fmt.Printf("isAnagram(%q, %q) = %v, want %v\n", test.s1, test.s2, got, test.want)
	}
}

// Exercise 3.12: Write a function that reports whether two strings are anagrams of each other, that is, they contain the same letters in a different order.
func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[rune]int)
	for _, v := range s1 {
		m[v]++
	}
	for _, v := range s2 {
		m[v]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
