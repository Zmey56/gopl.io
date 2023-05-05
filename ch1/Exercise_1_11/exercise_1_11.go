//“Exercise 1.11:
//Try fetchall with longer argument lists, such as samples
//from the top million web sites available at alexa.com.  How
//does the program behave if a web site just doesn’t respond?
//(Section 8.9 describes mechanisms for coping
//in such cases.)”

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func readCSVFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	records := readCSVFile("top-1m.csv")
	url := ""
	for _, numberURL := range records {
		if !strings.HasPrefix(numberURL[1], "http://") {
			url = "http://" + numberURL[1]
		} else {
			url = numberURL[1]
		}
		go fetch(url, ch)
	}

	for range records {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
