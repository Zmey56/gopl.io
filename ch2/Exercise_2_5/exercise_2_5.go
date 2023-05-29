//“Exercise 2.5:
//The expression x&(x-1) clears the rightmost non-zero bit of
//x.  Write a version of PopCount that counts bits by
//using this fact, and assess its performance.”

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

func PopCountByClearing(x uint64) int {
	result := 0
	for x != 0 {
		tmp := strconv.FormatInt(int64(x), 2)
		log.Println("1", x, " - ", tmp)
		log.Println("2", x&(x-1))
		x = x & (x - 1)
		tmp1 := strconv.FormatInt(int64(x), 2)
		log.Println("3", x, " - ", tmp1)

		result++
		log.Println("4", result)
	}
	return result

}

func main() {
	log.Println(PopCountByClearing(0x1234567890ABCDEF))
}
