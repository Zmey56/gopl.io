// Exercise 9.2:
//Rewrite the PopCount example from  Section 2.6.2 so that it
//initializes the lookup table using sync.Once the first time it  is needed.
//
//(Realistically, the cost of synchronization would be prohibitive for a
//small and highly optimized function like PopCount.)”

package main

import (
	"log"
	"sync"
)

var pc [256]byte
var once sync.Once

func populatePC() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func initialize() {
	once.Do(populatePC)
	log.Println("pc", pc)
}

func PopCount(x uint64) int {
	initialize()
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	x := uint64(12345)
	count := PopCount(x)
	log.Println("PopCount:", count)
}
