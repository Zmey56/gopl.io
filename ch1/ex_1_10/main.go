//“Exercise 1.10:
//Find a web site that produces a large amount of data.  Investigate
//caching by running fetchall twice in succession to see whether the reported time
//changes much.  Do you get the same content each time?  Modify
//fetchall to print its output to a file so it can be examined.”

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	file, err := os.OpenFile("data_url.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		ch <- fmt.Sprintf("Open %s: %v", file, err)
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	line := fmt.Sprintf("%.2f s %7d %s", secs, nbytes, url)

	resp.Body.Close()

	l, err := file.WriteString(line + "\n")

	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	log.Println(l, "bytes written successfully")

	file.Close()

	ch <- line
}
