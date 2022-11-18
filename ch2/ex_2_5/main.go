//“Exercise 2.5:
//The expression x&(x-1) clears the rightmost non-zero bit of
//x.  Write a version of PopCount that counts bits by
//using this fact, and assess its performance.”

package main

import (
	"fmt"
	"strconv"
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

	start_3 := time.Now()
	fmt.Println(PopCountLoop64(0x1234567890ABCDEF))
	fmt.Println("Time duration:", time.Since(start_3))

	start_4 := time.Now()
	fmt.Println(PopCountLoopClear(0x1234567890ABCDEF))
	fmt.Println("Time duration:", time.Since(start_4))
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

func PopCountLoop64(x uint64) int {
	count := 0
	for i := 0; i <= 63; i++ {
		count = count + int(pc[byte(x>>(i*8))])
	}
	return count
}

func PopCountLoopClear(x int64) int {
	count := 0
	for x != 0 {
		fmt.Println(strconv.FormatInt(x, 2))
		x = x & (x - 1)
		count++
	}
	return count
}
