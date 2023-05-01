//“Exercise 1.3:
//Experiment to measure the difference in running time between our
//potentially inefficient versions and the one that uses strings.Join.
//(Section 1.6 illustrates part of the time package,
//and Section 11.4 shows how to write benchmark tests
//for systematic performance evaluation.)”

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Println("First time: ", time.Since(start).Microseconds(), "ms")

	start1 := time.Now()

	s1, sep1 := "", ""
	for _, arg := range os.Args[1:] {
		s1 += sep1 + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Println("Second time: ", time.Since(start1).Microseconds(), "ms")

	start2 := time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println("First time: ", time.Since(start2).Microseconds(), "ms")

}
