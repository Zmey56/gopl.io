//“Exercise 3.12:
//Write a function that reports whether two strings are anagrams of each
//other, that is, they contain the same letters in a different order.”

package main

import (
	"fmt"
	"os"
	"sort"
)

type sortRunes []rune

func main() {
	for i := 1; i < len(os.Args)-1; i++ {
		fmt.Printf("  %v\n", areAnagrams(os.Args[i], os.Args[i+1]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func areAnagrams(wordOne, wordTwo string) bool {
	if len(wordOne) != len(wordTwo) {
		return false
	}

	if SortString(wordOne) == SortString(wordTwo) {
		return true
	} else {
		return false
	}
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
