//“Exercise 2.3:
//Rewrite PopCount to use a loop instead of a single expression.
//Compare the performance of the two versions.
//(Section 11.4 shows how to compare the
//performance of different implementations systematically.)”

package main

import (
	"fmt"
	"time"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	start := time.Now()
	fmt.Println(PopCount(0x1234567890ABCDEF))
	fmt.Println("Time duration:", time.Since(start))

	start_2 := time.Now()
	fmt.Println(PopCountLoop(0x1234567890ABCDEF))
	fmt.Println("Time duration:", time.Since(start_2))
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//!-

func PopCountLoop(x uint64) int {
	count := 0
	for i := 0; i <= 7; i++ {
		count = count + int(pc[byte(x>>(i*8))])
	}
	return count
}

//!-
