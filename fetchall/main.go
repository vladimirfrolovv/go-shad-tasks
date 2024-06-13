//go:build !solution

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	argsWithoutProg := os.Args[1:]
	start := time.Now()
	messages := make(chan string, len(argsWithoutProg))
	for _, url := range argsWithoutProg {
		go readUrl(url, messages)
	}

	for i := 0; i < len(argsWithoutProg); i++ {
		fmt.Println(<-messages)
	}

	fmt.Printf("%fs elapsed", time.Since(start).Seconds())
	close(messages)
}
func readUrl(url string, messages chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		messages <- "fetch: Get " + url + ": unsupported protocol scheme \"\"\n"

	} else {
		defer resp.Body.Close()
		elapsed := time.Since(start).Seconds()
		messages <- fmt.Sprintf("%fs\t%d\t%s\n", elapsed, resp.ContentLength, url)
	}

}
