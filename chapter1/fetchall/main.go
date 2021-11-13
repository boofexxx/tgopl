package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const path = "out.txt"

func main() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			go fetch(fmt.Sprintf("%d.txt", i), url, ch)
		} else {
			go fetch(fmt.Sprintf("%d.txt", i), "https://"+url, ch)
		}
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(path string, url string, ch chan<- string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stdout, "fetch: %v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(f, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
