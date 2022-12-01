//“Exercise 4.10:
//Modify issues to report the results in age categories,
//say less than a month old, less than a year old, and more
//than a year old.”

package main

import (
	"fmt"
	"gopl.io/ch4/github"
	_ "gopl.io/ch4/github"
	"log"
	"os"
	"time"
)

const (
	lessThanAMonth int = iota
	lessThanAYear
	moreThanAYear
)

func lessThanMonth(t time.Time) bool {
	return time.Since(t) <= time.Hour*24*30
}

func lessThanYear(t time.Time) bool {
	return time.Since(t) <= time.Hour*24*365
}

func main() {
	ageCategories := [3][]*github.Issue{}

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	// Exercise 4.10: Modify issues to report the results in age categories, say
	// less than a month old, less than a year old, and more than a year old.
	for _, item := range result.Items {
		switch {
		case lessThanMonth(item.CreatedAt):
			ageCategories[lessThanAMonth] = append(ageCategories[lessThanAMonth], item)
		case lessThanYear(item.CreatedAt):
			ageCategories[lessThanAYear] = append(ageCategories[lessThanAYear], item)
		default:
			ageCategories[moreThanAYear] = append(ageCategories[moreThanAYear], item)
		}
	}

	if len(ageCategories[lessThanAMonth]) != 0 {
		fmt.Println("--- Less Than A Month Old ---")
		for _, item := range ageCategories[lessThanAMonth] {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	if len(ageCategories[lessThanAYear]) != 0 {
		fmt.Println("--- Less Than A Year Old ---")
		for _, item := range ageCategories[lessThanAYear] {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	if len(ageCategories[moreThanAYear]) != 0 {
		fmt.Println("--- More Than A Year Old ---")
		for _, item := range ageCategories[moreThanAYear] {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
