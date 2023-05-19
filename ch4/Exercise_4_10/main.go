//Exercise 4.10:
//Modify issues to report the results in age categories,
//say less than a month old, less than a year old, and more
//than a year old.

package main

import (
	"fmt"
	"gopl.io/ch4/github"
	"log"
	"os"
	"time"
)

// !+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	now := time.Now()
	for _, item := range result.Items {
		timeDiff := now.Sub(item.CreatedAt)
		var category string
		if timeDiff < 30*24*time.Hour { // less than a month old
			category = "Less than a month old"
		} else if timeDiff < 365*24*time.Hour { // less than a year old
			category = "Less than a year old"
		} else { // more than a year old
			category = "More than a year old"
		}
		fmt.Printf("#%-5d %9.9s %.55s (%s)\n",
			item.Number, item.User.Login, item.Title, category)
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
