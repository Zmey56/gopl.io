//“Exercise 7.10:
//The sort.Interface type can be adapted to other uses.
//Write a function
//
//IsPalindrome(s sort.Interface) bool that reports whether
//the sequence s is a palindrome, in other words, reversing the
//sequence would not change it.
//
//Assume that the elements at indices i and j are
//equal if !s.Less(i, j) && !s.Less(j, i).”

package main

import "sort"

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !equal(s, i, j) {
			return false
		}
	}
	return true
}

func equal(s sort.Interface, i, j int) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

// Example usage
func main() {
	intSlice := sort.IntSlice([]int{1, 2, 3, 2, 1})
	strSlice := sort.StringSlice([]string{"hello", "world", "world", "hello"})

	isIntSlicePalindrome := IsPalindrome(intSlice)
	isStrSlicePalindrome := IsPalindrome(strSlice)

	println("Is IntSlice palindrome?", isIntSlicePalindrome) // Output: true
	println("Is StrSlice palindrome?", isStrSlicePalindrome) // Output: true
}
