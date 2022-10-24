// “Exercise 1.11
// Try fetchall with longer argument lists, such as samples
// from the top million web sites available at alexa.com.  How
// does the program behave if a web site just doesn’t respond?
// (Section 8.9 describes mechanisms for coping
// in such cases.)”
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

const http_prefix = "https://"

func main() {
	start := time.Now()
	ch := make(chan string)
	records := readCsvFile("top-1m.csv")
	for _, url := range records {
		go fetch111(url[1], ch)
	}
	for range records {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func readCsvFile(filePath string) [][]string {
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

func fetch111(url string, ch chan<- string) {
	start := time.Now()
	if strings.HasPrefix(url, http_prefix) != true {
		url = http_prefix + url
	}
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
	line := fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)

	resp.Body.Close()

	_, err = file.WriteString(line + "\n")

	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	//log.Println(l, "bytes written successfully")

	file.Close()

	ch <- line
}
