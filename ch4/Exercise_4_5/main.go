//“Exercise 4.5:
//Write an in-place function to eliminate adjacent duplicates
//in a []string slice.”

package main

import "fmt"

func main() {
	data := []string{"one", "one", "one", "three", "three"}
	fmt.Println(elemDublicas(data))
}

func elemDublicas(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	writeindex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != arr[writeindex] {
			writeindex++
			arr[writeindex] = arr[i]
		}
	}

	return arr[:writeindex+1]
}
