//“Exercise 2.4:
//Write a version of PopCount that counts bits by shifting
//its argument through 64 bit positions, testing the rightmost bit
//each time.  Compare its performance to the table-lookup version.”

package main

import "log"

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
func PopCountShifting(x uint64) int {
	result := 0
	for i := uint64(0); i < 64; i++ {
		log.Println(x)
		log.Println(x & 1)
		if x&1 != 0 {
			result++
		}
		x = x >> 1
	}
	return result
}

func main() {
	log.Println(PopCountShifting(1000))
}
