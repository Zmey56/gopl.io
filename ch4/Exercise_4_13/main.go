//“Exercise 4.13:
//The JSON-based web service of the Open Movie Database
//lets you search https://omdbapi.com/ for a movie by name and
//download its poster image.
//Write a tool poster that downloads the poster image for the
//movie named on the command line.”

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//const omdbAPIURL = "https://www.omdbapi.com/?apikey=YOUR_API_KEY&t="

type Movie struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	DVD        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}

func main() {
	apiKey := os.Getenv("OMDB_API_KEY")

	// Check if a movie name is provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide a movie name.")
		return
	}

	movieName := strings.Join(os.Args[1:], "+")

	downloadPoster(apiKey, movieName)

}

func downloadPoster(apiKey, movieName string) {
	// Build the URL for fetching movie data
	url := "http://www.omdbapi.com?t=" + movieName + "&apikey=" + apiKey

	// Send a GET request to the OMDb API
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch movie data in Get: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch movie data in Status: %s", resp.Status)
	}

	// Parse the JSON response
	var movie Movie

	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		log.Fatalf("Failed to parse movie data: %v", err)
	}

	// Check if the movie data is valid
	if movie.Poster == "N/A" {
		log.Fatalf("No poster available for the movie: %s", movieName)
	}

	// Download the poster image
	resp, err = http.Get(movie.Poster)
	if err != nil {
		log.Fatalf("Failed to download poster image: %v", err)
	}
	defer resp.Body.Close()

	// Create a file to save the poster image
	filePath := movieName + ".jpg"
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// Save the poster image to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalf("Failed to save poster image: %v", err)
	}

	fmt.Printf("Poster image downloaded successfully: %s\n", filePath)
}
