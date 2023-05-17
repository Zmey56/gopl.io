//“Exercise 4.3:
//Rewrite reverse to use an array pointer instead of a slice.”

package main

import "fmt"

func main() {

	tmp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverse(&tmp)
	fmt.Println(tmp)
}

func reverse(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
