//“Exercise 4.4:
//Write a version of rotate that operates in a single pass.”

package main

import "fmt"

func main() {
	//!+array
	a := []int{0, 1, 2, 3, 4, 5}
	result := rotate(a)
	fmt.Println(result) // "[5 4 3 2 1 0]"
	//!-array

}

func rotate(a []int) []int {
	var c []int
	for i := len(a) - 1; i >= 0; i-- {
		c = append(c, a[i])
	}
	return c
}
