//“Exercise 5.19:
//Use panic and recover to write a function that contains
//no return statement yet returns a non-zero value.”

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("A", a, err)
	}

	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("B", b, err)
	}

	result := devide(a, b)
	fmt.Println(result)
}

func devide(a, b int) (result float64) {
	defer func() {
		if r := recover(); r != nil {
			result = 11
		}
	}()

	result = float64(a / b)

	panic("Ooops! Somethings went wrong!!!!!!!")
}
