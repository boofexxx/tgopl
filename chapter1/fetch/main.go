package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		var (
			resp *http.Response
			err  error
		)
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			resp, err = http.Get(url)
		} else {
			resp, err = http.Get("https://" + url)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Status code:", resp.Status)
	}
}
