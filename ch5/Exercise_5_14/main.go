//“Exercise 5.14:
//Use the breadthFirst function to explore a different structure.
//
//For example, you could use the course dependencies from the
//topoSort example (a directed graph), the file system hierarchy
//on your computer (a tree), or a list of bus or subway routes
//downloaded from your city government’s web site (an undirected graph).

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	rootDir := "/Users/zmey56/Documents" // Set the root directory path

	err := breadthFirst(rootDir, visit)
	if err != nil {
		log.Fatal(err)
	}
}

func visit(path string) {
	fmt.Println(path)
}

func breadthFirst(root string, visit func(path string)) error {
	queue := []string{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		visit(current)

		infos, err := ioutil.ReadDir(current)
		if err != nil {
			return err
		}

		for _, info := range infos {
			if info.IsDir() {
				subdir := filepath.Join(current, info.Name())
				queue = append(queue, subdir)
			}
		}
	}

	return nil
}
