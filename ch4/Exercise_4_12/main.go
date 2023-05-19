//“Exercise 4.12:
//The popular web comic xkcd has a JSON interface.
//For example, a request to https://xkcd.com/571/info.0.json
//produces a detailed description of comic 571, one of many favorites.
//Download each URL (once!) and build an offline index.
//Write a tool xkcd that, using this index, prints the URL and
//transcript of each comic that matches a search term provided on the
//command line.”

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

const (
	baseURL  = "https://xkcd.com"
	jsonPath = "info.0.json"
)

func main() {
	tmpArg := []string{}
	for _, word := range os.Args[1:] {
		tmpArg = append(tmpArg, word)
	}

	searchTerm := strings.Join(tmpArg, " ")

	if searchTerm != "" {
		searchComics(searchTerm)
	} else {
		buildIndex()
	}
}

func buildIndex() {
	// Create a directory for storing the comic index
	indexDir := "./xkcd_index"
	err := os.MkdirAll(indexDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create index directory: %v", err)
	}

	// Download and save each comic's JSON data
	comicNotFound := 0
	for i := 0; ; i++ {
		comicURL := fmt.Sprintf("%s/%d/%s", baseURL, i, jsonPath)
		resp, err := http.Get(comicURL)
		if err != nil {
			log.Printf("Failed to get comic %d: %v", i, err)
			break
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Comic %d not found", i)
			comicNotFound++
			if comicNotFound > 5 {
				break
			}
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read comic %d response body: %v", i, err)
			continue
		}

		comic := Comic{}
		err = json.Unmarshal(body, &comic)
		if err != nil {
			log.Printf("Failed to unmarshal comic %d JSON: %v", i, err)
			continue
		}

		// Save the comic's JSON data to a file
		filePath := filepath.Join(indexDir, fmt.Sprintf("%d.json", i))
		err = os.WriteFile(filePath, body, os.ModePerm)
		if err != nil {
			log.Printf("Failed to save comic %d JSON: %v", i, err)
			continue
		}
	}

	fmt.Println("Comic index created successfully!")
}

func searchComics(searchTerm string) {
	// Load and search the comic index for matching transcripts
	indexDir := "./xkcd_index"
	files, err := os.ReadDir(indexDir)
	if err != nil {
		log.Fatalf("Failed to read index directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(indexDir, file.Name())
		body, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Failed to read comic JSON file: %v", err)
			continue
		}

		comic := Comic{}
		err = json.Unmarshal(body, &comic)
		if err != nil {
			log.Printf("Failed to unmarshal comic JSON: %v", err)
			continue
		}

		// Check if the search term is present in the comic transcript
		if strings.Contains(strings.ToLower(comic.Transcript), strings.ToLower(searchTerm)) {
			fmt.Printf("Comic URL: %s/%d\n", baseURL, comic.Num)
			fmt.Printf("Transcript: %s\n", comic.Transcript)
			fmt.Println("-------------------------")
		}
	}
}
