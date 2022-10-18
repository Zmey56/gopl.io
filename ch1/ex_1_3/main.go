//“Exercise 1.3:
//Experiment to measure the difference in running time between our
//potentially inefficient versions and the one that uses strings.Join.
//(Section 1.6 illustrates part of the time package,
//and Section 11.4 shows how to write benchmark tests
//for systematic performance evaluation.)”

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	start_1 := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	elapsed_1 := time.Since(start_1)
	log.Printf("First took %s", elapsed_1)

	start_2 := time.Now()
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	elapsed_2 := time.Since(start_2)
	log.Printf("Second took %s", elapsed_2)

	start_3 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	elapsed_3 := time.Since(start_3)
	log.Printf("Third took %s", elapsed_3)
}
