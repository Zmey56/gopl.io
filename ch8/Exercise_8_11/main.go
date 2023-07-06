//“Exercise 8.11:
//Following the approach of mirroredQuery in Section 8.4.4, implement a variant of fetch that
//requests several URLs concurrently.
//
//As soon as the first response arrives, cancel the other requests.”

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func mirroredFetch(urls []string) string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	responses := make(chan string, len(urls))

	for _, url := range urls {
		go func(url string) {
			select {
			case <-ctx.Done():
				return
			default:
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}

				select {
				case <-ctx.Done():
					return
				case responses <- string(body):
					cancel()
					return
				}
			}
		}(url)
	}

	return <-responses // Return the first response received
}

func main() {
	urls := []string{
		"http://www.acornpub.co.kr/book/go-programming",
		"http://www.williamspublishing.com/Books/978-5-8459-2051-5.html",
		"http://helion.pl/ksiazki/jezyk-go-poznaj-i-programuj-alan-a-a-donovan-brian-w-kernighan,jgopop.htm",
	}

	response := mirroredFetch(urls)
	fmt.Println(response)
}
