// Exercise 8.9: Write a version of du that computes and periodically displays separate totals
// for each of the root directories.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Create a channel for each subdirectory to collect file sizes.
	resultChans := make([]chan map[string]int64, len(roots))
	for i := range resultChans {
		resultChans[i] = make(chan map[string]int64)
	}

	// Traverse the file tree.
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(root, &n, resultChans[i])
	}

	go func() {
		n.Wait()
		for _, ch := range resultChans {
			close(ch)
		}
	}()

	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

loop:
	for {
		select {
		case size, ok := <-resultChans[0]:
			if !ok {
				break loop // resultChans was closed
			}
			for k, v := range size {
				fmt.Println(k, "-", v)
			}

		case <-tick:
			for _, ch := range resultChans {
				for size := range ch {
					for k, v := range size {
						fmt.Println(k, "-", v)
					}
				}
			}
		}
	}
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- map[string]int64) {
	defer n.Done()

	copiedData := make(map[string]int64)
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Add(1)
			go walkDir(subdir, n, fileSizes)
		} else {
			sizeFile, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du: %v\n", err)
				continue
			}
			copiedData[dir] = sizeFile.Size()
		}
	}

	fileSizes <- copiedData
}

func dirents(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
