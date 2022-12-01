//“Exercise 4.12:
//The popular web comic xkcd has a JSON interface.
//For example, a request to https://xkcd.com/571/info.0.json
//produces a detailed description of comic 571, one of many favorites.
// Download each URL (once!) and build an offline index.
// Write a tool xkcd that, using this index, prints the URL and
//transcript of each comic that matches a search term provided on the
//command line.”

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("usage: go run main.go 571")
		os.Exit(1)
	}

	url := "http://xkcd.com/" + args[0] + "/info.0.json"
	m, err := xkcd(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
	for k, v := range m {
		fmt.Printf("%s\t\t%-s\n", k, v)
	}
}

func xkcd(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	json.Unmarshal(b, &result)
	return result, nil
}
