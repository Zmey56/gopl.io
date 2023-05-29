//Exercise 5.15:
//Write variadic functions max and min, analogous to
//sum.
//What should these functions do when called with no arguments?
//Write variants that require at least one argument.”

package main

import "fmt"

func max(vals ...int) int {
	largeValue := 0
	for _, val := range vals {
		if val > largeValue {
			largeValue = val
		}
	}
	return largeValue
}

func min(vals ...int) int {
	littleValue := 0
	for i, val := range vals {
		if i == 0 {
			littleValue = val
			continue
		}
		if val < littleValue {
			littleValue = val
		}
	}
	return littleValue
}

func main() {
	//!+main
	fmt.Println(max())           //  "0"
	fmt.Println(min())           //  "0"
	fmt.Println(max(3))          //  "3"
	fmt.Println(min(3))          //  "3"
	fmt.Println(max(1, 2, 3, 4)) //  "4"
	fmt.Println(min(1, 2, 3, 4)) //  "1"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(max(values...)) // "4"
	fmt.Println(min(values...)) // "1"
	//!-slice
}
