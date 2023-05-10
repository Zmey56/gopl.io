//“Exercise 2.3:
//Rewrite PopCount to use a loop instead of a single expression.
//Compare the performance of the two versions.
//
//(Section 11.4 shows how to compare the
//performance of different implementations systematically.)”

package main

import (
	"log"
	"strconv"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	//log.Println("pc", pc)
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//log.Println(i, "pc[i] -", pc[i], "pc[i/2] - ", pc[i/2], "i/2 - ", i/2, "byte(i&1) - ", byte(i&1), "i&1 - ", i&1)
	}
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

// PopCountLoop returns the population count (number of set bits) of x in Loop.
func PopCountLoop(x uint64) int {
	tmp := strconv.FormatInt(int64(x), 2)
	result := 0
	for _, j := range tmp {
		if j == 49 {
			result = result + 1
		}
	}
	return result
}

func main() {
	log.Println(PopCount(0x1234567890ABCDEF))
	log.Println(PopCountLoop(0x1234567890ABCDEF))

}
