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
	filesName := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filesName)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filesName)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for _, j := range filesName[line] {
				fmt.Println(j)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, filesName map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		filesName[input.Text()] = append(filesName[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}
