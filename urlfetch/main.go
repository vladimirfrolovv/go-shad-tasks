//go:build !solution

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	for _, url := range argsWithoutProg {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Print("fetch: Get " + url + ": unsupported protocol scheme \"\"")
			os.Exit(1)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

}
