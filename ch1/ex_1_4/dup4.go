//“Exercise 1.4:
//Modify dup2 to print the names of all files in which
//each duplicated line occurs.”

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countsFile := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, countsFile, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
				continue
			}
			countLines(f, countsFile, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	for linef, nf := range countsFile {
		if nf > 0 {
			fmt.Printf("%s\n", linef)
		}
	}
}

func countLines(f *os.File, countsFile map[string]int, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, found := counts[input.Text()]; found {
			//fmt.Println(input.Text(), v, f.Name())
			countsFile[f.Name()]++
		}
		//fmt.Println(counts)
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
